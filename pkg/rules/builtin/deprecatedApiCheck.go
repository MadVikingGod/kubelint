package builtin

import (
	"fmt"

	"github.com/madvikinggod/kubelint/pkg/message"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func depractedAPICheck(newVersion yaml.TypeMeta) rules.KRule {
	return func(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
		tm := yaml.TypeMeta{
			Kind:       id.Kind,
			APIVersion: id.APIVersion,
		}
		return message.KMessage{
			RuleName: "DeprecatedAPICheck",
			Info:     fmt.Sprintf("%v should not be used, use %v", tm, newVersion),
			Id:       id,
			IsCrit:   true,
		}
	}
}

func init() {
	registerRule(depractedAPICheck(yaml.TypeMeta{"NetworkPolicy", "networking.k8s.io.v1"}), []yaml.TypeMeta{
		{"NetworkPolicy", "extensions/v1beta1"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{"PodSecurityPolicy", "extensions.v1beta1"}), []yaml.TypeMeta{
		{"PodSecurityPolicy", "policy.v1beta1"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{"DaemonSet", "apps.v1"}), []yaml.TypeMeta{
		{"DaemonSet", "extensions.v1beta1"},
		{"DaemonSet", "apps.v1beta2"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{"Deployment", "apps.v1"}), []yaml.TypeMeta{
		{"Deployment", "extensions.v1beta1"},
		{"Deployment", "apps.v1beta1"},
		{"Deployment", "apps.v1beta2"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{"StatefulSet", "apps.v1"}), []yaml.TypeMeta{
		{"StatefulSet", "apps.v1beta1"},
		{"StatefulSet", "apps.v1beta2"},
	})
	registerRule(depractedAPICheck(yaml.TypeMeta{"ReplicaSet", "apps.v1"}), []yaml.TypeMeta{
		{"ReplicaSet", "extensions.v1beta1"},
		{"ReplicaSet", "apps.v1beta1"},
		{"ReplicaSet", "apps.v1beta2"},
	})

}
