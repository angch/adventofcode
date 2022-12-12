package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord [2]int

var directions []coord = []coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func solve(grid map[coord]int, start coord) (int, int) {
	dist := map[coord]int{}
	queue := []coord{start}

	count := 0
	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]
		count++

		h := grid[this]
		for _, d := range directions {
			neighbour := coord{this[0] + d[0], this[1] + d[1]}
			h2, ok := grid[neighbour]
			if !ok {
				continue
			}
			if h2 <= h+1 {
				_, ok2 := dist[neighbour]
				if !ok2 {
					queue = append(queue, neighbour)
					dist[neighbour] = dist[this] + 1
				}
			}
			if grid[neighbour] == 'E' && 'z'-h <= 1 {
				return dist[this] + 1, count
			}
		}
	}
	return 999999, count
}

func day12(file string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := map[coord]int{}

	scanner := bufio.NewScanner(f)
	y := 0
	var start coord
	a := []coord{}
	for scanner.Scan() {
		t := scanner.Text()
		for x, c := range t {
			if c == 'S' {
				start = coord{x, y}
				c = 'a'
			}
			grid[coord{x, y}] = int(c)
			if c == 'a' {
				a = append(a, coord{x, y})
			}
		}
		y++
	}
	part1, count1 := solve(grid, start)
	part2 := part1
	for _, c := range a {
		p2, count2 := solve(grid, c)
		count1 += count2
		if p2 < part2 {
			part2 = p2
		}
	}
	fmt.Println("Counted", count1) // to be used for comparing against dijsktra later
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day12("test.txt")
	fmt.Println(part1, part2)
	if part1 != 31 || part2 != 29 {
		log.Fatal("Bad test")
	}
	fmt.Println(day12("input.txt"))
}
