package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world!!")
	fmt.Println(quote.Go())

	//At the same time
	var daysMonth = map[string]int{"Jan": 1, "Feb":2}
	fmt.Println(daysMonth)
}