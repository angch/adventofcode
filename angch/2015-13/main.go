package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func sumHappy(guests []string, happy map[string]map[string]int) int {
	happiness := 0

	for k, v := range guests {
		if k == 0 {
			happiness += happy[v][guests[len(guests)-1]]
			happiness += happy[v][guests[k+1]]
		} else if k == len(guests)-1 {
			happiness += happy[v][guests[k-1]]
			happiness += happy[v][guests[0]]
		} else {
			happiness += happy[v][guests[k-1]]
			happiness += happy[v][guests[k+1]]
		}
	}

	return happiness
}

func day13(file string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	happiness := make(map[string]map[string]int)
	guests := []string{}
	for scanner.Scan() {
		t := scanner.Text()
		a, gain, happy, b := "", "", 0, ""
		fmt.Sscanf(t, "%s would %s %d happiness units by sitting next to %s.", &a, &gain, &happy, &b)
		b = strings.TrimSuffix(b, ".")
		if gain == "lose" {
			happy = -happy
		}
		_, ok := happiness[a]
		if !ok {
			happiness[a] = make(map[string]int)
			guests = append(guests, a)
		}
		happiness[a][b] = happy
	}

	evalme := [][]string{
		{guests[0]},
	}

	combination := 0
	bestSoln := []string{}
	bestHappiness := -999999999

	for len(evalme) > 0 {
		eval := evalme[len(evalme)-1]
		evalme = evalme[:len(evalme)-1]

		if len(eval) == len(guests) {
			// // log.Println(eval)
			h := sumHappy(eval, happiness)
			// log.Println(h, happiness)
			if bestHappiness < h {
				bestHappiness = h
				bestSoln = eval
			}
			combination++
			continue
		}
		contains := make(map[string]bool)
		for _, v := range eval {
			contains[v] = true
		}

		for _, d := range guests {
			if !contains[d] {
				// Go gotcha: you need a deep copy of the slice
				k := make([]string, len(eval)+1)
				copy(k, eval)
				k[len(eval)] = d

				evalme = append(evalme, k)
			}
		}
	}
	log.Println(bestSoln)
	part1 := bestHappiness

	guests = append(guests, "me")
	evalme = [][]string{
		{guests[0]},
	}

	bestSoln = []string{}
	bestHappiness = -999999999

	for len(evalme) > 0 {
		eval := evalme[len(evalme)-1]
		evalme = evalme[:len(evalme)-1]

		if len(eval) == len(guests) {
			// // log.Println(eval)
			h := sumHappy(eval, happiness)
			// log.Println(h, happiness)
			if bestHappiness < h {
				bestHappiness = h
				bestSoln = eval
			}
			combination++
			continue
		}
		contains := make(map[string]bool)
		for _, v := range eval {
			contains[v] = true
		}

		for _, d := range guests {
			if !contains[d] {
				// Go gotcha: you need a deep copy of the slice
				k := make([]string, len(eval)+1)
				copy(k, eval)
				k[len(eval)] = d

				evalme = append(evalme, k)
			}
		}
	}
	log.Println(bestSoln)

	return part1, bestHappiness
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day13("test.txt"))
	fmt.Println(day13("input.txt"))
}
