package calc

import (
	"fmt"

	"github.com/reblex/macros/cmd/macros/internal/cli/base"
)

var CmdCalc = &base.Command{
	Name:        "calc",
	Usage:       "calc [flag]... calculation",
	Description: "Calculate Macros.",
	Help: `Overview:
    Calculates meal and energy related values based on profile settings.

Calculations:
    bmr                 Base metabolic rate
    calories            Total calories
    macros              Macronutrients
    all					Complete profile

Flags:
    -h, -help			Print this help message.
    `,
}

var (
	calcH bool // calc -h, --help flag
)

func init() {
	CmdCalc.Run = run

	CmdCalc.Flag.BoolVar(&calcH, "h", false, "")
	CmdCalc.Flag.BoolVar(&calcH, "help", false, "")
}

func run(cmd *base.Command, args []string) {
	if calcH {
		fmt.Println(CmdCalc.Help)
		return
	}
}
