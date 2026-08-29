package main

import (
	"bufio"
	"fmt"
	"guess-it-2/calculators"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateRange(regressionPrediction float64, residualStdDev float64, numbersLen int) (int, int) {
	lower, upper := calculators.PredictRange(regressionPrediction, residualStdDev, numbersLen)
	return int(math.Floor(lower)), int(math.Ceil(upper)) // rounding to lower decimal (math.Floor) and to upper decimal (math.Ceil)
}

func main() {
	var (
		userInput                []float64
		lower, upper, numbersLen int
	)
	scanner := bufio.NewScanner(os.Stdin) // "scanner" reads the user's input from CLI (using os.Stdin)

	// Read one input line at a time from user input
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" { // Bad input
			fmt.Println("Empty value is not accepted, please, give a number")
			return
		}

		num, err := strconv.ParseFloat(line, 64) // convert string to float64
		if err != nil {
			fmt.Println("Error in number from string conversion")
			return
		}

		// Add the first input to accepted history
		// 1st input range
		userInput = append(userInput, num) // add user's input to history
		numbersLen = len(userInput)        // update length of userInput's history

		if numbersLen <= 1 {
			lower = int(num - 50) // 1st input's range
			upper = int(num + 50)

			// Calculate the range after the 1st value of userInput
		} else {
			a, b, err := calculators.LinearRegressionLine(userInput)
			if err != nil {
				fmt.Println(err)
				return
			}

			regressionPrediction := (a * (float64(numbersLen))) + b                  // prediction of regression line for the next input (last numbersLen)
			residualStdDev, err := calculators.ResidualStdDeviation(userInput, a, b) // the diferrence between the actual y value and the predicted's value from regression line
			if err != nil {
				fmt.Println(err)
			}

			lower, upper = calculateRange(regressionPrediction, residualStdDev, numbersLen)
		}
		fmt.Println(lower, upper)
	}

	// Check if the scanner has completed successfully
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
