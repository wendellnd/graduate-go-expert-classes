package main

func main() {
	var x interface{} = "Wendell Nascimento"

	println(x.(string))

	res, ok := x.(int)
	println(res, ok)

	// Panic
	//res2 := x.(int)
	//println(res2)
}
