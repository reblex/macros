package bmr

import "math"

func Calculate(gender, standard string, weight, height float64, age int) (int, error) {
	switch gender {
	case "male":
		switch standard {
		case "metric":
			return calculateMaleMetric(weight, height, age)
		case "imperial":
			return calculateMaleImperial(weight, height, age)
		default:
			return 0, ValueTypeError{"Uncompatible measurement standard"}
		}
	case "female":
		switch standard {
		case "metric":
			return calculateFemaleMetric(weight, height, age)
		case "imperial":
			return calculateFemaleImperial(weight, height, age)
		default:
			return 0, ValueTypeError{"Uncompatible measurement standard"}
		}
	default:
		return 0, ValueTypeError{"Uncompatible gender"}
	}
}

func calculateMaleMetric(weight, height float64, age int) (int, error) {
	if weight < 0 || height < 0 || age < 0 {
		return 0, NegativeValueError{"Negative value not allowed."}
	}
	if weight == 0 || height == 0 || age == 0 {
		return 0, ZeroValueError{"Values of zero not allowed."}
	}

	return int(math.Round(66 + (13.7 * weight) + (5 * height) - (6.8 * float64(age)))), nil
}

func calculateFemaleMetric(weight, height float64, age int) (int, error) {
	if weight < 0 || height < 0 || age < 0 {
		return 0, NegativeValueError{"Negative value not allowed."}
	}
	if weight == 0 || height == 0 || age == 0 {
		return 0, ZeroValueError{"Values of zero not allowed."}
	}

	return int(math.Round(655 + (9.6 * weight) + (1.8 * height) - (4.7 * float64(age)))), nil
}

func calculateMaleImperial(weight, height float64, age int) (int, error) {
	if weight < 0 || height < 0 || age < 0 {
		return 0, NegativeValueError{"Negative value not allowed."}
	}
	if weight == 0 || height == 0 || age == 0 {
		return 0, ZeroValueError{"Values of zero not allowed."}
	}

	return int(math.Round(66 + (6.23 * weight) + (12.7 * height) - (6.8 * float64(age)))), nil
}

func calculateFemaleImperial(weight, height float64, age int) (int, error) {
	if weight < 0 || height < 0 || age < 0 {
		return 0, NegativeValueError{"Negative value not allowed."}
	}
	if weight == 0 || height == 0 || age == 0 {
		return 0, ZeroValueError{"Values of zero not allowed."}
	}

	return int(math.Round(655 + (4.35 * weight) + (4.7 * height) - (4.7 * float64(age)))), nil
}
