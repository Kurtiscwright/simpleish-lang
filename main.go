package main

import (
	"dev/simpleish-lang/token"
	"fmt"
)

func main() {
	result := token.Addition(5, 5)
	fmt.Println("Hello, World!")
	fmt.Printf("5+5 = %d \n", result)
}
