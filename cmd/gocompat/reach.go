package main

import (
	"fmt"
	"sort"

	compat "github.com/smola/gocompat"
	"gopkg.in/src-d/go-cli.v0"
)

func init() {
	app.AddCommand(&reachCommand{})
}

type reachCommand struct {
	cli.Command `name:"reach" short-desc:"List all types reachable from a package."`
	Positional  struct {
		Packages []string `positional-arg-name:"package" description:"Package to start from."`
	} `positional-args:"yes" required:"yes"`
}

func (c reachCommand) Execute(args []string) error {
	api, err := compat.ReachableFromPackages(c.Positional.Packages...)
	if err != nil {
		return err
	}

	var reachedString []string
	for r := range api.Reachable {
		str := fmt.Sprintf("\"%s\".%s", r.Pkg().Path(), r.Name())
		reachedString = append(reachedString, str)
	}

	sort.Strings(reachedString)

	for _, str := range reachedString {
		fmt.Println(str)
	}

	return nil
}
