package indicators

import (
	"fmt"
	"testing"
)

func TestSma(t *testing.T) {
	tables := []struct {
		lastPrices []float64
		result     float64
	}{
		{[]float64{22.27, 22.19, 22.08, 22.17, 22.18, 22.13, 22.23, 22.43, 22.24, 22.29}, 22.221},
	}

	for index, table := range tables {
		sma, err := SimpleMovingAverage(table.lastPrices)

		if err != nil {
			fmt.Println(err.Error())
		} else if sma != table.result {
			t.Errorf("Sma #%d was incorrect, got: %v, want %v.", index, sma, table.result)
		}
	}
}

func TestEma(t *testing.T) {
	tables := []struct {
		lastPrices  []float64
		previousEma float64
		counter     int
		result      float64
	}{
		{[]float64{22.2734, 22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933}, 0, 0, 22.22475},
		{[]float64{22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933, 22.1542}, 0, 0, 22.21283},
		{[]float64{23.3558, 24.0519, 23.753, 23.8324, 23.9516, 23.6338, 23.8225, 23.8722, 23.6537, 23.187, 23.0976, 23.326}, 0, 0, 23.628125},
		{[]float64{22.2734, 22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933, 22.1542, 22.3926, 22.3816, 22.6109, 23.3558, 24.0519, 23.753, 23.8324, 23.9516, 23.6338, 23.8225, 23.8722, 23.6537, 23.187, 23.0976, 23.326}, 0, 0, 22.897088},
		{[]float64{23.8324, 23.9516, 23.6338, 23.8225, 23.8722, 23.6537, 23.187, 23.0976, 23.326}, 0, 0, 23.597422},
	}

	for index, table := range tables {
		ema, err := ExponentialMovingAverage(table.lastPrices, table.previousEma, table.counter)

		if err != nil {
			fmt.Println(err.Error())
		} else if ema != table.result {
			t.Errorf("Ema #%d was incorrect, got: %v, want %v.", index, ema, table.result)
		}
	}
}
