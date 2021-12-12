package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func day12(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	edge := make(map[int][]int)
	node := make(map[string]int)
	nodeid := 3
	node["start"] = 1
	node["end"] = 2
	isSmall := make(map[int]bool)
	isSmall[1] = true
	isSmall[2] = true
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "-")
		// from, to := a[0], a[1]
		_, ok := node[a[0]]
		if !ok {
			node[a[0]] = nodeid
			if a[0][0] >= 'a' && a[0][0] <= 'z' {
				isSmall[nodeid] = true
			} else {
				isSmall[nodeid] = false
			}
			nodeid++
		}
		_, ok = node[a[1]]
		if !ok {
			node[a[1]] = nodeid
			if a[1][0] >= 'a' && a[1][0] <= 'z' {
				isSmall[nodeid] = true
			} else {
				isSmall[nodeid] = false
			}
			nodeid++
		}

		from, to := node[a[0]], node[a[1]]
		_, ok = edge[from]
		if !ok {
			edge[from] = make([]int, 0, 1)
		}
		_, ok = edge[to]
		if !ok {
			edge[to] = make([]int, 0, 1)
		}
		edge[to] = append(edge[to], from)
		edge[from] = append(edge[from], to)
	}

	part1, part2 := 0, 0
	eval := [][]int{{node["end"]}}

	for len(eval) > 0 {
		var path []int
		path, eval = eval[len(eval)-1], eval[:len(eval)-1]
		slicecnt := make([]int, nodeid)
		hasSmall := false
		for _, v := range path {
			slicecnt[v]++
			if slicecnt[v] > 1 {
				if isSmall[v] {
					hasSmall = true
				}
			}
		}

		end := path[len(path)-1]
		for _, k := range edge[end] {
			c := slicecnt[k]
			if k != 1 && k != 2 {
				if isSmall[k] {
					// Small cave
					if c > 0 && hasSmall {
						continue
					}
				}
			} else if c > 0 {
				continue
			}

			if k == 1 {
				if !hasSmall {
					part1++
				}
				part2++
				continue
			}
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
	t1 := time.Now()
	day12("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
