package main

import (
	"fmt"
)

func main() {
	var input = 10
	for i := 1; i <= input; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%2 == 0 && i%3 == 0 {
		fmt.Println(i, fizz+buzz)
	} else if i%2 == 0 {
		fmt.Println(i, fizz)
	} else if i%3 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}
