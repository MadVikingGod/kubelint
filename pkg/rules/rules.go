package rules

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

//type Rule interface {
//	Check(obj *unstructured.Unstructured) message.Message
//}

type Rule func(obj *unstructured.Unstructured) message.Message

type KRule func(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message
