package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var b []string = strings.Split(text, ";")
	var sum string
	var jav []float64

	for i := 0; i < len(b); i++ {
		sum = ""
		for j := 0; j < len(b[i]); j++ {
			if string(b[i][j]) != " " && string(b[i][j]) != "," && string(b[i][j]) != "\n" {
				sum = sum + string(b[i][j])
			}
			if string(b[i][j]) == "," {
				sum = sum + "."
			}

		}
		num21, _ := strconv.ParseFloat(sum, 64)
		jav = append(jav, num21)
	}

	fmt.Printf("%.4f\n", jav[0]/jav[1])
}
