package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

type Coord struct {
	X, Y int
}

var dir = []Coord{
	Coord{1, 0}, Coord{0, 1}, Coord{-1, 0}, Coord{0, -1},
}

type Path struct {
	Level    int
	Distance int
	Path     []Coord
	Last     Coord
	index    int // The index of the item in the heap.
}

type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func f(l, d int) int {
	return l + d*100
}

func (pq PriorityQueue) Less(i, j int) bool {
	// if pq[i].Distance == pq[j].Distance {
	// 	return pq[i].Level < pq[j].Level
	// }
	f1 := f(pq[i].Distance, pq[i].Level)
	f2 := f(pq[j].Distance, pq[j].Level)
	return f1 < f2
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// if pq[i].Level == pq[j].Level {
	// 	return len(pq[i].Path) < len(pq[j].Path)
	// }
	// return pq[i].Level < pq[j].Level
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

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Path, value string, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }
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

	// fmt.Println(board)
	start := Coord{0, 0}
	cur := start
	// lvl := 0
	path := Path{
		Level:    0,
		Path:     []Coord{start},
		Last:     start,
		Distance: distance(start, end),
	}
	// fmt.Println("Distance", path.Distance, "w", f(path.Level, path.Distance))
	eval := make(PriorityQueue, 0)
	eval.Push(&path)
	heap.Init(&eval)

	// var best *Path
	i := 0
	cost := make(map[Coord]int)
	cost[start] = 0
	came_from := make(map[Coord]Coord)

	for eval.Len() > 0 {
		// curpath := eval.Pop().(*Path)
		// fmt.Println("eval len", eval.Len())
		curpath := heap.Pop(&eval).(*Path)

		// cur = curpath.Path[len(curpath.Path)-1]
		cur = curpath.Last
		// fmt.Println("Curpath ", i, curpath.Level, curpath.Distance)
		// i++
		// fmt.Print("  ")
		// for _, c := range curpath.Path {
		// 	// fmt.Print(c, ":", board[c], ",")
		// 	fmt.Print(board[c], ",")
		// }
		// fmt.Println()
		if cur == end {
			// fmt.Println("end ", end, cur, curpath.Path)
			// best = curpath
			break
		}
		// currentCost := f(curpath.Level, curpath.Distance)
		currentCost := cost[curpath.Last]
		for _, c := range dir {
			c2 := Coord{cur.X + c.X, cur.Y + c.Y}
			_, ok := board[c2]
			if !ok {
				continue
			}

			// path2 := make([]Coord, len(curpath.Path))
			// copy(path2, curpath.Path)
			newCost := currentCost + board[c2]

			path := &Path{
				Level: newCost,
				// Path:     append(path2, c2),
				Last:     c2,
				Distance: distance(c2, end),
			}
			// predict := f(path.Level, path.Distance)

			old, ok2 := cost[c2]
			if ok2 {
				if newCost < old {
					cost[c2] = newCost
					came_from[c2] = curpath.Last
					heap.Push(&eval, path)
				}
				continue
			}
			cost[c2] = newCost
			came_from[c2] = curpath.Last
			// fmt.Println("push", c2, path.Level, path.Distance, f(path.Level, path.Distance), eval.Len())
			// eval.Push(&path)
			// if currentCost > predict {
			heap.Push(&eval, path)
			// }
		}
		// fmt.Println("eval", len(eval))
		i++
		if i%1000 == 0 {
			// fmt.Println("give up")
			// break
			// fmt.Println("eval", eval.Len(), len(cost), currentCost, currentCost+distance(cur, end)*9)
		}
		// break
	}
	// fmt.Println("best is", best)
	// lvl := -1
	// for _, c := range best.Path {
	// 	// fmt.Print(c, ":", board[c], ",")
	// 	fmt.Print(board[c], ",")
	// 	lvl += board[c]
	// }
	// fmt.Println()

	part1 := 0
	part1 = cost[end]
	// fmt.Println(cost)
	// fmt.Println(end)
	// for c := end; c != start; c = came_from[c] {
	// 	fmt.Println(c, board[c], cost[c])
	// }
	fmt.Println("Part 1", part1)
	// fmt.Println("Part 2", part2)
}

