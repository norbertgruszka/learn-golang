package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	numbers := make([]string, 4)
	for i, v := range ip {
		val, err := strconv.Itoa(int(v))
		if err != nil {
			fmt.Println("Coudln't convert %v", v)
			return ""
		}
    
	}
	return strings.Join(numbers, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
