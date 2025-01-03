package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func day10(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	mm := map[int]map[[2]int]bool{}

	y := 0
	for scanner.Scan() {
		t := scanner.Text()

		for x, k := range t {
			h := int(k - '0')
			_, ok := mm[h]
			if !ok {
				mm[h] = make(map[[2]int]bool)
			}
			mm[h][[2]int{x, y}] = true
		}
		y++
	}

	for v := range mm[0] {
		paths := map[[2]int]int{v: 1}
		for h := 1; h <= 9; h++ {
			npaths := map[[2]int]int{}
			for v := range paths {
				for p := range mm[h] {
					dx := max(p[0], v[0]) - min(p[0], v[0])
					dy := max(p[1], v[1]) - min(p[1], v[1])
					if dx+dy == 1 {
						npaths[p] += paths[v]
					}
				}
			}
			paths = npaths
		}
		part1 += len(paths)
		for _, v := range paths {
			part2 += v
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day10("test.txt")
	fmt.Println(part1, part2)
	if part1 != 36 || part2 != 81 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day10("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
