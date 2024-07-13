package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	currentOrder := 100
	var orderFormatted string

	if currentOrder < 10 {
		orderFormatted = "00" + strconv.Itoa(currentOrder)
	} else if currentOrder < 100 {
		orderFormatted = "0" + strconv.Itoa(currentOrder)
	} else {
		orderFormatted = strconv.Itoa(currentOrder)
	}
	fmt.Printf("INV-%v-0%v-%v\n", currentDate.Year(), int(currentDate.Month()), orderFormatted)
}

func FormatDate() {
	currentDate := time.Now()
	fmt.Printf("%v %v %v\n", currentDate.Day(), currentDate.Month(), currentDate.Year())
}
