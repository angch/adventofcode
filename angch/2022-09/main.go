package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord [2]int

var m = map[string]coord{"R": {1, 0}, "U": {0, -1}, "L": {-1, 0}, "D": {0, 1}}

func day9(file string) (int, int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ropes := make([]coord, 10)
	visit := make(map[int]map[coord]bool)
	visit[1] = map[coord]bool{}
	visit[9] = map[coord]bool{}
	for scanner.Scan() {
		t := scanner.Text()
		dir, dist_ := t[:1], t[2:]
		dist, _ := strconv.Atoi(dist_)

		for d := 0; d < dist; d++ {
			prev := coord{0, 0}
			for k, rope := range ropes {
				if visit[k] != nil {
					visit[k][rope] = true
				}
				if k == 0 {
					ropes[k][0] = ropes[k][0] + m[dir][0]
					ropes[k][1] = ropes[k][1] + m[dir][1]
					prev = ropes[k]
				}

				dx, dy := prev[0]-rope[0], prev[1]-rope[1]
				if dx < -1 || dx > 1 || dy < -1 || dy > 1 {
					if dx > 0 {
						ropes[k][0]++
					} else if dx < 0 {
						ropes[k][0]--
					}
					if dy > 0 {
						ropes[k][1]++
					} else if dy < 0 {
						ropes[k][1]--
					}
					if visit[k] != nil {
						visit[k][ropes[k]] = true
					}
				}
				prev = ropes[k]
			}
		}
	}
	return len(visit[1]), len(visit[9])
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
