package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	var rev []rune
	for i := utf8.RuneCountInString(s) - 1; i >= 0; i-- {
		rev = append(rev, r[i])
	}
	if string(rev) == s {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
