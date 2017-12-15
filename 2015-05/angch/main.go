package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"strings"

	"github.com/h2so5/goback/regexp"
)

func countVowels(s string) int {
	log.Println(s)
	re, err := regexp.Compile(`(?i)[aeiou]`)
	if err != nil {
		log.Fatal(err)
	}
	a := re.FindAllStringSubmatchIndex(s, -1)
	return len(a)
}

func countRepeats(s string) int {
	r := 0
	prev := '_'
	for _, c := range s {
		if c == prev {
			r++
		}
		prev = c
	}
	return r
}

func countPairsNoOverlap(s string) int {
	log.Println(s)
	re, err := regexp.Compile(`(..).*\k1`)
	if err != nil {
		log.Fatal(err)
	}
	a := re.FindAllStringSubmatchIndex(s, -1)
	return len(a)
}

func countRepeatNoOverlap(s string) int {
	r := 0
	prev := byte('_')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == prev {
			r++
			i++
			for ; i < len(s) && s[i] == prev; i++ {

			}
		}
		prev = c
	}
	return r
}

func countMids(s string) int {
	r := 0
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			r++
		}
	}
	return r
}

func countBad(s string) int {
	bads := []string{"ab", "cd", "pq", "xy"}
	c := 0
	for _, bad := range bads {
		if strings.Contains(s, bad) {
			c++
		}
	}
	return c
}

func advent05(s string) bool {
	good := false

	vowels := countVowels(s)
	if vowels >= 3 {
		good = true
	}

	rep := countRepeats(s)
	if good && rep >= 1 {
		good = true
	} else {
		good = false
	}

	bads := countBad(s)
	if bads > 0 {
		good = false
	}

	return good
}

func advent05b(s string) bool {
	good := false

	rep := countPairsNoOverlap(s)
	fmt.Println("pairs", s, rep)
	if rep >= 1 {
		good = true
	} else {
		good = false
	}

	mids := countMids(s)

	if good && mids < 1 {
		good = false
	} else {
		fmt.Println(s, mids)

	}

	return good
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	count := 0
	for _, s := range strings.Split(string(input), "\n") {
		fmt.Println(s)
		if advent05b(s) {
			count++
		}
	}
	fmt.Println(count)
}
