package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword() string {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var pass = ""

	for len(pass) < 8 {
		randomNum := rand.Int63n(int64(len(letters)))
		pass += string(letters[randomNum])
	}
	return pass
}

func GenerateInvoiceNumber() {
	currentDate := time.Now()
	currentOrder := 1
	fmt.Printf("INV-%v-%02d-%03d\n", currentDate.Year(), currentDate.Month(), currentOrder)
}

func FormatDate() {
	currentDate := time.Now()
	fmt.Printf("%v %v %v\n", currentDate.Day(), currentDate.Month(), currentDate.Year())
}

// Old solution for formatting GenerateInvoiceNumber int
// func padNumber(num int) string {
// 	numToString := strconv.Itoa(num)
// 	if num < 10 {
// 		return "00" + numToString
// 	} else if num < 100 {
// 		return "0" + numToString
// 	} else {
// 		return numToString
// 	}
// }
