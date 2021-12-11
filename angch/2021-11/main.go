package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct{ X, Y int }

var adj = []Coord{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func day11(filepath string) {
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

	part1, part2 := 0, 0

	for step := 1; ; step++ {
		flashed := make(map[Coord]bool)
		for c, _ := range board {
			board[c]++
		}
		again := true

		for again {
			again = false

			for c, v := range board {
				if v > 9 {
					flashed[c] = true
					board[c] = 0
					if step <= 100 {
						part1++
					}
					for _, d := range adj {
						c2 := Coord{c.X + d.X, c.Y + d.Y}
						_, ok := board[c2]
						if !ok {
							continue
						}
						if !flashed[c2] {
							board[c2]++
							again = true
						}
					}
				}
			}
		}
		if len(flashed) == len(board) {
			part2 = step
			break
		}
		// fmt.Println("After step", step, part1)
		// for y := 0; y < 10; y++ {
		// 	for x := 0; x < 10; x++ {
		// 		c := Coord{x, y}
		// 		v, _ := board[c]
		// 		fmt.Print(v)
		// 	}
		// 	fmt.Println()
		// }
	}

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	day11("test.txt")
	day11("input.txt")
}
