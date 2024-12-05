package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func day5(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rules := make(map[int][]int)
	printbefore := make(map[int][]int)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		p1, p2 := 0, 0
		fmt.Sscanf(t, "%d|%d", &p1, &p2)

		_, ok := rules[p1]
		if !ok {
			rules[p1] = make([]int, 0)
		}
		rules[p1] = append(rules[p1], p2)

		_, ok = printbefore[p2]
		if !ok {
			printbefore[p2] = make([]int, 0)
		}
		printbefore[p2] = append(printbefore[p2], p1)
	}
	// log.Println(rules)

	updates := [][]int{}
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Split(t, ",")
		u := []int{}
		for _, v := range f {
			i, err := strconv.Atoi(v)
			if err == nil {
				u = append(u, i)
			}
		}
		updates = append(updates, u)
	}
	// log.Println(updates)

	for _, update := range updates {
		ispart2 := false
	a:
		printed := make(map[int]int)
		good := true

		toprint := make(map[int]bool)
		for _, v := range update {
			toprint[v] = true
		}
		swap1, swap2 := 0, 0
		page := 0
	b:
		for k1, v := range update {
			if rules[v] != nil {
				for _, p := range rules[v] {
					if !toprint[p] {
						continue
					}
					if printed[p] > 0 {
						swap1 = printed[p]
						swap2 = k1
						good = false
						break b
					}
				}
			}
			page++
			printed[v] = page
		}
		if good {
			mid := len(update) / 2
			if !ispart2 {
				part1 += update[mid]
			} else {
				part2 += update[mid]
			}
		} else {
			// part2 reorder
			swap1--
			// log.Println(swap1, swap2)
			update[swap1], update[swap2] = update[swap2], update[swap1]
			// log.Println(update)
			ispart2 = true
			goto a // Meh
		}
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day5("test.txt")
	fmt.Println(part1, part2)
	if part1 != 143 || part2 != 123 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day5("input.txt"))
	fmt.Println("Elapsed time:", time.Since(t1))
}
