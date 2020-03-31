package builtin

import (
	"fmt"
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/madvikinggod/kubelint/pkg/rules"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var DefaultRules = map[string][]rules.Rule{}
var DefaultKRules = map[yaml.TypeMeta][]rules.KRule{}

func registerKRule(r rules.KRule, gvks []yaml.TypeMeta) {
	for _, gvk := range gvks {
		rls, found := DefaultKRules[gvk]
		if !found {
			DefaultKRules[gvk] = []rules.KRule{r}
			continue
		}
		DefaultKRules[gvk] = append(rls, r)
	}
}
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

type object interface {
	GroupVersionKind() schema.GroupVersionKind
	GetObjectKind() schema.ObjectKind
	GetNamespace() string
	GetName() string
}

func gvk(obj object) string {
	gvk := obj.GroupVersionKind()
	return gvk.GroupVersion().String() + "/" + gvk.Kind
}
func nName(obj object) string {
	return obj.GetNamespace() + "/" + obj.GetName()
}

type simpleMessage struct {
	name       string
	info       string
	gvk        string
	nName      string
	isCritical bool
}

func (s simpleMessage) String() string {
	return fmt.Sprintf("%s - %s - %s %s", s.name, s.info, s.gvk, s.nName)
}

func (s simpleMessage) IsCritical() bool {
	return s.isCritical
}

func (s *simpleMessage) addObjInfo(obj object) {
	s.gvk = gvk(obj)
	s.nName = nName(obj)
}

var scheme *runtime.Scheme

func getScheme() *runtime.Scheme {
	if scheme != nil {
		return scheme
	}
	scheme = runtime.NewScheme()
	appsv1.AddToScheme(scheme)
	return scheme
}
