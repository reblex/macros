package calc

import (
	"errors"
	"strconv"
	"testing"

	"github.com/reblex/macros/cmd/macros/internal/profile"
)

type calcCaloriesAndBmrTestValues struct {
	name string // Name of Test Case

	// Input
	profile profile.Profile // Profile for test
	weight  float64         //current weight for calculation

	// Output
	res        int  // Resulting calories/BMR
	shouldWork bool // Should the test case not result in an error?
}

// Test profiles
var (
	testProfile1 = profile.Profile{"Male Bulk", profile.ProfileData{"Male Bulk", "metric", 21, 190, "male", 200, 1.66, 66, 1.2, 5}}
	testProfile2 = profile.Profile{"Female Cut", profile.ProfileData{"Female Cut", "metric", 20, 172, "female", -250, 1.5, 40, 1, 6}}
)

/*
   calcBmr tests
*/
var calcBmrTestCases = []calcCaloriesAndBmrTestValues{
	{"Profile1_84kg", testProfile1, 84, 2024, true},
	{"Profile1_70kg", testProfile1, 70, 1832, true},
	{"Profile2_72kg", testProfile2, 72, 1562, true},
	{"Profile2_64kg", testProfile2, 64, 1485, true},
	{"Profile2_-2kg_err", testProfile2, -2, 0, false},
}

func TestCalcBmr(t *testing.T) {
	for _, tc := range calcBmrTestCases {
		t.Run(tc.name, func(t *testing.T) {
			bmr, err := calcBmr(tc.profile, tc.weight)

			// Returned error
			if tc.shouldWork && err != nil {
				t.Error(err)
			}

			// Incorrect result
			if bmr != tc.res && tc.shouldWork {
				errTxt := "Incorrect BMR result. exp:" + strconv.Itoa(tc.res) + ", res:" + strconv.Itoa(bmr)
				t.Error(errors.New(errTxt))
			}
		})
	}
}

/*
   calcCalories tests
*/
var calcCaloriesTestCases = []calcCaloriesAndBmrTestValues{
	{"Profile1_84kg", testProfile1, 84, 3560, true},
	{"Profile1_70kg", testProfile1, 70, 3241, true},
	{"Profile2_72kg", testProfile2, 72, 2093, true},
	{"Profile2_64kg", testProfile2, 64, 1978, true},
	{"Profile2_-2kg_err", testProfile2, -2, 0, false},
}

func TestCalcCalories(t *testing.T) {
	for _, tc := range calcCaloriesTestCases {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := calcCalories(tc.profile, tc.weight)

			// Returned error
			if tc.shouldWork && err != nil {
				t.Error(err)
			}

			// Incorrect result
			if calories != tc.res && tc.shouldWork {
				errTxt := "Incorrect Calorie result. exp:" + strconv.Itoa(tc.res) + ", res:" + strconv.Itoa(calories)
				t.Error(errors.New(errTxt))
			}
		})
	}
}
