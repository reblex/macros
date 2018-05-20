package calc

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/reblex/macros/cmd/macros/internal/cli/base"
	"github.com/reblex/macros/cmd/macros/internal/profile"
	"github.com/reblex/macros/pkg/bmr"
)

var CmdCalc = &base.Command{
	Name:        "calc",
	Usage:       "calc [flag]... <calculation> <weight>",
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
    help                Print this help message.
    `,
	CustomFlags: true,
}

var (
	calcH bool // calc -h, --help flag
)

type macros struct {
	Calories int
	Protein  int
	Fat      int
	Carbs    int
}

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

func calcCalories(p profile.Profile, weight float64) (int, error) {
	bmr, err := calcBmr(p, weight)

	if err != nil {
		return 0, err
	}

	calories := int(math.Round(float64(bmr)*p.Data.ActivityFactor)) + p.Data.CalorieConstant

	return calories, nil
}

func printCalories(p profile.Profile, args []string) {
	if len(args) < 2 {
		fmt.Println("Calories needs a weight argument.\n macros calc calories <weight>")
		os.Exit(1)
	}
	weight, _ := strconv.ParseFloat(args[1], 64)

	calories, _ := calcCalories(p, weight)

	fmt.Println("Calories:", calories)
}

func calcMacros(p profile.Profile, weight float64) macros {

	calories, _ := calcCalories(p, weight)

	if p.Data.Standard != "imperial" {
		weight *= 2.2046
	}

	protein := int(weight * p.Data.ProteinFactor)

	fat := int(float64(protein) * (float64(p.Data.FatPercentage) / 100))

	proteinCalories := int(protein * 4)
	fatCalories := int(fat * 9)
	carbCalories := int(calories - proteinCalories - fatCalories)

	carbs := int(carbCalories / 4)

	return macros{calories, protein, fat, carbs}
}

func printMacros(p profile.Profile, args []string) {
	if len(args) < 2 {
		fmt.Println("Macros need a weight argument.\n macros calc calories <weight>")
		os.Exit(1)
	}

	weight, _ := strconv.ParseFloat(args[1], 64)
	macros := calcMacros(p, weight)

	// TODO: Make this less messy.
	text := "Daily Macros\n"
	text += "    Calories: " + strconv.Itoa(macros.Calories) + " kcal\n"
	text += "    Protein: " + strconv.Itoa(macros.Protein) + "g\n"
	text += "    Fat: " + strconv.Itoa(macros.Fat) + "g\n"
	text += "    Carbs: " + strconv.Itoa(macros.Carbs) + "g\n"
	text += "Macros/Meal (" + strconv.Itoa(p.Data.MealsPerDay) + " meals/day)\n"
	text += "    Calories: " + strconv.Itoa(int(float64(macros.Calories)/float64(p.Data.MealsPerDay))) + " kcal\n"
	text += "    Protein: " + strconv.Itoa(int(float64(macros.Protein)/float64(p.Data.MealsPerDay))) + "g\n"
	text += "    Fat: " + strconv.Itoa(int(float64(macros.Fat)/float64(p.Data.MealsPerDay))) + "g\n"
	text += "    Carbs: " + strconv.Itoa(int(float64(macros.Carbs)/float64(p.Data.MealsPerDay))) + "g\n"

	fmt.Println(text)
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
	case "macros":
		printMacros(p, args)
	case "help":
		fmt.Println(cmd.Help)
	default:
		fmt.Println("Invalid calculation type.\n" + cmd.Usage + "\nFor more help type:\n macros calc help")
	}
}
