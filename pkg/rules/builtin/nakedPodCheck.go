package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var NakedPodCheckStr = "Pods should not be used directly. Use a deployment instead"
var NakedPodCheckMsg = simpleMessage{
	name:       "NakedPodCheck",
	info:       NakedPodCheckStr,
	isCritical: true,
}

func NakedPodCheck(obj *unstructured.Unstructured) message.Message {
	msg := NakedPodCheckMsg
	msg.gvk = gvk(obj)
	msg.nName = nName(obj)
	return msg
}

func init() {
	registerRule(NakedPodCheck, []string{"v1/pod"})
}
