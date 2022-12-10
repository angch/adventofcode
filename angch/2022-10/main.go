package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord [2]int

func mkinc(X *int, V *int, i int) func() int {
	return func() int {
		*X += i
		return i
	}
}
func mkrecord(X *int, V *int, i int) func() int {
	fmt.Println("mkrecord", i)
	return func() int {
		fmt.Println("Record at cycle before", i, *V, *X)
		*V += *X * i
		fmt.Println("Record at cycle", i, *V, *X)
		return *X
	}
}
func day10(file string) (int, int) {
	part1, part2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	X := map[int]int{}
	// V := 0
	cycle := 0
	trigger := make(map[int]bool)
	record := []int{
		20, 60, 100, 140, 180, 220,
	}
	for _, v := range record {
		trigger[v] = true
	}
	X[0] = 1
	cycle = 0

	fmt.Println("Trigger", trigger)
	for scanner.Scan() {
		t := scanner.Text()
		words := strings.Split(t, " ")
		n := 0
		if len(words) > 1 {
			n, _ = strconv.Atoi(words[1])
		}
		// cycle++
		// X[cycle] += X[cycle-1]
		prev := X[cycle]

		// fmt.Println("trigger", trigger, cycle, X, n, words[0])
		switch words[0] {
		case "noop":
			cycle += 1
			X[cycle] = prev
		case "addx":
			cycle += 2
			X[cycle-1] = prev
			X[cycle] = prev + n
			//   X[cycle+1]
		}
		// cycle++

		fmt.Println("AfterCycle", cycle, X[cycle], X[cycle]-X[cycle-1])
	}

	for cycle < 220 {
		cycle++
		X[cycle] += X[cycle-1]
		fmt.Println("AfterCycle", cycle, X[cycle], X[cycle]-X[cycle-1])
	}
	for k, v := range record {
		fmt.Println("Record", k, v, X[v-1])
		part1 += v * X[v-1]
	}

	grid := map[coord]rune{}
	for i := 0; i <= 240; i++ {
		y, x := i/40, i%40
		if x == 0 {
			fmt.Println()
		}
		a := '.'
		diff := X[i] - x
		if diff < 0 {
			diff = -diff
		}
		if diff < 2 {
			a = '#'
		}
		grid[coord{y, x}] = a
		fmt.Print(string(a))

	}
	fmt.Println()
	// part1 = V
	return part1, part2
}

func main() {
	// part1, part2 := day10("test.txt")
	// fmt.Println(part1, part2)
	// if part1 != 13 || part2 != 1 {
	// 	log.Fatal("Test failed")
	// }
	part1, part2 := day10("test2.txt")
	fmt.Println(part1, part2)
	fmt.Println(day10("input.txt"))
}
