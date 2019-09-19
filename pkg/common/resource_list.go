package common

import v1 "k8s.io/api/core/v1"

type On struct {
	Success Action
	Fail    Action
}

type KcResource interface {
	Exists(msg string) Action
	Create(msg string) Action
	Update(msg string) Action
	Branch(on On) Action
}

type KcConfigMap struct {
	Ref v1.ConfigMap
}

func (i KcConfigMap) Exists(msg string) Action {
	return &ExistsConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i KcConfigMap) Update(msg string) Action {
	return &UpdateConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i KcConfigMap) Create(msg string) Action {
	return &CreateConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i KcConfigMap) Branch(on On) Action {
	return &OnAction{
		ref:     &i.Ref,
		success: on.Success,
		fail:    on.Fail,
	}
}
