package main

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compare[T comparable](a, b T) bool {
	return a == b

	// Error -> Pode ser burlado usando o package constraints
	// return a > b
}

func main() {
	m := map[string]int{"Wendell": 1000, "João": 2000}
	println(Sum(m))

	m2 := map[string]float64{"Wendell": 1500.22, "João": 2333.66}
	println(Sum(m2))

	m3 := map[string]MyNumber{"Wendell": 1500, "João": 2333}
	println(Sum(m3))

	println(Compare(int(10), 10))
}
