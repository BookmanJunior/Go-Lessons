package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// TODO: 18, 21, 15, 5

func main() {
	// CountToTen()
	// Squared()
	// MTableOf3()
	// Fibonacci(10)
	// FizzBuzz()
	// fmt.Println(isPrime(7))
	// Factor(18)
	// result := SumOfNum(4)
	// fmt.Printf("Sum of %d: %d", 4, result)
	// IsPositive()
	// Factorial(5)
	// IsNumAPalindrome(222)
	// Pyramid(10)
	// InfiniteFactorial()
	// Sum()
	// AllBut2or3()
	// PrintPrimes()
	// SumOfRangeOfNumbers()
	IsDivisibleByNum()
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

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Factor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

func SumOfNum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
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

func InfiniteFactorial() {
	var num int
	prompt := "\nEnter a positive integer to get its factorials: "
	fmt.Print(prompt)
	fmt.Scan(&num)
	for num > 0 {
		Factorial(num)
		fmt.Print(prompt)
		fmt.Scan(&num)
	}
}

func Sum() {
	for {
		var n1, n2 float64
		prompt := "Enter two numbers separated by space: "
		fmt.Print(prompt)
		fmt.Scan(&n1, &n2)
		result := n1 + n2
		fmt.Printf("\n%v + %v = %v\n", n1, n2, result)
	}
}

func PrintPrimes() {
	for i := 1; i < 50; i++ {
		if isPrime(i) {
			fmt.Printf("%d is a prime number\n", i)
		}
	}
}

func AllBut2or3() {
	for i := 0; i < 30; i++ {
		if i%2 == 0 || i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func SumOfRangeOfNumbers() {
	for i := 1; i < 100; i++ {
		result := SumOfNum(i)

		if result > 200 {
			return
		}
		fmt.Printf("Sum of %d: %d\n", i, result)
	}
}

func IsDivisibleByNum() {
	var num int
	prompt := "Enter a number(Program exists if you enter a number divisible by 7): "
	fmt.Print(prompt)
	fmt.Scan(&num)
	for num%7 != 0 {
		fmt.Print(prompt)
		fmt.Scan(&num)
	}
}
