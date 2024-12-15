package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var dirmap = map[byte][2]int{
	'^': {0, -1},
	'>': {1, 0},
	'<': {-1, 0},
	'v': {0, 1},
}

type Board struct {
	board map[[2]int]byte
	obs   map[[2]int]bool
	boxes [][2]int
	robot [2]int
	maxx  int
	maxy  int
	halfx bool
}

func dumpboard(b Board) (int, int) {
	_ = dirmap
	display := false
	gps := 0
	for y := 0; y < b.maxy; y++ {
		for x := 0; x < b.maxx; x++ {
			coord := [2]int{x, y}

			if coord == b.robot {
				if display {
					fmt.Print("@")
				}
				continue
			}
			c := b.board[coord]
			if c == 'O' {
				gps += y*100 + x
			}
			if display {
				switch c {
				case '#', 'O':
					fmt.Print(string(c))
				default:
					fmt.Print(".")
				}
			}

		}
		if display {
			fmt.Println()
		}
	}
	return gps, 0
}

func dumpboard2(b Board, display bool) (int, int) {
	_ = dirmap
	// display := true
	gps := 0

	boxMap := map[[2]int]byte{}
	for _, v := range b.boxes {
		boxMap[v] = '['
		boxMap[[2]int{v[0] + 1, v[1]}] = ']'
	}

	for y := 0; y < b.maxy; y++ {
		for x := 0; x < b.maxx; x++ {
			coord := [2]int{x, y}
			coordhalf := [2]int{x / 2, y}
			// coordp1 := [2]int{x - 1, y}

			if coord == b.robot {
				if display {
					fmt.Print("@")
				}
				continue
			}

			if b.obs[coordhalf] {
				if display {
					fmt.Print("#")
				}
				continue
			}

			if boxMap[coord] != byte(0) {
				if display {
					fmt.Print(string(boxMap[coord]))
				}
				if boxMap[coord] == '[' {
					gps += y*100 + x
				}
				continue
			}

			if display {
				fmt.Print(".")
			}
		}
		if display {
			fmt.Println()
		}
	}
	fmt.Println("Robot:", b.robot)
	return 0, gps
}

func intheway(b Board, coord1 [2]int, d [2]int, dwidth bool) (bool, []int) {
	out := []int{}

	coord1 = [2]int{coord1[0] + d[0], coord1[1] + d[1]}
	coord1m := [2]int{coord1[0] - 1, coord1[1]}
	coord1m2 := [2]int{coord1[0] + 1, coord1[1]}

	coordhalf := [2]int{coord1[0] / 2, coord1[1]}
	coordhalf2 := [2]int{coord1[0]/2 + 1, coord1[1]}
	// coordhalf3 := [2]int{coord1m2[0] / 2, coord1[1]}

	if b.obs[coordhalf] {
		// fmt.Println("itw: obs")
		return true, []int{}
	}
	if dwidth {
		fmt.Println("itw: dwidth", coord1[0])
		if (coord1[0]%2 == 0 && b.obs[coordhalf]) ||
			(coord1[0]%2 == 1 && b.obs[coordhalf2]) {
			// fmt.Println("itw: obs dwidth")
			return true, []int{}
		}
	}

	for k, v := range b.boxes {
		if v == coord1 || v == coord1m {
			out = append(out, k)
		} else if dwidth {
			if v == coord1m2 {
				out = append(out, k)
			}
		}
	}
	// fmt.Println("itw", coord1, out)
	return false, out
}

