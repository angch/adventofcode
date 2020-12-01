package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func process1(input []int, sum int) {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == sum {
				log.Println(input[i], input[j], input[i]*input[j])
				return
			}
		}
	}
}

func process2(input []int, sum int) int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			foo := input[i] + input[j]
			if foo > sum {
				continue
			}
			for k := j + 1; k < len(input); k++ {
				if foo+input[k] == sum {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0
}

func process2clever(input []int, sum int) int {
	seen := make(map[int]bool, len(input))

	// Make a lookup map of the input we know exists,
	// and avoid a O(N^3) triple nested loop scan
	for _, v := range input {
		seen[v] = true
	}

	for i, v := range input {
		left := sum - v
		if left <= 0 {
			continue
		}
		for j := i + 1; j < len(input); j++ {
			if seen[left-input[j]] {
				a, b, c := v, input[j], left-input[j]
				return a * b * c
			}
		}
	}
	return 0
}

// Seen
// 1721 t
//  979 t

// seen2
// [123] = 1721, 979
//

func main() {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456}

	sum := 2020
	process1(input, sum)

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	a := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		i, _ := strconv.Atoi(l)
		a = append(a, i)
	}
	// process1(a, sum)

	log.Println(process2(input, sum))
	log.Println(process2(a, sum))

	log.Println(process2clever(input, sum))
	log.Println(process2clever(a, sum))
}
