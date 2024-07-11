package main

import (
	"errors"
)

const errorMessage = "a can't be less than b"

func FindNumberOfSegments(a, b int) (int, error) {
	if a < b {
		return 0, errors.New(errorMessage)
	}

	return a / b, nil
}

func FindRemainingSegments(a, b int) (int, error) {
	if a < b {
		return 0, errors.New(errorMessage)
	}

	return a % b, nil
}
