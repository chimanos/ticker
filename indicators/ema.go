package indicators

func sma(lastPrices []float64) float64 {
	var sumPrices float64 = 0

	for _, lastPrice := range lastPrices {
		sumPrices += lastPrice
	}

	return round(sumPrices / float64(len(lastPrices)))
}

func multiplier(n int) float64 {
	return round(float64((2 / (n + 1))))
}

func ema(prices []float64, previousEma float64, counter int) float64 {
	if counter == len(prices) {
		return round(previousEma)
	}

	if counter == 0 {
		return ema(prices, sma(prices), counter+1)
	}

	return ema(prices, round((multiplier(len(prices))*(prices[counter]-previousEma))+previousEma), counter+1)
}
