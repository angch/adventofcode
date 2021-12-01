package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	s, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(s), "\r\n")
	prev := 99999
	inc := 0
	inc2 := 0
	win1 := make([]int, 0)
	// win2 := make([]int, 0)
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		win1 = append(win1, n)
		if len(win1) > 3 {
			a := win1[len(win1)-4]
			// b := n
			if n > a {
				inc2++
			}
		}
		// win2 = win1
		// fmt.Println(n, line, win1)
		if n > prev {
			inc++
		}
		prev = n

	}
	fmt.Println(inc, inc2)
}
