package main

import "fmt"

func main() {
	fmt.Println(sum(10, 50, 10, 22))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
