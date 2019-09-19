package common

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Manifest interface {
	Add(res runtime.Object)
}

type ManifestImpl struct {
	ref *v1.ConfigMap
}

func NewManifest(cr *kc.Keycloak) *ManifestImpl {
	manifestConfigMap := &v1.ConfigMap{
		ObjectMeta: v12.ObjectMeta{
			Name:      "manifest",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{},
	}

	return &ManifestImpl{
		ref: manifestConfigMap,
	}
}

func (i *ManifestImpl) Add(res runtime.Object) {
}
