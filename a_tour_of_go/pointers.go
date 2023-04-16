package main

import "fmt"

func main() {
	i, j := 42, 50

	p := &i
	*p = 21
	fmt.Println(i)

	q := &j
	*q = *q / 37
	fmt.Println(j)
}
