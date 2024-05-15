package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Number int
}

// Interfaces são implementatas automaticamente aos tipos que respeitarem sua descrição
// Interfaces só implementam métodos e não propriedades
type Person interface {
	Deactivate()
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

func Deactivation(person Person) {
	person.Deactivate()
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
	Deactivation(&client1)

	fmt.Println(client1.Name)
	fmt.Println(client1.Age)
	fmt.Println(client1.Active)
}
