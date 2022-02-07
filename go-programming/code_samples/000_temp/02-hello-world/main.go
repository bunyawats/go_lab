package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello")
	fmt.Println(hello())
	foo()
}

func hello() string {
	return quote.Hello()
}

func foo() {
	fmt.Println("I am fool")
}
