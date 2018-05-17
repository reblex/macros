package calc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/reblex/macros/cmd/macros/internal/cli/base"
	"github.com/reblex/macros/cmd/macros/internal/profile"
	"github.com/reblex/macros/pkg/bmr"
)

var CmdCalc = &base.Command{
	Name:        "calc",
	Usage:       "calc [flag]... calculation <weight>",
	Description: "Calculate Macros.",
	Help: `Overview:
    Calculates meal and energy related values based on profile settings.

Usage:
    calc [flag]... calculation <weight>

Flags:
    -h, -help           Print this help message.

Calculations:
    bmr                 Base metabolic rate
    calories            Total calories((bmr * activity factor) + calorie constant)
    macros              Macronutrients
    all                 Complete profile
    help                Print this help message.
    `,
	CustomFlags: true,
}

var (
	calcH bool // calc -h, --help flag
)

func init() {
	CmdCalc.Run = run
	CmdCalc.Flag.BoolVar(&calcH, "h", false, "")
	CmdCalc.Flag.BoolVar(&calcH, "help", false, "")
}

func calcBmr(p profile.Profile, weight float64) (int, error) {
	bmr, err := bmr.Calculate(p.Data.Gender, p.Data.Standard, weight, p.Data.Height, p.Data.Age)

	if err != nil {
		return 0, err
	}

	return bmr, nil
}

func printBmr(p profile.Profile, args []string) {
	if len(args) < 2 {
		fmt.Println("BMR needs weight argument.\n macros calc bmr <weight>")
		os.Exit(1)
	}
	weight, _ := strconv.ParseFloat(args[1], 64)

	if bmr, err := calcBmr(p, weight); err == nil {
		fmt.Println("BMR: ", bmr)
	} else {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func printCalories(p profile.Profile, args []string) {
	if len(args) < 2 {
		fmt.Println("Calories need a weight argument.\n macros calc calories <weight>")
		os.Exit(1)
	}
	weight, _ := strconv.ParseFloat(args[1], 64)

	bmr, err := calcBmr(p, weight)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	calories := int(float64(bmr)*p.Data.ActivityFactor) + p.Data.CalorieConstant
	fmt.Println("Calories:", calories)
}

func run(cmd *base.Command, args []string) {
	if calcH || len(args) < 1 {
		fmt.Println(CmdCalc.Help)
		return
	}
	var p profile.Profile
	p.Name = base.Settings.MainProfile
	p.Load("config/profiles.json")

	switch args[0] {
	case "bmr":
		printBmr(p, args)
	case "calories":
		printCalories(p, args)
	case "help":
		fmt.Println(cmd.Help)
	default:
		fmt.Println("Invalid calculation type.\n" + cmd.Usage + "\nFor more help type:\n macros calc help")
	}
}
