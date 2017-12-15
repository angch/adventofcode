package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	s, _ := ioutil.ReadFile("input.txt")
	c := 0
	for k, v := range s {
		if v == '(' {
			c++
		}
		if v == ')' {
			c--
		}
		if c == -1 {
			log.Println(k + 1)
			return
		}
	}
	fmt.Println(c)
}
