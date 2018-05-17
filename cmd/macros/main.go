package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/reblex/macros/cmd/macros/internal/cli/base"
	"github.com/reblex/macros/cmd/macros/internal/cli/commands/calc"
)

var (
	BaseFlags flag.FlagSet
)

func init() {
	BaseFlags := base.Flags
	BaseFlags.Parse(os.Args[1:])
	base.Commands = []*base.Command{
		calc.CmdCalc,
	}
	base.Settings.Load("config/settings.json")
}

func main() {
	BaseFlags.Usage = base.Usage
	// flag.Parse()
	BaseFlags.Parse(os.Args[1:])

	args := BaseFlags.Args()
	if len(args) < 1 {
		// base.Usage()
		fmt.Println("TODO: Main usage")
	}

	if args[0] == "help" || args[0] == "-h" || args[0] == "-help" {
		fmt.Println("TODO: Main help")
		return
	}

	if base.FlagH {
		fmt.Println("TODO: Base help")
		return
	}

	for _, cmd := range base.Commands {
		if cmd.Name == args[0] {
			cmd.Flag.Usage = func() { fmt.Println("usage:", cmd.Usage) }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			os.Exit(0)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown subcommand %q\nRun 'macros help' for usage.\n", args[0])
	os.Exit(1)
}
