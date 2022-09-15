package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	var cheak int
	fmt.Scanln(&a)
	rs := []rune(a)

	if 5 <= len(a) {
		for i := 0; i < len(rs); i++ {
			if unicode.Is(unicode.Latin, rs[i]) || unicode.IsDigit(rs[i]) {
				//			fmt.Println(string(rs[i]))
				cheak++
			}
		}

		//fmt.Println(cheak, len(a)-)
		if cheak == len(a) {
			fmt.Println("Ok")
		} else {
			fmt.Println("Wrong password")
		}
	} else {
		fmt.Println("Wrong password")
	}
}
