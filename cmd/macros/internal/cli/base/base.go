package base

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

type applicationSettings struct {
	MainProfile string `json:"mainProfile"`
}

var (
	FlagH bool // macros -h, -help flag

	Flags flag.FlagSet
	Usage func()

	// All available commands
	Commands []*Command

	//System-wide settings
	Settings applicationSettings
)

func init() {
	Flags := flag.NewFlagSet("", flag.ExitOnError)
	Flags.BoolVar(&FlagH, "h", false, "")
	Flags.BoolVar(&FlagH, "help", false, "")
}

func getProjectPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/reblex/macros/"
}

// Print command usage
func (c *Command) PrintUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n", c.Usage)
	fmt.Fprintf(os.Stderr, "Run 'macros help %s' for details.\n", c.Name)
}

func (s *applicationSettings) Load(path string) error {
	raw, err := ioutil.ReadFile(getProjectPath() + path)

	if err != nil {
		return err
	}

	json.Unmarshal(raw, &Settings)

	return nil
}

func (s *applicationSettings) Save(path string) error {
	data, _ := json.MarshalIndent(Settings, "", "    ")

	ioutil.WriteFile(getProjectPath()+path, data, 0644)

	return nil
}

// Run command
func Run(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}
