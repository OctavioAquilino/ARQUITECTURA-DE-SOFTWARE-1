package main

import (
	"errors"
	"fmt"
)

func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No puedo dividir por 0")
	}
	return a / b, nil
}

func main() {
	var a int
	var b int
	a = 10
	b = 2
	div, _ := division(a, b)

	fmt.Println("Divsion:", div)
}
