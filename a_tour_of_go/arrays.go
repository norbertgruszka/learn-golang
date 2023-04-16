package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	exampleSlice := []bool{true, false}
	fmt.Println(exampleSlice)

	var s []int
	s = append(s, 1)
	s = append(s, 2)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow {
		fmt.Println(v)
	}

	m := make(map[string]Vertex)
	m["foo"] = Vertex{
		10.3, 12.9,
	}
	fmt.Println(m)

	value, ok := m["foo"]
	if ok {
		fmt.Printf("Value %s in map\n", value)
	}
}
