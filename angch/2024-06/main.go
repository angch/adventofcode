package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"time"
)

func sim(board []byte, maxX2 int, maxX int, maxY int, guard [2]int, obstruct [2]int) int64 {
	dir := [2]int{0, -1}
	dirIndex := 0

	visited2 := make([]byte, len(board))
	guardoff := guard[0] + guard[1]*maxX2
	visited2[guardoff] = byte(1 << dirIndex)

	for {
		guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}

		if guard2[0] < 0 || guard2[0] >= maxX || guard2[1] < 0 || guard2[1] >= maxY {
			// escaped
			return 0
		}
		offset := guard2[0] + guard2[1]*maxX2
		if board[offset] != '#' && guard2 != obstruct {
			guard = guard2
			if (visited2[offset] & (1 << dirIndex)) != 0 {
				return 1
			}
			visited2[offset] |= 1 << dirIndex
		} else if board[offset] == '#' || guard2 == obstruct {
			// turn right
			dir = [2]int{-dir[1], dir[0]}
			dirIndex = (dirIndex + 1) % 4
		}
	}
}

func day6(file string) (part1, part2 int64) {
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
	for i, c := range board {
		if board[i] == '\n' {
			if maxX == 0 {
				maxX = i + 1
			}
			y++
		}
		if c == '^' {
			x := i % maxX
			y = i / maxX
			guard = [2]int{x, y}
			board[i] = '.'
			break
		}
	}
	maxX2 := maxX
	maxX--
	maxY = len(board) / (maxX)
	dir := [2]int{0, -1}

	origguard := guard
	// origd := dir

	visited := make(map[[2]int]bool)
	visited[guard] = true

	workers := 32
	wg := sync.WaitGroup{}
	wg.Add(workers)
	work := make(chan [2]int, workers*16)
	// atomic.StoreInt32(&part2, 0)
	for i := 0; i < workers; i++ {
		go func() {
			for v := range work {
				atomic.AddInt64(&part2, sim(board, maxX2, maxX, maxY, origguard, v))
			}
			wg.Done()
		}()
	}

	for {
		guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}
		if guard2[0] < 0 || guard2[0] >= maxX || guard2[1] < 0 || guard2[1] >= maxY {
			// escaped
			break
		}
		offset := guard2[0] + guard2[1]*maxX2
		if board[offset] != '#' {
			guard = guard2
			visited[guard] = true
			work <- guard
		} else if board[offset] == '#' {
			// turn right
			dir = [2]int{-dir[1], dir[0]}
		}
	}
	part1 = int64(len(visited))

	close(work)
	wg.Wait()
	part2--

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