func day15(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	board := map[[2]int]byte{}
	robot := [2]int{0, 0}
	y := 0
	maxx := 0
	obs := map[[2]int]bool{}
	boxes := [][2]int{}
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		for x, c := range t {
			coord := [2]int{x, y}
			if c == '@' {
				robot = coord
				continue
			}
			if c == '.' {
				continue
			}
			board[coord] = byte(c)
			if c == '#' {
				obs[[2]int{x, y}] = true
			}
			if c == 'O' {
				boxes = append(boxes, coord)
			}
		}
		maxx = len(t)
		y++
	}
	b := Board{
		board: board,
		robot: robot,
		maxx:  maxx,
		maxy:  y,
		halfx: false,
		obs:   obs,
		boxes: boxes,
	}
	b2 := Board{
		board: map[[2]int]byte{},              // not used
		robot: [2]int{robot[0] * 2, robot[1]}, // fullx
		maxx:  maxx * 2,                       // fullx
		maxy:  y,
		halfx: true,
		obs:   obs,   // halfx
		boxes: boxes, // fullx
	}
	for k := range b2.boxes {
		b2.boxes[k][0] *= 2
	}

	dumpboard(b)
	tm := 0
	for scanner.Scan() {
		moves := scanner.Text()

	a:
		for m, move := range moves {
			_ = m
			if false {
				fmt.Println(m)
				dumpboard(b)
			}

			d := dirmap[byte(move)]

			coord1 := [2]int{b.robot[0] + d[0], b.robot[1] + d[1]}
			// Empty
			bb := b.board[coord1]

			switch bb {
			case byte(0):
				b.robot = coord1
				continue a
			case '#':
				continue a
			}

			l1 := coord1
			r1 := coord1
			for b.board[r1] == 'O' {
				r1 = [2]int{r1[0] + d[0], r1[1] + d[1]}
			}
			if b.board[r1] == '#' {
				// Can't move.
				// fmt.Println("Can't move O")
				continue a
			}
			b.board[r1] = 'O'
			delete(b.board, l1)
			b.robot = coord1
			// fmt.Println("Move block O")
		}
		// fmt.Println("Moves", moves)
		prevmove := ""
		debug := true
	b:
		for _, move := range moves {
			tm++
			if debug && tm >= 108 {
				// bufio.NewReader(os.Stdin).ReadBytes('\n')
				time.Sleep(100 * time.Millisecond)
				fmt.Print("\033[H\033[2J")
				fmt.Println("Part2", tm, string(prevmove), string(move))
				// fmt.Println("Part2", m, string(move))
				_ = prevmove
				dumpboard2(b2, true)
			}
			prevmove = string(move)
			d := dirmap[byte(move)]

			coord1 := [2]int{b2.robot[0] + d[0], b2.robot[1] + d[1]}
			// Empty

			// coord1m := [2]int{coord1[0] - 1, coord1[1]}
			coordhalf := [2]int{coord1[0] / 2, coord1[1]}

			// fmt.Println(coord1, coord1m, coordhalf)
			if b2.obs[coordhalf] {
				// blocked
				fmt.Println("Blocked", string(move))
				continue
			}

			obstructed, itw := intheway(b2, b2.robot, d, false)
			if !obstructed && len(itw) == 0 {
				fmt.Println("no block", string(move))
				b2.robot = coord1
				if debug {
					// dumpboard2(b2, true)
				}
				continue b
			}
			fmt.Println("Blocked by boxes", itw, obstructed)

			boxes := map[int]bool{}
			for _, k := range itw {
				boxes[k] = true
			}
		c:
			for _, box := range itw {
				obstructed, itw = intheway(b2, b2.boxes[box], d, true)
				if obstructed {
					fmt.Println("Can't move boxes", string(move))
					// dumpboard2(b2, true)
					continue b
				}
				if len(itw) == 0 {
					// Ok, move the boxes
					continue c
				} else {
					// More boxes in the way
					itw2 := []int{}
					for _, k := range itw {
						if !boxes[k] {
							boxes[k] = true
							itw2 = append(itw2, k)
						}
					}
					// itw = []int{}
					itw = itw2
					fmt.Println("Again", boxes)
					goto c // again
				}
			}
			fmt.Println("Moving boxes", boxes)
			for k := range boxes {
				b2.boxes[k][0] += d[0]
				b2.boxes[k][1] += d[1]
			}
			b2.robot = coord1
		}
	}
	part1, _ = dumpboard(b)
	_, part2 = dumpboard2(b2, true)
	log.Println("Part2", part2, b2.robot)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	logperf := false
	if logperf {
		pf, _ := os.Create("default.pgo")
		err := pprof.StartCPUProfile(pf)
		if err != nil {
			log.Fatal(err)
		}
		defer pf.Close()
	}
	t1 := time.Now()
	// part1, part2 := day15("test.txt")
	// fmt.Println(part1, part2)
	// if part1 != 2028 {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// part1, part2 := day15("test3.txt")
	// fmt.Println(part1, part2)
	// if part1 != 908 {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	// part1, part2 := day15("test2.txt")
	// fmt.Println(part1, part2)
	// if part1 != 10092 || part2 != 9021 {
	// 	log.Fatal("Test failed ", part1, part2)
	// }
	fmt.Println(day15("input.txt"))
	// too low: 1451343
	// too low: 1453002
	if logperf {
		pprof.StopCPUProfile()
	}

	fmt.Println("Elapsed time:", time.Since(t1))
}
