package builtin

import (
	"fmt"
	"github.com/madvikinggod/kubelint/pkg/message"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"strings"
)

func ImageTagMissingCheckDeployment(obj *unstructured.Unstructured) message.Message {
	scheme := getScheme()
	deploy := &appsv1.Deployment{}
	err := scheme.Convert(obj, deploy, nil)
	if err != nil {
		return imageTagMissingConvertError("apps.v1/Deployment", obj)
	}

	for _, container := range deploy.Spec.Template.Spec.Containers {
		if msg := imageTagMissingCheckContainer(container); msg != nil {
			msg.addObjInfo(obj)
			return *msg
		}
	}
	return nil
}

func imageTagMissingCheckContainer(container corev1.Container) *simpleMessage {
	image := container.Image
	s := strings.Split(image, ":")
	if len(s) < 2 {
		return &simpleMessage{
			name:       "ImageTagMissing",
			info:       fmt.Sprintf("No tag detected on contaner %s", container.Name),
			isCritical: false,
		}
	}
	if s[1] == "latest" {
		return &simpleMessage{
			name:       "ImageTagMissing",
			info:       fmt.Sprintf("'latest' tag on contaner %s", container.Name),
			isCritical: false,
		}
	}
	return nil
}

func imageTagMissingConvertError(typ string, obj *unstructured.Unstructured) message.Message {
	return simpleMessage{
		name:       "ImageTagMissing",
		info:       "Could not convert object to " + typ,
		gvk:        gvk(obj),
		nName:      nName(obj),
		isCritical: true,
	}
}
