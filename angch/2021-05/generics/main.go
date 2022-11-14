package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/angch/adventofcode/angch/vector"
)

type Counts struct {
	Part1 int
	Part2 int
}

func day5(filepath string) (int, int) {
	file, err := os.Open(filepath)
	if err != nil {
		return -1, -1
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]vector.Line[int], 0)

	for scanner.Scan() {
		t := scanner.Text()
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Sscanf(t, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		line := vector.NewLine(x1, y1, x2, y2)
		lines = append(lines, line)
	}

	board := make(map[vector.Point[int]]Counts)

	part1, part2 := 0, 0
	for _, l := range lines {
		points := l.Points()
		ispart1 := l.IsRightAngled()
		for _, p := range points {
			c := board[p]
			if ispart1 {
				c.Part1++
				if c.Part1 == 2 {
					part1++
				}
			}
			c.Part2++
			if c.Part2 == 2 {
				part2++
			}
			board[p] = c
		}
	}
	return part1, part2
}

func main() {
	p1, p2 := day5("../test.txt")
	fmt.Println("Test", p1, p2)
	p1, p2 = day5("../input.txt")
	fmt.Println("Input", p1, p2)
}
