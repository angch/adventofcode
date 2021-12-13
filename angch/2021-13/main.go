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
		xy := ""
		cord := 0
		c := Coord{0, 0}
		instr = strings.ReplaceAll(instr, "fold along ", "")
		a := strings.Split(instr, "=")
		cord, _ = strconv.Atoi(a[1])
		xy = a[0]
		if xy == "y" {
			c.Y = cord
		} else {
			c.X = cord
		}

		board2 := board
		for k, _ := range board2 {
			if xy == "y" {
				if k.Y >= cord {
					y := cord - (k.Y - cord)
					board[Coord{k.X, y}] = true
					delete(board, k)
				}
			} else {
				if k.X >= cord {
					x := cord - (k.X - cord)
					board[Coord{x, k.Y}] = true
					delete(board, k)
				}
			}
		}

		if part1 == 0 {
			for y := 0; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if board[Coord{x, y}] {
						part1++
					}
				}
			}
		}
	}
	fmt.Println("Part 1", part1)

	minX, minY := maxX, maxY
	maxX, maxY = 0, 0
	for c := range board {
		if c.X < minX {
			minX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.X > maxX {
			maxX = c.X
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if board[Coord{x, y}] {
				fmt.Print("##")
			} else {
				fmt.Print("..")
			}

		}
		fmt.Println()
	}
	// fmt.Println(board)

}

func main() {
	// day13("test.txt")

	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day13("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
