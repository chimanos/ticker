package indicators

import (
	"log"
	"os"
)

func macd(lastPrices []float64) float64 {
	if len(lastPrices) < 26 {
		log.Printf("The array of the last prices is too short, got: %d, want: 26.", len(lastPrices))
		os.Exit(1)
	}

	var pEma float64 = ema(lastPrices[len(lastPrices)-12:len(lastPrices)], 0, 0)
	var xpEma float64 = ema(lastPrices[len(lastPrices)-26:len(lastPrices)], 0, 0)

	return round(pEma - xpEma)
}

func signalLine(lastPrices []float64) float64 {
	return ema(lastPrices[len(lastPrices)-9:len(lastPrices)], 0, 0)
}

func isTrendUp(macd float64) bool {
	return (macd > 0)
}

func isTrendDown(macd float64) bool {
	return (macd < 0)
}
