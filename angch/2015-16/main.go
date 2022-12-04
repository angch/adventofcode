package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func score(amounts []int, ingredients [][]int) (int, int) {
	calories := 0
	sc := 1
	verbose := false
	if false && amounts[0] == 44 {
		verbose = true
	}
	for i := 0; i < 4; i++ {
		s := 0
		for j := 0; j < len(amounts); j++ {
			if verbose {
				log.Println("a * i = ", amounts[j], ingredients[j][i], amounts[j]*ingredients[j][i])
			}
			s += amounts[j] * ingredients[j][i]
		}
		if s < 0 {
			return 0, 0
		}
		if verbose {
			log.Println("i", i, "s", s)
		}
		sc *= s
	}
	for j := 0; j < len(amounts); j++ {
		calories += amounts[j] * ingredients[j][4]
	}
	if amounts[1] == 44 {
		// log.Fatal("x")
	}
	return sc, calories
}

func day16(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	want := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for scanner.Scan() {
		t := scanner.Text()
		l, r, _ := strings.Cut(t, ": ")
		props := strings.Split(r, ", ")

		nope1 := false
		nope2 := false
		for _, v := range props {
			prop := strings.Split(v, ": ")
			i, _ := strconv.Atoi(prop[1])
			k2, ok := want[prop[0]]
			if ok {
				if k2 != i {
					nope1 = true
				}

				switch prop[0] {
				case "cats", "trees":
					if k2 > i {
						nope2 = true
					}
				case "pomeranians", "goldfish":
					if k2 < i {
						nope2 = true
					}
				default:
					if k2 != i {
						nope2 = true
					}
				}

			}
		}
		if !nope1 {
			// log.Println("part1", l)
			fmt.Sscanf(l, "Sue %d", &part1)
		}
		if !nope2 && nope1 {
			// log.Println("part2", l)
			fmt.Sscanf(l, "Sue %d", &part2)
		}

	}
	// fmt.Println(ingredients)

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// fmt.Println(day15("test.txt"))
	fmt.Println(day16("input.txt"))
}
