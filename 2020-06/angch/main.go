package main

import (
	"bufio"
	"log"
	"os"
)

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	group := make([]string, 0)
	n := 0
	sum, sum2 := 0, 0
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			group = append(group, l)
		} else {
			n++
			ans := make(map[rune]int)
			for _, v := range group {
				for _, v2 := range v {
					ans[v2]++
				}
			}
			// log.Println(n, len(ans))
			sum += len(ans)
			for _, v := range ans {
				if v == len(group) {
					sum2++
				}
			}
			group = make([]string, 0)
		}
	}
	n++
	ans := make(map[rune]int)
	for _, v := range group {
		for _, v2 := range v {
			ans[v2]++
		}
	}
	// log.Println(n, len(ans))
	sum += len(ans)
	for _, v := range ans {
		if v == len(group) {
			sum2++
		}
	}

	return sum, sum2
}
func main() {
	log.Println(do("test.txt"))
	log.Println(do("input.txt"))
}
