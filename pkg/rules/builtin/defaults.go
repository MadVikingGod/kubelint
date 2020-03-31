package builtin

import (
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/madvikinggod/kubelint/pkg/rules"
)

var DefaultRules = map[yaml.TypeMeta][]rules.Rule{}

func registerRule(r rules.Rule, gvks []yaml.TypeMeta) {
	for _, gvk := range gvks {
		rls, found := DefaultRules[gvk]
		if !found {
			DefaultRules[gvk] = []rules.Rule{r}
			continue
		}
		DefaultRules[gvk] = append(rls, r)
	}
}
