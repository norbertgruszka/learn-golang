package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("it's a string!", v)
	default:
		fmt.Println("I dunno")
	}
}

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	do(s)
}
