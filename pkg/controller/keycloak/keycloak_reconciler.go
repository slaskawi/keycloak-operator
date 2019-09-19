package keycloak

import (
	"fmt"
	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"github.com/keycloak/keycloak-operator/pkg/common"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Reconciler interface {
	Reconcile(CurrentState common.CurrentState) []common.Action
}

type KeycloakReconciler struct {

}

func (reconciler *KeycloakReconciler) Reconcile(CurrentState common.CurrentState, cr v1alpha1.Keycloak) []common.Action {
	var d []common.Action = []common.Action{}

	if CurrentState.DemoConfigMap == nil {
		log.Info(fmt.Sprint("Let's create a CM!"))

		cm := v1.ConfigMap{
			ObjectMeta: v12.ObjectMeta{
				Name:      "demo-config-map",
				Namespace: cr.Namespace,
			},
			Data: map[string]string{
				"password": "secret",
			},
		}

		action := &common.CreateConfigMapAction{
			Ref: &cm,
		}

		//Alternatively we can replace it with nice DSL
		//action := keycloak.DemoConfigMap(&cr).Create("Let's create a CM");
		d = append(d, action)

	} else {
		log.Info(fmt.Sprint("There is no need to create a CM"))
	}

	return d
}

