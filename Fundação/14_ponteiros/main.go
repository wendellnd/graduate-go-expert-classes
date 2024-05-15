package main

func main() {
	a := 10
	println(&a)

	var pointer *int = &a
	println(pointer)

	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)

	*b = 30
	println(a)
}
