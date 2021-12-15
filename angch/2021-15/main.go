package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sync"
	"time"
)

type Coord struct {
	X, Y int
}

var dir = []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Path struct {
	Level    int
	Distance int
	Last     Coord
	index    int
}

type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func f(l, d int) int {
	return l + d*10
}

func (pq PriorityQueue) Less(i, j int) bool {
	f1 := f(pq[i].Distance, pq[i].Level)
	f2 := f(pq[j].Distance, pq[j].Level)
	return f1 < f2
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Path)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func distance(a, b Coord) int {
	d := a.X - b.X
	if d < 0 {
		d = -d
	}
	e := a.Y - b.Y
	if e < 0 {
		e = -e
	}
	return d + e
}

func solve(board map[Coord]int, start Coord, end Coord) int {
	path := Path{
		Level:    0,
		Last:     start,
		Distance: distance(start, end),
	}

	eval := make(PriorityQueue, 0)
	eval.Push(&path)
	heap.Init(&eval)

	cost := make(map[Coord]int)
	cost[start] = 0
	// came_from := make(map[Coord]Coord)

	for eval.Len() > 0 {
		curpath := heap.Pop(&eval).(*Path)
		cur := curpath.Last
		if cur == end {
			break
		}
		currentCost := cost[curpath.Last]
		for _, c := range dir {
			c2 := Coord{cur.X + c.X, cur.Y + c.Y}
			_, ok := board[c2]
			if !ok {
				continue
			}

			newCost := currentCost + board[c2]
			path := &Path{
				Level:    newCost,
				Last:     c2,
				Distance: distance(c2, end),
			}

			_, ok2 := cost[c2]
			if ok2 {
				continue
			}
			cost[c2] = newCost
			// came_from[c2] = curpath.Last
			heap.Push(&eval, path)
		}
	}
	return cost[end]
}

func day15(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	board := make(map[Coord]int)
	y := 0
	end := Coord{0, 0}
	for scanner.Scan() {
		t := scanner.Text()
		for x, v := range t {
			end = Coord{y, x}
			board[end] = int(v) - int('0')
		}
		y++
	}
	start := Coord{0, 0}
	part1, part2 := 0, 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(board map[Coord]int, start, end Coord) {
		part1 = solve(board, start, end)
		wg.Done()
	}(board, start, end)

	maxX := end.X + 1
	maxY := end.Y + 1

	board2 := make(map[Coord]int)
	for c, l := range board {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				c2 := Coord{X: c.X + i*maxX, Y: c.Y + j*maxY}
				l2 := l + i + j
				for l2 > 9 {
					l2 -= 9
				}
				board2[c2] = l2
			}
		}
	}
	end2 := Coord{5*maxX - 1, 5*maxY - 1}

	wg.Add(1)
	go func(board map[Coord]int, start, end Coord) {
		part2 = solve(board, start, end)
		wg.Done()
	}(board2, start, end2)
	wg.Wait()

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day15b("test.txt")
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day15("input.txt") // 402 too high 399 too high
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
