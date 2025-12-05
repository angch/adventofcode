package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/angch/adventofcode/angch/span"
)

func day5(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	spans := span.NewSpans[bool]()
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		l, r := 0, 0
		fmt.Sscanf(t, "%d-%d", &l, &r)
		spans = spans.AddCompress(l, r, true)
	}

	for scanner.Scan() {
		t := scanner.Text()
		i, _ := strconv.Atoi(t)
		if spans.Contains(i) {
			part1++
		}
	}

	for _, v := range spans {
		part2 += v.To - v.From + 1
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day5("test.txt")
	fmt.Println(part1, part2)
	if part1 != 3 || part2 != 14 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day5("input.txt")
	fmt.Println(part1, part2)
	fmt.Println("Elapsed time:", time.Since(t1))
}
