package main

import "fmt"

func main() {
	salaries := map[string]int{
		"Wesley": 1000,
		"João":   2000,
		"Maria":  3000,
	}

	fmt.Println(salaries["Wesley"])
	delete(salaries, "Wesley")

	salaries["Wes"] = 1500
	fmt.Println(salaries["Wes"])

	sal := make(map[string]int)
	sal["Wesley"] = 1000

	for name, salary := range salaries {
		fmt.Printf("O salário de %s é %d \n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("O salário é %d \n", salary)
	}
}
