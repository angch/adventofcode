package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/angch/adventofcode/angch/vector"
)

func flood(board map[vector.Point[int]]int, c vector.Point[int]) int {
	basin := make(map[vector.Point[int]]bool)
	basin[c] = true
	eval := []vector.Point[int]{c}

	for len(eval) > 0 {
		var c1 vector.Point[int]
		c1, eval = eval[len(eval)-1], eval[:len(eval)-1]
		for _, d := range vector.CompassDirectionsInt {
			c2 := c1.Add(d)
			v2, ok := board[c2]
			if !ok {
				continue
			}
			v1 := board[c1]
			if v2 > v1 && v2 < 9 {
				if !basin[c2] {
					basin[c2] = true
					eval = append(eval, c2)
				}
			}
		}
	}
	return len(basin)
}

func day9(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make(map[vector.Point[int]]int)
	c := vector.Point[int]{0, 0}
	for scanner.Scan() {
		t := scanner.Text()
		for _, v := range t {
			board[c] = int(v - '0')
			c[0]++
		}
		c[0], c[1] = 0, c[1]+1
	}

	areas := make([]int, 0)
	part1 := 0

next:
	for c, v := range board {
		for _, d := range vector.CompassDirectionsInt {
			c2 := c.Add(d)
			v2, ok := board[c2]
			if !ok {
				continue
			}
			if v > v2 {
				continue next
			}
		}
		area := flood(board, c)
		areas = append(areas, area)
		part1 += v + 1
	}

	fmt.Println("Part 1", part1)
	sort.Ints(areas)
	part2 := 1
	for k := 0; k < 3; k++ {
		part2 *= areas[len(areas)-1-k]
	}
	fmt.Println("Part 2", part2)
}

func main() {
	day9("../test.txt")
	day9("../input.txt")
}
