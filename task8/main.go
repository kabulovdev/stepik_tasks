package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	var b []string = strings.Split(buf, " ")

	var a []string = strings.Split(b[1], ":")

	vaqt, err := strconv.Atoi(a[0])

	if err != nil {
		fmt.Println(err)
	}
	if vaqt < 13 {
		fmt.Println(buf)
	} else {
		var vaqt1 []string = strings.Split(b[0], "-")

		mt, errr := strconv.Atoi(vaqt1[2])
		if errr != nil {
			fmt.Println(errr)
		}
		mt++
		var blast string
		i := 0
		for {
			if i > 7 {
				break
			} else {
				blast = blast + string(b[1][i])
			}
			i++
		}

		lasttime := vaqt1[0] + "-" + vaqt1[1] + "-" + strconv.Itoa(mt) + " " + blast
		fmt.Println(lasttime)
	}

}
