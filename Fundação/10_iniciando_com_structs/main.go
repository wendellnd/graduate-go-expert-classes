package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client1 := Client{
		Name:   "Wendell",
		Age:    21,
		Active: true,
	}

	fmt.Println(client1.Name)
	fmt.Println(client1.Age)
	fmt.Println(client1.Active)
}
