package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v", ip) //@todo
}

func main() {
	addrs := map[string]IPAddr{
		"lookback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs{
		fmt.Printf("%v, %v\n", n, a)
	}
}
