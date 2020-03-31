package linter

import (
	"fmt"
	"io"
	"os"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/madvikinggod/kubelint/pkg/message"
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
	objs, err := readObjects(os.Stdin)
	if err != nil {
		return err
	}

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

func readObjects(reader io.Reader) ([]*yaml.RNode, error) {
	r := kio.ByteReader{
		Reader: reader,
	}
	return r.Read()
}

func (l *linter) lintObjects(objs []*yaml.RNode) []message.Message {
	msgs := []message.Message{}
	for _, obj := range objs {
		id, err := obj.GetMeta()
		if err != nil {
			msgs = append(msgs, message.SimpleMessage{
				Name:   "Linting",
				Info:   "Could not get objects metadata:\n" + obj.MustString(),
				IsCrit: true,
			})
			continue
		}
		tm := yaml.TypeMeta{
			Kind:       id.Kind,
			APIVersion: id.APIVersion,
		}
		rules, found := l.cfg.Rules[tm]
		if !found {
			msgs = append(msgs, message.KMessage{
				RuleName: "Linting",
				Info:     "No rule found",
				Id:       id.GetIdentifier(),
				IsCrit:   false,
			})
			continue
		}
		for _, rule := range rules {
			msg := rule(obj, id.GetIdentifier())
			if msg != nil {
				msgs = append(msgs, msg)
			}
		}
	}
	return msgs
}
