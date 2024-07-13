package main

func Calculate(num1, num2 int, action func(num1, num2 int) int) int {
	return action(num1, num2)
}

func Execute(input bool, boolReader func(input bool)) {
	boolReader(input)
}

func Apply(num int, action func(num2 int) int) int {
	return action(num)
}
