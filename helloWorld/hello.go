package main

import "fmt"

import "rsc.io/quote/v4"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
