package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	formattedMonth := strings.Split(currentDate.Format(time.RFC3339), "-")[1]
	currentOrder := 1
	formattedOrder := padNumber(currentOrder)
	fmt.Printf("INV-%v-%v-%v\n", currentDate.Year(), formattedMonth, formattedOrder)
}

func FormatDate() {
	currentDate := time.Now()
	fmt.Printf("%v %v %v\n", currentDate.Day(), currentDate.Month(), currentDate.Year())
}

func padNumber(num int) string {
	numToString := strconv.Itoa(num)
	if num < 10 {
		return "00" + numToString
	} else if num < 100 {
		return "0" + numToString
	} else {
		return numToString
	}
}
