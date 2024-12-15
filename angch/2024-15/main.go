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
			if [2]int{x, y} == b.robot {
				if display {
					fmt.Print("@")
				}
				continue
			}
			c := b.board[[2]int{x, y}]
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
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		for x, c := range t {
			if c == '@' {
				robot = [2]int{x, y}
				continue
			}
			if c == '.' {
				continue
			}
			board[[2]int{x, y}] = byte(c)
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
	}
	b2 := Board{
		board: board,
		robot: [2]int{robot[0] * 2, robot[1]}, // robot's x is twice
		maxx:  maxx,
		maxy:  y,
		halfx: true,
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
			_ = m
			if true {
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
	part1, part2 := day15("test.txt")
	fmt.Println(part1, part2)
	if part1 != 2028 {
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
