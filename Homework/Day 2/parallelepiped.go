package main

func FindParallelepipedsVolume(a, b, c int) int {
	return a * b * c
}

func FindParallelepipedsSurfaceArea(a, b, c int) int {
	return 2 * (a*b + b*c + a*c)
}
