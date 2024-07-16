package main

import (
	"fmt"
	"strings"
)

func main() {
	// PrintGreeting()
	// PrintWeather()
	// PrintTrafficLight()
	// GetGrade()
	// GetDiscount()
	// GetTemperatureDescription()
	// CheckNumber(100)
	// CheckAge(10)
	// CheckPassword("1234")
	// Add(5, 5)
	// CompareString("Hello", "World")
	// fmt.Println(Max(5, 20))
	// fmt.Println(operation(5, 10))
	// fmt.Println(concat("Hello", "World"))
}

func PrintGreeting() {
	fmt.Print("Type Morning or Evening: ")
	var dayType string
	var result string
	fmt.Scan(&dayType)

	if strings.ToLower(dayType) == "morning" {
		result = "Good Morning!"
	} else if strings.ToLower(dayType) == "evening" {
		result = "Good Evening!"
	} else {
		result = "Wrong Input"
	}

	fmt.Println(result)
}

func PrintWeather() {
	fmt.Print("Type Sunny or Cloudy: ")
	var weatherType string
	var result string
	fmt.Scan(&weatherType)

	if strings.ToLower(weatherType) == "sunny" {
		result = "It's sunny!"
	} else if strings.ToLower(weatherType) == "cloudy" {
		result = "It's cloudy!"
	} else {
		result = "Wrong Input"
	}

	fmt.Println(result)

}

func PrintTrafficLight() {
	fmt.Print("Choose traffics current state - Red, Yellow or Green: ")
	var lightColor string
	var result string
	fmt.Scan(&lightColor)

	if strings.ToLower(lightColor) == "red" {
		result = "Stop!"
	} else if strings.ToLower(lightColor) == "yellow" {
		result = "Attention!"
	} else if strings.ToLower(lightColor) == "green" {
		result = "You may cross!"
	} else {
		result = "Wrong Input"
	}

	fmt.Println(result)
}

func GetGrade() {
	fmt.Print("Enter score: ")
	var score int
	var result string
	fmt.Scan(&score)

	switch {
	case score >= 85:
		result = "A"
	case score >= 70:
		result = "B"
	case score < 70:
		result = "C"
	default:
		result = "Wrong Input"
	}

	fmt.Println(result)
}

func GetDiscount() {
	fmt.Print("Enter amount: ")
	var amount int
	var result int
	fmt.Scan(&amount)

	switch {
	case amount > 1000:
		result = 10
	case amount > 500 && amount >= 100:
		result = 5
	case amount <= 500:
		result = 0
	default:
		fmt.Println("Wrong input")
		return
	}

	fmt.Printf("You get %d discount\n", result)
}

func GetTemperatureDescription() {
	fmt.Print("Enter degree in C: ")
	var temperature int
	var result string
	fmt.Scan(&temperature)

	if temperature < 10 {
		result = "Cold"
	} else if temperature <= 25 {
		result = "Warm"
	} else {
		result = "Hot"
	}

	fmt.Println(result)
}

func CheckNumber(num int) {
	if num == 0 {
		fmt.Println("Zero")
		return
	}

	if num > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}

func CheckAge(age uint) {
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

func CheckPassword(password string) {
	const pass string = "1234"

	if password == pass {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Password is incorrect")
	}
}

func Add(n1, n2 int) int {
	return n1 + n2
}

func CompareString(str1, str2 string) string {
	if str1 == str2 {
		return "String are equal"
	}

	return "Strings are not equal"
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func operation(n1, n2 int) int {
	return n1 - n2
}

func concat(str1, str2 string) string {

	if str2 != "" {
		return str1 + " " + str2
	}

	return str1 + str2
}
