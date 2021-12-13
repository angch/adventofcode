package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	X, Y int
}

func day13(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	maxX, maxY := 0, 0

	board := make(map[Coord]bool)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		x, y := 0, 0
		fmt.Sscanf(t, "%d,%d", &x, &y)
		board[Coord{x, y}] = true
		if y > maxY {
			maxY = y
		}
		if x > maxX {
			maxX = x
		}
	}

	part1 := 0
	for scanner.Scan() {
		instr := scanner.Text()
		instr = strings.ReplaceAll(instr, "fold along ", "")
		a := strings.Split(instr, "=")
		fold, _ := strconv.Atoi(a[1])
		xy := a[0]

		if xy == "y" {
			for c := range board {
				if c.Y >= fold {
					y := fold - (c.Y - fold)
					board[Coord{c.X, y}] = true
					delete(board, c)
				}
			}
		} else {
			for c := range board {
				if c.X >= fold {
					x := fold - (c.X - fold)
					board[Coord{x, c.Y}] = true
					delete(board, c)
				}
			}
		}
		if part1 == 0 {
			part1 = len(board)
		}
	}
	fmt.Println("Part 1", part1)

	maxX, maxY = 0, 0
	for c := range board {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if board[Coord{x, y}] {
				fmt.Print("##")
			} else {
				fmt.Print("..")
			}

		}
		fmt.Println()
	}
}

func main() {
	// day13("test.txt")

	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day13("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
