package cmds

import (
	"github.com/jeckbjy/cli"
)

func Commands() []*cli.Command {
	return []*cli.Command{
		cli.NewCmd("clean", "", onClean),
		cli.NewCmd("xcode", "", onGen),
		cli.NewCmd("vs", "", onGen),
		cli.NewCmd("ninja", "", onGen),
	}
}
