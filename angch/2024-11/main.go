package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// func split(i int) (int, int) {
// 	counts := 0
// 	mult := 1
// 	for j := i; j > 0; j /= 10 {
// 		counts++
// 		mult *= 10
// 	}
// 	if counts%2 == 1 {
// 		return i, 0
// 	}

// 	return l, r
// }

var maxl = 0

var cycles = map[int][]int{
	0:     {1},       // -1 1
	1:     {2024},    // -2 1
	2024:  {20, 24},  // -3 2
	20:    {2, 0},    // -4 4
	24:    {2, 4},    // -4 4
	2:     {4048},    // -6
	4:     {8096},    // -7
	4048:  {40, 48},  // -8
	8096:  {80, 96},  // -9
	40:    {4, 0},    // -10
	48:    {4, 8},    // -10
	80:    {8, 0},    // -10
	96:    {9, 6},    // -10
	8:     {16192},   // -12
	9:     {18144},   // -13
	16192: {161, 92}, // -14
	18144: {181, 44}, // -15
	161:   {16, 1},   // -16
	92:    {9, 2},    // -16

	44: {4, 4},  // -16
	16: {32384}, // -18
	18: {36336}, // -19

}

// 20 24
// 2 0 2 4
// 0 0 0 0

func dostone(stones *list.List) {
	for i := stones.Front(); i != nil; i = i.Next() {
		if i.Value == 0 {
			i.Value = 1
			continue
		}

		ten := fmt.Sprintf("%d", i.Value.(int))
		isEvenTen := len(ten)%2 == 0
		if isEvenTen {
			l, _ := strconv.Atoi(ten[:len(ten)/2])
			r, _ := strconv.Atoi(ten[len(ten)/2:])
			i.Value = l
			i = stones.InsertAfter(r, i)
			if len(ten) > maxl {
				maxl = len(ten)
				fmt.Println("Max length", maxl)
			}
			continue
		}
		i.Value = i.Value.(int) * 2024
	}
}

func dostone2(stones map[int]uint) map[int]uint {
	out := make(map[int]uint)
	for stone, count := range stones {
		if stone == 0 {
			out[1] += count
			continue
		}

		ten := fmt.Sprintf("%d", stone)
		isEvenTen := len(ten)%2 == 0
		if isEvenTen {
			l, _ := strconv.Atoi(ten[:len(ten)/2])
			r, _ := strconv.Atoi(ten[len(ten)/2:])
			out[l] += count
			out[r] += count
			continue
		}

		out[stone*2024] += count
	}
	return out
}

func day11(file string, parttwo bool) (part1, part2 uint) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	stones := list.New()
	stones2 := make(map[int]uint)
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		for _, v := range f {
			i, _ := strconv.Atoi(v)
			stones.PushBack(i)
			stones2[i]++
		}
	}

	for times := range 25 {
		// dostone(stones)
		stones2 = dostone2(stones2)
		fmt.Println(times+1, len(stones2), maxl)
	}
	// part1 = len(stones2)
	for _, v := range stones2 {
		part1 += v

	}
	// if !parttwo {
	// 	return
	// }
	for times := range 75 - 25 {
		// dostone(stones)
		stones2 = dostone2(stones2)
		fmt.Println(times+26, len(stones2), maxl)
	}
	for _, v := range stones2 {
		part2 += v
	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day11("test.txt", false)
	fmt.Println(part1, part2)
	if part1 != 55312 {
		log.Fatal("Test failed ", part1, part2)
	}
	// too low  207957094518723
	// too low  207961465400585
	// too high 315840075824712
	fmt.Println(day11("input.txt", true))
	fmt.Println("Elapsed time:", time.Since(t1))
}
