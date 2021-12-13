package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const startId, endId = 1, 2

func day12(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	edge := make(map[int][]int)
	node := make(map[string]int)
	nodeid := endId + 1
	node["start"] = startId
	node["end"] = endId
	isSmall := make(map[int]bool)
	isSmall[startId] = true
	isSmall[endId] = true
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "-")
		for _, n := range a {
			// Enumerate all seen nodes
			nid, ok := node[n]
			if !ok {
				node[n] = nodeid
				isSmall[nodeid] = n[0] >= 'a' && n[0] <= 'z'
				nid = nodeid
				nodeid++
			}
			_, ok = edge[nid]
			if !ok {
				edge[nid] = make([]int, 0, 1)
			}
		}
		from, to := node[a[0]], node[a[1]]
		edge[to] = append(edge[to], from)
		edge[from] = append(edge[from], to)
	}

	part1, part2 := 0, 0
	// We start at the end
	eval := [][]int{{endId}}

	for len(eval) > 0 {
		var path []int
		path, eval = eval[len(eval)-1], eval[:len(eval)-1]
		visitedCount := make([]int, nodeid)
		hasSmall := false
		for _, v := range path {
			visitedCount[v]++
			if visitedCount[v] > 1 {
				if isSmall[v] {
					hasSmall = true
				}
			}
		}

		end := path[len(path)-1]
		for _, k := range edge[end] {
			c := visitedCount[k]
			if k != 1 && k != 2 {
				if isSmall[k] && c > 0 && hasSmall {
					continue
				}
			} else if c > 0 {
				continue
			}

			// ... and end at the start
			if k == startId {
				if !hasSmall {
					part1++
				}
				part2++
				continue
			}

			// path appended to eval needs to be a copy
			path2 := make([]int, len(path), len(path)+1)
			copy(path2, path)
			path2 = append(path2, k)
			eval = append(eval, path2)
		}
	}

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day12("test.txt")
	// day12("test2.txt")
	t1 := time.Now() // For the purpose of running on machines without `time` aka Windows
	day12("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
