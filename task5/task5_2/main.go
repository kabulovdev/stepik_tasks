package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scanln(&a)

	var text []string = strings.Split(a, "")
	fmt.Println(strings.Join(text, "*"))
}
