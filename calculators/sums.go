package calculators

func sums(data []float64) (float64, float64, float64, float64, float64) {
	var sumX, sumY, sumXY, sumXX, sumYY float64

	for i, num := range data {
		sumX += float64(i)
		sumY += num
		sumXY += float64(i) * num
		sumXX += float64(i) * float64(i)
		sumYY += num * num
	}
	return sumX, sumY, sumXY, sumXX, sumYY
}
