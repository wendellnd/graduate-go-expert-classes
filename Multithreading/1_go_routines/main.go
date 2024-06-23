package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s is running\n", name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("task1")
	go task("task2")
	time.Sleep(15 * time.Second)
}
