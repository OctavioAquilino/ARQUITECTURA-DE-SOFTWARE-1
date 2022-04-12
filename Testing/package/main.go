package main

import (
	"fmt"
	"ucc/package_a"
	"ucc/package_b"
)

func main() {

	fmt.Println(package_a.Hello())
	fmt.Println(package_b.Hello())
}
