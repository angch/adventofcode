package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	checksum := 0
	twos, threes := 0, 0
	for scanner.Scan() {
		t := scanner.Text()

		counts := make(map[rune]int, 0)
		for _, v := range t {
			counts[v]++
		}
		two, three := 0, 0
		for _, v := range counts {
			switch v {
			case 2:
				two++
			case 3:
				three++
			}
		}
		if two > 0 {
			twos++
		}
		if three > 0 {
			threes++
		}
	}
	checksum = twos * threes
	fmt.Println(checksum)
}

func diff(a, b string) (int, []rune) {
	if len(a) != len(b) {
		log.Fatal("diff lengths")
	}
	same := make([]rune, 0)
	d := 0
	for k := range a {
		if a[k] != b[k] {
			d++
		} else {
			same = append(same, rune(a[k]))
		}
	}
	return d, same
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prev := make([]string, 0)

	for scanner.Scan() {
		t := scanner.Text()

		for _, v := range prev {
			d, same := diff(t, v)
			//log.Println(t, v, d, same)
			if d == 1 {
				fmt.Println(string(same))
			}
		}
		prev = append(prev, t)
	}
}
