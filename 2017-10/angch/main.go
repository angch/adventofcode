package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	rope := make([]int, 256)
	for i := range rope {
		rope[i] = i
	}

	cur, skip := 0, 0

	inputs := "157,222,1,2,177,254,0,228,159,140,249,187,255,51,76,30"
	//inputs := "3,4,1,5"
	//inputs = "1,2,3"
	//inputs = "1,2,4"
	//inputs = ""
	//inputs = "AoC 2017"

	lengths := make([]int, len(inputs))
	for i, c := range inputs {
		lengths[i] = int(c)
	}
	for _, c := range []int{17, 31, 73, 47, 23} {
		lengths = append(lengths, c)
		//fmt.Println(rope[0] * rope[1])
	}

	if false {
		// Part1
		lengths1 := strings.Split(inputs, ",")
		lengths = make([]int, len(lengths1))
		for i, s := range lengths1 {
			lengths[i], _ = strconv.Atoi(s)
		}
		rope, cur, skip = knot(rope, lengths, cur, skip, 1)
		fmt.Println(rope[0] * rope[1])
		return
	}

	rope, cur, skip = knot(rope, lengths, cur, skip, 64)

	hash := make([]int, 16)
	for i, r := 0, 0; i < len(rope); r++ {
		c := 0
		for r := 0; r < 16; r++ {
			c ^= rope[i]
			i++
		}
		hash[r] = c
	}

	for hex := 0; hex < 16; hex++ {
		fmt.Printf("%02x", hash[hex])
	}
}

func knot(rope []int, lengths []int, cur, skip int, rounds int) ([]int, int, int) {
	for r := 0; r < rounds; r++ {
		for _, length := range lengths {
			j := length - 1
			for i := 0; i <= j; i++ {
				rope[(cur+i)%len(rope)], rope[(cur+j)%len(rope)] = rope[(cur+j)%len(rope)], rope[(cur+i)%len(rope)]
				j--
			}
			cur = (cur + length + skip) % len(rope)
			skip++
		}
	}
	return rope, cur, skip
}
