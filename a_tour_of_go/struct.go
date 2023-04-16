package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{X: 1, Y: 2}
	v2 = Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2)
}
