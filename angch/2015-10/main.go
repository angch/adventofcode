package main

import (
	"fmt"
	"strconv"
)

func day10dumb(i string, part1, part2 int) (int, int) {
	rpart1 := 0
	for ; part2 > 0; part2-- {
		prev, count := -1, 0
		o := ""
		part1--
		if part1 == -1 {
			rpart1 = len(i)
		}
		for j := 0; j < len(i); j++ {
			if prev == -1 {
				prev = int(i[j]) - '0'
				count = 1
				continue
			}
			if int(i[j])-'0' == prev {
				count++
			} else {
				o = o + strconv.Itoa(count) + strconv.Itoa(prev)
				count = 1
				prev = int(i[j]) - '0'
			}
		}
		o = o + strconv.Itoa(count) + strconv.Itoa(prev)

		i = o
		fmt.Println(len(i), part2)
	}
	return rpart1, len(i)
}

func day10(in string, part1, part2 int) (int, int) {
	rpart1 := 0
	i := []byte(in)
	for ; part2 > 0; part2-- {
		prev, count := -1, 0
		o := make([]byte, 0, len(i)*2)
		part1--
		if part1 == -1 {
			rpart1 = len(i)
		}
		for j := 0; j < len(i); j++ {
			if prev == -1 {
				prev = int(i[j]) - '0'
				count = 1
				continue
			}
			if int(i[j])-'0' == prev {
				count++
			} else {
				o = append(o, []byte(strconv.Itoa(count))...)
				o = append(o, []byte(strconv.Itoa(prev))...)
				count = 1
				prev = int(i[j]) - '0'
			}
		}
		o = append(o, []byte(strconv.Itoa(count))...)
		o = append(o, []byte(strconv.Itoa(prev))...)

		i = o
		// fmt.Println(len(i), part2)
	}
	return rpart1, len(i)
}

func main() {
	// fmt.Println(day10("1", 4))
	fmt.Println(day10("1321131112", 40, 50))
}
