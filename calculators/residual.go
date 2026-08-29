package calculators

import (
	"errors"
	"math"
)

// Calculates how far actual values deviate from the regression line
func ResidualStdDeviation(data []float64, a float64, b float64) (float64, error) {
	var res, residualSum, residualStdDev float64
	n := float64(len(data))

	if n <= 1 {
		return 0, errors.New("Not enough values to calculate Residual Standard Deviation")
	}

	for i, num := range data {
		res = num - ((a * float64(i)) + b)
		res = res * res    // for giving the same number (range) for negative OR positive numbers (eg. -10 * -10 == 100)
		residualSum += res // sum of residuals
	}

	residualStdDev = math.Sqrt(residualSum / n) // mean (average) of residualSum | use of Sqrt for finding real standard deviation value
	return residualStdDev, nil
}
