package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	compat "github.com/smola/gocompat"
	"gopkg.in/src-d/go-cli.v0"
)

func init() {
	app.AddCommand(&compareCommand{})
}

type compareCommand struct {
	cli.Command    `name:"compare" short-desc:"List all symbols reachable from a package."`
	Exclude        []string `long:"exclude" description:"excluded change type"`
	ExcludePackage []string `long:"exclude-package" description:"excluded package"`
	ExcludeSymbol  []string `long:"exclude-symbol" description:"excluded symbol" unquote:"false"`
	Go1Compat      bool     `long:"go1compat" description:"Based on Go 1 promise of compatibility. Equivalent to --exclude=SymbolAdded --exclude=FieldAdded --exclude=MethodAdded"`
	Positional     struct {
		From     string   `positiona-arg-name:"from" description:"from git reference"`
		To       string   `positiona-arg-name:"to" description:"to git reference"`
		Packages []string `positional-arg-name:"package" description:"Package to start from."`
	} `positional-args:"yes" required:"yes"`
}

func (c compareCommand) Execute(args []string) error {
	excluded := make(map[compat.ChangeType]bool)
	for _, e := range c.Exclude {
		c, err := compat.ChangeTypeFromString(e)
		if err != nil {
			return err
		}

		excluded[c] = true
	}

	if c.Go1Compat {
		excluded[compat.SymbolAdded] = true
		excluded[compat.FieldAdded] = true
		excluded[compat.MethodAdded] = true
	}

	head, err := getHEAD()
	if err != nil {
		return err
	}

	defer func() {
		gitCheckout(head)
	}()

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

	changed := false
	changes := compat.Compare(fromReachable, toReachable)
	for _, change := range changes {
		if excluded[change.Type] {
			continue
		}

		exclude := false
		for _, pkg := range c.ExcludePackage {
			prefix := fmt.Sprintf(`"%s"`, pkg)
			if strings.HasPrefix(change.Symbol, prefix) {
				exclude = true
				break
			}
		}
		if exclude {
			continue
		}

		for _, sym := range c.ExcludeSymbol {
			if change.Symbol == sym {
				exclude = true
				break
			}
		}
		if exclude {
			continue
		}

		changed = true
		fmt.Println(change)
	}

	if changed {
		return fmt.Errorf("found backwards incompatible changes")
	}

	return nil
}

func getHEAD() (string, error) {
	headBytes, err := ioutil.ReadFile(".git/HEAD")
	if err != nil {
		return "", err
	}

	head := string(headBytes)
	head = strings.TrimSpace(head)
	if strings.HasPrefix(head, "ref: refs/heads/") {
		return head[len("ref: refs/heads/"):], nil
	}

	if strings.HasPrefix(head, "ref: ") {
		return head[len("ref: "):], nil
	}

	return head, nil
}

func gitCheckout(ref string) error {
	cmd := exec.Command("git", "checkout", ref)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
