package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord [2]int

var directions = []coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func day8(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := make(map[coord]int)
	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, v := range t {
			grid[coord{x, y}] = int(v) - '0'
		}
		y++
	}

	for c := range grid {
		score, donepart1 := 1, false
		for _, d := range directions {
			c1, h := c, grid[c]
			dist, isblocked := 0, false
			for {
				c1 = coord{c1[0] + d[0], c1[1] + d[1]}
				h2, ok := grid[c1]
				if !ok {
					break
				}
				dist++
				if h2 >= h {
					isblocked = true
					break
				}
			}
			score *= dist
			if !isblocked && !donepart1 {
				donepart1 = true
				part1++
			}
		}
		if score > part2 {
			part2 = score
		}
	}
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
