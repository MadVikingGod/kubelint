package builtin

import (
	"fmt"
	"github.com/madvikinggod/kubelint/pkg/message"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)
// https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/
func depractedAPICheck( newVersion string) rules.Rule {
	return func(obj *unstructured.Unstructured) message.Message {
		return message.SimpleMessage{
			Name:   "DeprecatedAPICheck",
			Info:   fmt.Sprintf("%s should not be used, use %s", gvk(obj), newVersion),
			Gvk:    gvk(obj),
			NName:  nName(obj),
			IsCrit: true,
		}
	}
}

func init() {
	registerRule(depractedAPICheck("networking.k8s.io.v1/NetworkPolicy"), []string{
		"extensions.v1beta1/NetworkPolicy",
	})
	registerRule(depractedAPICheck("extensions.v1beta1/PodSecurityPolicy"), []string{
		"policy.v1beta1/PodSecurityPolicy",
	})
	registerRule(depractedAPICheck("apps.v1/DaemonSet"), []string{
		"extensions.v1beta1/DaemonSet",
		"apps.v1beta2/DaemonSet",
	})
	registerRule(depractedAPICheck("apps.v1/Deployment"), []string{
		"extensions.v1beta1/Deployment",
		"apps.v1beta1/Deployment",
		"apps.v1beta2/Deployment",
	})
	registerRule(depractedAPICheck("apps.v1/StatefulSet"), []string{
		"apps.v1beta1/StatefulSet",
		"apps.v1beta2/StatefulSet",
	})
	registerRule(depractedAPICheck("apps.v1/ReplicaSet"), []string{
		"extensions.v1beta1/ReplicaSet",
		"apps.v1beta1/ReplicaSet",
		"apps.v1beta2/ReplicaSet",
	})

}
