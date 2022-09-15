package main

import (
	"fmt"
)

func main() {
	value1, value2, operation := readTask()
	typeValue1 := fmt.Sprintf("%T", value1)
	typeValue2 := fmt.Sprintf("%T", value2)
	if typeValue1 == "float64" && typeValue2 == "float64" {
		num1, _ := value1.(float64)
		num2, _ := value2.(float64)
		switch operation {
		case "-":
			result := num1 - num2
			fmt.Printf("%.4f", float64(result))
		case "+":
			result := num1 + num2
			fmt.Printf("%.4f", float64(result))
		case "*":
			result := num1 * num2
			fmt.Printf("%.4f", float64(result))
		case "/":
			result := num1 / num2
			fmt.Printf("%.4f", float64(result))
		}

	} else if typeValue1 != "float64" || typeValue2 != "float64" {
		if typeValue1 != "float64" {
			typeValue := value1
			fmt.Printf("value=%v: %T", typeValue, typeValue)
			return
		}
		if typeValue2 != "float64" {
			typeValue := value2
			fmt.Printf("value=%v: %T", typeValue, typeValue)
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}
