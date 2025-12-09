package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type pairDist struct {
	p1    int
	p2    int
	d     int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*pairDist

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].d < pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*pairDist)
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

// dist doesn't need actual dist, just enough for sorting
func dist(a, b [3]int) int {
	x1, y1, z1 := a[0]-b[0], a[1]-b[1], a[2]-b[2]
	x1, y1, z1 = x1*x1, y1*y1, z1*z1
	return x1 + y1 + z1
}

func day8(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	points := [][3]int{}
	circuits := [][]int{}
	pointInCircuit := map[int]int{}

	for scanner.Scan() {
		t := scanner.Text()
		x, y, z := 0, 0, 0
		fmt.Sscanf(t, "%d,%d,%d", &x, &y, &z)
		points = append(points, [3]int{x, y, z})
	}
	shortest := make(PriorityQueue, 0, len(points)*(len(points)))
	heap.Init(&shortest)

	maxConn := len(points)
	if maxConn == 20 {
		maxConn = 10 // the test is shorter than full run
	}

	t1 := time.Now()
	for k1, v1 := range points {
		for k2 := k1 + 1; k2 < len(points); k2++ {
			v2 := points[k2]
			pd := pairDist{k1, k2, dist(v1, v2), 0}
			heap.Push(&shortest, &pd)
		}
	}
	t2 := time.Now()

	count := 0
	k := -1
	for {
		k++
		pd := heap.Pop(&shortest).(*pairDist)
		if k == maxConn {
			lens := []int{}
			for _, v := range circuits {
				lens = append(lens, len(v))
			}
			// log.Println(lens)
			sort.Sort(sort.Reverse(sort.IntSlice(lens)))
			part1 = lens[0] * lens[1] * lens[2]
			// fmt.Println("part 1 done")
			// break
		}

		if pointInCircuit[pd.p1] > 0 && pointInCircuit[pd.p2] > 0 &&
			pointInCircuit[pd.p1] == pointInCircuit[pd.p2] {
			// Already in, do nothing
		} else if pointInCircuit[pd.p1] > 0 && pointInCircuit[pd.p2] > 0 {
			nc := pointInCircuit[pd.p1] - 1
			nc2 := pointInCircuit[pd.p2] - 1
			circuits[nc] = append(circuits[nc], circuits[nc2]...)
			sort.Ints(circuits[nc])
			for _, v2 := range circuits[nc2] {
				pointInCircuit[v2] = nc + 1
			}
			circuits[nc2] = []int{}
			count++
		} else if pointInCircuit[pd.p1] > 0 {
			nc := pointInCircuit[pd.p1] - 1
			circuits[nc] = append(circuits[nc], pd.p2)
			sort.Ints(circuits[nc])
			pointInCircuit[pd.p2] = nc + 1
			count++
		} else if pointInCircuit[pd.p2] > 0 {
			nc := pointInCircuit[pd.p2] - 1
			circuits[nc] = append(circuits[nc], pd.p1)
			sort.Ints(circuits[nc])
			pointInCircuit[pd.p1] = nc + 1
			count++
		} else {
			nc := len(circuits)
			circuits = append(circuits, []int{pd.p1, pd.p2})
			pointInCircuit[pd.p1] = nc + 1
			pointInCircuit[pd.p2] = nc + 1
			count++
		}

		nonzero := 0
		lens := 0
		for k3 := range circuits {
			if len(circuits[k3]) > 0 {
				nonzero++
				lens += len(circuits[k3])
			}
		}

		if nonzero == 1 && len(pointInCircuit) == len(points) {
			// fmt.Println("Found", len(pointInCircuit), len(points))
			part2 = points[pd.p1][0] * points[pd.p2][0]
			break
		}
	}

	t3 := time.Now()
	fmt.Println("Time", t2.Sub(t1), t3.Sub(t2))
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day8("test.txt")
	fmt.Println(part1, part2)
	if part1 != 40 || part2 != 25272 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day8("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
