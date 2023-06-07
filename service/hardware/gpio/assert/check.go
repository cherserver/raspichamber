package assert

import "fmt"

func IsPercentFrom0To100(value uint8) error {
	if value < 0 || value > 100 {
		return fmt.Errorf("value '%v' is not a valid percent value. Must be 0-100", value)
	}

	return nil
}

func IsAngleFrom0To90(value uint8) error {
	if value < 0 || value > 90 {
		return fmt.Errorf("value '%v' is not a valid angle value. Must be 0-90", value)
	}

	return nil
}
