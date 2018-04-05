package base

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	// Runs the command with remaining args.
	Run func(cmd *Command, args []string)

	// Name of command
	Name string

	// One line usage description
	Usage string

	// Short description of command.
	Description string

	// Help text.
	Help string

	// Specific set of flags for this command.
	Flag flag.FlagSet

	// True indicates that the command parses
	// it's own flags.
	CustomFlags bool
}

var (
	FlagH bool // macros -h, -help flag

	Flags flag.FlagSet
	Usage func()

	// All available commands
	Commands []*Command
)

func init() {
	Flags := flag.NewFlagSet("", flag.ExitOnError)
	Flags.BoolVar(&FlagH, "h", false, "")
	Flags.BoolVar(&FlagH, "help", false, "")
}

// Print command usage
func (c *Command) PrintUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n", c.Usage)
	fmt.Fprintf(os.Stderr, "Run 'macros help %s' for details.\n", c.Name)
}

// Run command
func Run(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	err := cmd.Run()

	if err != nil {
		fmt.Errorf("%v", err)
	}
}
