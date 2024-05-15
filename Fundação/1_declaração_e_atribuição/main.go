package main

var (
	b int     = 10
	c string  = "Wendell"
	d float64 = 1.2
)

func main() {
	a := "Hello, World!"
	println(a)
	println(c)
	println(d)

	x()
}

// x pode acessar o valor de b,c e d, mas não o valor de a
func x() {
	print(b)
}
