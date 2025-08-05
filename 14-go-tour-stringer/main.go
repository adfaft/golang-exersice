package main

// ref: https://go.dev/tour/methods/18

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string {
	result := ""
	for i, x := range v {
		result += fmt.Sprintf("%v", x)
		if i < len(v)-1 {
			result += "."
		}
	}

	return result
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
