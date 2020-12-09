package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func do(fileName string, window int) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		n := 0
		fmt.Sscanf(l, "%d", &n)
		numbers = append(numbers, n)
	}

	invalid := make([]int, 0)
	for i := window; i < len(numbers); i++ {

		ok := false
	jj:
		for j := i - window; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					ok = true
					break jj
				}
			}
		}
		if !ok {
			invalid = append(invalid, numbers[i])
		}
	}

	inv := invalid[0]
	min, max := 0, 0
ii:
	for i := 0; i < len(numbers); i++ {
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if inv == sum {
				min, max = 9999999999, -1
				for k := i; k <= j; k++ {
					if numbers[k] < min {
						min = numbers[k]
					}
					if numbers[k] > max {
						max = numbers[k]
					}
				}
				break ii
			}
		}
	}
	ret2 := min + max

	return inv, ret2
}

func main() {
	log.Println(do("test.txt", 5))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt", 25))
}
