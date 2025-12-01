package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func day1(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	n := 50
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		dir, diststr := t[0], t[1:]
		dist, _ := strconv.Atoi(diststr)
		old := part2
		switch dir {
		case 'R':
			n += dist
			for n >= 100 {
				n -= 100
				part2++
			}
		case 'L':
			// Damn off by one
			if n == 0 {
				part2--
			}
			n -= dist
			for n < 0 {
				n += 100
				part2++
			}
			if n == 0 {
				part2++
			}
		}
		if n == 0 {
			part1++
		}
		log.Printf("%s%3d %2d %1d %1d", string(dir), dist, n, part1, part2-old)
	}
	// if n == 0 {
	// 	part2++
	// }
	// fmt.Println(counts)
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day1("test.txt")
	fmt.Println(part1, part2)
	if part1 != 3 || part2 != 6 {
		log.Fatal("Test failed ", part1, part2)
	}
	part1, part2 = day1("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
	// 6225
}
