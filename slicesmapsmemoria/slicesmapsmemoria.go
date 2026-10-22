package sliceexemplo

func SliceTest() {
	base := make([]int, 3, 4)
	a := append(base, 100)
	b := append(base, 200)

	println(base)
	println(a[3])
	println(b[0])
}

func MapTest() {
	m := map[string]int{"test": 1, "test2": 2}

	v, ok := m["test"]

	println(v, ok)
}
