package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//log.SetOutput(ioutil.Discard)
	test1a, test1b := advent06("test1.txt")
	fmt.Println(test1a, test1b)
	if test1a != 5 {
		log.Fatal("Fail part 1")
	}
	if test1b != 4 {
		log.Fatal("Fail part 2")
	}
	part1, part2 := advent06("input.txt")
	fmt.Println(part1, part2)
}

func key(ints []int) (sum string) {
	for _, i := range ints {
		sum += "," + strconv.Itoa(i)
	}
	return
}

func advent06(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	program := make([]int, 0)

	for scanner.Scan() {
		t := scanner.Text()
		words := strings.Split(t, "\t")
		for _, w := range words {
			i := 0
			fmt.Sscanf(w, "%d", &i)
			program = append(program, i)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	seen := make(map[string]int, 0)
	seen[key(program)] = 1
	count := 1

	fmt.Println(program)
	for {
		max, maxi := -1, -1
		for k, v := range program {
			if v > max {
				maxi, max = k, v
			}
		}
		program[maxi] = 0
		for i := 1; i <= max; i++ {
			program[(i+maxi)%len(program)]++
		}

		key_ := key(program)
		cnt2, ok := seen[key_]
		if ok {
			return count, count - cnt2
		}
		seen[key_] = count
		count++
	}
}
