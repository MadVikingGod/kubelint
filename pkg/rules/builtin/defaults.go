package builtin

import (
	"fmt"
	"github.com/madvikinggod/kubelint/pkg/rules"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var DefaultRules = map[string][]rules.Rule{}

func registerRule(r rules.Rule, gvks []string) {
	for _, gvk := range gvks {
		rls, found := DefaultRules[gvk]
		if !found {
			DefaultRules[gvk] = []rules.Rule{r}
			continue
		}
		DefaultRules[gvk] = append(rls, r)
	}
}

func gvk(obj *unstructured.Unstructured) string {
	return obj.GetAPIVersion() + "/" + obj.GetKind()
}
func nName(obj *unstructured.Unstructured) string {
	return obj.GetNamespace() + "/" + obj.GetName()
}

type simpleMessage struct {
	name string
	info string
	gvk string
	nName string
	isCritical bool
}

func (s simpleMessage) String() string {
	return fmt.Sprintf("%s - %s - %s %s", s.name, s.info, s.gvk, s.nName)
}

func (s simpleMessage) IsCritical() bool {
	   return s.isCritical
}

