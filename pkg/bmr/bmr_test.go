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
