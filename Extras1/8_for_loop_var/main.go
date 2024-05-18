package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	// New feature from 1.22
	for i := range 10 {
		fmt.Println(i)
	}

	// The same result from the for above
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for _, v := range values {
		// fix to versions before 1.22
		//v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
