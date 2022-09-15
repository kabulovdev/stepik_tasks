package main

import (
	"fmt"
	"strconv"
)

func adding(a string, b string) int64 {
	var sum1 string
	var sum2 string
	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			sum1 = sum1 + string(a[i])
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			sum2 = sum2 + string(b[i])
		}
	}

	tes2, _ := strconv.Atoi(sum1)
	tes1, _ := strconv.Atoi(sum2)
	tes := tes2 + tes1
	return int64(tes)
}

func main() {

	var word1 string
	var word2 string

	fmt.Scanln(word1, word2)

	result := adding(word1, word2)
	fmt.Println(result)

}
