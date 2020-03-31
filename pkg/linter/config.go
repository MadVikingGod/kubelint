package linter

import (
	"github.com/madvikinggod/kubelint/pkg/rules"
	"github.com/madvikinggod/kubelint/pkg/rules/builtin"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// Config is used to control high level functionality, like fail on warnings or which rules are checked..
type Config struct {
	Rules map[yaml.TypeMeta][]rules.KRule
}

// DefaultConfig provides a default config, the one used if no flags are toggeled.
func DefaultConfig() Config {
	return Config{
		Rules: builtin.DefaultKRules,
	}
}
