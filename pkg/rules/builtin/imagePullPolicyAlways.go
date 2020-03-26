package builtin

import (
	"fmt"
	"github.com/madvikinggod/kubelint/pkg/message"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ImagePullPolicyAlways struct{}

var imagePullPolicyAlwaysStr = "imagePullPolicy should be set to Always"

func (i ImagePullPolicyAlways) Check(obj *unstructured.Unstructured) message.Message {
	// spec.template.spec.[init]containers[].imagePullPolicy == Always

	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	if !found || len(containers) == 0 || err != nil {
		msg := simpleMessage{
			name:       "ImagePullPolicyAlways",
			info:       "object did not have a container",
			gvk:        gvk(obj),
			nName:      nName(obj),
			isCritical: true,
		}
		return &msg
	}

	for _, inter := range containers {
		container := inter.(map[string]interface{})
		if container["imagePullPolicy"] != "Always" {
			msg := simpleMessage{
				name:       "ImagePullPolicyAlways",
				info:       fmt.Sprintf("container %s has an image pull policy of %s, it should be Alaways", container["name"], container["imagePullPolicy"]),
				gvk:        gvk(obj),
				nName:      nName(obj),
				isCritical: true,
			}
			return msg
		}
	}

	containers, found, err = unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "initContainers")
	for _, inter := range containers {
		container := inter.(map[string]interface{})
		if container["imagePullPolicy"] != "Always" {
			msg := simpleMessage{
				name:       "ImagePullPolicyAlways",
				info:       fmt.Sprintf("container %s has an image pull policy of %s, it should be Alaways", container["name"], container["imagePullPolicy"]),
				gvk:        gvk(obj),
				nName:      nName(obj),
				isCritical: true,
			}
			return msg
		}
	}

	return nil
}

var _ rules.Rule = ImagePullPolicyAlways{}

func init() {
	registerRule(ImagePullPolicyAlways{}, []string{
		"apps.v1/Deployment",
		"apps.v1/ReplicaSet",
		"apps.v1/DaemonSet",
		"apps.v1/StatefulSet",
		"apps.v1/Job",
	})
}
