package indicators

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func round(aNumber float64) float64 {
	roundedNbStr := fmt.Sprintf("%.6f", aNumber)
	roundedNb, err := strconv.ParseFloat(roundedNbStr, 64)

	if err != nil {
		log.Printf("Unable to parse \"%s\" to float \n", roundedNbStr)
		log.Println(err)
		os.Exit(1)
	}

	return roundedNb
}