func day15b(filepath string) {
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
	maxX := end.X + 1
	maxY := end.Y + 1

	board2 := make(map[Coord]int)
	cp := 0
	for c, l := range board {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				c2 := Coord{X: c.X + i*maxX, Y: c.Y + j*maxY}
				l2 := l + i + j
				for l2 > 9 {
					l2 -= 9
				}
				cp++
				// _, ok := board[c2]
				// if ok {
				// 	log.Fatal("dupe", c2)
				// }
				board2[c2] = l2
				// if cp != len(board2) {
				// 	fmt.Println("x", c, c2, i, j)
				// } else {
				// 	fmt.Println(" ", c, c2, i, j)
				// }

				// fmt.Println(cp, len(board2))
			}
		}
	}
	end = Coord{5*maxX - 1, 5*maxY - 1}
	// fmt.Println(end)
	// fmt.Println(len(board), len(board2), cp)
	board = board2
	// fmt.Println("xxx")
	start := Coord{0, 0}
	cur := start
	// lvl := 0
	path := Path{
		Level:    0,
		Path:     []Coord{start},
		Last:     start,
		Distance: distance(start, end),
	}
	// fmt.Println("Distance", path.Distance, "w", f(path.Level, path.Distance))
	eval := make(PriorityQueue, 0)
	eval.Push(&path)
	heap.Init(&eval)

	// var best *Path
	i := 0
	cost := make(map[Coord]int)
	cost[start] = 0
	came_from := make(map[Coord]Coord)

	for eval.Len() > 0 {
		// curpath := eval.Pop().(*Path)
		// fmt.Println("eval len", eval.Len())
		curpath := heap.Pop(&eval).(*Path)

		// cur = curpath.Path[len(curpath.Path)-1]
		cur = curpath.Last
		// fmt.Println("Curpath ", i, curpath.Level, curpath.Distance)
		// i++
		// fmt.Print("  ")
		// for _, c := range curpath.Path {
		// 	// fmt.Print(c, ":", board[c], ",")
		// 	fmt.Print(board[c], ",")
		// }
		// fmt.Println()
		if cur == end {
			// fmt.Println("end ", end, cur, curpath.Path)
			// best = curpath
			break
		}
		// currentCost := f(curpath.Level, curpath.Distance)
		currentCost := cost[curpath.Last]
		for _, c := range dir {
			c2 := Coord{cur.X + c.X, cur.Y + c.Y}
			_, ok := board[c2]
			if !ok {
				continue
			}

			// path2 := make([]Coord, len(curpath.Path))
			// copy(path2, curpath.Path)
			newCost := currentCost + board[c2]

			path := &Path{
				Level: newCost,
				// Path:     append(path2, c2),
				Last:     c2,
				Distance: distance(c2, end),
			}
			// predict := f(path.Level, path.Distance)

			old, ok2 := cost[c2]
			if ok2 {
				if newCost < old {
					cost[c2] = newCost
					came_from[c2] = curpath.Last
					heap.Push(&eval, path)
				}
				continue
			}
			cost[c2] = newCost
			came_from[c2] = curpath.Last
			// fmt.Println("push", c2, path.Level, path.Distance, f(path.Level, path.Distance), eval.Len())
			// eval.Push(&path)
			// if currentCost > predict {
			heap.Push(&eval, path)
			// }
		}
		// fmt.Println("eval", len(eval))
		i++
		if i%1000 == 0 {
			// fmt.Println("give up")
			// break
			// fmt.Println("eval", eval.Len(), len(cost), currentCost, currentCost+distance(cur, end)*9)
		}
		// break
	}
	// fmt.Println("best is", best)
	// lvl := -1
	// for _, c := range best.Path {
	// 	// fmt.Print(c, ":", board[c], ",")
	// 	fmt.Print(board[c], ",")
	// 	lvl += board[c]
	// }
	// fmt.Println()

	part2 := 0
	part2 = cost[end]
	// fmt.Println(cost)
	// fmt.Println(end)
	// for c := end; c != start; c = came_from[c] {
	// 	fmt.Println(c, board[c], cost[c])
	// }
	// fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day15b("test.txt")
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day15("input.txt")  // 402 too high 399 too high
	day15b("input.txt") // 402 too high 399 too high
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
