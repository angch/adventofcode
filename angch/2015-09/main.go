package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Route struct {
	From string
	To   string
	Dist int
}

type Path struct {
	Dest []string
	Dist int
}

func check(p Path, distance map[string]map[string]int) {
	prev := ""
	cost := 0
	for _, v := range p.Dest {
		if prev != "" {
			cost += distance[prev][v]
			log.Printf(" . %s to %s = %d (%d)", prev, v, distance[prev][v], cost)
		}
		prev = v
	}
	if cost != p.Dist {
		log.Println(p.Dest)
		log.Fatalf("Expect %d, got %d", cost, p.Dist)
	}
}

func day09(file string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	routes := make([]Route, 0)
	distance := make(map[string]map[string]int)
	dest := make(map[string]bool)
	dests := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		from, to, dist := "", "", 0
		fmt.Sscanf(t, "%s to %s = %d", &from, &to, &dist)
		// log.Println(from, to, dist)
		routes = append(routes, Route{from, to, dist})
		_, ok := distance[from]
		if !ok {
			distance[from] = make(map[string]int)
		}
		distance[from][to] = dist
		dest[to] = true

		_, ok = distance[to]
		if !ok {
			distance[to] = make(map[string]int)
		}
		distance[to][from] = dist
		dest[from] = true

	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Dist > routes[j].Dist
	})

	evalme := make([]Path, 0)
	for k := range dest {
		evalme = append(evalme, Path{[]string{k}, 0})
	}

	for k, v := range dest {
		if v {
			dests = append(dests, k)
		}
	}
	sort.Strings(dests)

	bestSoln := Path{Dist: 999999999}
	bestSoln2 := Path{Dist: 0}
	combination := 0

	for len(evalme) > 0 {
		eval := evalme[len(evalme)-1]
		evalme = evalme[:len(evalme)-1]

		// check(eval, distance)

		if len(eval.Dest) == len(dests) {
			// log.Println(eval)
			if bestSoln.Dist >= eval.Dist {
				bestSoln = eval
			}
			if bestSoln2.Dist <= eval.Dist {
				bestSoln2 = eval
			}
			combination++
			continue
		}
		contains := make(map[string]bool)
		for _, v := range eval.Dest {
			contains[v] = true
		}

		for _, d := range dests {
			if !contains[d] {
				newdist := distance[eval.Dest[len(eval.Dest)-1]][d]
				// log.Printf("New dist from %s to %s = %d\n", eval.Dest[len(eval.Dest)-1], d, newdist)

				// Go gotcha: you need a deep copy of the slice
				k := make([]string, len(eval.Dest)+1)
				copy(k, eval.Dest)
				k[len(eval.Dest)] = d

				newpath := Path{
					k,
					eval.Dist + newdist,
				}
				// log.Printf("Old dist is %d, new dist is %d\n", eval.Dist, newpath.Dist)
				evalme = append(evalme, newpath)
				// check(eval, distance)
				// log.Println("x")
				// check(newpath, distance)
			}
		}

		//FIXME
		// log.Println(eval)

	}
	// log.Println("Best is", bestSoln)
	// log.Println("Total combination", combination)
	cost := 0
	prev := ""
	for _, v := range bestSoln.Dest {
		if prev != "" {
			cost += distance[prev][v]
			// log.Printf("%s to %s = %d", prev, v, distance[prev][v])
		}
		prev = v
	}
	// log.Println("Recalc", cost)

	// dists := 0
	// for _, r := range routes {
	// 	dists += r.Dist
	// }
	// log.Println(dest)
	// dists += -(routes[0].Dist)
	// log.Println(len(distance))
	return bestSoln.Dist, bestSoln2.Dist
}

func main() {
	// fmt.Println(day09("test.txt"))
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day09("input.txt"))
}
