package main

import (
	"github.com/jeckbjy/cli"
	"github.com/jeckbjy/zmake/cmds"
)

// premake https://github.com/premake/premake-core
func main() {
	app := cli.New()

	app.AddHeader("zmake 1.0.0, a build script generator")

	app.AddFooter("For additional information, see https://github.com/jeckbjy/zmake")

	app.AddCommands(cmds.Commands())
	app.Run()
}
