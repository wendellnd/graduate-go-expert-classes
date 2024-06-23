package main

import "time"

func worker(workerId int, data chan int) {
	for x := range data {
		println("Worker", workerId, "received", x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workersQuantity := 10

	for i := 0; i < workersQuantity; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
