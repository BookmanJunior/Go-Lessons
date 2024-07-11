package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userInputType, err := checkType(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(input, "is of type", userInputType)
}

func checkType(userInput string) (string, error) {
	float, floatErr := parseFloat(userInput)
	num, numErr := strconv.Atoi(userInput)
	bool, errBool := strconv.ParseBool(userInput)

	switch {
	case floatErr == nil:
		return float, nil
	case numErr == nil:
		return reflect.TypeOf(num).String(), nil
	case errBool == nil:
		return reflect.TypeOf(bool).String(), nil
	default:
		return reflect.TypeOf(userInput).String(), nil
	}

}
