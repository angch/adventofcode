package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func countexit(mm map[[2]int]byte) (count int) {
	for _, v := range mm {
		if v == '#' {
			count++
		}
	}
	return
}

func day6(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mm := make(map[[2]int]byte)
	guard := [2]int{0, 0}

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, v := range t {
			if v != 'p' {
				mm[[2]int{x, y}] = byte(v)
			}
			if v == '^' {
				guard = [2]int{x, y}
			}
		}
		y++
	}
	// log.Println(mm)
	dir := [2]int{0, -1}

	orig := guard
	origd := dir

	visited := make(map[[2]int]bool)
	visited[guard] = true
	for {
		guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}
		if mm[guard2] == '.' || mm[guard2] == '^' {
			guard = guard2
			visited[guard] = true
		} else if mm[guard2] == '#' {
			// turn right
			dir = [2]int{dir[1], -dir[0]}
			dir = [2]int{dir[1], -dir[0]}
			dir = [2]int{dir[1], -dir[0]}
		} else {
			break
		}
	}
	part1 = len(visited)

a:
	for v := range visited {
		guard = orig
		dir = origd

		visited2 := make(map[[2]int]bool)
		looping := 0

		for {
			guard2 := [2]int{guard[0] + dir[0], guard[1] + dir[1]}
			if guard2 != v && mm[guard2] == '.' || mm[guard2] == '^' {
				guard = guard2

				if visited2[guard] {
					looping++
					if looping > 400 {
						part2++
						continue a
					}
				} else {
					looping = 0
				}
				visited2[guard] = true
			} else if mm[guard2] == '#' || guard2 == v {
				// turn right
				dir = [2]int{dir[1], -dir[0]}
				dir = [2]int{dir[1], -dir[0]}
				dir = [2]int{dir[1], -dir[0]}
			} else {
				break
			}
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day6("test.txt")
	fmt.Println(part1, part2)
	if part1 != 41 || part2 != 6 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day6("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
