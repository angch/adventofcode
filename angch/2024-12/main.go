package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type region struct {
	Coords    map[[2]int]byte
	Perimeter int
	Area      int
	Side      int
	Plant     byte
}

func day12(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	mm := map[[2]int]byte{}

	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, c := range t {
			mm[[2]int{x, y}] = byte(c)
		}
		y++
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for k, plant := range mm {
		r := region{Coords: map[[2]int]byte{k: plant}}
		checks := [][2]int{k}
		delete(mm, k)
		for len(checks) > 0 {
			c := checks[len(checks)-1]
			checks = checks[:len(checks)-1]

			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if mm[dc] == plant {
					r.Coords[dc] = plant
					delete(mm, dc)
					checks = append(checks, dc)
				}
			}
		}
		r.Area = len(r.Coords)

		// Perimeter
		miny, maxy := 10000, -1
		minx, maxx := 10000, -1
		for c := range r.Coords {
			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if r.Coords[dc] == 0 {
					r.Perimeter++
				}
			}
			miny, maxy = min(miny, c[1]), max(maxy, c[1])
			minx, maxx = min(minx, c[0]), max(maxx, c[0])
		}

		// Sides
		// Using scanlining, we count the changes in edges from row to row
		// and column to column, giving us the number of sides.
		// Top to bottom
		prevchanges := map[int]int{}
		for y1 := miny; y1 <= maxy; y1++ {
			prev := byte(0)
			changes := map[int]int{}
			// Mark rising and falling edges
			for x1 := minx; x1 <= maxx; x1++ {
				if prev == 0 {
					if r.Coords[[2]int{x1, y1}] == plant {
						prev = plant
						changes[x1] = 1
					}
				} else {
					if r.Coords[[2]int{x1, y1}] != plant {
						prev = 0
						changes[x1] = -1
					}
				}
			}
			if prev > 0 {
				changes[maxx+1] = -1
			}
			// Every different change from row to row denotes a side
			for v, change := range changes {
				if prevchanges[v] != change {
					r.Side++
				}
			}
			prevchanges = changes
		}

		// Left to right
		prevchanges = map[int]int{}
		for x1 := minx; x1 <= maxx; x1++ {
			prev := byte(0)
			changes := map[int]int{}
			// Mark rising and falling edges
			for y1 := miny; y1 <= maxy; y1++ {
				if prev == 0 {
					if r.Coords[[2]int{x1, y1}] == plant {
						prev = plant
						changes[y1] = 1
					}
				} else {
					if r.Coords[[2]int{x1, y1}] != plant {
						prev = 0
						changes[y1] = -1
					}
				}
			}
			if prev > 0 {
				changes[maxy+1] = -1
			}
			// Every different change from row to row denotes a side
			for v, change := range changes {
				if prevchanges[v] != change {
					r.Side++
				}
			}
			prevchanges = changes
		}
		part1 += r.Area * r.Perimeter
		part2 += r.Area * r.Side
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day12("test0.txt")
	fmt.Println(part1, part2)
	if part1 != 140 || part2 != 80 {
		log.Fatal("Test failed ", part1, part2)
	}
	part1, part2 = day12("test1.txt")
	fmt.Println(part1, part2)
	if part1 != 772 || part2 != 436 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day12("test2.txt")
	fmt.Println(part1, part2)
	if part1 != 692 || part2 != 236 {
		log.Fatal("Test failed ", part1, part2)
	}

	// This is the important edge case test.
	part1, part2 = day12("test3.txt")
	fmt.Println(part1, part2)
	if part1 != 1184 || part2 != 368 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day12("test.txt")
	fmt.Println(part1, part2)
	if part1 != 1930 || part2 != 1206 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day12("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
