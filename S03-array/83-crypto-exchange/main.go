package main

import "fmt"

func main() {
	const (
		ETH = 9 - iota
		WAN
		ICX
	)
	rates := [...]float64{
		ETH: 25.5,  //ethereum
		WAN: 120.5, //wanchain
		ICX: 20,
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
