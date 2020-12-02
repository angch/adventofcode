package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseAndCheck(input string) (int, int, rune, string, bool, bool) {
	from, to, c, pass, count := 0, 0, ' ', "", 0
	fmt.Sscanf(input, "%d-%d %c: %s\n", &from, &to, &c, &pass)

	for _, v := range pass {
		if v == c {
			count++
		}
	}
	valid1 := count >= from && count <= to
	if from-1 < 0 || to-1 < 0 || from-1 >= len(pass) || to-1 >= len(pass) {
		// Yes, this is an assert, Ivan.
		log.Fatal("Out of bounds")
	}
	valid2 := (pass[from-1] == byte(c)) != (pass[to-1] == byte(c))

	return from, to, c, pass, valid1, valid2
}

func main() {
	// Test inputs
	inputs := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	for _, input := range inputs {
		log.Println(parseAndCheck(input))
	}

	// Actual inputs
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count1, count2 := 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		_, _, _, _, valid1, valid2 := parseAndCheck(l)
		if valid1 {
			count1++
		}
		if valid2 {
			count2++
		}
	}
	log.Println(count1, count2)
}
