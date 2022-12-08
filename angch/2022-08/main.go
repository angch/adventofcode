package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var verbose = false

type coord [2]int

func day8(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// curdir, cmd := "", ""
	grid := make(map[coord]int)
	y := 0
	for scanner.Scan() {
		t := scanner.Text()

		for x, v := range t {
			c := coord{x, y}
			grid[c] = int(v) - '0'
		}
		y++

		if verbose {

		}
	}
	if verbose {
		fmt.Println(grid)
	}

	bestscore := 0
	for y1 := 0; y1 < y; y1++ {
		for x1 := 0; x1 < y; x1++ {
			c := coord{x1, y1}
			h := grid[c]
			score := 1

			dist := 0
			no := false
			for x2 := x1 - 1; x2 >= 0; x2-- {
				c2 := coord{x2, y1}
				dist++
				if grid[c2] >= h {
					no = true
					break
				}
			}
			// log.Println("dist is", dist)
			score *= dist
			// log.Println("dist is", dist, score)

			done1 := false
			if !no {
				part1++
				// log.Println("found", c, h, "left")
				done1 = true
			}

			no = false
			dist = 0
			for x2 := x1 + 1; x2 < y; x2++ {
				c2 := coord{x2, y1}
				dist++
				if grid[c2] >= h {
					no = true
					break
				}
			}
			score *= dist
			// log.Println("dist is", dist, score)
			if !no && !done1 {
				part1++
				// log.Println("found", c, h, "right")
				done1 = true
			}
			no = false
			dist = 0
			for y2 := y1 - 1; y2 >= 0; y2-- {
				c2 := coord{x1, y2}
				dist++
				if grid[c2] >= h {
					no = true
					break
				}
				// dist++
			}
			score *= dist
			// log.Println("dist is", dist, score)
			if !no && !done1 {
				part1++
				done1 = true
				// continue
			}

			no = false
			dist = 0
			for y2 := y1 + 1; y2 < y; y2++ {
				c2 := coord{x1, y2}
				dist++
				if grid[c2] >= h {
					no = true
					break
				}
			}
			// log.Println("dist is", dist, score)
			score *= dist
			if !no && !done1 {
				part1++
				// continue
				done1 = true
			}

			// log.Println("score is", score, x1, y1)
			if score > bestscore {
				bestscore = score
			}
		}
	}
	part2 = bestscore

	return part1, part2
}

func main() {
	part1, part2 := day8("test.txt")
	fmt.Println(part1, part2)
	if part1 != 21 || part2 != 8 {
		log.Fatal("Test failed")
	}
	fmt.Println(day8("input.txt"))
}
