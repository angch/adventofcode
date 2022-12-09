package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

var weapons = []item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []item{
	{"NoArmor", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []item{
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

type person struct {
	HP     int
	Damage int
	Armor  int
}

func Battle(person1, person2 person) bool {
	for {
		hit := person1.Damage - person2.Armor
		if hit < 1 {
			hit = 1
		}
		person2.HP -= hit
		if person2.HP <= 0 {
			return true
		}
		hit = person2.Damage - person1.Armor
		if hit < 1 {
			hit = 1
		}
		person1.HP -= hit
		if person1.HP <= 0 {
			return false
		}
	}
}

func day21(file string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	_, r, _ := strings.Cut(scanner.Text(), ": ")
	hp, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	_, r, _ = strings.Cut(scanner.Text(), ": ")
	damage, _ := strconv.Atoi(r)
	scanner.Scan()
	_, r, _ = strings.Cut(scanner.Text(), ": ")
	armor, _ := strconv.Atoi(r)
	boss := person{hp, damage, armor}

	hp1 := 100
	best, worst := 99999999, 0

	for _, v := range weapons {
		for _, v2 := range armors {
			mask := 0
			for mask < (1 << len(rings)) {
				bitcount := 0
				for m := mask; m > 0; m >>= 1 {
					if m&1 == 1 {
						bitcount++
					}
				}
				if bitcount > 2 {
					mask++
					continue
				}

				cost, d, a := v.Cost+v2.Cost, v.Damage+v2.Damage, v.Armor+v2.Armor
				for i, v3 := range rings {
					if mask&(1<<i) != 0 {
						cost += v3.Cost
						d += v3.Damage
						a += v3.Armor
					}
				}
				win := Battle(person{hp1, d, a}, boss)
				if win {
					if cost < best {
						best = cost
					}
				} else {
					if cost > worst {
						worst = cost
					}
				}
				mask++
			}
		}
	}

	return best, worst
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	fmt.Println(Battle(person{8, 5, 5}, person{12, 7, 2}))
	fmt.Println(Battle(person{100, 8, 3}, person{103, 9, 2}))

	fmt.Println(day21("test.txt"))
	fmt.Println(day21("input.txt"))
}
