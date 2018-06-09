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
	list                List all profiles
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

func init() {
	CmdProfile.Run = run
	CmdProfile.Flag.BoolVar(&profH, "h", false, "")
	CmdProfile.Flag.BoolVar(&profH, "help", false, "")
}

func run(cmd *base.Command, args []string) {
	if profH || len(args) < 1 {
		fmt.Println(CmdProfile.Help)
		return
	}
	var p profile.Profile
	p.Name = base.Settings.MainProfile
	p.Load("config/profiles.json")

	switch args[0] {
	case "list":
		listProfiles()
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

func getProfileNames() []string {
	var p profile.Profile
	profiles := p.GetProfiles("config/profiles.json")

	names := make([]string, len(profiles))

	for i, profile := range profiles {
		names[i] = profile.Name
	}

	return names
}

func isProfile(p string) bool {
	isProfile := false
	profiles := getProfileNames()

	for _, profile := range profiles {
		if profile == p {
			isProfile = true
			break
		}
	}
	return isProfile
}

func listProfiles() {
	profiles := getProfileNames()

	for i := 0; i < len(profiles); i++ {
		fmt.Println(profiles[i])
	}
}

func selectProfile(args []string) {
	if len(args) < 1 {
		fmt.Println(CmdProfile.Help)
		return
	}

	if base.Settings.MainProfile == args[0] {
		fmt.Printf("The profile '%v' is already selected.\n", args[0])
	} else if isProfile(args[0]) == false {
		fmt.Printf("A profile with the name '%v' does not exist.\n", args[0])
		fmt.Println("To see a list of profile, use:")
		fmt.Println("  macros profile list")
	} else {
		base.Settings.MainProfile = args[0]
		base.Settings.Save("config/settings.json")
		fmt.Printf("Selecting profile '%v'.\n", args[0])
	}
}

func createProfile(args []string) {
	return
}

func editProfile(args []string) {
	return
}
