package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	s, _ := ioutil.ReadFile("input.txt")
	c := 0
	part2 := true
	for k, v := range s {
		if v == '(' {
			c++
		}
		if v == ')' {
			c--
		}
		if c == -1 && part2 {
			fmt.Println("Part 2", k+1)
			part2 = false
		}
	}
	fmt.Println("Part 1", c)
}
