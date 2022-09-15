package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var taym string
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var b []string = strings.Split(text, ",")
	//fmt.Println(b[0], b[1])

	firstTime, err := time.Parse("02.01.2006 15:04:05", b[0])
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 18; i++ {
		taym = taym + string(b[1][i])
	}
	firstTime2, err := time.Parse("02.01.2006 15:04:05", taym)
	if err != nil {
		panic(err)
	}

	res := firstTime.After(firstTime2)
	if res == true {
		future := firstTime
		fmt.Println(future.Sub(firstTime2))
	} else {
		future := firstTime2
		fmt.Println(future.Sub(firstTime))
	}

}
