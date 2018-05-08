package indicators

func AverageGain(lastGain float64, lastGains []float64) (float64, error) {
	currentGain := 0.0

	for _, lastGain := range lastGains {
		currentGain += lastGain
	}

	return round((lastGain*13 + currentGain) / 14)
}

func AverageLoss(lastLoss float64, lastLosses []float64) (float64, error) {
	currentLoss := 0.0

	for _, lastLoss := range lastLosses {
		currentLoss += lastLoss
	}

	return round((lastLoss*13 + currentLoss) / 14)
}

func Rsi(averageGain float64, averageLoss float64) (float64, error) {
	rs := (averageGain / averageLoss)
	return round(100 - (100 / (1 + rs)))
}

func OverSold(rsi float64) bool {
	return (rsi < 30)
}

func OverBought(rsi float64) bool {
	return (rsi > 70)
}
