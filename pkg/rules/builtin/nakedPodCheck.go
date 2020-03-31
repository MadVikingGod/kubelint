package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

var NakedPodCheckStr = "Pods should not be used directly. Use a deployment instead"
var NakedPodCheckMsg = simpleMessage{
	name:       "NakedPodCheck",
	info:       NakedPodCheckStr,
	isCritical: true,
}

func NakedPodCheck(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
	return message.KMessage{
		RuleName: "NakedPodCheck",
		Info:     "Pods should not be used directly, apps/v1 Deployments are recommended",
		Id:       id,
		IsCrit:   true,
	}
}

func init() {

	registerRule(NakedPodCheck, []yaml.TypeMeta{
		{"Pod", "v1"},
	})
}
