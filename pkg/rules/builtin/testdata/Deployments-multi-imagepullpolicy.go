package testdata

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var MultiImagePullPolicyYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: multiImagePullPolicy
  namespace: multiImagePullPolicy
spec:
  template:
    spec:
      containers:
        - name: multiImagePullPolicyContainer
          imagePullPolicy: Always
        - name: multiImagePullPolicyContainer-fail
          imagePullPolicy: Never

`

func MultiImagePullPolicyUnstructured() *unstructured.Unstructured {
	scheme := runtime.NewScheme()
	appsv1.AddToScheme(scheme)

	d := &appsv1.Deployment{}
	yaml.Unmarshal([]byte(MultiImagePullPolicyYaml), d)

	o := &unstructured.Unstructured{}
	scheme.Convert(d, o, nil)

	return o
}

