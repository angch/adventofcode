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

	board := make(map[Coord]bool)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		x, y := 0, 0
		fmt.Sscanf(t, "%d,%d", &x, &y)
		board[Coord{x, y}] = true
	}

	part1 := 0
	maxX, maxY := 0, 0
	for scanner.Scan() {
		a := strings.Split(strings.TrimPrefix(scanner.Text(), "fold along "), "=")
		fold, _ := strconv.Atoi(a[1])
		fold2 := fold * 2

		if a[0] == "y" {
			maxY = 0
			for c := range board {
				if c.Y >= fold {
					delete(board, c)
					c.Y = fold2 - c.Y
					board[c] = true
				}
				if c.Y > maxY {
					maxY = c.Y
				}
			}
		} else {
			maxX = 0
			for c := range board {
				if c.X >= fold {
					delete(board, c)
					c.X = fold2 - c.X
					board[c] = true
				}
				if c.X > maxX {
					maxX = c.X
				}
			}
		}
		if part1 == 0 {
			part1 = len(board)
		}
	}
	fmt.Println("Part 1", part1)

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
