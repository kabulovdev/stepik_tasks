package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Scanln(&a, &b)

	fmt.Println(math.Hypot(a, b))
}
