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

	edge := make(map[string][]string)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "-")
		from, to := a[0], a[1]
		_, ok := edge[from]
		if !ok {
			edge[from] = make([]string, 0, 1)
		}
		_, ok = edge[to]
		if !ok {
			edge[to] = make([]string, 0, 1)
		}
		edge[to] = append(edge[to], from)
		edge[from] = append(edge[from], to)
	}

	part1, part2 := 0, 0
	eval := [][]string{{"end"}}

	for len(eval) > 0 {
		var path []string
		path, eval = eval[len(eval)-1], eval[:len(eval)-1]
		slicecnt := make(map[string]int)
		hasSmall := false
		for _, v := range path {
			slicecnt[v]++
			if slicecnt[v] > 1 {
				if v[0] >= 'a' && v[0] <= 'z' {
					hasSmall = true
				}
			}
		}

		end := path[0]
		for _, k := range edge[end] {
			c := slicecnt[k]
			if k != "start" && k != "end" {
				if k[0] >= 'a' && k[0] <= 'z' {
					// Small cave
					if c > 0 && hasSmall {
						continue
					}
				}
			} else if c > 0 {
				continue
			}
			path2 := append([]string{k}, path...)

			if k == "start" {
				if !hasSmall {
					part1++
				}
				part2++
				continue
			}
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
