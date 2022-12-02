package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var win = map[string]int{
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

var shapescore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var response = map[string][]string{
	"A": {"Z", "X", "Y"},
	"B": {"X", "Y", "Z"},
	"C": {"Y", "Z", "X"},
}

func day2(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		lr := strings.Split(t, " ")
		l, r := lr[0], lr[1]

		roundscore := win[l+r] + shapescore[r]
		part1 += roundscore

		i, j := 0, 0
		switch r {
		case "X":
			i, j = 0, 0
		case "Y":
			i, j = 1, 3
		case "Z":
			i, j = 2, 6
		}
		resp := response[l][i]
		part2 += shapescore[resp] + j
	}
	return part1, part2
}

func main() {
	part1, part2 := day2("test.txt")
	if part1 != 15 || part2 != 12 {
		log.Fatal("Test failed")
	}
	fmt.Println(part1, part2)
	fmt.Println(day2("input.txt"))
}
