package compat

import "fmt"

//go:generate stringer -type=ChangeType

type ChangeType int

const (
	_ ChangeType = iota
	TypeChanged
	FieldAdded
	FieldDeleted
	FieldChangedType
	SignatureChanged
	MethodAdded
	MethodDeleted
	MethodSignatureChanged
	InterfaceChanged
)

type Change struct {
	Type   ChangeType
	Symbol string
}

func (c Change) String() string {
	return fmt.Sprintf("%s %s", c.Symbol, c.Type.String())
}
