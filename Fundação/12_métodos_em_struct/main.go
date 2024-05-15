package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Number int
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func (c *Client) Deactivate() {
	c.Active = false
}

func main() {
	client1 := Client{
		Name:   "Wendell",
		Age:    21,
		Active: true,
		Address: Address{
			Street: "Rua Tal Tal Tal",
			City:   "Cidade",
			State:  "SP",
			Number: 10,
		},
	}

	client1.Deactivate()

	fmt.Println(client1.Name)
	fmt.Println(client1.Age)
	fmt.Println(client1.Active)
}
