package message

import "fmt"

type Msg struct {
	Message string
	IsError bool
}

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
