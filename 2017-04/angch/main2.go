package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
			letters := strings.Split(w, "")
			sort.Strings(letters)

			w2 := strings.Join(letters, "")

			_, ok := a[w2]
			if ok {
				continue b
			}
			a[w2] = true
		}
		count++
		//fmt.Printf("%#v\n", words)
	}
	fmt.Println(count, count2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
