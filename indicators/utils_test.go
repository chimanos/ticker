package indicators

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	tables := []struct {
		numberToRound float64
		result        float64
	}{
		{10.51423657896542, 10.514237},
		{10.51423647896542, 10.514236},
		{0, 0},
		{-0.0000001, 0},
	}

	for index, table := range tables {
		roundedNumber, err := round(table.numberToRound)

		if err != nil {
			fmt.Println(err.Error())
		} else if roundedNumber != table.result {
			t.Errorf("Rounded Number #%d was incorrect, got: %v, want %v.", index, roundedNumber, table.result)
		}
	}
}
