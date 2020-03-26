package linter

import (
	"github.com/madvikinggod/kubelint/pkg/rules"
	"github.com/madvikinggod/kubelint/pkg/rules/builtin"
)

type LinterConfig struct {
	Rules map[string][]rules.Rule
}

func DefaultConfig() LinterConfig {
	return LinterConfig{
		Rules: builtin.DefaultRules,
	}
}
