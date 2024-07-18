package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// CountToTen()
	// Squared()
	// MTableOf3()
	// Fibonacci(10)
	// FizzBuzz()
	// isPrime(7)
	// Factor(18)
	// SumOfNum(4)
	// IsPositive()
	// Factorial(5)
	// IsNumAPalindrome(222)
	Pyramid(10)
}

func CountToTen() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Squared() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * 2)
	}
}

func MTableOf3() {
	for i := 1; i <= 10; i++ {
		result := 3 * i
		fmt.Printf("%dx%d=%d\n", 3, i, result)
	}
}

func Fibonacci(n int) {
	n1, n2 := 1, 1

	for i := 2; i < n; i++ {
		result := n1 + n2
		n1, n2 = n2, result
		fmt.Println(result)
	}

	// Array version
	sequence := [10]int{1, 1}

	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-2] + sequence[i-1]
	}

	fmt.Println(sequence)
}

func GreatestCommonD() {

}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) {
	if n <= 1 {
		fmt.Printf("%d is not a prime number", n)
		return
	}

	if n == 2 {
		fmt.Printf("%d is a prime number", n)
		return
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d is not a prime number", n)
			return
		}
	}

	fmt.Printf("%d is a prime number", n)
}

func Factor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

func SumOfNum(n int) {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Sum of %d: %d", n, sum)
}

func IsPositive() {
	var num int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&num)
	for num < 0 {
		fmt.Printf("%d is not a positive integer.\n Try again: ", num)
		fmt.Scan(&num)
	}
	fmt.Printf("You entered: %d", num)
}

func Factorial(n int) {
	if n > 1000 {
		fmt.Println("Enter a number less than 1000")
		return
	}

	sum := 1

	for i := 2; i <= n; i++ {
		sum *= i
	}

	fmt.Printf("Factorial of %d is: %d", n, sum)
}

func IsNumAPalindrome(n int) {
	if n < 10 {
		fmt.Printf("%d is a palindrome", n)
		return
	}
	nToS := strconv.Itoa(n)
	var nReversed string

	for i := len(nToS) - 1; i >= 0; i-- {
		nReversed += string(nToS[i])
	}

	if nToS == nReversed {
		fmt.Printf("%d is a palindrome", n)
		return
	}
	fmt.Printf("%d is not a palindrome", n)
}

func Pyramid(n int) {
	for i := 1; i <= n; i += 2 {
		mean := int(math.Floor(float64(n-i) / 2))
		emptySpaces := strings.Repeat(" ", mean)
		fmt.Printf("%v%v%v\n", emptySpaces, strings.Repeat("*", i), emptySpaces)
	}
}
