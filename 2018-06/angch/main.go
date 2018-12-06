package main

// 7720
// 22875

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x, y int
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := make([]Coord, 0)
	tl := Coord{999, 999}
	br := Coord{}

	fudge := 0
	for scanner.Scan() {
		t := scanner.Text()
		c := Coord{}
		fmt.Sscanf(t, "%d, %d", &c.x, &c.y)
		c.x += fudge
		c.y += fudge
		//log.Println(c)
		coords = append(coords, c)
		if c.x < tl.x {
			tl.x = c.x
		}
		if c.y < tl.y {
			tl.y = c.y
		}
		if c.x > br.x {
			br.x = c.x
		}
		if c.y > br.y {
			br.y = c.y
		}
	}
	fmt.Println(tl, br)

	// We'll optimize for maximal prettiness, not operational efficiency.
	viz := make([][]byte, br.y+2+fudge+fudge)
	viz2 := make([][]byte, br.y+2+fudge+fudge)
	for y, _ := range viz {
		viz[y] = make([]byte, br.x+2+fudge+fudge)
		viz2[y] = make([]byte, br.x+2+fudge+fudge)
	}

	for k, c := range coords {
		viz[c.y][c.x] = byte(k + 1)
	}

	area := 0
	if false {
		// Cellular automata!
		//visualize(viz)
		// Damn, this can't do Part 2
		for {
			//visualize(viz)
			flips := 0
			for y, vizy := range viz {
				for x, _ := range vizy {
					viz2[y][x] = viz[y][x]
				}
			}

			for y, cc := range viz {
				for x, cell := range cc {
					if cell == 0 {
						neighbour, count := countNeighbour(viz, x, y)
						if count < 0 {
							viz2[y][x] = 0xff
							flips++
						} else if count > 0 {
							viz2[y][x] = neighbour
							flips++
						} else {

						}
					} else {
						viz2[y][x] = cell
					}
				}
			}
			//log.Println(flips, viz2)
			//visualize(viz2)
			//break
			if flips == 0 {
				break
			}
			for y, vizy := range viz {
				for x, _ := range vizy {
					viz[y][x] = viz2[y][x]
				}
			}
		}
	} else {
		for y, cc := range viz {
			//visualize(viz)
			for x, _ := range cc {

				closestdist := 999999
				closest := -1
				clash := false
				sumdist := 0
				for k, v := range coords {
					dx := v.x - x
					dy := v.y - y
					if dx < 0 {
						dx = -dx
					}
					if dy < 0 {
						dy = -dy
					}
					d := dx + dy

					sumdist += d
					if d < closestdist {
						closestdist = d
						closest = k
						clash = false
					} else if d == closestdist {
						clash = true
					}
				}
				if sumdist < 10000 {
					area++
				}
				if !clash {
					viz[y][x] = byte(closest + 1)
				} else {
					viz[y][x] = 0xff
				}
			}
		}
	}

	//visualize(viz)

	counts := make([]int, 256)
	for y, vizy := range viz {
		if false {
			if y == 0 || y == len(viz)-1 {
				continue
			}
		}
		for x, vizyz := range vizy {
			if false {
				if x == 0 || x == len(vizy)-1 {
					continue
				}
			}
			counts[vizyz]++
		}
	}
	if true {
		for x, _ := range viz {
			w := len(viz) - 1
			counts[viz[0][x]] = 0
			counts[viz[x][0]] = 0
			counts[viz[w][x]] = 0
			counts[viz[x][w]] = 0
		}
		// for y, _ := range []int{0, len(viz) - 1} {
		// 	for x, _ := range []int{0, len(viz[0]) - 1} {

		// 		for yy :=
		// 		vizyx := viz[y][x]
		// 		fmt.Printf("%02x is inf\n", vizyx)
		// 		counts[vizyx] = 0
		// 	}
		// }
	} else {
		for k, c := range coords {
			if c.x == tl.x || c.y == tl.y || c.x == br.x || c.y == br.y {
				fmt.Printf("%02x is inf\n", k+1)
				counts[k+1] = 0
			}
		}
	}

	maxCount, maxByte := 0, 0
	for k, v := range counts {
		if v > 0 {
			if false && k-1 < len(coords) {
				fmt.Printf("%02x %d %d %d\n", k, v, coords[k-1].x, coords[k-1].y)
			} else {
				fmt.Printf("%02x %d\n", k, v)
			}
		}
		if maxCount < v && k != 0xff {
			maxCount = v
			maxByte = k
		}
	}
	fmt.Printf("%02x %d %d", maxByte, maxCount, area)
}
func visualize(viz [][]byte) {
	for y, vizy := range viz {
		for x, _ := range vizy {
			fmt.Printf("%02x ", viz[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func countNeighbour(viz [][]byte, x, y int) (byte, int) {
	n, count := byte(0xff), 0

	matrix := []Coord{
		Coord{-1, 0},
		Coord{0, -1},
		Coord{1, 0},
		Coord{0, 1},
	}

	for _, m := range matrix {
		yy := m.y + y
		xx := m.x + x
		if yy < 0 {
			continue
		}
		if yy >= len(viz) {
			continue
		}
		if xx < 0 {
			continue
		}
		if xx >= len(viz[yy]) {
			continue
		}
		if viz[yy][xx] > 0 {
			if count > 0 {
				if viz[yy][xx] == n {
					count++
					continue
				} else {
					return 0, -1
				}
			}
			if viz[yy][xx] != 0xff {
				n = viz[yy][xx]
				count++
			}
		}
	}
	return n, count
}
