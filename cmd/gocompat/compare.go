package main

import (
	"fmt"
	"os"
	"os/exec"

	compat "github.com/smola/gocompat"
	"gopkg.in/src-d/go-cli.v0"
)

func init() {
	app.AddCommand(&compareCommand{})
}

type compareCommand struct {
	cli.Command `name:"compare" short-desc:"List all symbols reachable from a package."`
	Positional  struct {
		From     string   `positiona-arg-name:"from" description:"from git reference"`
		To       string   `positiona-arg-name:"to" description:"to git reference"`
		Packages []string `positional-arg-name:"package" description:"Package to start from."`
	} `positional-args:"yes" required:"yes"`
}

func (c compareCommand) Execute(args []string) error {
	if err := gitCheckout(c.Positional.From); err != nil {
		return err
	}

	fromReachable, err := compat.ReachableFromPackages(c.Positional.Packages...)
	if err != nil {
		return err
	}
	if err := gitCheckout(c.Positional.To); err != nil {
		return err
	}

	toReachable, err := compat.ReachableFromPackages(c.Positional.Packages...)
	if err != nil {
		return err
	}

	changes := compat.Compare(fromReachable, toReachable)
	for _, change := range changes {
		fmt.Println(change)
	}
	return nil
}

func gitCheckout(ref string) error {
	cmd := exec.Command("git", "checkout", ref)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
