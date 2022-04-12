package main

import (
	"fmt"
)

func main() {

	for pos, char := range "chorizo" {
		defer fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

}
