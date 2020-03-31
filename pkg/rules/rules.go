package rules

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type Rule func(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message
