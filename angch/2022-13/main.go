package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func parse(t string) []any {
	foo := []any{}
	err := json.Unmarshal([]byte(t), &foo)
	if err != nil {
		log.Fatal(err)
	}
	return foo
}

var spaces = "                                             "

func compare(a, b []any, a3, b3 string, depth int) (bool, bool) {
	a1, _ := json.Marshal(a)
	b1, _ := json.Marshal(b)
	fmt.Printf("- Compare %s vs %s\n", string(a1), string(b1))
	tab := string(spaces[depth*2])

	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if i >= len(b) {
			fmt.Printf("%s - Right side ran out of items, so inputs are not in the right order\n", tab)
			return false, false
		}

		left, right := a[i], b[j]

		leftInt, leftIsInt := left.(float64) // Yes, not int
		rightInt, rightIsInt := right.(float64)
		if leftIsInt && rightIsInt {
			fmt.Printf("%s - Compare %v vs %v\n", tab, left, right)
			if leftInt < rightInt {
				fmt.Printf("%s   - Left side is smaller, so inputs are in the right order\n", tab)
				return true, false
			}
			if leftInt > rightInt {
				fmt.Printf("%s   - Right side is smaller, so inputs are not in the right order\n", tab)
				return false, false
			}
			continue
		}
		if !leftIsInt && !rightIsInt {
			// Both are lists
			f, same := compare(left.([]any), right.([]any), a3, b3, depth+1)
			if !same {
				return f, false
			}
			continue
		}

		// One or more is a list
		leftList := []any{}
		rightList := []any{}
		if leftIsInt {
			fmt.Printf("%s - Mixed types; convert left\n", tab)
			leftList = []any{left}
		} else {
			leftList = left.([]any)
		}
		if rightIsInt {
			fmt.Printf("%s - Mixed types; convert right\n", tab)
			rightList = []any{right}
		} else {
			rightList = right.([]any)
		}
		f, same := compare(leftList, rightList, a3, b3, depth+1)
		if !same {
			return f, false
		}
	}
	fmt.Printf("%s   - Left side ran out of items, so inputs are in the right order\n", tab)
	return true, false
}

func day13(file string) (int, int) {
	part1, part2 := 0, 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0

	two := parse("[[2]]")
	six := parse("[[6]]")
	two2, six2 := 1, 2

	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		if t == "" {
			continue
		}
		i++
		a := parse(t)
		scanner.Scan()
		t2 := scanner.Text()
		b := parse(t2)
		fmt.Printf("== Pair %d ==\n", i)
		comp, _ := compare(a, b, t, t2, 0)
		if comp {
			part1 += i
		}

		comp, _ = compare(six, a, t, t2, 0)
		if !comp {
			six2++
		}
		comp, _ = compare(six, b, t, t2, 0)
		if !comp {
			six2++
		}
		comp, _ = compare(two, a, t, t2, 0)
		if !comp {
			two2++
		}
		comp, _ = compare(two, b, t, t2, 0)
		if !comp {
			two2++
		}

		fmt.Printf("%s\n%s\n", t, t2)
		fmt.Println(comp)
		fmt.Println()
	}

	part2 = two2 * six2
	return part1, part2
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day13("test.txt")

	fmt.Println(part1, part2)
	if part1 != 13 || part2 != 140 {
		log.Fatal("Bad test")
	}
	fmt.Println(day13("input.txt")) // 6586 6462 5604 5985 7273 6536 6590 780 4346 6568
	// 20394 too high
}
