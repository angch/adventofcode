package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type monkey struct {
	i           int
	items       []int
	op1         string
	op2         int
	testdiv     int
	trueaction  int
	falseaction int
}

func day11(file string) (int, int) {
	p1, _ := day11part1(file, 20)
	p2, _ := day11part1(file, 10000)
	// p2 := 0
	return p1, p2
}
func day11part1(file string, r int) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	monkeys := []monkey{}
	divSum := 1
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		m := monkey{}
		fmt.Sscanf(t, "Monkey %d:", &m.i)

		if t == "" {
			continue
		}

		scanner.Scan()
		t = scanner.Text()
		_, r, _ := strings.Cut(t, ": ")
		items := strings.Split(r, ", ")
		m.items = make([]int, len(items))
		for i, item := range items {
			fmt.Sscanf(item, "%d", &m.items[i])
		}
		scanner.Scan()
		t = scanner.Text()
		_, r, _ = strings.Cut(t, " old ")
		words := strings.Split(r, " ")
		fmt.Sscanf(r, "%s %d", &m.op1, &m.op2)
		if len(words) > 1 && words[1] == "old" {
			m.op2 = -999
		}
		scanner.Scan()
		t = scanner.Text()
		fmt.Sscanf(t, "  Test: divisible by %d", &m.testdiv)
		scanner.Scan()
		t = scanner.Text()
		fmt.Sscanf(t, " If true: throw to monkey %d", &m.trueaction)
		scanner.Scan()
		t = scanner.Text()
		fmt.Sscanf(t, " If false: throw to monkey %d", &m.falseaction)
		monkeys = append(monkeys, m)
		divSum *= m.testdiv
	}
	fmt.Printf("%+v\n", monkeys)
	monkeyCount := make([]int, len(monkeys))
	// zero := big.NewInt(0)
	verbose := true
	if r > 100 {
		verbose = false
	}
	for rounds := 0; rounds < r; rounds++ {
		for k, m := range monkeys {
			if verbose {
				fmt.Printf("Monkey %d\n", m.i)
			}
			for k, v := range m.items {
				monkeyCount[m.i]++
				if verbose {
					fmt.Println(" Monkey inspects item with level", v)
				}
				worry := v
				if m.op1 == "+" {
					if m.op2 == -999 {
						// worry.Add(&worry, &worry)
						worry += worry
					} else {
						worry += m.op2
						// 	o := big.NewInt(int(m.op2))
						// 	worry.Add(&worry, o)
					}
				} else if m.op1 == "*" {
					if m.op2 == -999 {
						worry *= worry
						// worry.Mul(&worry, &worry)
					} else {
						// o := big.NewInt(int(m.op2))
						// worry.Mul(&worry, o)
						worry *= m.op2
					}
				}
				if verbose {
					fmt.Printf("  Worry level is now %d for item %d\n", worry, k)
				}
				// three := big.NewInt(3)
				if r < 100 {
					worry /= 3
				}
				if verbose {
					fmt.Printf("  Worry level is now %d for item %d\n", worry, k)
				}
				target := m.falseaction

				if worry%m.testdiv == 0 {
					target = m.trueaction
				}

				worry %= divSum
				monkeys[target].items = append(monkeys[target].items, worry)
				if verbose {
					fmt.Printf("  Throwing item %d to monkey %d\n", worry, target)
				}
			}
			monkeys[k].items = []int{}
		}
		// fmt.Println("Round", rounds, monkeyCount)

	}

	for _, m := range monkeys {
		fmt.Println("Monkey", m.i, m.items)
	}
	fmt.Println(monkeyCount)
	sort.Ints(monkeyCount)
	part1 = monkeyCount[len(monkeyCount)-1] * monkeyCount[len(monkeyCount)-2]

	return part1, part2
}

// func day11part2(file string) (int, int) {
// 	part1, part2 := 0, 0
// 	f, _ := os.Open(file)
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	monkeys := []monkey{}
// 	for scanner.Scan() {
// 		t := scanner.Text()
// 		_ = t
// 		m := monkey{}
// 		fmt.Sscanf(t, "Monkey %d:", &m.i)

// 		if t == "" {
// 			continue
// 		}

// 		scanner.Scan()
// 		t = scanner.Text()
// 		_, r, _ := strings.Cut(t, ": ")
// 		items := strings.Split(r, ", ")
// 		m.items = make([]big.Int, len(items))
// 		for i, item := range items {
// 			fmt.Sscanf(item, "%d", &m.items[i])
// 		}
// 		scanner.Scan()
// 		t = scanner.Text()
// 		_, r, _ = strings.Cut(t, " old ")
// 		words := strings.Split(r, " ")
// 		fmt.Sscanf(r, "%s %d", &m.op1, &m.op2)
// 		if len(words) > 1 && words[1] == "old" {
// 			m.op2 = -999
// 		}
// 		scanner.Scan()
// 		t = scanner.Text()
// 		fmt.Sscanf(t, "  Test: divisible by %d", &m.testdiv)
// 		scanner.Scan()
// 		t = scanner.Text()
// 		fmt.Sscanf(t, " If true: throw to monkey %d", &m.trueaction)
// 		scanner.Scan()
// 		t = scanner.Text()
// 		fmt.Sscanf(t, " If false: throw to monkey %d", &m.falseaction)
// 		monkeys = append(monkeys, m)

// 	}
// 	fmt.Printf("%+v\n", monkeys)
// 	monkeyCount := make([]int, len(monkeys))
// 	zero := big.NewInt(0)
// 	for rounds := 0; rounds < 1000; rounds++ {
// 		for k, m := range monkeys {
// 			// fmt.Printf("Monkey %d\n", m.i)

// 			for k, v := range m.items {
// 				_ = k
// 				monkeyCount[m.i]++
// 				// fmt.Println(" Monkey inspects item with level", v)
// 				worry := v
// 				if m.op1 == "+" {
// 					if m.op2 == -999 {
// 						worry.Add(&worry, &worry)

// 					} else {

// 						o := big.NewInt(int(m.op2))
// 						worry.Add(&worry, o)
// 					}
// 				} else if m.op1 == "*" {
// 					if m.op2 == -999 {
// 						// worry *= worry
// 						worry.Mul(&worry, &worry)
// 					} else {
// 						o := big.NewInt(int(m.op2))
// 						worry.Mul(&worry, o)
// 						// worry *= uint(m.op2)
// 					}
// 				}
// 				// fmt.Printf("  Worry level is now %d for item %d\n", worry, k)
// 				three := big.NewInt(3)
// 				worry.Div(&worry, three)
// 				// fmt.Printf("  Worry level is now %d for item %d\n", worry, k)
// 				target := m.falseaction
// 				td := big.NewInt(m.testdiv)
// 				td.Mod(&worry, td)
// 				if td.Cmp(zero) == 0 {
// 					target = m.trueaction
// 				}
// 				monkeys[target].items = append(monkeys[target].items, worry)
// 				// fmt.Printf("  Throwing item %d to monkey %d\n", worry, target)
// 			}
// 			monkeys[k].items = []big.Int{}

// 		}
// 		fmt.Println(rounds, monkeyCount)
// 	}

// 	for _, m := range monkeys {
// 		fmt.Println("Monkey", m.i, m.items)
// 	}
// 	fmt.Println(monkeyCount)
// 	sort.Ints(monkeyCount)
// 	part1 = monkeyCount[len(monkeyCount)-1] * monkeyCount[len(monkeyCount)-2]

//		return part1, part2
//	}
func main() {
	part1, part2 := day11("test.txt")
	fmt.Println(part1, part2)
	fmt.Println(day11("input.txt"))
}
