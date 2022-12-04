package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type aoc1 int

type aoc2 struct {
	v         int
	ignoreRed bool
}

func (a *aoc1) UnmarshalJSON(b []byte) error {
	// int
	i, err := strconv.Atoi(string(b))
	if err == nil {
		*a = aoc1(i)
		return nil
	}

	// Array
	k2 := make([]aoc1, 0)
	err = json.Unmarshal(b, &k2)
	if err == nil {
		sum := 0
		for _, v := range k2 {
			sum += int(v)
		}
		*a = aoc1(sum)
	}

	// object
	k1 := make(map[string]aoc1, 0)
	err = json.Unmarshal(b, &k1)
	if err == nil {
		sum := 0
		for _, v := range k1 {
			sum += int(v)
		}
		*a = aoc1(sum)
		return nil
	}
	return nil
}

func (a *aoc2) UnmarshalJSON(b []byte) error {
	if string(b) == "\"red\"" {
		a.v = 0
		a.ignoreRed = true
		return nil
	}

	// int
	i, err := strconv.Atoi(string(b))
	if err == nil {
		a.v = i
		return nil
	}

	// Array
	k2 := make([]aoc2, 0)
	err = json.Unmarshal(b, &k2)
	if err == nil {
		sum := 0
		for _, v := range k2 {
			sum += int(v.v)
		}
		a.v = sum
	}

	// object
	k1 := make(map[string]aoc2, 0)
	// log.Println("object", string(b))
	err = json.Unmarshal(b, &k1)
	if err == nil {
		sum := 0
		mult := 1
		for _, v := range k1 {
			if v.ignoreRed {
				mult = 0
			}
			sum += int(v.v)
		}
		a.v = sum * mult
		return nil
	}
	return nil
}

func parse(t []byte) int {
	// k1 := make(map[string]*json.RawMessage, 0)
	k2 := make([]float64, 0)
	err := json.Unmarshal(t, &k2)
	if err == nil {
		sum := 0
		for _, v := range k2 {
			sum += int(v)
		}
		return sum
	}

	// err = json.Unmarshal(t, &k1)
	// if err == nil {
	// 	sum := 0
	// 	for _, v := range k1 {
	// 	}
	// 	return sum
	// }

	return 0
}

func day12(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := []byte(scanner.Text())
		var p1 aoc1
		_ = json.Unmarshal(t, &p1)
		part1 += int(p1)

		var p2 aoc2
		_ = json.Unmarshal(t, &p2)
		part2 += int(p2.v)
	}

	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(day12("input.txt"))
}
