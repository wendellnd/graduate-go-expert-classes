package main

import "fmt"


type ID int

var( 
	a = 10
	b ID = 10
)

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	// myArray[3] = 40

	for index, item := range myArray {
		fmt.Println(index, "-", item)
	}
}
