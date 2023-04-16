package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

var python, c, java bool
var i, j = 1, 2
var (
	ToBe   bool    = false
	MaxInt uint64  = 1<<64 - 1
	iFloat float64 = float64(i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func naked_return(sum int) (x, y int) {
	x = sum * 4
	y = sum + 2
	return
}

func main() {
	fmt.Println("Hello, here's your random number", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(3, 2))
	fmt.Println(swap("World", "Hello"))
	fmt.Println(naked_return(2))
	fmt.Println(c, python, java)
	fmt.Println(i, j)

	for k := 0; k < 10; k++ {
		fmt.Println(k)
	}

	sum := 1
	for sum < 100 {
		sum += sum
		if sum == 128 {
			fmt.Println(sum)
		}
	}
	fmt.Println(sum)

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("running on Linux")
	}
}
