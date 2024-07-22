package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// 1
	// var acc float64 = 100000
	// for i := 0; i < 10; i++ {
	// 	var depositAmount float64 = 100000
	// 	err := deposit(&acc, depositAmount)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Printf("Deposited %v to account. Current balance: %v\n", depositAmount, acc)
	// }

	// 3
	// var account1 float64 = 500000
	// var account2 float64 = 0
	// fmt.Println("Before:")
	// fmt.Printf("Account 1 balance: %v\n", account1)
	// fmt.Printf("Account 1 balance: %v\n", account2)

	// err := transfer(&account1, 1000, &account2)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("After:")
	// fmt.Printf("Account 1 balance: %v\n", account1)
	// fmt.Printf("Account 1 balance: %v\n", account2)

	// 4
	// var amount float64 = 100000
	// for i := 0; i < 20; i++ {
	// 	err := simpleInterest(&amount, 5)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Println(amount)
	// }

	// 5
	// currentRate := 74.5
	// fmt.Println(convertCurrencyToRubles(currentRate, 1000))
	// updateCurrencyRate(&currentRate, 85)
	// fmt.Println(convertCurrencyToRubles(currentRate, 3500))

}

func deposit(balance *float64, amount float64) error {
	updatedBalance := math.Round((*balance+amount)*100) / 100
	if updatedBalance > 500000 {
		return errors.New("Reached savings limit")
	}
	*balance = updatedBalance
	return nil
}

func withdraw(balance *float64, amount float64) error {
	if *balance < amount {
		return fmt.Errorf("Amount can't exceed balance")
	}
	*balance -= amount
	return nil
}

func statement(balance *float64) {
	fmt.Printf("Current balance: %v", *balance)
}

func transfer(balance *float64, transferAmount float64, receiver *float64) error {
	var limit float64 = 100000

	if transferAmount > limit {
		return fmt.Errorf("%v exceeds allowed transfer amount", transferAmount)
	}

	err := withdraw(balance, transferAmount)

	if err != nil {
		return err
	}

	err = deposit(receiver, transferAmount)

	if err != nil {
		return err
	}
	return nil
}

func simpleInterest(balance *float64, monthlyRate float64) error {
	var limit float64 = 150000
	rate := (monthlyRate / 100) / 12
	updatedAmount := *balance * rate

	if updatedAmount > limit {
		return errors.New("Reached deposit limit")
	}
	deposit(balance, updatedAmount)
	return nil
}

func updateCurrencyRate(rate *float64, newRate float64) {
	if newRate < 70 {
		fmt.Println("Rate is too low.")
		return
	}

	*rate = newRate
}

func convertCurrencyToRubles(rate float64, amount float64) float64 {
	return math.Round((amount/rate)*100) / 100
}
