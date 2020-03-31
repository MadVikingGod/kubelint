package testdata

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var SSNoImagePullPolicyYaml = `apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: noImagePullPolicy
  namespace: noImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: noImagePullPolicyContainer
`

func SSNoImagePullPolicyUnstructured() *unstructured.Unstructured {
	scheme := runtime.NewScheme()
	appsv1.AddToScheme(scheme)

	d := &appsv1.StatefulSet{}
	yaml.Unmarshal([]byte(SSNoImagePullPolicyYaml), d)

	o := &unstructured.Unstructured{}
	scheme.Convert(d, o, nil)

	return o
}
