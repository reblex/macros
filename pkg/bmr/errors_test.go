package bmr

import (
	"testing"
)

func TestNegativeValueError(t *testing.T) {
	err := NegativeValueError{"Negative Value"}
	if s := err.Error(); s != "Negative Value" {
		t.Error("Unfunctional Error Type")
	}
}

func TestZeroValueError(t *testing.T) {
	err := ZeroValueError{"Zero Value"}
	if s := err.Error(); s != "Zero Value" {
		t.Error("Unfunctional Error Type")
	}
}

func TestValueTypeError(t *testing.T) {
	err := ValueTypeError{"Value Type Error"}
	if s := err.Error(); s != "Value Type Error" {
		t.Error("Unfunctional Error Type")
	}
}
