package builtin

import (
	"fmt"
	"strings"

	"github.com/madvikinggod/kubelint/pkg/message"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func ippCheck(container *yaml.RNode) error {
	n, _ := container.Pipe(yaml.LookupCreate(yaml.ScalarNode, "name"))
	name := strings.TrimSpace(n.MustString())
	i, _ := container.Pipe(yaml.LookupCreate(yaml.ScalarNode, "imagePullPolicy"))
	ipp := strings.TrimSpace(i.MustString())
	if ipp != "Always" {
		return fmt.Errorf("container %s has an ImagePullPolocy of %s, it should be Always", name, ipp)
	}

	return nil
}

func ImagePullPolicyAlwaysCheck(obj *yaml.RNode, id yaml.ResourceIdentifier) message.Message {
	containers, err := obj.Pipe(yaml.Lookup("spec", "template", "spec", "containers"))
	if err != nil {
		return message.KMessage{
			RuleName: "ImagePullPolicyAlways",
			Info:     "Could not find containers",
			ID:       id,
			IsCrit:   true,
		}
	}
	err = containers.VisitElements(ippCheck)
	if err != nil {
		return message.KMessage{
			RuleName: "ImagePullPolicyAlways",
			Info:     err.Error(),
			ID:       id,
			IsCrit:   true,
		}
	}

	containers, err = obj.Pipe(yaml.LookupCreate(yaml.SequenceNode, "spec", "template", "spec", "initContainers"))
	if err != nil {
		return message.KMessage{
			RuleName: "ImagePullPolicyAlways",
			Info:     "Could not find initContainers",
			ID:       id,
			IsCrit:   true,
		}
	}
	err = containers.VisitElements(ippCheck)
	if err != nil {
		return message.KMessage{
			RuleName: "ImagePullPolicyAlways",
			Info:     err.Error(),
			ID:       id,
			IsCrit:   true,
		}
	}

	return nil
}

func init() {

	registerRule(ImagePullPolicyAlwaysCheck, []yaml.TypeMeta{
		{APIVersion: "apps/v1", Kind: "Deployment"},
		{APIVersion: "apps/v1", Kind: "StatefulSet"},
		{APIVersion: "apps/v1", Kind: "ReplicaSet"},
		{APIVersion: "apps/v1", Kind: "DaemonSet"},
		{APIVersion: "batch/v1", Kind: "Job"},
	})
}
