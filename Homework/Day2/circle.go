package main

const pi = 3.14

func FindCirclesDiameterLength(d int) float64 {
	return float64(d) * pi
}

func FindCirclesCircumference(r int) float64 {
	return 2 * pi * float64(r)
}

func FindCircleArea(r int) float64 {
	return pi * float64(r)
}
