package main

import "math"

func FindCubesVolume(a int) int {
	return int(math.Pow(float64(a), 3))
}

func FindCubesSurfaceArea(a int) int {
	return 6 * a
}
