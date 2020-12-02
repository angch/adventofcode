package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(i string) (int, int, rune, string, bool, bool) {
	from, to := 0, 0
	var c rune
	var pass string
	fmt.Sscanf(i, "%d-%d %c: %s\n", &from, &to, &c, &pass)

	count := 0
	for _, v := range pass {
		if v == c {
			count++
		}
	}
	valid := false
	if count >= from && count <= to {
		valid = true
	}

	valid2 := false
	cnt := 0
	if pass[from-1] == byte(c) {
		cnt++
	}
	if pass[to-1] == byte(c) {
		cnt++
	}
	if cnt == 1 {
		valid2 = true
	}

	return from, to, c, pass, valid, valid2
}

func main() {
	inputs := []string{
		"1-3 a: abcde", "1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	for _, input := range inputs {
		log.Println(parse(input))
	}

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// a := make([]int, 0)
	count, count2 := 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		_, _, _, _, valid, valid2 := parse(l)
		if valid {
			count++
		}
		if valid2 {
			count2++
		}

	}
	log.Println(count, count2)
}
