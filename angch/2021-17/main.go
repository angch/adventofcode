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
	step := 0
	c := Coord{0, 0}
	maxY := 0
	// seen := make(map[Coord]bool)
	for {
		step++

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

		// fmt.Println("Step", step, c)

		if c.y >= target.c1.y && c.y <= target.c2.y && c.x >= target.c1.x && c.x <= target.c2.x {
			// fmt.Println("S", step, maxY)
			return maxY
		}
		// _, ok := seen[c]
		// if ok {
		// 	return -1
		// }
		// seen[c] = true

		if c.y < target.c1.y {
			break
		}
	}
	return -1
}

func day17(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// bs := BitStream{}
	scanner.Scan()
	t := scanner.Text()
	//target area: x=20..30, y=-10..-5
	a := strings.Split(t, ",")
	// fmt.Println(t)
	b := strings.Split(a[0], "=")
	c := strings.Split(a[1], "=")

	x1, x2 := 0, 0
	y1, y2 := 0, 0
	fmt.Sscanf(b[1], "%d..%d", &x1, &x2)
	fmt.Sscanf(c[1], "%d..%d", &y1, &y2)
	// fmt.Println(x1, x2, y1, y2)

	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	// v := Coord{7, 2}
	target := Area{Coord{x1, y1}, Coord{x2, y2}}
	// fmt.Println("done", do(v, target))
	// v = Coord{6, 3}
	// fmt.Println("done", do(v, target))
	// v = Coord{9, 0}
	// fmt.Println("done", do(v, target))
	// v = Coord{17, -4}
	// fmt.Println("done", do(v, target))
	// test := []Coord{{23, -10}, {25, -9}, {27, -5}, {29, -6}, {22, -6}, {21, -7}, {9, 0}, {27, -7}, {24, -5},
	// 	{25, -7}, {26, -6}, {25, -5}, {6, 8}, {11, -2}, {20, -5}, {29, -10}, {6, 3}, {28, -7},
	// 	{8, 0}, {30, -6}, {29, -8}, {20, -10}, {6, 7}, {6, 4}, {6, 1}, {14, -4}, {21, -6},
	// 	{26, -10}, {7, -1}, {7, 7}, {8, -1}, {21, -9}, {6, 2}, {20, -7}, {30, -10}, {14, -3},
	// 	{20, -8}, {13, -2}, {7, 3}, {28, -8}, {29, -9}, {15, -3}, {22, -5}, {26, -8}, {25, -8},
	// 	{25, -6}, {15, -4}, {9, -2}, {15, -2}, {12, -2}, {28, -9}, {12, -3}, {24, -6}, {23, -7},
	// 	{25, -10}, {7, 8}, {11, -3}, {26, -7}, {7, 1}, {23, -9}, {6, 0}, {22, -10}, {27, -6},
	// 	{8, 1}, {22, -8}, {13, -4}, {7, 6}, {28, -6}, {11, -4}, {12, -4}, {26, -9}, {7, 4},
	// 	{24, -10}, {23, -8}, {30, -8}, {7, 0}, {9, -1}, {10, -1}, {26, -5}, {22, -9}, {6, 5},
	// 	{7, 5}, {23, -6}, {28, -10}, {10, -2}, {11, -1}, {20, -9}, {14, -2}, {29, -7}, {13, -3},
	// 	{23, -5}, {24, -8}, {27, -9}, {30, -7}, {28, -5}, {21, -10}, {7, 9}, {6, 6}, {21, -5},
	// 	{27, -10}, {7, 2}, {30, -9}, {21, -8}, {22, -7}, {24, -9}, {20, -6}, {6, 9}, {29, -5},
	// 	{8, -2}, {27, -8}, {30, -5}, {24, -7},
	// }
	// for _, v := range test {
	// 	out := do(v, target)
	// 	if out == -1 {
	// 		fmt.Println("wrong", v)
	// 	}
	// }
	// fmt.Println("c", len(test))
	// return

	best := -1000
	bestN := Coord{0, 0}
	count := 0
	for x1 := 0; x1 < target.c2.x*2; x1++ {
		for y1 := target.c1.y; y1 < 1000; y1++ {
			v := Coord{x1, y1}
			b := do(v, target)
			if b >= 0 {
				count++
				if best < b {
					best = b
					bestN = Coord{x1, y1}
				}
			}
		}
	}
	// fmt.Println("Best", best, bestN, count)

	part1, part2 := best, count
	// bs.shift = 0
	// bs.buffer = []byte(t)
	// }
	// p, versions, _ := bs.ReadPacket()
	// part1, part2 := versions, p.Num
	_ = bestN
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day14("test.txt")
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day17("input.txt")
	// day17("test.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
