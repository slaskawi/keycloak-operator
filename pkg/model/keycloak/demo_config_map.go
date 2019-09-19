package keycloak

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"github.com/keycloak/keycloak-operator/pkg/common"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DemoConfigMap(cr *kc.Keycloak) common.KcConfigMap {
	cm := v1.ConfigMap{
		ObjectMeta: v12.ObjectMeta{
			Name:      "demo-config-map",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"password": "secret",
		},
	}

	return common.KcConfigMap{
		Ref: cm,
	}
}
