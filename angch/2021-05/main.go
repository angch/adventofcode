package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	X, Y int
}
type Line struct {
	From Coord
	To   Coord
}
type Counts struct {
	Part1 int
	Part2 int
}

func day5(filepath string) {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]Line, 0)

	for scanner.Scan() {
		t := scanner.Text()
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Sscanf(t, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		line := Line{
			From: Coord{X: x1, Y: y1},
			To:   Coord{X: x2, Y: y2},
		}
		lines = append(lines, line)
	}

	board := make(map[Coord]Counts)

	for _, l := range lines {
		x, y := l.From.X, l.From.Y
		dx, dy, length := 0, 0, 0

		if l.From.X > l.To.X {
			dx = -1
			length = l.From.X - l.To.X
		} else if l.From.X < l.To.X {
			dx = 1
			length = l.To.X - l.From.X
		}

		if l.From.Y > l.To.Y {
			dy = -1
			length = l.From.Y - l.To.Y
		} else if l.From.Y < l.To.Y {
			dy = 1
			length = l.To.Y - l.From.Y
		}

		for ; length >= 0; x, y, length = x+dx, y+dy, length-1 {
			c := board[Coord{x, y}]
			if dx == 0 || dy == 0 {
				c.Part1++
			}
			c.Part2++
			board[Coord{x, y}] = c
		}
	}
	part1, part2 := 0, 0
	for _, v := range board {
		if v.Part1 > 1 {
			part1++
		}
		if v.Part2 > 1 {
			part2++
		}
	}
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	day5("test.txt")
	day5("input.txt")
}
