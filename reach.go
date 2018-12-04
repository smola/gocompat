package compat

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/packages"
)

// API
type API struct {
	// Packages included directly in the API.
	Packages []*types.Package
	// Reachable is the set of all objects that are reachable from the API.
	Reachable map[types.Object]bool
}

// NewAPI creates an empty API.
func NewAPI() *API {
	return &API{
		Reachable: make(map[types.Object]bool),
	}
}

// ReachableFromPackages gets an API given list of package paths.
func ReachableFromPackages(pkgs ...string) (*API, error) {
	return reachableFromPackages(false, pkgs...)
}

func reachableFromPackages(tests bool, pkgs ...string) (*API, error) {
	r := newReachability()
	if err := r.FromPackages(tests, pkgs...); err != nil {
		return nil, err
	}

	return r.API, nil
}

type reachability struct {
	API  *API
	seen map[types.Type]bool
}

func newReachability() *reachability {
	return &reachability{
		API:  NewAPI(),
		seen: make(map[types.Type]bool),
	}
}

func (r *reachability) FromPackages(tests bool, pkgs ...string) error {
	conf := &packages.Config{
		Mode:  packages.LoadTypes,
		Tests: tests,
	}

	loadedPackages, err := packages.Load(conf, pkgs...)
	if err != nil {
		return err
	}

	for _, pkg := range loadedPackages {
		r.API.Packages = append(r.API.Packages, pkg.Types)
		if err := r.fromPackage(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (r *reachability) fromPackage(pkg *packages.Package) error {
	if len(pkg.Errors) > 0 {
		return fmt.Errorf("has errors: %s", pkg.Errors[0])
	}

	scope := pkg.Types.Scope()
	for _, name := range scope.Names() {
		if !ast.IsExported(name) {
			continue
		}

		obj := scope.Lookup(name)
		if err := r.reachFromObject(obj); err != nil {
			return err
		}
	}

	return nil
}

func (r *reachability) reachFromObject(obj types.Object) error {
	if obj.Pkg() == nil {
		return nil
	}

	if obj.Parent() != nil {
		if r.API.Reachable[obj] {
			return nil
		}
		r.API.Reachable[obj] = true
	}

	return r.reachFromType(obj.Type())
}

func (r *reachability) reachFromType(typ types.Type) error {
	if r.seen[typ] {
		return nil
	}
	r.seen[typ] = true

	switch typ := typ.(type) {
	case *types.Named:
		if err := r.reachFromObject(typ.Obj()); err != nil {
			return err
		}

		for i := 0; i < typ.NumMethods(); i++ {
			m := typ.Method(i)
			if err := r.reachFromObject(m); err != nil {
				return err
			}
		}

		return nil
	case *types.Pointer:
		return r.reachFromType(typ.Elem())
	case *types.Struct:
		for i := 0; i < typ.NumFields(); i++ {
			f := typ.Field(i)
			if err := r.reachFromObject(f); err != nil {
				return err
			}
		}

		return nil
	case *types.Signature:
		if err := r.reachFromType(typ.Params()); err != nil {
			return err
		}

		return r.reachFromType(typ.Results())
	case *types.Tuple:
		for i := 0; i < typ.Len(); i++ {
			f := typ.At(i)
			if err := r.reachFromObject(f); err != nil {
				return err
			}
		}

		return nil
	case *types.Slice:
		return r.reachFromType(typ.Elem())
	case *types.Map:
		if err := r.reachFromType(typ.Key()); err != nil {
			return err
		}

		return r.reachFromType(typ.Elem())
	case *types.Basic:
		return nil
	case *types.Array:
		return r.reachFromType(typ.Elem())
	case *types.Chan:
		return r.reachFromType(typ.Elem())
	case *types.Interface:
		for i := 0; i < typ.NumMethods(); i++ {
			f := typ.Method(i)
			if err := r.reachFromObject(f); err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("unhandled type: %s", reflect.TypeOf(typ))
	}
}
