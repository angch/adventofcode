package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type monkey struct {
	items       []int
	op1         string
	op2         int
	testdiv     int
	trueaction  int
	falseaction int
}

func day11(file string) (int, int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	monkeys := []monkey{}
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		m := monkey{}
		scanner.Scan()
		_, r, _ := strings.Cut(scanner.Text(), ": ")
		items := strings.Split(r, ", ")
		m.items = make([]int, len(items))
		for i, item := range items {
			fmt.Sscanf(item, "%d", &m.items[i])
		}
		scanner.Scan()
		_, r, _ = strings.Cut(scanner.Text(), " old ")
		words := strings.Split(r, " ")
		fmt.Sscanf(r, "%s %d", &m.op1, &m.op2)
		if len(words) > 1 && words[1] == "old" {
			m.op2 = -999
		}
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &m.testdiv)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), " If true: throw to monkey %d", &m.trueaction)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), " If false: throw to monkey %d", &m.falseaction)
		monkeys = append(monkeys, m)
	}
	part1 := make(chan int)
	go func() {
		part1 <- do(monkeys, 20, true)
	}()
	part2 := make(chan int)
	go func() {
		part2 <- do(monkeys, 10000, false)
	}()
	return <-part1, <-part2
}

func do(monkeysInput []monkey, r int, div3 bool) int {
	div := 1
	monkeys := make([]monkey, len(monkeysInput))
	for k, v := range monkeysInput {
		div *= v.testdiv
		monkeys[k] = monkeysInput[k]
	}
	monkeyCount := make([]int, len(monkeys))

	for rounds := 0; rounds < r; rounds++ {
		for k, m := range monkeys {
			for _, worry := range m.items {
				monkeyCount[k]++
				op2 := m.op2
				if m.op2 == -999 {
					op2 = worry
				}
				if m.op1 == "+" {
					worry += op2
				} else if m.op1 == "*" {
					worry *= op2
				}

				if div3 {
					worry /= 3
				} else {
					worry %= div
				}

				target := m.falseaction
				if worry%m.testdiv == 0 {
					target = m.trueaction
				}
				monkeys[target].items = append(monkeys[target].items, worry)
			}
			monkeys[k].items = []int{}
		}
	}
	sort.Ints(monkeyCount)
	return monkeyCount[len(monkeyCount)-1] * monkeyCount[len(monkeyCount)-2]
}

func main() {
	part1, part2 := day11("test.txt")
	fmt.Println(part1, part2)
	if part1 != 10605 || part2 != 2713310158 {
		panic("Wrong answer")
	}
	fmt.Println(day11("input.txt"))
}
