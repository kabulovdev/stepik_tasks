package main

import (
	"fmt"
)

func main() {
	var a string
	var b string

	fmt.Scanln(&a)

	for i := 1; i < len(a); i += 2 {
		b = b + string(a[i])
	}
	fmt.Println(b)
}
