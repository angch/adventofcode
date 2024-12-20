package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"time"
)

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type boardstate struct {
	index int
	mm    map[[2]int]byte
	moves [][2]int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*boardstate

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return len(pq[i].moves) < len(pq[j].moves)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*boardstate)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func solve(mm map[[2]int]byte, start, end [2]int, maxx, maxy int) []boardstate {
	cur := start
	move := 0
	_ = move

	pq := make(PriorityQueue, 1)
	pq[0] = &boardstate{0, mm, [][2]int{cur}}
	heap.Init(&pq)

	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*boardstate)

	a:
		for _, d := range dirs {
			newpos := [2]int{e.moves[len(e.moves)-1][0] + d[0], e.moves[len(e.moves)-1][1] + d[1]}
			if newpos == end {
				return []boardstate{{0, e.mm, append(e.moves, newpos)}}
			}
			c := mm[newpos]
			if c == '#' {
				// it's a wall
				continue
			}
			if c == '.' {
				for _, v := range e.moves {
					if newpos == v {
						continue a
					}
				}
				heap.Push(&pq, &boardstate{0, e.mm, append(e.moves, newpos)})
			}
		}
	}
	return nil
}

func day20(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	mm := make(map[[2]int]byte)
	start := [2]int{}
	end := [2]int{}
	maxx := 0
	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		maxx = len(t)
		for x, c := range t {
			if c == 'S' {
				start = [2]int{x, y}
				mm[[2]int{x, y}] = '.'
				continue
			} else if c == 'E' {
				end = [2]int{x, y}
				mm[[2]int{x, y}] = '.'
				continue
			} else if c == '#' {
				mm[[2]int{x, y}] = '#'
				continue
			}
			mm[[2]int{x, y}] = '.'
		}
		y++
	}
	_ = start
	_ = end
	out := solve(mm, start, end, maxx, y)
	// log.Printf("%+v\n", out[0].moves)
	nocheat := len(out[0].moves) - 1
	part1 = nocheat

	moves := out[0].moves

	path := map[[2]int]int{}
	for k, move := range moves {
		path[move] = k
	}

	counts := make(map[int]int)
	shortcuts := make(map[[4]int]int)
	for k, move := range moves {
		for _, d1 := range dirs {
			newpos1 := [2]int{move[0] + d1[0], move[1] + d1[1]}
			if mm[newpos1] == '#' {
				for _, d2 := range dirs {
					newpos2 := [2]int{newpos1[0] + d2[0], newpos1[1] + d2[1]}
					s, ok := path[newpos2]
					if s > k+2 && ok {

						startpos := move
						endpos := newpos2
						saved := s - k - 2

						first, ok := shortcuts[[4]int{startpos[0], startpos[1], endpos[0], endpos[1]}]
						if ok {
							if first <= saved {
								continue
							}
						}
						// fmt.Println("Found shortcut", startpos, endpos, saved)

						shortcuts[[4]int{startpos[0], startpos[1], endpos[0], endpos[1]}] = saved

					}
				}
			}
		}
	}
	for _, v := range shortcuts {
		counts[v]++
		if v >= 100 {
			part2++
		}
	}
	// fmt.Println(counts)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day20("test.txt")
	fmt.Println(part1, part2)
	if part1 != 84 || part2 != 0 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day20("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
