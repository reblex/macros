package bmr

import (
	"testing"
)

// Test case struct for funcitons calculating
// BMR with specific gender and standard.
type calculateSpecific struct {
	name string // Test case name

	// Input
	weight float64 // Input weight
	height float64 // Input height
	age    int     // Input age

	// Output
	calories int   // Resulting calories
	err      error // Resulting error
}

type calculateVariables struct {
	name string // Test case name

	// Input
	gender   string
	standard string
	weight   float64 // Input weight
	height   float64 // Input height
	age      int     // Input age

	// Output
	calories int   // Resulting calories
	err      error // Resulting error
}

/*
   Validate the results of a BMR calculation, comparing expected results with actual results.
*/
func validateCalculate(t *testing.T, expCals int, expErr error, resCals int, resErr error) string {
	// Validate err
	if resErr == nil {
		if expErr != nil {
			return "Unexpected error"
		}
	} else {
		switch resErr.(type) {
		case NegativeValueError:
			if _, ok := (expErr).(NegativeValueError); !ok {
				return "Wrong error type returned"
			}
		case ZeroValueError:
			if _, ok := (expErr).(ZeroValueError); !ok {
				return "Wrong error type returned"
			}
		case ValueTypeError:
			if _, ok := (expErr).(ValueTypeError); !ok {
				return "Wrong error type returned"
			}
		}
	}

	// Validate calories
	if resCals != expCals {
		return "Wrong amount of calories"
	}
	return ""
}

/*
	Calculate BMR.
*/
var calculateTests = []calculateVariables{
	{"Functional Input", "male", "metric", 80.0, 190.0, 24, 1949, nil},
	{"Other Functional Input", "female", "imperial", 176.0, 75.0, 24, 1660, nil},
	{"Negative value error", "male", "metric", -1.0, 174.3, 34, 0, NegativeValueError{}},
	{"Zero value error", "male", "metric", 64.3, 0.0, 12, 0, ZeroValueError{}},
	{"Non compatible gender", "dragon", "metric", 153.1, 68.4, 30, 0, ValueTypeError{}},
	{"Non compatible measurment standard", "female", "smoots", 153.1, 68.4, 30, 0, ValueTypeError{}},
}

func TestCalculate(t *testing.T) {
	for _, tc := range calculateTests {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := Calculate(tc.gender, tc.standard, tc.weight, tc.height, tc.age)
			calcErr := validateCalculate(t, tc.calories, tc.err, calories, err)
			if calcErr != "" {
				t.Error(calcErr)
			}
		})
	}
}

/*
   Calculate Male BMR using metric measurment standard.
*/
var calculateMaleMetricTests = []calculateSpecific{
	{"Functional Input", 80.0, 190.0, 24, 1949, nil},
	{"Other Functional Input", 120.3, 185.1, 29, 2442, nil},
	{"Negative value error", -1.0, 174.3, 34, 0, NegativeValueError{}},
	{"Zero value error", 64.3, 0.0, 12, 0, ZeroValueError{}},
}

func TestCalculateMaleMetric(t *testing.T) {
	for _, tc := range calculateMaleMetricTests {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := calculateMaleMetric(tc.weight, tc.height, tc.age)
			calcErr := validateCalculate(t, tc.calories, tc.err, calories, err)
			if calcErr != "" {
				t.Error(calcErr)
			}
		})
	}
}

/*
   Calculate Male BMR using imperial measurment standard.
*/
var calculateMaleImperialTests = []calculateSpecific{
	{"Functional Input", 176.0, 75.0, 24, 1952, nil},
	{"Other Functional Input", 153.1, 68.4, 30, 1684, nil},
	{"Negative value error", -1.0, 78.1, 34, 0, NegativeValueError{}},
	{"Zero value error", 48.2, 0.0, 12, 0, ZeroValueError{}},
}

func TestCalculateMaleImperial(t *testing.T) {
	for _, tc := range calculateMaleImperialTests {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := calculateMaleImperial(tc.weight, tc.height, tc.age)
			calcErr := validateCalculate(t, tc.calories, tc.err, calories, err)
			if calcErr != "" {
				t.Error(calcErr)
			}
		})
	}
}

/*
   Calculate Female BMR using metric measurment standard.
*/
var calculateFemaleMetricTests = []calculateSpecific{
	{"Functional Input", 65.0, 170.0, 22, 1482, nil},
	{"Other Functional Input", 72.5, 174.1, 32, 1514, nil},
	{"Negative value error", -1.0, 78.1, 34, 0, NegativeValueError{}},
	{"Zero value error", 48.2, 0.0, 12, 0, ZeroValueError{}},
}

func TestCalculateFemaleMetric(t *testing.T) {
	for _, tc := range calculateFemaleMetricTests {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := calculateFemaleMetric(tc.weight, tc.height, tc.age)
			calcErr := validateCalculate(t, tc.calories, tc.err, calories, err)
			if calcErr != "" {
				t.Error(calcErr)
			}
		})
	}
}

/*
   Calculate Female BMR using imperial measurment standard.
*/
var calculateFemaleImperialTests = []calculateSpecific{
	{"Functional Input", 176.0, 75.0, 24, 1660, nil},
	{"Other Functional Input", 153.1, 68.4, 30, 1501, nil},
	{"Negative value error", -1.0, 78.1, 34, 0, NegativeValueError{}},
	{"Zero value error", 48.2, 0.0, 12, 0, ZeroValueError{}},
}

func TestCalculateFemaleImperial(t *testing.T) {
	for _, tc := range calculateFemaleImperialTests {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := calculateFemaleImperial(tc.weight, tc.height, tc.age)
			calcErr := validateCalculate(t, tc.calories, tc.err, calories, err)
			if calcErr != "" {
				t.Error(calcErr)
			}
		})
	}
}
