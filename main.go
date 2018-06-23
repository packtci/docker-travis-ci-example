package main

import "fmt"

func average(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum = sum + val
	}
	return sum / len(numbers)
}

func main() {
	avg := average([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	})
	fmt.Printf("average of the numbers is %v\n", avg)
}
