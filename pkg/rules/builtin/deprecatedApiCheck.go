package builtin

import (
	"fmt"

	"github.com/madvikinggod/kubelint/pkg/message"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func depractedAPICheck(newVersion yaml.TypeMeta) rules.Rule {
	return func(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
		tm := yaml.TypeMeta{
			Kind:       id.Kind,
			APIVersion: id.APIVersion,
		}
		return message.KMessage{
			RuleName: "DeprecatedAPICheck",
			Info:     fmt.Sprintf("%v should not be used, use %v", tm, newVersion),
			ID:       id,
			IsCrit:   true,
		}
	}
}

func init() {
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "networking.k8s.io/v1", Kind: "NetworkPolicy"}), []yaml.TypeMeta{
		{APIVersion: "extensions/v1beta1", Kind: "NetworkPolicy"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "extensions/v1beta1", Kind: "PodSecurityPolicy"}), []yaml.TypeMeta{
		{APIVersion: "policy/v1beta1", Kind: "PodSecurityPolicy"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "apps/v1", Kind: "DaemonSet"}), []yaml.TypeMeta{
		{APIVersion: "extensions/v1beta1", Kind: "DaemonSet"},
		{APIVersion: "apps/v1beta2", Kind: "DaemonSet"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"}), []yaml.TypeMeta{
		{APIVersion: "extensions/v1beta1", Kind: "Deployment"},
		{APIVersion: "apps/v1beta1", Kind: "Deployment"},
		{APIVersion: "apps/v1beta2", Kind: "Deployment"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "apps/v1", Kind: "StatefulSet"}), []yaml.TypeMeta{
		{APIVersion: "apps/v1beta1", Kind: "StatefulSet"},
		{APIVersion: "apps/v1beta2", Kind: "StatefulSet"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{APIVersion: "apps/v1", Kind: "ReplicaSet"}), []yaml.TypeMeta{
		{APIVersion: "extensions/v1beta1", Kind: "ReplicaSet"},
		{APIVersion: "apps/v1beta1", Kind: "ReplicaSet"},
		{APIVersion: "apps/v1beta2", Kind: "ReplicaSet"},
	})

}
