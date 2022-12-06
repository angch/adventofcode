package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	levenshtein "github.com/ka-weihe/fast-levenshtein"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	depth    int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func calc(input, target string) int {
	if true {
		distance := levenshtein.Distance(input, target)
		return -distance
	}

	diff := len(target) - len(input)

	if diff > 0 {
		diff = -diff
	}

	count := 0
	a := len(input)
	if len(target) < a {
		a = len(target)
	}
	for i := 0; i < a; i++ {
		if input[i] == target[i] {
			count++
		}
	}
	diff = count + diff

	return diff
}

func getReplacements(replacesFromTo []rep, target string) map[string]bool {
	replacements := make(map[string]bool)
	for _, v := range replacesFromTo {
		from, to := v.From, v.To
		matches := from.FindAllStringIndex(target, -1)
		for _, v2 := range matches {
			left := target[:v2[0]]
			right := target[v2[1]:]
			out := left + to + right
			replacements[out] = true
		}
	}
	return replacements
}

type rep struct {
	From *regexp.Regexp
	To   string
}

func day19(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	replacesFromTo := []rep{}
	replacesToFrom := []rep{}
	target := ""
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		t2 := strings.Split(t, " => ")
		if len(t2) == 1 {
			target = t2[0]
			break
		}
		r := rep{From: regexp.MustCompile(t2[0]), To: t2[1]}
		replacesFromTo = append(replacesFromTo, r)
		r2 := rep{From: regexp.MustCompile(t2[1]), To: t2[0]}
		replacesToFrom = append(replacesToFrom, r2)
	}
	replacements := getReplacements(replacesFromTo, target)
	part1 = len(replacements)

	// debug := getReplacements(replacesFromTo, "HH")
	// log.Println(debug)

	if false {
		pq := make(PriorityQueue, 0)
		first := &Item{
			value:    "e",
			priority: calc("e", target),
			depth:    0,
		}
		heap.Init(&pq)
		heap.Push(&pq, first)

		done := make(map[string]bool)
		done["e"] = true
		prevTime := time.Now()
		operations := 0

		// answer := make(chan )
		for pq.Len() > 0 {
			q := heap.Pop(&pq).(*Item)
			// fmt.Printf("%.2d:%s ", item.priority, item.value)

			if q.value == target {
				part2 = q.depth
				log.Println(q.value)
				break
			}
			// queue = queue[1:]
			if time.Since(prevTime) > 2*time.Second {
				prevTime = time.Now()
				log.Println(q.depth, calc(q.value, target), len(pq), len(done), q.value)
				// log.Println(q.depth, q.value)
			}
			operations++
			replacements := getReplacements(replacesFromTo, q.value)
			for k := range replacements {
				if !done[k] {
					done[k] = true
					item := &Item{
						value:    k,
						priority: calc(k, target) - q.depth,
						depth:    q.depth + 1,
					}
					heap.Push(&pq, item)
				}
			}
		}
		log.Println("Calcs done:", operations)
	} else {
		pq := make(PriorityQueue, 0)
		first := &Item{
			value:    target,
			priority: -len(target),
			depth:    0,
		}
		heap.Init(&pq)
		heap.Push(&pq, first)

		done := make(map[string]bool)
		done["e"] = true
		prevTime := time.Now()
		operations := 0

		// answer := make(chan )
		// log.Println("x")
	a:
		for pq.Len() > 0 {
			q := heap.Pop(&pq).(*Item)

			if q.value == "e" {
				part2 = q.depth
				log.Println("found", q.value)
				break
			}
			// queue = queue[1:]
			if time.Since(prevTime) > 2*time.Second {
				prevTime = time.Now()
				log.Println(q.depth, calc(q.value, target), len(pq), len(done), q.value)
				// log.Println(q.depth, q.value)
			}
			operations += 1
			replacements := getReplacements(replacesToFrom, q.value)
			// log.Println(q.value, replacements)
			for k := range replacements {
				if k == "e" {
					part2 = q.depth + 1
					// log.Println("found", k)
					break a
				}
				if !done[k] {
					done[k] = true
					item := &Item{
						value:    k,
						priority: -len(k),
						depth:    q.depth + 1,
					}
					heap.Push(&pq, item)
				}
			}
		}
		// log.Println("Calcs done:", operations)
	}
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// fmt.Println(day19("test.txt"))
	// fmt.Println(day19("test2.txt"))
	// fmt.Println(day19("test3.txt"))
	fmt.Println(day19("input.txt"))
}
