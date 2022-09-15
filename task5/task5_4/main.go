package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	var x int
	var sum string
	fmt.Scanln(&a)
	var b []string = strings.Split(a, "")
	for i := 0; i < len(b); i++ {

		n, err := strconv.ParseInt(b[i], 10, 64)
		if err == nil {

			x = int(n)

			s2 := strconv.Itoa(x * x)
			sum = sum + (s2)
		}
	}
	fmt.Println(sum)

}
