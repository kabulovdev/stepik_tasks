package main

import "fmt"

func main() {
	var a int
	var b int
	_, err := fmt.Scanln(&a, &b)
	if err != nil && a == 0 || b == 0 {
		fmt.Println("ошибка")
	} else {
		fmt.Println(a / b)
	}
	// package main уже объявлен, нужные импорты уже есть
}
