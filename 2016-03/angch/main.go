package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func isValid(a []int) bool {
	if len(a) < 3 {
		log.Fatal(a) // Yes, pickfire, this is an assert()
	}

	sort.Ints(a) // lazy
	// log.Println(a)
	return a[0]+a[1] > a[2]
}

func main() {
	log.Println(isValid([]int{5, 10, 25}) == false)

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a := make([]int, 3, 3)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Sscanf(l, "%d %d %d", &a[0], &a[1], &a[2])
		if isValid(a) {
			count++
		}
	}
	log.Println(count)

	// Part B
	file.Seek(0, 0) // Rewind
	scanner = bufio.NewScanner(file)

	b := make([][]int, 3, 3)
	for i := range b {
		b[i] = make([]int, 3, 3)
	}
	countb := 0
c:
	for {
		i := 0
		for ; i < 3; i++ {
			if !scanner.Scan() {
				break c
			}
			l := scanner.Text()
			fmt.Sscanf(l, "%d %d %d", &(b[0][i]), &(b[1][i]), &(b[2][i]))
		}

		for i--; i >= 0; i-- {
			if isValid(b[i]) {
				countb++
			}
		}
	}
	log.Println(countb) // 1826
}
