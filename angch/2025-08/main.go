package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"time"
)

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
	_ = circuits

	for scanner.Scan() {
		t := scanner.Text()
		x, y, z := 0, 0, 0
		fmt.Sscanf(t, "%d,%d,%d", &x, &y, &z)
		points = append(points, [3]int{x, y, z})
	}
	// log.Printf("%+v\n", points)

	type pairDist struct {
		p1 int
		p2 int
		d  int
	}
	shortest := []pairDist{}

	maxConn := len(points)
	if maxConn == 20 {
		maxConn = 10 // the test is shorter than full run
	}
	shortEntries := maxConn * 5

	t1 := time.Now()
	for k1, v1 := range points {
		for k2 := k1 + 1; k2 < len(points); k2++ {
			v2 := points[k2]

			dist1 := dist(v1, v2)
			if len(shortest) == shortEntries {
				if dist1 > shortest[len(shortest)-1].d {
					continue
				}
			}

			pd := pairDist{k1, k2, dist1}
			shortest = append(shortest, pd)
			slices.SortFunc(shortest, func(a, b pairDist) int {
				if a.d > b.d {
					return 1
				}
				if a.d < b.d {
					return -1
				}
				return 0
			})

			if len(shortest) > shortEntries {
				shortest = shortest[:shortEntries]
			}
		}
	}
	t2 := time.Now()

	// fmt.Printf("%+v\n", shortest)
	// log.Fatal("a")
	count := 0
	for k, pd := range shortest {
		// log.Printf("%d %d %d %f %+v\n", k, pd.p1, pd.p2, pd.d, pointInCircuit)

		// log.Printf("%d/%d: b4 %+v %+v %+v\n", k, count, circuits, pointInCircuit, len(points)-len(pointInCircuit))
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
			// fmt.Println("Same circuit", points[pd.p1], points[pd.p2])
		} else if pointInCircuit[pd.p1] > 0 && pointInCircuit[pd.p2] > 0 {
			// log.Fatal("Combine!")

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
	// log.Printf("circuits %+v %v\n", circuits, len(circuits[0]))
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
