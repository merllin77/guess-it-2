package calculators

import "math"

// Dynamically calculates a prediction range for the next number
func PredictRange(regressionPrediction float64, residualStdDev float64, n int) (float64, float64) {
	var lower, upper, multiplier float64

	// After 100 values, the regression line is stable enough.
	// Round the prediction and always use a ±6 range (total width: 12).
	if n > 100 {
		regressionPrediction = math.Round(regressionPrediction) // round the decimal to nearest int
		lower = regressionPrediction - 6                        // fixed range +-6
		upper = regressionPrediction + 6
		return lower, upper
	}

	// Dynamically narrowing range if Residual SD < 35
	if residualStdDev < 35 {
		switch {

		case n == 2:
			lower = regressionPrediction - 20 // fixed range +-20
			upper = regressionPrediction + 20
			return lower, upper

		case n > 2 && n <= 20:
			multiplier = 0.5

		case n > 20 && n <= 50:
			multiplier = 0.4

		case n > 50 && n <= 75:
			multiplier = 0.3

			// default: 			// We check above n > 100 and returning the values
			// 	multiplier = 0.2
		}

	} else {
		// Standard multiplier if Residual SD >= 35
		multiplier = 0.4
	}

	residualStdDev *= multiplier
	lower = regressionPrediction - residualStdDev
	upper = regressionPrediction + residualStdDev
	return lower, upper
}
