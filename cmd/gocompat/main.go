package main

import (
	"gopkg.in/src-d/go-cli.v0"
)

var (
	version string
	build   string
)

var app = cli.New("gocompat", version, build, "Check API backwards compatibility.")

func main() {
	app.RunMain()
}
