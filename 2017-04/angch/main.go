package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
b:
	for scanner.Scan() {
		t := scanner.Text()
		count2++
		words := strings.Split(t, " ")
		a := make(map[string]bool, 0)
		for _, w := range words {
			_, ok := a[w]
			if ok {
				continue b
			}
			a[w] = true
		}
		count++
		//fmt.Printf("%#v\n", words)
	}
	fmt.Println(count, count2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
