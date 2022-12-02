package main

import (
	"bufio"
	"fmt"
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
	count1, count2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// elf := make([]int, 0)
	// cal := 0
	score := 0
	score2 := 0
	for scanner.Scan() {
		t := scanner.Text()
		lr := strings.Split(t, " ")
		l, r := lr[0], lr[1]

		roundscore := win[l+r] + shapescore[r]
		score += roundscore

		i, j := 0, 0
		switch r {
		case "X":
			i, j = 0, 0
		case "Y":
			i, j = 1, 3
		case "Z":
			i, j = 2, 6
		}
		// fmt.Println(i)

		resp := response[l][i]
		// log.Println(l, "x", r, roundscore, resp)
		score2 += shapescore[resp] + j

		// if t == "" {
		// 	elf = append(elf, -cal)
		// 	cal = 0
		// 	continue
		// }
		// c, _ := strconv.Atoi(t)
		// cal += c
	}
	count1 = score
	count2 = score2
	// _ = co
	return count1, count2
	// sort.Ints(elf)
	// count1 = elf[0]
	// count2 = elf[0] + elf[1] + elf[2]
}

func main() {
	fmt.Println(day2("test.txt"))
	fmt.Println(day2("input.txt"))
}
