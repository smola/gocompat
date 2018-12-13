package compat

import (
	"fmt"
	"strings"
)

//go:generate stringer -type=ChangeType

type ChangeType int

const (
	_ ChangeType = iota

	// PackageDeleted means that a package that was added explicitly to the API
	// was deleted. This is not used when a package that is reachable but not
	// explicitly added to the API is deleted. In that case, a package deletion
	// would be seen as type changes that render the package not reachable
	// anymore.
	PackageDeleted

	// TopLevelDeclarationAdded means that a top-level declaration was added to
	// a package. This can be either a type, var, const or func. It does not
	// include method declarations.
	TopLevelDeclarationAdded

	// TopLevelDeclarationDeleted means that a top-level declaration is deleted
	// from a package. It does not include method declarations.
	TopLevelDeclarationDeleted

	// DeclarationTypeChanged means that the type of declaration changed. For
	// example, a declaration changed from being a var to a const. Note that
	// changing between type definitions and type aliases is covered here too.
	DeclarationTypeChanged

	// TypeChanged means that an entity (type declaration, field, parameter,
	// return type) has changed its type.
	//
	// This does not cover when there is a change in the type underlying a type
	// name. For example, TypeChanged will be signaled if a var changes from
	// MyTypeA to MyTypeB, but it will not if MyType itself changes from being
	// a struct to an interface.
	//
	// A change in the direction of a channel is also considered a type change.
	//
	// See https://golang.org/ref/spec#Types
	TypeChanged

	// FieldAdded means that a new exported field was added to a struct.
	FieldAdded

	// FieldDeleted means that a previously exported field was deleted from a
	// struct.
	FieldDeleted

	// FieldChangedType means that a field in a struct changed its type.
	// See TypeChanged.
	FieldChangedType

	// SignatureChanged means that a function declaration changed its signature.
	SignatureChanged

	// MethodAdded means that a new exported method declaration was added.
	MethodAdded

	// MethodDeleted means that a previously exported method declaration was
	// deleted.
	MethodDeleted

	// MethodSignatureChanged means that the signature of a method has changed.
	// Receiver name, parameter names and return type names are ignored.
	MethodSignatureChanged

	// InterfaceChanged means that any function, exported or not, was added,
	// deleted or had its signature changed.
	InterfaceChanged
)

type Change struct {
	Type   ChangeType
	Symbol string
}

func (c Change) String() string {
	return fmt.Sprintf("%s %s", c.Symbol, c.Type.String())
}

func init() {
	for i := 0; i < len(_ChangeType_index)-1; i++ {
		s, e := _ChangeType_index[i], _ChangeType_index[i+1]
		name := _ChangeType_name[s:e]
		name = strings.ToLower(name)
		lookupChangeType[name] = ChangeType(i + 1)
	}
}

var lookupChangeType = make(map[string]ChangeType)

// ChangeTypeFromString converts a string representation of the ChangeType to
// its numeric value.
func ChangeTypeFromString(s string) (ChangeType, error) {
	c, ok := lookupChangeType[strings.ToLower(s)]
	if !ok {
		return 0, fmt.Errorf("invalid change type: %s", s)
	}

	return c, nil
}
