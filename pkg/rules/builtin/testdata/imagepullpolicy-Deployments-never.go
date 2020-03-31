package testdata

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var NeverImagePullPolicyYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: neverImagePullPolicy
  namespace: neverImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: neverImagePullPolicyContainer
          imagePullPolicy: Never
`

func NeverImagePullPolicyUnstructured() *unstructured.Unstructured {
	scheme := runtime.NewScheme()
	appsv1.AddToScheme(scheme)

	d := &appsv1.Deployment{}
	yaml.Unmarshal([]byte(NeverImagePullPolicyYaml), d)

	o := &unstructured.Unstructured{}
	scheme.Convert(d, o, nil)

	return o
}
