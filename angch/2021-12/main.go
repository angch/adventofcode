package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct{ X, Y int }

type Edge struct {
	From, To string
}

// func visit0(nodeEdge map[string]map[string]bool, visited map[string]int, start, end string) [][]string {
// 	log.Println("visit", start, end)
// 	if start == end {
// 		log.Println(start, "-", end, "returning nil")
// 		return nil
// 	}

// 	paths := make([][]string, 0)

// 	for k, _ := range nodeEdge[start] {
// 		visited[k]++
// 		log.Println("k is", k)

// 		if k == end {
// 			// log.Println(start, end, "subpaths is", subpaths)
// 			path := []string{start, k}
// 			paths = append(paths, path)
// 			log.Println(start, end, "paths is", paths)
// 		} else {
// 			subpaths := visit0(nodeEdge, visited, k, end)
// 			if len(subpaths) > 0 {
// 				log.Println("paths is", paths)
// 				for _, p := range subpaths {
// 					path := []string{start}
// 					path = append(path, p...)
// 					// path = append(path, end)
// 					paths = append(paths, path)
// 					log.Println("path is", path)
// 				}
// 			}
// 		}
// 	}
// 	log.Println(start, "-", end, "returning", paths)

// 	return paths
// }

func slicecount(s []string, e string) int {
	c := 0
	for _, a := range s {
		if a == e {
			c++
		}
	}
	return c
}

func visit1(nodeEdge map[string]map[string]bool, start, end string) [][]string {
	// log.Println("visit2", start, end, paths)

	eval := make([][]string, 0)
	eval = [][]string{[]string{end}}
	out := make([][]string, 0)

	for len(eval) > 0 {
		path := eval[0]
		eval = eval[1:]

		end := path[0]
		for k, _ := range nodeEdge[end] {
			c := slicecount(path, k)
			if k != "start" && k != "end" {
				if strings.ToUpper(k) != k {
					// Small cave
					if c > 0 {
						// fmt.Println(" skip", k, path, c)
						continue
					}

				} else {
					// if c > 20 {
					// 	continue
					// }
				}
			} else {
				if c > 0 {
					continue
				}
			}
			path2 := []string{k}
			path2 = append(path2, path...)

			if k == start {
				out = append(out, path2)
			}
			eval = append(eval, path2)
		}
	}

	return out
}

func visit2(nodeEdge map[string]map[string]bool, start, end string) [][]string {
	// log.Println("visit2", start, end, paths)

	eval := make([][]string, 0)
	eval = [][]string{[]string{end}}
	out := make([][]string, 0)

	for len(eval) > 0 {
		path := eval[0]
		eval = eval[1:]

		end := path[0]
		for k, _ := range nodeEdge[end] {
			c := slicecount(path, k)
			if k != "start" && k != "end" {
				if strings.ToUpper(k) != k {
					// Small cave
					if c > 0 {
						// fmt.Println(" skip", k, path, c)
						// hack hack
						smallcount := make(map[string]int)
						hasSmall := false
						for _, p := range path {
							if strings.ToUpper(p) != p {
								smallcount[p]++
								if smallcount[p] > 1 {
									hasSmall = true
									break
								}
							}
						}
						if hasSmall {
							continue
						}
						// continue
					}

				} else {
					// if c > 20 {
					// 	continue
					// }
				}
			} else {
				if c > 0 {
					continue
				}
			}
			path2 := []string{k}
			path2 = append(path2, path...)

			if k == start {
				out = append(out, path2)
			}
			eval = append(eval, path2)
		}
	}

	return out
}

func day12(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	edges := make(map[Edge]bool)
	nodeEdge := make(map[string]map[string]bool)
	nodeEdge2 := make(map[string]map[string]bool)

	for scanner.Scan() {
		t := scanner.Text()
		a := strings.Split(t, "-")
		e := Edge{a[0], a[1]}
		edges[e] = true
		_, ok := nodeEdge[a[0]]
		if !ok {
			nodeEdge[a[0]] = make(map[string]bool)
		}
		_, ok = nodeEdge[a[1]]
		if !ok {
			nodeEdge[a[1]] = make(map[string]bool)
		}
		_, ok = nodeEdge2[a[1]]
		if !ok {
			nodeEdge2[a[1]] = make(map[string]bool)
		}
		_, ok = nodeEdge2[a[0]]
		if !ok {
			nodeEdge2[a[0]] = make(map[string]bool)
		}

		nodeEdge[a[0]][a[1]] = true

		_, ok = nodeEdge2[a[1]]
		if !ok {
			nodeEdge2[a[1]] = make(map[string]bool)
		}
		nodeEdge2[a[1]][a[0]] = true
		if a[1] != "end" {
			nodeEdge[a[0]][a[1]] = true
		}

		nodeEdge2[a[0]][a[1]] = true
	}
	// fmt.Println(edges)
	// fmt.Println(nodeEdge)
	// visited := make(map[string]int)
	// a := visit(nodeEdge, visited, "start", "end")
	visited := make(map[string]bool)
	visited["end"] = true
	part1 := make(chan int)
	part2 := make(chan int)
	go func() {
		a := visit1(nodeEdge2, "start", "end")
		part1 <- len(a)
	}()
	go func() {
		a := visit2(nodeEdge2, "start", "end")
		part2 <- len(a)
	}()
	// fmt.Println(a)

	fmt.Println("Part 1", <-part1)
	fmt.Println("Part 2", <-part2)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// day12("test.txt")
	// day12("test2.txt")
	day12("input.txt")
}
