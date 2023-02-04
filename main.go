package main

import (
	"fmt"

	"github.com/Erickype/GoMicroservices/geometry"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world!!")
	fmt.Println(quote.Go())

	//At the same time
	var daysMonth = map[string]int{"Jan": 1, "Feb":2}
	fmt.Println(daysMonth)

	//uso of local package
	var length, width float64 = 3, 4
	res := geometry.Area(length, width)

	fmt.Printf("Area: %.2f\n", res)

	res = geometry.Diagonal(length, width)

	fmt.Printf("Diagonal: %.2f\n", res)
}