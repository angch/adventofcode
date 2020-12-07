package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := make(map[string]map[string]int)

	for scanner.Scan() {
		l := scanner.Text()
		l = strings.ReplaceAll(l, " bags", "")
		l = strings.ReplaceAll(l, " bag", "")
		l = strings.ReplaceAll(l, ".", "")
		sp := strings.Split(l, " contain ")

		right := strings.Split(sp[1], ", ")

		items := make(map[string]int)
		for _, v := range right {
			if v == "no other" {
				continue
			}
			num, s1, s2 := 0, "", ""
			fmt.Sscanf(v, "%d %s %s", &num, &s1, &s2)
			items[s1+" "+s2] = num
		}
		rules[sp[0]] = items
	}

	canContain := make(map[string]int)
	canContain["shiny gold"] = 1
a:
	for k, v := range rules {
		for k2, v2 := range canContain {
			if v[k2] >= v2 && canContain[k] == 0 {
				canContain[k] = v2
				goto a
			}
		}
	}
	return len(canContain) - 1, countme(rules, "shiny gold") - 1
}

func countme(rules map[string]map[string]int, c string) int {
	count := 0
	for k, v := range rules[c] {
		count += v * countme(rules, k)
	}
	return count + 1
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
