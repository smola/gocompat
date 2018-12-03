package compat

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/packages"
)

// ReachableFromPackages gets all top-level exported types.Object reachable from
// the given list of package paths.
func ReachableFromPackages(pkgs ...string) ([]types.Object, error) {
	r := newReachability()
	if err := r.FromPackages(pkgs...); err != nil {
		return nil, err
	}

	return r.GetReachable(), nil
}

type reachability struct {
	seen      map[types.Type]bool
	reachable map[types.Object]bool
}

func newReachability() *reachability {
	return &reachability{
		seen:      make(map[types.Type]bool),
		reachable: make(map[types.Object]bool),
	}
}

func (r *reachability) GetReachable() []types.Object {
	result := make([]types.Object, 0, len(r.reachable))
	for obj := range r.reachable {
		result = append(result, obj)
	}

	return result
}

func (r *reachability) FromPackages(pkgs ...string) error {
	conf := &packages.Config{
		Mode:  packages.LoadTypes,
		Tests: false,
	}

	loadedPackages, err := packages.Load(conf, pkgs...)
	if err != nil {
		return err
	}

	for _, pkg := range loadedPackages {
		if err := r.fromPackage(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (r *reachability) fromPackage(pkg *packages.Package) error {
	if len(pkg.Errors) > 0 {
		return fmt.Errorf("has errors")
	}

	scope := pkg.Types.Scope()
	for _, name := range scope.Names() {
		if !ast.IsExported(name) {
			continue
		}

		obj := scope.Lookup(name)
		if err := r.reachFromObject(pkg.Types, obj); err != nil {
			return err
		}
	}

	return nil
}

func (r *reachability) reachFromObject(pkg *types.Package, obj types.Object) error {
	if r.reachable[obj] {
		return nil
	}
	r.reachable[obj] = true

	if !obj.Exported() {
		return nil
	}

	return r.reachFromType(pkg, obj.Type())
}

func (r *reachability) reachFromType(pkg *types.Package, typ types.Type) error {
	if r.seen[typ] {
		return nil
	}
	r.seen[typ] = true

	switch typ := typ.(type) {
	case *types.Named:
		obj := typ.Obj()
		opkg := obj.Pkg()
		if opkg == nil {
			return nil
		}

		if opkg != pkg {
			return r.reachFromObject(opkg, obj)
		}

		if !typ.Obj().Exported() {
			return nil
		}

		for i := 0; i < typ.NumMethods(); i++ {
			m := typ.Method(i)
			if err := r.reachFromType(pkg, m.Type()); err != nil {
				return err
			}
		}

		return nil
	case *types.Pointer:
		return r.reachFromType(pkg, typ.Elem())
	case *types.Struct:
		for i := 0; i < typ.NumFields(); i++ {
			f := typ.Field(i)
			if !f.Exported() {
				continue
			}

			if err := r.reachFromType(pkg, f.Type()); err != nil {
				return err
			}
		}

		return nil
	case *types.Signature:
		if err := r.reachFromType(pkg, typ.Params()); err != nil {
			return err
		}

		return r.reachFromType(pkg, typ.Results())
	case *types.Tuple:
		for i := 0; i < typ.Len(); i++ {
			f := typ.At(i)
			if err := r.reachFromType(pkg, f.Type()); err != nil {
				return err
			}
		}

		return nil
	case *types.Slice:
		return r.reachFromType(pkg, typ.Elem())
	case *types.Map:
		if err := r.reachFromType(pkg, typ.Key()); err != nil {
			return err
		}

		return r.reachFromType(pkg, typ.Elem())
	case *types.Basic:
		return nil
	default:
		return fmt.Errorf("unhandled type: %f", reflect.TypeOf(typ))
	}
}
