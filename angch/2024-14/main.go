package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"time"
)

type Robot struct {
	pos [2]int
	v   [2]int
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func day14(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	parse := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	robots := []Robot{}
	_ = robots
	for scanner.Scan() {
		t := scanner.Text()
		n := parse.FindStringSubmatch(t)
		// log.Println(n[1:])
		r := Robot{}
		x, _ := strconv.Atoi(n[1])
		y, _ := strconv.Atoi(n[2])
		dx, _ := strconv.Atoi(n[3])
		dy, _ := strconv.Atoi(n[4])
		r.pos[0] = x
		r.pos[1] = y
		r.v[0] = dx
		r.v[1] = dy
		_ = t
		robots = append(robots, r)
	}
	// fmt.Println(robots)

	w, h := 101, 103
	if file == "test.txt" {
		w, h = 11, 7
	}
	t := 100
	qcount := [4]int{}
	for k := range robots {
		newx := mod(robots[k].pos[0]+t*robots[k].v[0], w)
		newy := mod(robots[k].pos[1]+t*robots[k].v[1], h)
		q := 0
		// fmt.Println(k, robots[k].pos, robots[k].v, w/2, h/2)
		if newx == w/2 || newy == h/2 {
			// fmt.Println("skip", k)
			continue
		}
		if newx > w/2 {
			q++
		}
		if newy > h/2 {
			q += 2
		}
		// fmt.Println("q", k, q)
		qcount[q]++
	}
	// log.Println("test", -1%3)
	part1 = 1
	for _, v := range qcount {
		// log.Println(v)
		part1 *= v
	}
	if file == "test.txt" {
		return
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for t := range 50000 {
		// fmt.Println(t)
		board := make(map[[2]int]int)
		board2 := make(map[[2]int]int)
		heat := 0
		for k := range robots {
			newx := mod(robots[k].pos[0]+t*robots[k].v[0], w)
			newy := mod(robots[k].pos[1]+t*robots[k].v[1], h)
			board[[2]int{newx, newy}]++
			board2[[2]int{newx, newy}]++
			board2[[2]int{newx, newy}] *= 2

			// Boost surrounding area, so we can score the entire image
			// based on a "hotter" image with more likely real image
			// because a real image will have more robots in the same area
			for _, d := range dirs {
				x1, y2 := newx+d[0], newy+d[1]
				x1 = mod(x1, w)
				y2 = mod(y2, h)
				board2[[2]int{x1, y2}] *= 2
			}
		}
		for _, v := range board2 {
			heat += v
		}
		if heat > 3000 {
			fmt.Println(t, heat)
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					if board[[2]int{x, y}] > 0 {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			return
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	logperf := false
	if logperf {
		pf, _ := os.Create("default.pgo")
		err := pprof.StartCPUProfile(pf)
		if err != nil {
			log.Fatal(err)
		}
		defer pf.Close()
	}
	t1 := time.Now()
	part1, part2 := day14("test.txt")
	fmt.Println(part1, part2)
	if part1 != 12 || part2 != 0 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day14("input.txt"))
	if logperf {
		pprof.StopCPUProfile()
	}

	fmt.Println("Elapsed time:", time.Since(t1))
}
