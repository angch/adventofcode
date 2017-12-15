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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	count2 := 0
	for scanner.Scan() {
		t := scanner.Text()
		count2++
		words := strings.Split(t, "\t")
		min := 99999
		max := -999999
		log.Println(words)
		for _, w := range words {
			i, _ := strconv.Atoi(w)
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
			log.Println(min, max)
		}
		diff := max - min
		count += diff

	}
	fmt.Println(count, count2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
