package indicators

import (
	"fmt"
	"strconv"
)

func round(aNumber float64) (float64, error) {
	roundedNbStr := fmt.Sprintf("%.6f", aNumber)
	roundedNb, err := strconv.ParseFloat(roundedNbStr, 64)

	if err != nil {
		return 0.0, err
	}

	return roundedNb, nil
}
