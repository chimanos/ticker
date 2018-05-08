package indicators

func SimpleMovingAverage(lastPrices []float64) (float64, error) {
	sum := 0.0

	for _, lastPrice := range lastPrices {
		sum += lastPrice
	}

	roundedSum, err := round(sum / float64(len(lastPrices)))

	return roundedSum, err
}

func multiplier(n int) (float64, error) {
	return round(float64((2 / (n + 1))))
}

func ExponentialMovingAverage(prices []float64, previousEma float64, counter int) (float64, error) {
	if counter == len(prices) {
		return round(previousEma)
	}

	if counter == 0 {
		sma, err := SimpleMovingAverage(prices)

		if err != nil {
			return 0.0, nil
		}

		return ExponentialMovingAverage(prices, sma, counter+1)
	}

	multiplier, err := multiplier(len(prices))

	if err != nil {
		return 0.0, nil
	}

	roundedNumber, err := round((multiplier * (prices[counter] - previousEma)) + previousEma)

	if err != nil {
		return 0.0, nil
	}

	return ExponentialMovingAverage(prices, roundedNumber, counter+1)
}
