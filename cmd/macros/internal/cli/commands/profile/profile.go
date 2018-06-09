package profile

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

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
	case "examine":
		examineProfile(args[1:])
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

func examineProfile(args []string) {
	if len(args) < 1 {
		fmt.Println("examine requires profile argument.")
		fmt.Println(" macros profile examine <profile>")
		return
	}

	if isProfile(args[0]) == false {
		fmt.Printf("A profile with the name '%v' does not exist.\n", args[0])
		fmt.Println("To see a list of profile, use:")
		fmt.Println("  macros profile list")
		return
	}

	p := profile.Profile{}
	p.Name = args[0]
	p.Load("config/profiles.json")

	//https://play.golang.org/p/0Uqf8Qi0lC
	value := reflect.ValueOf(p.Data)

	for i := 0; i < value.NumField(); i++ {
		key := value.Type().Field(i).Name
		val := value.Field(i)
		fmt.Printf("%v: %v\n", key, val)
	}
}

func createProfile(args []string) {
	p := profile.Profile{}
	pd := &profile.ProfileData{}
	reader := bufio.NewReader(os.Stdin)

	//https://play.golang.org/p/0Uqf8Qi0lC
	ptr := reflect.ValueOf(pd)
	value := ptr.Elem()

	// Loop profiledata struct and set values for all fields
	for i := 0; i < value.NumField(); i++ {
		key := value.Type().Field(i).Name
		val := value.Field(i)
		// fmt.Printf("%v: %v\n", key, val)

		fmt.Printf("%v: ", key)
		newVal, _ := reader.ReadString('\n')
		newVal = strings.TrimSuffix(newVal, "\n")

		// Set struct field values based on type
		switch val.Type().Name() {
		case "string":
			val.SetString(newVal)
		case "float64":
			newVal, _ := strconv.ParseFloat(newVal, 64)
			val.SetFloat(newVal)
		case "int":
			newVal, _ := strconv.ParseInt(newVal, 10, 64)
			val.SetInt(newVal)
		case "bool":
			newVal, _ := strconv.ParseBool(newVal)
			val.SetBool(newVal)
		}

		if key == "Name" {
			p.Name = val.String()
		}
	}

	p.Data = *pd
	p.Save("config/profiles.json")
}

func editProfile(args []string) {
	if len(args) < 1 {
		fmt.Println("edit requires profile argument.")
		fmt.Println(" macros profile edit <profile>")
		return
	}

	if isProfile(args[0]) == false {
		fmt.Printf("A profile with the name '%v' does not exist.\n", args[0])
		fmt.Println("To see a list of profile, use:")
		fmt.Println("  macros profile list")
		return
	}

	p := profile.Profile{}
	p.Name = args[0]
	p.Load("config/profiles.json")

	pdNewData := p.Data
	pdNew := &pdNewData
	reader := bufio.NewReader(os.Stdin)

	//https://play.golang.org/p/0Uqf8Qi0lC
	ptr := reflect.ValueOf(pdNew)
	value := ptr.Elem()

	fmt.Println("Enter a new value and press enter to accept.")
	fmt.Println("Keep blank and press enter to keep current value.")

	// Loop profiledata struct and set values for all fields
	for i := 0; i < value.NumField(); i++ {
		key := value.Type().Field(i).Name
		val := value.Field(i)
		// fmt.Printf("%v: %v\n", key, val)

		fmt.Printf("%v: %v -> ", key, val)
		newVal, _ := reader.ReadString('\n')
		newVal = strings.TrimSuffix(newVal, "\n")

		if newVal != "" {
			// Set struct field values based on type
			switch val.Type().Name() {
			case "string":
				val.SetString(newVal)
			case "float64":
				newVal, _ := strconv.ParseFloat(newVal, 64)
				val.SetFloat(newVal)
			case "int":
				newVal, _ := strconv.ParseInt(newVal, 10, 64)
				val.SetInt(newVal)
			case "bool":
				newVal, _ := strconv.ParseBool(newVal)
				val.SetBool(newVal)
			}

			if key == "Name" {
				p.Name = val.String()
			}
		}
	}

	p.Data = *pdNew
	slot, _ := p.GetSlot("config/profiles.json")
	p.SaveToSlot("config/profiles.json", slot)
}
