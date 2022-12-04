package main

import (
	"fmt"
	"log"
)

func hasInc(s string) bool {
	inc := 0
	prev := rune(0)
	for _, v := range s {
		if v == prev+1 {
			inc++
			if inc == 2 {
				return true
			}

		} else {
			inc = 0
		}
		prev = v
	}
	return false
}

func isValid(s string) bool {
	inc := 0
	hasPair, prevPair := 0, rune(0)
	hasInc := false
	prev := rune(0)
	for _, v := range s {
		switch v {
		case 'i', 'o', 'l':
			return false
		}
		if v == prevPair {
			hasPair++
			prevPair = 0
		} else {
			prevPair = v
		}

		if v == prev+1 {
			inc++
			if inc == 2 {
				hasInc = true
			}
		} else {
			inc = 0
		}
		prev = v
	}
	return hasPair >= 2 && hasInc
}

func incPassword(s []byte) []byte {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'z' {
			s[i] = 'a'
		} else {
			s[i]++
			break
		}
	}
	return s
}

func incValid(t string) string {
	test := 0
	s := []byte(t)
	for {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == 'z' {
				s[i] = 'a'
			} else {
				s[i]++
				break
			}
		}
		// fmt.Println(string(s))
		if isValid(string(s)) {
			return string(s)
		}
		test++
		if test > 1000000 {
			log.Fatal("no")
		}
	}
}

func day11(in string) (string, string) {
	var part1, part2 string
	part1 = incValid(in)
	part2 = incValid(part1)
	return part1, part2
}

func main() {
	// hijklmmn
	fmt.Println(day11("hijklmmn"))
	fmt.Println(day11("hepxcrrq"))
}
