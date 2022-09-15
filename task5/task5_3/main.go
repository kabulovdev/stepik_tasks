package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var max string = "0"
	fmt.Scanln(&a)
	var b []string = strings.Split(a, "")
	for i := 0; i < len(b); i++ {

		if b[i] > max {
			max = b[i]
		}
	}
	fmt.Println(max)
}
