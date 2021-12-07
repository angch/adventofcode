package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var f_memo = []int{0}

func f(a int) int {
	d := 0
	if len(f_memo) > a {
		return f_memo[a]
	}
	for i, d := len(f_memo), f_memo[len(f_memo)-1]; a > 0; a-- {
		d += i
		f_memo = append(f_memo, d)
		i++
	}
	return d
}

func day7(filepath string) {
	t, _ := ioutil.ReadFile(filepath)
	pos := make([]int, 0)
	x := strings.Split(string(t), ",")
	for _, v := range x {
		a, _ := strconv.Atoi(v)
		pos = append(pos, a)
	}

	least1, least2 := math.MaxInt, math.MaxInt
	best1, best2 := -1, -1
	for target := 0; target < len(pos); target++ {
		fuel1, fuel2 := 0, 0
		for _, v := range pos {
			d := v - target
			if d < 0 {
				d = -d
			}
			fuel1 += d
			fuel2 += f(d)
		}
		if fuel1 < least1 {
			least1, best1 = fuel1, target
		}
		if fuel2 < least2 {
			least2, best2 = fuel2, target
		}
	}
	fmt.Println("Part 1", least1, best1)
	fmt.Println("Part 2", least2, best2)
}

func main() {
	// day7("test.txt")
	day7("input.txt")
}
