package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// isValid counts the number of valid chains of adapters recursively.
func isValid(num []int, k int) int {
	count := 0
	i := 0
	for _, v := range num {
		if v-i > 3 {
			return 0
		}
		i = v
	}
	count++
	a := make([]int, len(num))
	for i := k; i < len(num)-1; i++ {
		if num[i]-num[i-1] >= 3 {
			continue
		}
		copy(a[:i], num[:i])
		copy(a[i:], num[i+1:]) // Shift a[i+1:] left one index.
		a = a[:len(num)-1]
		count += isValid(a, i)
	}
	return int(count)
}

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)
	max := -1
	for scanner.Scan() {
		l := scanner.Text()
		n := 0
		fmt.Sscanf(l, "%d", &n)
		numbers = append(numbers, n)
		if n > max {
			max = n
		}
	}

	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, max+3)
	jolt := 0
	diff := make(map[int]int)
	for _, v := range numbers {
		if v == 0 {
			continue
		}
		d := v - jolt
		diff[d]++
		jolt = v
	}

	// oneRuns stores the length of runs of ones we see.
	// Prelude to using non-Bruteforce
	oneRuns := make([]int, 0)
	isRun := false
	j := 0
	for i := 1; i < len(numbers); i++ {
		if isRun {
			if numbers[i]-numbers[i-1] > 1 {
				isRun = false
				oneRuns = append(oneRuns, i-j)
				j = -1
			}
		} else {
			if numbers[i]-numbers[i-1] == 1 {
				isRun = true
				j = i
				continue
			}
		}
	}

	count := 0
	bruteFoce := false
	// Use bruteforce to confirm if answer is correct,
	// and to look for patterns
	if bruteFoce {
		count = isValid(numbers, 1)
	} else {
		count = 1
		for _, d := range oneRuns {
			count *= m(d)
		}
	}

	return diff[1] * diff[3], count
}

// multipliers
var mult = []int{
	0, //0
	1, //1
	2, //2
	4, //3
	7, //4
	13, 24, 44, 81, 149, 274, 504,
}

func m(i int) int {
	// This is the progression, but we'll skip that in favour of precalculated mult
	//((((((4*2-1)*2-1)*2-2)*2-4)*2-7)*2-13)*2-(13+7+4)

	return mult[i]
}

func main() {
	log.Println(do("test.txt"))
	log.Println(do("test2.txt"))
	log.Println(do("test3.txt"))
	log.Println(do("input.txt"))
}
