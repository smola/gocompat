package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"sort"

	"golang.org/x/tools/go/loader"
	"gopkg.in/src-d/go-cli.v0"
	"gopkg.in/src-d/go-log.v1"
)

func init() {
	app.AddCommand(&reachCommand{})
}

type reachCommand struct {
	cli.Command `name:"reach" short-desc:"List all symbols reachable from a package."`
	Positional  struct {
		Packages []string `positional-arg-name:"package" description:"Package to start from."`
	} `positional-args:"yes" required:"yes"`
}

func (c reachCommand) Execute(args []string) error {
	var conf loader.Config
	_, err := conf.FromArgs(c.Positional.Packages, false)
	if err != nil {
		log.Errorf(err, "creating load configuration")
		return err
	}

	prog, err := conf.Load()
	if err != nil {
		log.Errorf(err, "loading packages")
		return err
	}

	seen := make(map[string]bool)
	reachable := make(map[types.Object]bool)
	for _, pkg := range prog.InitialPackages() {
		if err := reachFromPackage(seen, reachable, pkg); err != nil {
			log.Errorf(err, "processing package: %s", pkg.Pkg.Name)
			return err
		}
	}

	var reachedString []string
	for symbol := range reachable {
		str := fmt.Sprintf("\"%s\".%s", symbol.Pkg().Path(), symbol.Name())
		reachedString = append(reachedString, str)
	}

	sort.Strings(reachedString)

	for _, str := range reachedString {
		fmt.Println(str)
	}

	return nil
}
func reachFromPackage(seen map[string]bool, reachable map[types.Object]bool, pkg *loader.PackageInfo) error {
	if !pkg.TransitivelyErrorFree {
		return fmt.Errorf("has errors")
	}

	scope := pkg.Pkg.Scope()
	for _, name := range scope.Names() {
		if !ast.IsExported(name) {
			continue
		}

		obj := scope.Lookup(name)
		reachFromObject(seen, reachable, pkg.Pkg, obj)
	}

	return nil
}

func reachFromObject(seen map[string]bool, reachable map[types.Object]bool, pkg *types.Package, obj types.Object) {
	if reachable[obj] {
		return
	}
	reachable[obj] = true

	if !obj.Exported() {
		return
	}

	reachFromType(seen, reachable, pkg, obj.Type())
}

func reachFromType(seen map[string]bool, reachable map[types.Object]bool, pkg *types.Package, typ types.Type) {
	if seen[typ.String()] {
		return
	}
	seen[typ.String()] = true

	switch typ := typ.(type) {
	case *types.Named:
		obj := typ.Obj()
		opkg := obj.Pkg()
		if opkg == nil {
			return
		}

		if opkg != pkg {
			reachFromObject(seen, reachable, opkg, obj)
			return
		}

		if !typ.Obj().Exported() {
			return
		}

		for i := 0; i < typ.NumMethods(); i++ {
			m := typ.Method(i)
			reachFromType(seen, reachable, pkg, m.Type())
		}

	case *types.Pointer:
		reachFromType(seen, reachable, pkg, typ.Elem())
	case *types.Struct:
		for i := 0; i < typ.NumFields(); i++ {
			f := typ.Field(i)
			if !f.Exported() {
				continue
			}
			reachFromType(seen, reachable, pkg, f.Type())
		}
	case *types.Signature:
		reachFromType(seen, reachable, pkg, typ.Params())
		reachFromType(seen, reachable, pkg, typ.Results())
	case *types.Tuple:
		for i := 0; i < typ.Len(); i++ {
			f := typ.At(i)
			reachFromType(seen, reachable, pkg, f.Type())
		}
	case *types.Slice:
		reachFromType(seen, reachable, pkg, typ.Elem())
	case *types.Map:
		reachFromType(seen, reachable, pkg, typ.Key())
		reachFromType(seen, reachable, pkg, typ.Elem())
	case *types.Basic:
	default:
		log.Warningf("unhandled type: %f", reflect.TypeOf(typ))
	}
}
