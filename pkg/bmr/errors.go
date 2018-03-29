package bmr

import "fmt"

// Value is below zero
type NegativeValueError struct {
	message string
}

func (e NegativeValueError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// Value is zero
type ZeroValueError struct {
	message string
}

func (e ZeroValueError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// Unusable measurment type.
type ValueTypeError struct {
	message string
}

func (e ValueTypeError) Error() string {
	return fmt.Sprintf("%s", e.message)
}
