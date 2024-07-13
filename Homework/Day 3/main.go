package main

import "fmt"

func main() {
	result := Calculate(5, 10, adder)
	fmt.Println(result)
	DisplaySeprator()
	multiplier := Multiplier(5)
	fmt.Println(multiplier(10))
	DisplaySeprator()
	fmt.Println(GeneratePassword())
	DisplaySeprator()
	GenerateInvoiceNumber()
	DisplaySeprator()
	FormatDate()
}
