package testdata

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)



func LatestImageTagUnstructured() *unstructured.Unstructured {
	scheme := runtime.NewScheme()
	appsv1.AddToScheme(scheme)

	d := &appsv1.Deployment{}
	yaml.Unmarshal([]byte(LatestImageTagYaml), d)

	o := &unstructured.Unstructured{}
	scheme.Convert(d, o, nil)

	return o
}
