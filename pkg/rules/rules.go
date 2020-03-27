package rules

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//type Rule interface {
//	Check(obj *unstructured.Unstructured) message.Message
//}

type Rule func(obj *unstructured.Unstructured) message.Message