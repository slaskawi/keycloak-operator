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
		Ref: &i.Ref,
		Msg: msg,
	}
}

func (i KcConfigMap) Update(msg string) Action {
	return &UpdateConfigMapAction{
		Ref: &i.Ref,
		Msg: msg,
	}
}

func (i KcConfigMap) Create(msg string) Action {
	return &CreateConfigMapAction{
		Ref: &i.Ref,
		Msg: msg,
	}
}

func (i KcConfigMap) Branch(on On) Action {
	return &OnAction{
		Ref:     &i.Ref,
		Success: on.Success,
		Fail:    on.Fail,
	}
}
