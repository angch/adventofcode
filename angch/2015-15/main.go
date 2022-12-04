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

func day15(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ingredients := [][]int{}
	for scanner.Scan() {
		t := scanner.Text()
		_, r, _ := strings.Cut(t, ": ")
		props := strings.Split(r, ", ")

		ingredientProp := []int{}
		for _, v := range props {
			prop := strings.Split(v, " ")
			i, _ := strconv.Atoi(prop[1])
			ingredientProp = append(ingredientProp, i)
		}
		ingredients = append(ingredients, ingredientProp)
	}
	// fmt.Println(ingredients)

	counts := 100
	amounts := make([]int, len(ingredients))
	best := 0
	best2 := 0
a:
	for {
		c := 0
		for i := 1; i < len(amounts); i++ {
			amounts[i]++
			if amounts[i] > 100 {
				amounts[i] = 0
				if i == len(amounts)-1 {
					break a
				}
			} else {
				break
			}
		}
		for i := 1; i < len(amounts); i++ {
			c += amounts[i]
		}
		amounts[0] = counts - c
		sc, cal := score(amounts, ingredients)
		if sc > best {
			best = sc
			// log.Println(amounts, sc)
		}
		if cal == 500 {
			if sc > best2 {
				best2 = sc
			}
		}
	}
	part1 = best
	part2 = best2

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day15("test.txt"))
	fmt.Println(day15("input.txt"))
}
