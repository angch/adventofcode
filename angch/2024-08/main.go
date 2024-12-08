package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func day8(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ants := map[byte][][2]int{}
	antinodes := map[[2]int]bool{}
	antinodes2 := map[[2]int]bool{}
	y := 0
	maxX := 0
	for scanner.Scan() {
		t := scanner.Text()
		maxX = len(t)
		for x, v := range t {
			if v != '.' {
				if _, ok := ants[byte(v)]; !ok {
					ants[byte(v)] = [][2]int{}
				}
				ants[byte(v)] = append(ants[byte(v)], [2]int{x, y})
				antinodes2[[2]int{x, y}] = true
			}
		}
		y++
	}

	for _, ant := range ants {
		for k2, v2 := range ant {
			for k3, v3 := range ant {
				if k3 == k2 {
					continue
				}
				dx, dy := v2[0]-v3[0], v2[1]-v3[1]
				x1, y1 := v2[0], v2[1]
				c := 0
				for {
					x1 += dx
					y1 += dy
					if x1 >= 0 && x1 < maxX && y1 >= 0 && y1 < y {
						if c == 0 {
							antinodes[[2]int{x1, y1}] = true
						}
						antinodes2[[2]int{x1, y1}] = true
						c++
						continue
					}
					break
				}
				dx, dy = -dx, -dy
				x1, y1 = v3[0], v3[1]
				c = 0
				for {
					x1 += dx
					y1 += dy
					if x1 >= 0 && x1 < maxX && y1 >= 0 && y1 < y {
						if c == 0 {
							antinodes[[2]int{x1, y1}] = true
						}
						antinodes2[[2]int{x1, y1}] = true
						c++
						continue
					}
					break
				}
			}
		}
	}
	part1 = len(antinodes)
	part2 = len(antinodes2)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day8("test.txt")
	fmt.Println(part1, part2)
	if part1 != 14 || part2 != 34 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day8("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
