package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12, 14, 16}

	fmt.Printf("len=%d, cap=%d %v \n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v \n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	// O slice dobra a sua capacidade ao adicionar um item
	// Pois ele cria um array por trás dos panos com o tamanho original do slice
	s = append(s, 18)
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s)
}
