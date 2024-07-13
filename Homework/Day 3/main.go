package main

import (
	"math"
)

func main() {
	// PrintGreeting()
	// DisplaySeparator()
	// NotifyStart()
	// PrintNumber(50)
	// GreetUser("Azam")
	// ToggleLight(true)
	// fmt.Println(adder(5, 5))
	// fmt.Println(concatenator("Hello ", "World"))
	// fmt.Println(isPositive(10))
	// fmt.Println(isServiceActive())
	// fmt.Println(isEven(9))
	// result := Calculate(5, 10, adder)
	// fmt.Println(result)
	// multiplier := Multiplier(5)
	// fmt.Println(multiplier(10))
	// stringRepeater := StringRepeater(5)
	// fmt.Println(stringRepeater("Hello"))
	// conditionalRepeater := ConditionalRepeater(true)
	// conditionalRepeater(":)")
	// Execute(false, ToggleLight)
	// fmt.Println(Apply(2, Pow4))
	// fmt.Println(GeneratePassword())
	GenerateInvoiceNumber()
	// FormatDate()
}

func Pow4(num int) int {
	return int(math.Pow(float64(num), 4))
}
