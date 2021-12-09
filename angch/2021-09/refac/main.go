package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coord struct{ X, Y int }

var adj = []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func flood(board map[Coord]int, c Coord) int {
	basin := make(map[Coord]bool)
	basin[c] = true
	eval := []Coord{c}

	for len(eval) > 0 {
		var c1 Coord
		c1, eval = eval[len(eval)-1], eval[:len(eval)-1]
		for _, d := range adj {
			c2 := Coord{c1.X + d.X, c1.Y + d.Y}
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

	board := make(map[Coord]int)
	c := Coord{0, 0}
	for scanner.Scan() {
		t := scanner.Text()
		for _, v := range t {
			board[c] = int(v - '0')
			c.X++
		}
		c = Coord{0, c.Y + 1}
	}
	// fmt.Println(board)

	areas := make([]int, 0)
	part1 := 0

next:
	for c, v := range board {
		for _, d := range adj {
			c2 := Coord{c.X + d.X, c.Y + d.Y}
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
	day9("test.txt")
	day9("input.txt")
}
