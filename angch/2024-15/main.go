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
	'^': [2]int{0, -1},
	'>': [2]int{1, 0},
	'<': [2]int{-1, 0},
	'v': [2]int{0, 1},
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

func dumpboard2(b Board) (int, int) {
	_ = dirmap
	display := true
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
			coordp1 := [2]int{x - 1, y}

			if coord == b.robot || coordp1 == b.robot {
				if display {
					fmt.Print("@")
				}
				continue
			}

			if b.obs[coordhalf] {
				fmt.Print("#")
				continue
			}

			if boxMap[coord] != byte(0) {
				fmt.Print(string(boxMap[coord]))
				continue
			}

			fmt.Print(".")
		}
		if display {
			fmt.Println()
		}
	}
	return gps, 0
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
				fmt.Println("Can't move O")
				continue a
			}
			b.board[r1] = 'O'
			delete(b.board, l1)
			b.robot = coord1
			fmt.Println("Move block O")
		}
	b:
		for m, move := range moves {
			if true {
				fmt.Println(m)
				dumpboard2(b2)
			}
			d := dirmap[byte(move)]

			coord1 := [2]int{b.robot[0] + d[0], b.robot[1] + d[1]}
			// Empty

			coord1m := [2]int{coord1[0] - 1, coord1[1]}
			coordhalf := [2]int{coord1[0] / 2, coord1[1]}

			if b.obs[coordhalf] {
				// blocked
				fmt.Println("Blocked", string(move))
				continue
			}

			intheway := []int{}
			for k, v := range b.boxes {
				if v == coord1 || v == coord1m {
					intheway = append(intheway, k)
				}
			}
			if len(intheway) == 0 {
				fmt.Println("no blockj", string(move))
				b2.robot = coord1
				continue b
			}
		}
	}
	part1, part2 = dumpboard(b)
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
	part1, part2 := day15("test3.txt")
	fmt.Println(part1, part2)
	if part1 != 2028+1 {
		log.Fatal("Test failed ", part1, part2)
	}
	part1, part2 = day15("test2.txt")
	fmt.Println(part1, part2)
	if part1 != 10092 || part2 != 9021 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day15("input.txt"))
	// 7790
	if logperf {
		pprof.StopCPUProfile()
	}

	fmt.Println("Elapsed time:", time.Since(t1))
}
