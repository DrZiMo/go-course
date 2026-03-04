package convertions

import (
	"errors"
	"strconv"
)

func StringToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to convert strings to float")
		}

		floats[stringIndex] = floatPrice
	}

	return floats, nil
}
