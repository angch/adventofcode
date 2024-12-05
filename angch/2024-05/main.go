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
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		p1, p2 := 0, 0
		_, _ = fmt.Sscanf(t, "%d|%d", &p1, &p2)
		rules[p1] = append(rules[p1], p2)
	}

	updates := [][]int{}
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Split(t, ",")
		u := make([]int, 0, 4)
		for _, v := range f {
			i, _ := strconv.Atoi(v)
			u = append(u, i)
		}
		updates = append(updates, u)
	}

	for _, update := range updates {
		tries := 0
	a:
		for ; ; tries++ {
			printed := make(map[int]int)
			page := 0

			for k1, v := range update {
				for _, p := range rules[v] {
					if printed[p] > 0 {
						swap1 := printed[p] - 1
						update[swap1], update[k1] = update[k1], update[swap1]
						continue a
					}
				}
				page++
				printed[v] = page
			}
			break
		}
		mid := update[len(update)/2]
		if tries == 0 {
			part1 += mid
		} else {
			part2 += mid
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
