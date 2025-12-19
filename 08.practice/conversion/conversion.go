package conversion

import (
	"errors"
	"strconv"
	"strings"
)

func StringsToFloat(lines []string) ([]float64, error) {
	gg := []float64{}

	for _, val := range lines {
		val = strings.TrimSpace(val)
		temp, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return []float64{}, errors.New("failed to convert string to float")
		}
		gg = append(gg, temp)
	}
	return gg, nil

}
