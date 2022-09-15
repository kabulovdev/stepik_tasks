package main

import "fmt"

func test(x1 *int, x2 *int) {
	var a int = *x1
	*x1 = *x2
	*x2 = a
	fmt.Println(*x1, *x2)
	// здесь ваш код
}

func main() {

	var a int
	var b int

	fmt.Scanln(&a, &b)

	test(&a, &b)

}
