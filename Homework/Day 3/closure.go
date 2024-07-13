package main

import (
	"fmt"
	"strings"
)

func Multiplier(factor int) func(num int) int {
	return func(num int) int {
		return num * factor
	}
}

func StringRepeater(n int) func(input string) string {
	return func(input string) string {
		return strings.Repeat(input, n)
	}
}

func ConditionalRepeater(condition bool) func(input string) {
	return func(input string) {
		if condition {
			fmt.Println(input)
		}
	}
}
