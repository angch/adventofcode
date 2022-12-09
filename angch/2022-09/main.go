package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord [2]int

var m = map[string]coord{
	"R": {1, 0},
	"U": {0, -1},
	"L": {-1, 0},
	"D": {0, 1},
}

func sim(dir string, dist int, ropes []coord, visit map[coord]bool) {
	visit[coord{0, 0}] = true
	for d := 0; d < dist; d++ {
		prev := coord{0, 0}
		for k, rope := range ropes {
			if k == 0 {
				ropes[k][0] = ropes[k][0] + m[dir][0]
				ropes[k][1] = ropes[k][1] + m[dir][1]
				prev = ropes[k]
			}

			dx, dy := prev[0]-rope[0], prev[1]-rope[1]
			adx := dx
			if adx < 0 {
				adx = -adx
			}
			ady := dy
			if ady < 0 {
				ady = -ady
			}
			if adx > 1 || ady > 1 || adx+ady > 2 {
				if dx > 0 {
					dx = 1
				} else if dx < 0 {
					dx = -1
				}
				if dy > 0 {
					dy = 1
				} else if dy < 0 {
					dy = -1
				}
				ropes[k][0] = ropes[k][0] + dx
				ropes[k][1] = ropes[k][1] + dy
				if k == len(ropes)-1 {
					visit[ropes[k]] = true
				}
			}
			prev = ropes[k]
		}
	}
}
func day9(file string) (int, int) {
	part1, part2 := 0, 0

	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ropes1 := make([]coord, 2)
	visit1 := map[coord]bool{}
	ropes2 := make([]coord, 10)
	visit2 := map[coord]bool{}
	for scanner.Scan() {
		t := scanner.Text()
		dir, dist_ := t[:1], t[2:]
		dist, _ := strconv.Atoi(dist_)
		sim(dir, dist, ropes1, visit1)
		sim(dir, dist, ropes2, visit2)
	}
	part1 = len(visit1)
	part2 = len(visit2)

	return part1, part2
}

func main() {
	part1, part2 := day9("test.txt")
	fmt.Println(part1, part2)
	if part1 != 13 || part2 != 1 {
		log.Fatal("Test failed")
	}
	part1, part2 = day9("test2.txt")
	fmt.Println(part1, part2)
	if part1 != 88 || part2 != 36 {
		log.Fatal("Test failed")
	}
	fmt.Println(day9("input.txt"))
}
