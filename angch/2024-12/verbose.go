package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day12verbose(file string) (part1, part2 int) {
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
	for scanner.Scan() {
		t := scanner.Text()
		for x, c := range t {
			mm[[2]int{x, y}] = byte(c)
			mm2[[2]int{x, y}] = byte(c)
		}
		y++
	}

	verbose := false
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for k, plant := range mm {
		thisregion := region{Coords: map[[2]int]byte{k: plant}, Plant: plant}
		checks := [][2]int{k}
		delete(mm, k)
		for len(checks) > 0 {
			c := checks[len(checks)-1]
			checks = checks[:len(checks)-1]

			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if mm[dc] == plant {
					thisregion.Coords[dc] = plant
					delete(mm, dc)
					checks = append(checks, dc)
				}
			}
		}
		thisregion.Area = len(thisregion.Coords)

		miny, maxy := 10000, -1
		minx, maxx := 10000, -1
		for c := range thisregion.Coords {
			for _, v2 := range dirs {
				dc := [2]int{c[0] + v2[0], c[1] + v2[1]}
				if mm2[dc] != plant {
					thisregion.Perimeter++
				}
			}
			miny, maxy = min(miny, c[1]), max(maxy, c[1])
			minx, maxx = min(minx, c[0]), max(maxx, c[0])
		}

		// Scanlining, we count the changes in edges from row to row and column to column, giving
		// us the number of sides.
		// Top to bottom
		prevchanges := map[int]int{}
		for y1 := miny; y1 <= maxy; y1++ {
			prev := byte(0)
			changes := map[int]int{}
			for x1 := minx; x1 <= maxx; x1++ {
				if prev == 0 {
					if thisregion.Coords[[2]int{x1, y1}] == plant {
						prev = plant
						changes[x1] = 1
					}
				} else {
					if thisregion.Coords[[2]int{x1, y1}] != plant {
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
			prevchanges = changes
		}

		// Left to right
		prevchanges = map[int]int{}
		for x1 := minx; x1 <= maxx; x1++ {
			prev := byte(0)
			changes := map[int]int{}
			for y1 := miny; y1 <= maxy; y1++ {
				if prev == 0 {
					if thisregion.Coords[[2]int{x1, y1}] == plant {
						prev = plant
						changes[y1] = 1
					}
				} else {
					if thisregion.Coords[[2]int{x1, y1}] != plant {
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
	for _, r := range regions {
		if verbose {
			fmt.Println("Region", string(r.Plant), r.Area, r.Perimeter, r.Side)
		}
		part1 += r.Area * r.Perimeter
		part2 += r.Area * r.Side

		if verbose && r.Side > 20 {
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
