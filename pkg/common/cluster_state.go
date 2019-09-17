package common

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ClusterAction interface {
	Run(client.Client, logr.Logger) error
}

type ClusterState struct {
	Actions []ClusterAction
	Client  client.Client
	Logger  logr.Logger
}

type Create struct {
	Obj runtime.Object
}

type Ensure struct {
	Fn func(client.Client) error
}

func NewClusterState(c client.Client, logger logr.Logger) *ClusterState {
	return &ClusterState{
		Actions: []ClusterAction{},
		Client:  c,
		Logger:  logger,
	}
}

func (i *ClusterState) Put(a ClusterAction) {
	i.Actions = append(i.Actions, a)
}

func (i *ClusterState) RunAll() error {
	for _, action := range i.Actions {
		err := action.Run(i.Client, i.Logger)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Create) Run(c client.Client, l logr.Logger) error {
	name := i.Obj.(v1.Object).GetName()
	l.Info(fmt.Sprint(fmt.Sprintf("Cluster Action CREATE: %v", name)))

	err := c.Create(context.TODO(), i.Obj)
	if err != nil {
		// If the object already exist still try to update it
		if errors.IsAlreadyExists(err) {
			return c.Update(context.TODO(), i.Obj)
		}
	}
	return err
}

func (i Ensure) Run(c client.Client, l logr.Logger) error {
	l.Info("Cluster Action ENSURE")
	return i.Fn(c)
}
