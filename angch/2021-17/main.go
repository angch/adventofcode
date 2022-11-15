package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Coord struct {
	x, y int
}
type Area struct {
	c1, c2 Coord
}

func do(v Coord, target Area) int {
	c := Coord{0, 0}
	maxY := 0
	// seen := make(map[Coord]bool)
	for {
		c.x += v.x
		c.y += v.y
		if c.y > maxY {
			maxY = c.y
		}
		if v.x > 0 {
			v.x -= 1
		} else if v.x < 0 {
			v.x += 1
		}
		v.y -= 1
		if c.y >= target.c1.y && c.y <= target.c2.y && c.x >= target.c1.x && c.x <= target.c2.x {
			return maxY
		}
		if c.y < target.c1.y {
			return -1
		}
	}
}

func day17(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	t := scanner.Text()
	a := strings.Split(t, ",")
	b := strings.Split(a[0], "=")
	c := strings.Split(a[1], "=")

	x1, x2 := 0, 0
	y1, y2 := 0, 0
	fmt.Sscanf(b[1], "%d..%d", &x1, &x2)
	fmt.Sscanf(c[1], "%d..%d", &y1, &y2)

	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	target := Area{Coord{x1, y1}, Coord{x2, y2}}
	best := target.c1.y
	bestV := Coord{0, 0}
	count := 0
	mostY := 0
	repeat := make(map[Coord]bool)
	for x1 := 0; x1 < target.c2.x*2; x1++ {
		for y1 := target.c1.y; y1 < -target.c2.y*2; y1++ {
			v := Coord{x1, y1}
			b := do(v, target)
			if b >= 0 {
				count++
				repeat[v] = true
				if best < b {
					best, bestV = b, v
				}
				if y1 > mostY {
					mostY = y1
				}
			}
		}
	}

	part1, part2, _, _ := best, count, bestV, mostY
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2, len(repeat))

	if false {
		fmt.Println("x,y")
		for c := range repeat {
			fmt.Println(c.x, ",", c.y)
		}
		// fmt.Println("Part 2", mostY)
	}
}

func main() {
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day17("input.txt")
	// day17("test.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
