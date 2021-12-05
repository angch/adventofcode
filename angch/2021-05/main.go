package main

import (
	"bufio"
	"fmt"
	"os"
)

// var filepath = "input.txt"

var filepath = "test.txt"
var debug = false

type Coord struct {
	X, Y int
}
type Line struct {
	From Coord
	To   Coord
}

func part1() {
	file, _ := os.Open(filepath)
	// file, _ := os.Open("input.txt")
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
	_ = lines

	board := make(map[int]map[int]int)
	for _, l := range lines {
		// log.Println(l)
		if l.From.X == l.To.X {
			y1, y2 := l.From.Y, l.To.Y
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				_, ok := board[y]
				if !ok {
					board[y] = make(map[int]int)
				}
				board[y][l.From.X]++
			}
			// fmt.Println(board)
		} else if l.From.Y == l.To.Y {
			x1, x2 := l.From.X, l.To.X
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				_, ok := board[l.From.Y]
				if !ok {
					board[l.From.Y] = make(map[int]int)
				}
				board[l.From.Y][x]++
			}
			// fmt.Println(board)
		} else {
			count := l.From.X - l.To.X
			if count < 0 {
				count = -count
			}
			x1, x2 := l.From.X, l.To.X
			dx := 1
			if x2 < x1 {
				dx = -1
			}
			y1, y2 := l.From.Y, l.To.Y
			dy := 1
			if y2 < y1 {
				dy = -1
			}
			for x, y := x1, y1; count >= 0; x, y, count = x+dx, y+dy, count-1 {
				_, ok := board[y]
				if !ok {
					board[y] = make(map[int]int)
				}
				board[y][x]++
			}
		}
	}
	count := 0
	for _, v := range board {
		for _, v2 := range v {
			if v2 > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}

}

func main() {
	part1()
	// part2()
}
