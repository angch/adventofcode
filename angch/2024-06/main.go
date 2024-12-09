package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func day6(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	board, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY := 0, 0
	guard := [2]int{0, 0}

	y := 0
	xc := 0
	yc := 0
	for i, c := range board {
		if board[i] == '\n' {
			if maxX == 0 {
				maxX = i
			}
			y++
			xc = 0
			yc++
		}
		if c == '^' {
			x := i % (maxX + 1)
			y = i / (maxX + 1)
			guard = [2]int{x, y}
			board[i] = '.'
		}
		xc++
	}
	maxY = len(board) / (maxX)
	dir := [2]int{0, -1}

	orig := guard
	origd := dir

	visited := make(map[[2]int]bool)
	visited[guard] = true
	maxX2 := maxX + 1

	for {
		guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}
		offset := guard2[0] + guard2[1]*maxX2
		if guard2[0] < 0 || guard2[0] >= maxX || guard2[1] < 0 || guard2[1] >= maxY {
			// escaped
			break
		}
		if board[offset] != '#' {
			guard = guard2
			// fmt.Println(guard)
			visited[guard] = true
		} else if board[offset] == '#' {
			// turn right
			dir = [2]int{-dir[1], dir[0]}
		}
	}
	part1 = len(visited)
a:
	for v := range visited {
		guard = orig
		dir = origd
		dirIndex := 0

		visited2 := make([]byte, len(board))
		guardoff := guard[0] + guard[1]*maxX2
		visited2[guardoff] = byte(1 << dirIndex)

		for {
			guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}
			offset := guard2[0] + guard2[1]*maxX2
			if guard2[0] < 0 || guard2[0] >= maxX || guard2[1] < 0 || guard2[1] >= maxY {
				// escaped
				break
			}
			if board[offset] != '#' && guard2 != v {
				guard = guard2
				if (visited2[offset] & (1 << dirIndex)) != 0 {
					part2++
					continue a
				}
				visited2[offset] |= 1 << dirIndex
			} else if board[offset] == '#' || guard2 == v {
				// turn right
				dir = [2]int{-dir[1], dir[0]}
				dirIndex = (dirIndex + 1) % 4
			} else {
				break
			}
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	logperf := false
	if logperf {
		pf, _ := os.Create("cpu.pprof")
		pprof.StartCPUProfile(pf)
		defer pf.Close()
	}
	t1 := time.Now()
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	if part1 != 41 || part2 != 6 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day6("input.txt"))
	if logperf {
		pprof.StopCPUProfile()
	}
	fmt.Println("Elapsed time:", time.Since(t1))
}
