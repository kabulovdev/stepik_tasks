package main

import "fmt"

// считайте что fmt уже импортирован и main объявлен
func test(x1 *int, x2 *int) {
	var a int = *x1 * *x2
	fmt.Println(a)
	// здесь пишите ваш код
}

func main() {

	var a int
	var b int
	fmt.Scanln(&a, &b)
	test(&a, &b)

}
