package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Coord struct {
	X, Y int
}

var adj = []Coord{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func flood(board [][]int, x, y int) int {
	basin := make(map[Coord]bool)
	basin[Coord{x, y}] = true

	eval := make([]Coord, 0)
	eval = append(eval, Coord{x, y})

	for len(eval) > 0 {
		var k Coord
		k, eval = eval[len(eval)-1], eval[:len(eval)-1]
		for c2 := range adj {
			dx, dy := adj[c2].X, adj[c2].Y
			c := board[k.Y][k.X]
			if k.X+dx < 0 || k.X+dx >= len(board[k.Y]) {
				continue
			}
			if k.Y+dy < 0 || k.Y+dy >= len(board) {
				continue
			}
			v := board[k.Y+dy][k.X+dx]
			if v > c && v != 9 {
				coord := Coord{k.X + dx, k.Y + dy}
				if !basin[coord] {
					basin[coord] = true
					eval = append(eval, coord)
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
			n, _ := strconv.Atoi(string(v))
			line = append(line, n)
		}
		board = append(board, line)
	}
	// fmt.Println(board)

	low := make([]int, 0)
	areas := make([]int, 0)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			v := board[y][x]
			bad := false
			for _, d := range adj {
				if x+d.X < 0 || x+d.X >= len(board[y]) {
					continue
				}
				if y+d.Y < 0 || y+d.Y >= len(board) {
					continue
				}
				if board[y+d.Y][x+d.X] <= v {
					bad = true
					break
				}
			}
			if !bad {
				area := flood(board, x, y)
				areas = append(areas, area)

				low = append(low, v)
			}
		}
	}
	part1 := 0
	for _, v := range low {
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
