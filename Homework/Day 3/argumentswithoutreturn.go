package main

import "fmt"

func PrintNumber(num int) {
	fmt.Println(num)
}

func GreetUser(name string) {
	fmt.Printf("Welcome %v\n", name)
}

func ToggleLight(condition bool) {
	if condition {
		fmt.Println("Lights on")
	} else {
		fmt.Println("Lights off")
	}
}
