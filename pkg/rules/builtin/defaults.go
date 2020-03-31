package builtin

import (
	"fmt"
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/madvikinggod/kubelint/pkg/rules"
)

var DefaultRules = map[yaml.TypeMeta][]rules.KRule{}

func registerRule(r rules.KRule, gvks []yaml.TypeMeta) {
	for _, gvk := range gvks {
		rls, found := DefaultRules[gvk]
		if !found {
			DefaultRules[gvk] = []rules.KRule{r}
			continue
		}
		DefaultRules[gvk] = append(rls, r)
	}
}

type simpleMessage struct {
	name       string
	info       string
	isCritical bool
}

func (s simpleMessage) String() string {
	return fmt.Sprintf("%s - %s ", s.name, s.info)
}

func (s simpleMessage) IsCritical() bool {
	return s.isCritical
}
