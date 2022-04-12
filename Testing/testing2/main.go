package main

import (
	"fmt"
)

func Calculate(num int) int {
	return num + 2
}

func main() {
	fmt.Println("Suma de 10:", Calculate(10))
}
