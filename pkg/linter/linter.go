package linter

import (
	"fmt"
	"io"

	"github.com/madvikinggod/kubelint/pkg/message"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type linter struct {
	cfg Config
}

func NewLinter(cfg Config) *linter {
	return &linter{
		cfg: cfg,
	}
}

func (l *linter) Run(src io.Reader) error {

	//read each yaml file into a "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	objs := []*unstructured.Unstructured{}

	msgs := l.lintObjects(objs)

	fail := 0
	for _, msg := range msgs {
		fmt.Println(msg.String())
		if msg.IsCritical() {
			fail++
		}
	}
	if fail > 0 {
		return fmt.Errorf("Encountered %d errors", fail)
	}

	return nil
}

func (l *linter) lintObjects(objs []*unstructured.Unstructured) []message.Message {
	msgs := []message.Message{}
	for _, obj := range objs {
		gvk := obj.GetAPIVersion() + "/" + obj.GetKind()
		rules, found := l.cfg.Rules[gvk]
		if !found {
			namespacedName := fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName())
			msg := message.SimpleMessage{
				Name:   "NoRuleFound",
				Info:   "Did not find any rules for",
				Gvk:    gvk,
				NName:  namespacedName,
				IsCrit: false,
			}
			msgs = append(msgs, msg)
		}
		for _, rule := range rules {
			msg := rule.Check(obj)
			if msg != nil {
				msgs = append(msgs, msg)
			}
		}
	}
	return msgs
}
