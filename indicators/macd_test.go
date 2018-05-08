package indicators

import (
	"fmt"
	"testing"
)

func TestMacd(t *testing.T) {
	tables := []struct {
		lastPrices  []float64
		bottomRange int
		topRange    int
		result      float64
	}{
		{[]float64{22.2734, 22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933, 22.1542, 22.3926, 22.3816, 22.6109, 23.3558, 24.0519, 23.753, 23.8324, 23.9516, 23.6338, 23.8225, 23.8722, 23.6537, 23.187, 23.0976, 23.326}, 12, 26, 0.731037},
		{[]float64{22.2734, 22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933, 22.1542, 22.3926, 22.3816, 22.6109}, 12, 26, 0.0},
	}

	for index, table := range tables {
		macd, err := Macd(table.lastPrices, table.bottomRange, table.topRange)
		if err != nil {
			fmt.Printf(err.Error())
		} else if macd != table.result {
			t.Errorf("MACD #%d was incorrect, got: %v, want %v.", index, macd, table.result)
		}
	}
}

func TestSignal(t *testing.T) {
	tables := []struct {
		lastPrices []float64
		result     float64
	}{
		{[]float64{22.2734, 22.194, 22.0847, 22.1741, 22.184, 22.1344, 22.2337, 22.4323, 22.2436, 22.2933, 22.1542, 22.3926, 22.3816, 22.6109, 23.3558, 24.0519, 23.753, 23.8324, 23.9516, 23.6338, 23.8225, 23.8722, 23.6537, 23.187, 23.0976, 23.326}, 23.597422},
	}

	for index, table := range tables {
		signalLine, err := SignalLine(table.lastPrices, 9)

		if err != nil {
			fmt.Println(err.Error())
		} else if signalLine != table.result {
			t.Errorf("Signal Line #%d was incorrect, got: %v, want %v.", index, signalLine, table.result)
		}
	}
}
