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
	ref *v1.ConfigMap
	msg string
}

type CreateConfigMapAction struct {
	ref *v1.ConfigMap
	msg string
}

type UpdateConfigMapAction struct {
	ref *v1.ConfigMap
	msg string
}

type OnAction struct {
	ref     *v1.ConfigMap
	success Action
	fail    Action
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
		return i.msg, runner.lastError
	}

	selector := client.ObjectKey{
		Name:      i.ref.Name,
		Namespace: i.ref.Namespace,
	}

	return i.msg, runner.client.Get(context.TODO(), selector, i.ref)
}

func (i *CreateConfigMapAction) run(runner *ActionRunner) (string, error) {
	// Don't continue if there was a previous error
	if runner.lastError != nil {
		return i.msg, runner.lastError
	}

	return i.msg, runner.client.Create(context.TODO(), i.ref)
}

func (i *UpdateConfigMapAction) run(runner *ActionRunner) (string, error) {
	// Don't continue if there was a previous error
	if runner.lastError != nil {
		return i.msg, runner.lastError
	}

	return i.msg, runner.client.Update(context.TODO(), i.ref)
}

func (i *OnAction) run(runner *ActionRunner) (string, error) {
	if runner.lastError != nil {
		runner.lastError = nil
		return i.fail.run(runner)
	} else {
		runner.lastError = nil
		return i.success.run(runner)
	}
}
