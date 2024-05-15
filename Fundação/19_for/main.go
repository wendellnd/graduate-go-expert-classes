package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		println(index, "-", number)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	/*
		for {
			println("Infinity")
		}
	*/
}
