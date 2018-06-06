package profile

import (
	"fmt"

	"github.com/reblex/macros/cmd/macros/internal/cli/base"
	"github.com/reblex/macros/cmd/macros/internal/profile"
)

var CmdProfile = &base.Command{
	Name:        "profile",
	Usage:       "macros profile [flag]... <command> [arg]",
	Description: "Handle profile settings.",
	Help: `Overview:
    Handles profiles and profile settings.

Usage:
    macros profile [flag]... <command> [arg]

Flags:
    -h, -help           Print this help message

Commands:
    select              Select a profile to use
    create            	Create a new profile
    edit                Edit a current profile settings
	help                Print this help message
    `,
	CustomFlags: true,
}

var (
	profH bool // profile -h, -help flag
)

func run(cmd *base.Command, args []string) {
	if profH || len(args) < 1 {
		fmt.Println(CmdProfile.Help)
		return
	}
	var p profile.Profile
	p.Name = base.Settings.MainProfile
	p.Load("config/profiles.json")

	switch args[0] {
	case "select":
		selectProfile(args[1:])
	case "create":
		createProfile(args[1:])
	case "edit":
		editProfile(args[1:])
	case "help":
		fmt.Println(cmd.Help)
	default:
		fmt.Println("Invalid command.\n" + cmd.Usage + "\nFor more help type:\n macros profile help")
	}
}

func selectProfile(args []string) {
	return
}

func createProfile(args []string) {
	return
}

func editProfile(args []string) {
	return
}
