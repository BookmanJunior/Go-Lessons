package main

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func parseFloat(userInput string) (string, error) {
	var isDot = isDot(userInput)

	if !isDot {
		return "", errors.New("Not a float")
	}

	isFloat, err := strconv.ParseFloat(userInput, 64)
	inputType := reflect.TypeOf(isFloat).String()

	if err != nil {
		return inputType, errors.New("Not a float")
	}

	return inputType, nil
}

func isDot(userInput string) bool {
	arrOfInput := strings.Split(userInput, " ")[0]

	for i := 0; i < len(arrOfInput); i++ {
		if string(arrOfInput[i]) == "." {
			return true
		}
	}
	return false
}
