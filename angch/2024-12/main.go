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

	regions := []region{}
	mm := map[[2]int]byte{}
	mm2 := map[[2]int]byte{}

	y := 0
	// maxx := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, c := range t {
			mm[[2]int{x, y}] = byte(c)
			mm2[[2]int{x, y}] = byte(c)
		}
		y++
		// maxx = len(t)
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// a:
	for k, v := range mm {
		first := v

		thisregion := region{Coords: map[[2]int]byte{k: first}, Plant: first}
		checks := [][2]int{k}
		delete(mm, k)
		for len(checks) > 0 {
			c := checks[len(checks)-1]
			checks = checks[:len(checks)-1]

			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if mm[dc] == first {
					// thisregion.Coords = append(thisregion.Coords, dc)
					thisregion.Coords[dc] = first
					delete(mm, dc)
					checks = append(checks, dc)
				}
			}
			// fmt.Println("ds", len(checks), first, checks, c)
		}
		thisregion.Area = len(thisregion.Coords)

		miny, maxy := 10000, -1
		minx, maxx := 10000, -1
		for c, _ := range thisregion.Coords {
			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if mm2[dc] != first {
					thisregion.Perimeter++
				}
			}
			miny = min(miny, c[1])
			maxy = max(maxy, c[1])
			minx = min(minx, c[0])
			maxx = max(maxx, c[0])
		}

		// Scanlining
		// Top to bottom
		prevchanges := map[int]int{}
		for y1 := miny; y1 <= maxy; y1++ {
			prev := byte(0)
			changes := map[int]int{}
			for x1 := minx; x1 <= maxx; x1++ {
				if prev == 0 {
					if thisregion.Coords[[2]int{x1, y1}] == first {
						prev = first
						changes[x1] = 1
					}
				} else {
					if thisregion.Coords[[2]int{x1, y1}] != first {
						prev = 0
						changes[x1] = -1
					}
				}
			}
			if prev > 0 {
				changes[maxx+1] = -1
			}
			for v, change := range changes {
				if prevchanges[v] != change {
					thisregion.Side++
				}
			}
			log.Println("x", changes, minx, miny, maxx, maxy)
			prevchanges = changes
		}

		// Left to right
		prevchanges = map[int]int{}
		for x1 := minx; x1 <= maxx; x1++ {
			prev := byte(0)
			changes := map[int]int{}
			for y1 := miny; y1 <= maxy; y1++ {
				if prev == 0 {
					if thisregion.Coords[[2]int{x1, y1}] == first {
						prev = first
						changes[y1] = 1
					}
				} else {
					if thisregion.Coords[[2]int{x1, y1}] != first {
						prev = 0
						changes[y1] = -1
					}
				}
			}
			if prev > 0 {
				changes[maxy+1] = -1
			}
			for v, change := range changes {
				if prevchanges[v] != change {
					thisregion.Side++
				}
			}
			prevchanges = changes
		}

		regions = append(regions, thisregion)
	}
	_ = regions
	for _, r := range regions {
		fmt.Println("Region", string(r.Plant), r.Area, r.Perimeter, r.Side)
		part1 += r.Area * r.Perimeter
		part2 += r.Area * r.Side

		if r.Side > 20 {
			minx, miny, maxx, maxy := 10000, 10000, -1, -1
			for c, _ := range r.Coords {
				miny = min(miny, c[1])
				maxy = max(maxy, c[1])
				minx = min(minx, c[0])
				maxx = max(maxx, c[0])
			}

			for y1 := miny; y1 <= maxy; y1++ {
				for x1 := minx; x1 <= maxx; x1++ {
					if r.Coords[[2]int{x1, y1}] == r.Plant {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
		}
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
	// Not: 912944
	fmt.Println("Elapsed time:", time.Since(t1))
}
