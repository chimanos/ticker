package indicators

import (
	"testing"
)

func TestAverageGain(t *testing.T) {
	tables := []struct {
		lastGain  float64
		lastGains []float64
		result    float64
	}{
		{0, []float64{0, 0.06, 0, 0.72, 0.5, 0.27, 0.33, 0.42, 0.24, 0, 0.14, 0, 0.67, 0}, 0.239286},
		{0.2392857142857143, []float64{0}, 0.222194},
	}

	for index, table := range tables {
		averageGain := AverageGain(table.lastGain, table.lastGains)

		if averageGain != table.result {
			t.Errorf("Average Gain #%d was incorrect, got: %v, want %v.", index, averageGain, table.result)
		}
	}
}

func TestAverageLoss(t *testing.T) {
	tables := []struct {
		lastLoss   float64
		lastLosses []float64
		result     float64
	}{
		{0, []float64{0.25, 0, 0.54, 0, 0, 0, 0, 0, 0, 0.20, 0, 0.42, 0, 0}, 0.100714},
	}

	for index, table := range tables {
		averageLoss := AverageLoss(table.lastLoss, table.lastLosses)

		if averageLoss != table.result {
			t.Errorf("Average Loss #%d was incorrect, got: %v, want %v.", index, averageLoss, table.result)
		}
	}
}

func TestRsi(t *testing.T) {
	tables := []struct {
		averageGain float64
		averageLoss float64
		result      float64
	}{
		{0.239286, 0.100714, 70.378235},
	}

	for index, table := range tables {
		rsi := Rsi(table.averageGain, table.averageLoss)

		if rsi != table.result {
			t.Errorf("RSI #%d was incorrect, got: %v, want %v.", index, rsi, table.result)
		}
	}
}
