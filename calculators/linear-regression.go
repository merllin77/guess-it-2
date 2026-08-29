package calculators

import "errors"

// Calculates the Linear Regession Line of given data
func LinearRegressionLine(data []float64) (float64, float64, error) {
	var a, b, n, sumX, sumY, sumXY, sumXX float64
	n = float64(len(data))

	if n <= 1 {
		return 0, 0, errors.New("Not enough values to calculate Linear Regression Line")
	}

	sumX, sumY, sumXY, sumXX, _ = sums(data)

	a = ((n * sumXY) - (sumX * sumY)) / ((n * sumXX) - (sumX * sumX))
	b = (sumY - (a * sumX)) / n
	return a, b, nil
}
