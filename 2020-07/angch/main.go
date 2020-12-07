package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Foo struct {
	Num  int
	Item string
}

func do(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := make(map[string][]Foo)
	rules2 := make(map[string]map[string]int)

	for scanner.Scan() {
		l := scanner.Text()

		l = strings.ReplaceAll(l, " bags", "")
		l = strings.ReplaceAll(l, " bag", "")
		l = strings.ReplaceAll(l, ".", "")
		sp := strings.Split(l, " contain ")

		left := sp[0]
		right := strings.Split(sp[1], ", ")

		items := make([]Foo, 0)
		items2 := make(map[string]int)
		for _, v := range right {
			if v == "no other" {
				continue
			}
			num := 0
			s1, s2 := "", ""
			fmt.Sscanf(v, "%d %s %s", &num, &s1, &s2)
			s := s1 + " " + s2

			items = append(items, Foo{num, s})
			items2[s] = num
		}

		rules[left] = items
		rules2[left] = items2
	}

	canContain := make(map[string]int)
	canContain["shiny gold"] = 1
a:
	for k, v := range rules2 {
		c := false
		for k2, v2 := range canContain {
			if v[k2] >= v2 {
				c = true
			}
		}
		if c && canContain[k] == 0 {
			canContain[k] = 1
			goto a
		}
	}

	log.Println(rules, len(canContain)-1)

	count := countme(rules, "shiny gold")
	log.Println("Count ", count-1)

	return 0, 0
}

func countme(rules map[string][]Foo, c string) int {
	count := 0
	for _, v := range rules[c] {
		count += v.Num * countme(rules, v.Item)
		log.Println(c, "contains", v.Num, v.Item, count)
	}
	log.Println(" ", c, "contains ", count+1)
	return count + 1
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
