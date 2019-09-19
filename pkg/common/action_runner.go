package common

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ActionRunner struct {
	client    client.Client
	logger    logr.Logger
	lastError error
}

type Action interface {
	run(runner *ActionRunner) (string, error)
}

type DesiredClusterState []Action

type ExistsConfigMapAction struct {
	Ref *v1.ConfigMap
	Msg string
}

type CreateConfigMapAction struct {
	Ref *v1.ConfigMap
	Msg string
}

type UpdateConfigMapAction struct {
	Ref *v1.ConfigMap
	Msg string
}

type OnAction struct {
	Ref     *v1.ConfigMap
	Success Action
	Fail    Action
}

func NewActionRunner(client client.Client, logger logr.Logger) *ActionRunner {
	return &ActionRunner{
		client: client,
		logger: logger,
	}
}

func (i *ActionRunner) RunAll(actions []Action) error {
	for index, action := range actions {
		msg, err := action.run(i)
		if err != nil {
			i.lastError = err
			i.logger.Info(fmt.Sprintf("(%d) %-15s %s", index, "FAILED", msg))
			continue
		}

		i.lastError = nil
		i.logger.Info(fmt.Sprintf("(%d) %-15s %s", index, "SUCCESS", msg))
	}

	return i.lastError
}

func (i *ExistsConfigMapAction) run(runner *ActionRunner) (string, error) {
	// Don't continue if there was a previous error
	if runner.lastError != nil {
		return i.Msg, runner.lastError
	}

	selector := client.ObjectKey{
		Name:      i.Ref.Name,
		Namespace: i.Ref.Namespace,
	}

	return i.Msg, runner.client.Get(context.TODO(), selector, i.Ref)
}

func (i *CreateConfigMapAction) run(runner *ActionRunner) (string, error) {
	// Don't continue if there was a previous error
	if runner.lastError != nil {
		return i.Msg, runner.lastError
	}

	return i.Msg, runner.client.Create(context.TODO(), i.Ref)
}

func (i *UpdateConfigMapAction) run(runner *ActionRunner) (string, error) {
	// Don't continue if there was a previous error
	if runner.lastError != nil {
		return i.Msg, runner.lastError
	}

	return i.Msg, runner.client.Update(context.TODO(), i.Ref)
}

func (i *OnAction) run(runner *ActionRunner) (string, error) {
	if runner.lastError != nil {
		runner.lastError = nil
		return i.Fail.run(runner)
	} else {
		runner.lastError = nil
		return i.Success.run(runner)
	}
}
