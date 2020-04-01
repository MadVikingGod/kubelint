package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func NakedPodCheck(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
	return message.KMessage{
		RuleName: "NakedPodCheck",
		Info:     "Pods should not be used directly, apps/v1 Deployments are recommended",
		ID:       id,
		IsCrit:   true,
	}
}

func init() {

	registerRule(NakedPodCheck, []yaml.TypeMeta{
		{APIVersion: "v1", Kind: "Pod"},
	})
}
