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

func process2(input []int, sum int) {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			foo := input[i] + input[j]
			if foo > sum {
				continue
			}
			for k := j; k < len(input); k++ {
				if foo+input[k] == sum {
					log.Println(input[i], input[j], input[k], input[i]*input[j]*input[k])
					return
				}
			}
		}
	}
}

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
	process1(a, sum)
	process2(input, sum)
	process2(a, sum)
}
