package compat

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"
)

func Compare(a, b *API) []Change {
	var changes []Change
	checked := make(map[types.Object]bool)

	//TODO: check added symbols to package
	for _, apkg := range a.Packages {
		ascope := apkg.Scope()
		found := false
		for _, bpkg := range b.Packages {
			if apkg.Path() == bpkg.Path() {
				bscope := bpkg.Scope()
				changes = append(changes, compareScopes(checked, ascope, bscope)...)
				found = true
				break
			}
		}

		if !found {
			//TODO: removed package
		}
	}

	for aobj := range a.Reachable {
		found := false
		for bobj := range b.Reachable {
			if aobj.Pkg().Path() == bobj.Pkg().Path() && aobj.Name() == bobj.Name() {
				changes = append(changes, CompareObjects(aobj, bobj)...)
				found = true
				break
			}
		}

		if !found {
			//TODO: not reachable anymore!
		}
	}

	return changes
}

func compareScopes(checked map[types.Object]bool, a, b *types.Scope) []Change {
	//TODO
	return nil
}

// CompareObjects compares two objects and reports backwards incompatible changes.
func CompareObjects(a, b types.Object) []Change {
	if a, ok := a.(*types.TypeName); ok && a.IsAlias() {
		return compareAliases(a, a, b.(*types.TypeName))
	}

	return compareTypes(a, a.Type(), b.Type())
}

func symbolName(parent types.Object, children ...types.Object) string {
	str := fmt.Sprintf(`"%s".%s`, parent.Pkg().Path(), parent.Name())
	for _, child := range children {
		str += fmt.Sprintf(".%s", child.Name())
	}
	return str
}

func compareTypes(obj types.Object, a, b types.Type) []Change {
	var changes []Change
	aUnderlying := a.Underlying()
	bUnderlying := b.Underlying()

	if basicType(aUnderlying) != basicType(bUnderlying) {
		changes = append(changes, Change{
			Type:   TypeChanged,
			Symbol: symbolName(obj),
		})
		return changes
	}

	switch a := aUnderlying.(type) {
	case *types.Struct:
		changes = append(changes, compareStruct(obj, a, bUnderlying.(*types.Struct))...)
	case *types.Signature:
		if !signaturesAreEqual(a, bUnderlying.(*types.Signature)) {
			changes = append(changes, Change{
				Type:   SignatureChanged,
				Symbol: symbolName(obj),
			})
		}
	case *types.Interface:
		b := bUnderlying.(*types.Interface)
		if len(compareMethods(false, obj, a, b)) != 0 {
			changes = append(changes, Change{
				Type:   InterfaceChanged,
				Symbol: symbolName(obj),
			})
		}
	case *types.Basic, *types.Map, *types.Slice, *types.Array:
		if basicType(a) != basicType(bUnderlying) {
			changes = append(changes, Change{
				Type:   TypeChanged,
				Symbol: symbolName(obj),
			})
		}
	default:
		panic(fmt.Sprintf("unhandled type: %s", reflect.TypeOf(a)))
	}

	switch a := a.(type) {
	case *types.Named:
		changes = append(changes, compareMethods(true, obj, a, b.(*types.Named))...)
	}

	return changes
}

func compareAliases(obj types.Object, a, b *types.TypeName) []Change {
	aType := a.Type()
	bType := b.Type()

	if basicType(aType) != basicType(bType) {
		return []Change{{
			Type:   TypeChanged,
			Symbol: symbolName(obj),
		}}
	}

	return nil
}

func basicType(a fmt.Stringer) string {
	str := a.String()
	idx := strings.Index(str, "{")
	if idx != -1 {
		str = str[:idx]
	}

	idx = strings.Index(str, "(")
	if idx != -1 {
		str = str[:idx]
	}

	return str
}

func compareStruct(obj types.Object, a, b *types.Struct) []Change {
	//TODO: report field order changes
	//TODO: report struct tag changes

	var changes []Change

	for i := 0; i < a.NumFields(); i++ {
		aField := a.Field(i)
		if !aField.Exported() {
			continue
		}
		found := false
		for j := 0; j < b.NumFields(); j++ {
			bField := b.Field(j)
			if aField.Name() == bField.Name() {
				found = true
				if basicType(aField.Type().Underlying()) != basicType(bField.Type().Underlying()) {
					changes = append(changes, Change{
						Type:   FieldChangedType,
						Symbol: symbolName(obj, aField),
					})
				}
			}
		}

		if !found {
			changes = append(changes, Change{
				Type:   FieldDeleted,
				Symbol: symbolName(obj, aField),
			})
		}
	}

	for i := 0; i < b.NumFields(); i++ {
		bField := b.Field(i)
		found := false
		for j := 0; j < a.NumFields(); j++ {
			aField := a.Field(j)
			found = found || aField.Name() == bField.Name()
		}

		if !found {
			changes = append(changes, Change{
				Type:   FieldAdded,
				Symbol: symbolName(obj, bField),
			})
		}
	}

	return changes
}

type methoder interface {
	NumMethods() int
	Method(int) *types.Func
}

func signaturesAreEqualForFunc(a, b *types.Func) bool {
	asig := a.Type().(*types.Signature)
	bsig := b.Type().(*types.Signature)
	return signaturesAreEqual(asig, bsig)
}

func signaturesAreEqual(a, b *types.Signature) bool {
	return tuplesAreEqual(a.Params(), b.Params()) &&
		tuplesAreEqual(a.Results(), b.Results())
}

func tuplesAreEqual(a, b *types.Tuple) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := 0; i < a.Len(); i++ {
		ael := a.At(i)
		bel := b.At(i)
		if basicType(ael.Type()) != basicType(bel.Type()) {
			return false
		}
	}

	return true
}

func compareMethods(exportedOnly bool, obj types.Object, a, b methoder) []Change {
	var changes []Change
	for i := 0; i < a.NumMethods(); i++ {
		aMethod := a.Method(i)
		if exportedOnly && !aMethod.Exported() {
			continue
		}
		found := false
		for j := 0; j < b.NumMethods(); j++ {
			bMethod := b.Method(j)
			if aMethod.Name() == bMethod.Name() {
				found = true
				if !signaturesAreEqualForFunc(aMethod, bMethod) {
					changes = append(changes, Change{
						Type:   MethodSignatureChanged,
						Symbol: symbolName(obj, aMethod),
					})
				}
			}
		}

		if !found {
			changes = append(changes, Change{
				Type:   MethodDeleted,
				Symbol: symbolName(obj, aMethod),
			})
		}
	}

	for i := 0; i < b.NumMethods(); i++ {
		bMethod := b.Method(i)
		if exportedOnly && !bMethod.Exported() {
			continue
		}
		found := false
		for j := 0; j < a.NumMethods(); j++ {
			aMethod := a.Method(j)
			found = found || aMethod.Name() == bMethod.Name()
		}

		if !found {
			changes = append(changes, Change{
				Type:   MethodAdded,
				Symbol: symbolName(obj, bMethod),
			})
		}
	}
	return changes
}
