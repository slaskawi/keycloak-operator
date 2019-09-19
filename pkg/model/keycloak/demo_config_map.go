package keycloak

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DemoConfigMap(cr *kc.Keycloak) v1.ConfigMap {
	return v1.ConfigMap{
		ObjectMeta: v12.ObjectMeta{
			Name:      "demo-config-map",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"password": "secret",
		},
	}
}
