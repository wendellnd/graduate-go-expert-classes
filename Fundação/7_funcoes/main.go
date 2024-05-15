package main

import "fmt"

func main() {
	fmt.Println(sum(10, 50))
}

func sum(a, b int) (int, bool) {
	result := a + b
	return result, result >= 50
}
