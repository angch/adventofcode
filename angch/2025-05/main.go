package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	spans := [][2]int{}
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		if t == "" {
			break
		}
		l_, r_, ok := strings.Cut(t, "-")
		_ = ok
		l, _ := strconv.Atoi(l_)
		r, _ := strconv.Atoi(r_)
		spans = append(spans, [2]int{l, r})
	}
	// log.Println(spans)

	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		i, _ := strconv.Atoi(t)
		for _, v := range spans {
			if i >= v[0] && i <= v[1] {
				part1++
				break
			}
		}
	}

	s2 := span.NewSpans[bool]()
	for _, v := range spans {
		s2 = s2.AddCompress(v[0], v[1], true)
		log.Printf("%+v\n", s2)
	}
	for _, v := range s2 {
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
