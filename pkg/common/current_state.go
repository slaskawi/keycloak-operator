package common

import (
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type CurrentState struct {
	DemoConfigMap *v1.ConfigMap
	KeycloakDeployment *v12.StatefulSet
}