package indicators

import (
	"errors"
)

func Macd(lastPrices []float64, bottomRange int, topRange int) (float64, error) {
	if len(lastPrices) < topRange {
		return 0.0, errors.New("The array of the last prices is too short. Its length must be greater than topRange\n")
	}

	pEma, err := ExponentialMovingAverage(lastPrices[len(lastPrices)-bottomRange:len(lastPrices)], 0, 0)

	if err != nil {
		return 0.0, err
	}

	xpEma, err := ExponentialMovingAverage(lastPrices[len(lastPrices)-topRange:len(lastPrices)], 0, 0)

	if err != nil {
		return 0.0, err
	}

	roundedNumber, err := round(pEma - xpEma)

	if err != nil {
		return 0.0, err
	}

	return roundedNumber, nil
}

func SignalLine(lastPrices []float64, signalRange int) (float64, error) {
	return ExponentialMovingAverage(lastPrices[len(lastPrices)-signalRange:len(lastPrices)], 0, 0)
}

func Bullish(macd float64) bool {
	return (macd > 0)
}

func Bearish(macd float64) bool {
	return (macd < 0)
}
