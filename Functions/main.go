package main

import "fmt"

type TransformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, createTransformer(2))
	tripled := transformNumbers(&numbers, createTransformer(3))

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, tranform TransformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, tranform(val))
	}

	return dNumbers
}

func createTransformer(factor int) TransformFn {
	return func(number int) int {
		return number * factor
	}
}
