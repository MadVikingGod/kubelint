package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/message"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type NakedPodCheck struct{}

var NakedPodCheckStr = "Pods should not be used directly. Use a deployment instead"
var NakedPodCheckMsg = simpleMessage{
	name: "NakedPodCheck",
	info: NakedPodCheckStr,
	isCritical:true,
}

func (n NakedPodCheck) Check(obj *unstructured.Unstructured) message.Message {
	msg := NakedPodCheckMsg
	msg.gvk = gvk(obj)
	msg.nName = nName(obj)
	return msg
}

var _ rules.Rule = NakedPodCheck{}

func init() {
	registerRule(NakedPodCheck{}, []string{"v1/pod"})
}
