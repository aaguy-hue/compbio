package util

func CalculatePValue(score, k int) float32 {
	// TODO: This is a placeholder for the actual p-value calculation.
	// The formula for calculating the p-value can vary based on the context and the statistical model used.
	// Here, we return a simple calculation as an example.
	return float32(score) / float32(k)
}

func CalculateEValue(score, k, dbSize int) float32 {
	// TODO: This is a placeholder for the actual e-value calculation.
	// The formula for calculating the e-value can vary based on the context and the statistical model used.
	// Here, we return a simple calculation as an example.
	return float32(dbSize) * CalculatePValue(score, k)
}
