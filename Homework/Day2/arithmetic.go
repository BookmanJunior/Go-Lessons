package main

import "math"

func FindMean(a, b int) float64 {
	return float64(a+b) / 2
}

func FindGeometricMean(a, b int) float64 {
	return math.Sqrt(float64(a * b))
}

func FindSumOfSquares(a, b int) int {
	return (a * a) + (b * b) + (a*b)*2 - (2 * (a * b))
}

func FindDifferenceOfSquares(a, b int) int {
	return (a + b) * (a - b)
}

func FindMultiplicationOfSquares(a, b int) int {
	return (a * a) * (b * b)
}

func FindDivisionOfSquares(a, b int) int {
	return (a * a) / (b * b)
}

func ConvertCmToM(l int) int {
	return int(math.Floor(float64(l) / 100))
}

func ConvertKgToT(m int) int {
	return int(math.Floor(float64(m) / 1000))
}

func ConvertByteToKb(b int) int {
	return int(math.Floor(float64(b) / 1024))
}
