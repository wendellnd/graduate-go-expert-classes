package main

import "fmt"


type ID int

var( 
	a = 10
	b ID = 10
)

func main() {
	fmt.Printf("O tipo de de a é: %T \n", a)
	fmt.Printf("O valor de a é: %v \n", a)

	fmt.Printf("O tipo de de b é: %T \n", b)
}
