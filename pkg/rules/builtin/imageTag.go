package builtin

import (
	"fmt"
	"strings"

	"github.com/madvikinggod/kubelint/pkg/message"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func tagCheck(container *yaml.RNode) error {
	n, _ := container.Pipe(yaml.LookupCreate(yaml.ScalarNode, "name"))
	name := strings.TrimSpace(n.MustString())
	i, _ := container.Pipe(yaml.LookupCreate(yaml.ScalarNode, "image"))
	image := strings.TrimSpace(i.MustString())

	imageParts := strings.Split(image, ":")

	if len(imageParts) < 2 {
		return fmt.Errorf("container %s has no image tag", name)
	}

	if imageParts[1] == "latest" {
		return fmt.Errorf("container %s has an image tag of latest", name)
	}

	return nil
}

func ImageTagCheck(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
	containers, err := obj.Pipe(yaml.Lookup("spec", "template", "spec", "containers"))
	if err != nil {
		return message.KMessage{
			RuleName: "ImageTagCheck",
			Info:     "Could not find containers",
			ID:       id,
			IsCrit:   true,
		}
	}

	err = containers.VisitElements(tagCheck)
	if err != nil {
		return message.KMessage{
			RuleName: "ImageTagCheck",
			Info:     err.Error(),
			ID:       id,
			IsCrit:   true,
		}
	}

	return nil
}

func init() {
	registerRule(ImageTagCheck, []yaml.TypeMeta{
		{APIVersion: "apps/v1", Kind: "Deployment"},
		{APIVersion: "apps/v1", Kind: "StatefulSet"},
		{APIVersion: "apps/v1", Kind: "ReplicaSet"},
		{APIVersion: "apps/v1", Kind: "DaemonSet"},
		{APIVersion: "batch/v1", Kind: "Job"},
	})
}
