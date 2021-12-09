package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coord struct{ X, Y int }

var adj = []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func flood(board [][]int, c Coord) int {
	basin := make(map[Coord]bool)
	basin[c] = true
	eval := []Coord{c}

	for len(eval) > 0 {
		var c1 Coord
		c1, eval = eval[len(eval)-1], eval[:len(eval)-1]
		for _, d := range adj {
			c2 := Coord{c1.X + d.X, c1.Y + d.Y}
			if c2.X < 0 || c2.X >= len(board[c1.Y]) {
				continue
			}
			if c2.Y < 0 || c2.Y >= len(board) {
				continue
			}
			v1 := board[c1.Y][c1.X]
			v2 := board[c2.Y][c2.X]
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

	board := [][]int{}
	for scanner.Scan() {
		t := scanner.Text()
		line := make([]int, 0)
		for _, v := range t {
			line = append(line, int(v-'0'))
		}
		board = append(board, line)
	}
	// fmt.Println(board)

	areas := make([]int, 0)
	part1 := 0

	for y := 0; y < len(board); y++ {
	next:
		for x := 0; x < len(board[y]); x++ {
			v := board[y][x]
			for _, d := range adj {
				c2 := Coord{x + d.X, y + d.Y}
				if c2.X < 0 || c2.X >= len(board[y]) {
					continue
				}
				if c2.Y < 0 || c2.Y >= len(board) {
					continue
				}
				if v > board[c2.Y][c2.X] {
					continue next
				}
			}
			area := flood(board, Coord{x, y})
			areas = append(areas, area)
			part1 += v + 1
		}
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
