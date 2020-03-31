package message

import (
	"fmt"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type Message interface {
	String() string
	IsCritical() bool
}

type SimpleMessage struct {
	Name   string
	Info   string
	Gvk    string
	NName  string
	IsCrit bool
}

func (s SimpleMessage) String() string {
	return fmt.Sprintf("%s - %s - %s %s", s.Name, s.Info, s.Gvk, s.NName)
}

func (s SimpleMessage) IsCritical() bool {
	return s.IsCrit
}

type KMessage struct {
	RuleName string
	Info     string
	Id       yaml.ResourceIdentifier
	IsCrit   bool
}

func (k KMessage) String() string {
	return fmt.Sprintf("%s - %s - %s", k.RuleName, k.Info, k.Id)
}
func (k KMessage) IsCritical() bool {
	return k.IsCrit
}
