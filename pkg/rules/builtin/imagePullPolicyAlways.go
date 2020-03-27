package builtin

import (
	"fmt"
	"github.com/madvikinggod/kubelint/pkg/message"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func imagePullPolicyConvertError(typ string, obj *unstructured.Unstructured) message.Message {
	return simpleMessage{
		name:       "ImagePullPolicyAlways",
		info:       "Could not convert object to " + typ,
		gvk:gvk(obj),
		nName:nName(obj),
		isCritical: true,
	}
}

func ImagePullPolicyAlwaysCheck(obj *unstructured.Unstructured) message.Message {
	// A deployment is used here to stand for an object with the path
	// spec.template.spec.containers and
	// spec.temaptle.spec.initContainers
	deploy := &appsv1.Deployment{}
	scheme := getScheme()
	err := scheme.Convert(obj, deploy, nil)
	if err != nil {
		return imagePullPolicyConvertError("apps.v1/Deployment", obj)
	}

	for _,container := range deploy.Spec.Template.Spec.Containers {
		if msg:= imagePullPolicyAlwaysCheckContainer(container); msg != nil {
			msg.addObjInfo(obj)
			return *msg
		}
	}
	for _,container := range deploy.Spec.Template.Spec.InitContainers {
		if msg:= imagePullPolicyAlwaysCheckContainer(container); msg != nil {
			msg.addObjInfo(obj)
			return *msg
		}
	}
	return nil
}



func imagePullPolicyAlwaysCheckContainer(container corev1.Container) *simpleMessage {
	if container.ImagePullPolicy != corev1.PullAlways {
		msg := simpleMessage{
			name:       "ImagePullPolicyAlways",
			info:       fmt.Sprintf("container %s has an image pull policy of %s, it should be Alaways", container.Name, container.ImagePullPolicy),
			isCritical: true,
		}
		return &msg
	}
	return nil
}


func init() {
	registerRule(ImagePullPolicyAlwaysCheck,[]string{
		"apps.v1/Deployment",
		"apps.v1/StatefulSet",
		"apps.v1/ReplicaSet",
		"apps.v1/DaemonSet",
		"batch.v1/Job",
	})
}
