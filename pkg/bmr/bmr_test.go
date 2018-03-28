package bmr

import (
	"testing"
)

/*
   Validate the results of a BMR calculation, comparing expected results with actual results.
*/
func validateCalculate(t *testing.T, expCals int, expErr error, resCals int, resErr error) {
	// Validate calories
	if resCals != expCals {
		t.Error("Wrong amount of calories")
	}

	// Validate err
	if resErr == nil {
		if expErr != nil {
			t.Error("Unexpected error")
		}
	} else {
		switch resErr.(type) {
		case NegativeValueError:
			if _, ok := (expErr).(NegativeValueError); !ok {
				t.Error("Wrong error type returned")
			}
		case ZeroValueError:
			if _, ok := (expErr).(ZeroValueError); !ok {
				t.Error("Wrong error type returned")
			}
		}
	}
}

/*
   Calculate Male BMR using metric measurment standard.
*/
var calculateMaleMetricTests = []struct {
	// Input
	weight float64 // Input weight
	height float64 // Input height
	age    int     // Input age

	// Output
	calories int   // Resulting calories
	err      error // Resulting error
}{
	{80.0, 190.0, 24, 1949, nil},               // Functional input
	{120.3, 185.1, 29, 2442, nil},              // Other functional input
	{-1.0, 174.3, 34, 0, NegativeValueError{}}, // Negative value error
	{64.3, 0.0, 12, 0, ZeroValueError{}},       // Zero value error
}

func TestCalculateMaleMetric(t *testing.T) {
	for _, tc := range calculateMaleMetricTests {
		calories, err := calculateMaleMetic(tc.weight, tc.height, tc.age)
		validateCalculate(t, tc.calories, tc.err, calories, err)
	}
}

/*
   Calculate Male BMR using imperial measurment standard.
*/
var calculateMaleImperialTests = []struct {
	// Input
	weight float64 // Input weight
	height float64 // Input height
	age    int     // Input age

	// Output
	calories int   // Resulting calories
	err      error // Resulting error
}{
	{176.0, 75.0, 24, 1952, nil},              // Functional input
	{153.1, 68.4, 30, 1684, nil},              // Other functional input
	{-1.0, 78.1, 34, 0, NegativeValueError{}}, // Negative value error
	{48.2, 0.0, 12, 0, ZeroValueError{}},      // Zero value error
}

func TestCalculateMaleImperial(t *testing.T) {
	for _, tc := range calculateMaleMetricTests {
		calories, err := calculateMaleImperial(tc.weight, tc.height, tc.age)
		validateCalculate(t, tc.calories, tc.err, calories, err)
	}
}

/*
   Calculate Female BMR using metric measurment standard.
*/
var calculateFemaleMetricTests = []struct {
	// Input
	weight float64 // Input weight
	height float64 // Input height
	age    int     // Input age

	// Output
	calories int   // Resulting calories
	err      error // Resulting error
}{
	{65.0, 170.0, 22, 1482, nil},              // Functional input
	{72.5, 174.1, 32, 1514, nil},              // Other functional input
	{-1.0, 78.1, 34, 0, NegativeValueError{}}, // Negative value error
	{48.2, 0.0, 12, 0, ZeroValueError{}},      // Zero value error
}

func TestCalculateFemaleMetric(t *testing.T) {
	for _, tc := range calculateMaleMetricTests {
		calories, err := calculateFemaleMetric(tc.weight, tc.height, tc.age)
		validateCalculate(t, tc.calories, tc.err, calories, err)
	}
}
