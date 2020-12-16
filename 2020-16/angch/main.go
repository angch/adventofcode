package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[string]map[int]bool)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		r, a1, a2, b1, b2 := "", 0, 0, 0, 0

		l1 := strings.Split(l, ":")
		fmt.Sscanf(l1[1], "%d-%d or %d-%d", &a1, &a2, &b1, &b2)
		r = l1[0]
		// log.Println(r)
		// log.Println(r, a1, a2, b1, b2)
		rng := make(map[int]bool)
		for i := a1; i <= a2; i++ {
			rng[i] = true
		}
		for i := b1; i <= b2; i++ {
			rng[i] = true
		}
		rules[r] = rng
	}
	mytick := []int{}
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		f := strings.Split(l, ",")
		if len(f) < 2 {
			continue
		}
		for _, v := range f {
			i, _ := strconv.Atoi(v)
			mytick = append(mytick, i)
		}
	}
	_ = mytick

	// log.Println(rules)
	invalidCount := 0
	checkTick := make([][]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		f := strings.Split(l, ",")
		if len(f) < 2 {
			continue
		}
		invalidtick := false
		// a:
		values := make([]int, 0)
		for _, v := range f {
			i, _ := strconv.Atoi(v)
			// log.Println(i)

			invalid := true
			for _, rule := range rules {
				if rule[i] {
					invalid = false
					break
				}
			}
			if invalid {
				// log.Println("invalid", i)
				invalidCount += i
				invalidtick = true
			}
			values = append(values, i)
		}
		if !invalidtick {
			// log.Println("add")
			checkTick = append(checkTick, values)
		}
	}
	ret1 = invalidCount
	// log.Println("mt", checkTick)

	possibleFields := make(map[string]map[int]bool)
	for k, rule := range rules {
		field := 0
		possibleFields[k] = make(map[int]bool)
	f:
		for ; field < len(mytick); field++ {
			for _, tick := range checkTick {
				if !rule[tick[field]] {
					continue f
				}
			}
			possibleFields[k][field] = true
		}
	}

	// log.Println(possibleFields)
	done := len(rules)

	fieldMap := make(map[string]int)
	for done > 0 {
	a:
		for k := range rules {
			if len(possibleFields[k]) == 1 {
				field := 0
				for v := range possibleFields[k] {
					field = v
				}

				// log.Println(k, field)
				fieldMap[k] = field

				for k2 := range possibleFields {
					if k2 == k {
						continue
					}
					delete(possibleFields[k2], field)
				}
				done--
				delete(rules, k)
				break a
			}
		}
	}
	// log.Println(fieldMap)
	ret2 = 1
	for k, v := range fieldMap {
		if strings.Contains(k, "departure") {
			// log.Println(v)
			ret2 *= mytick[v]
		}
	}

	return ret1, ret2
}
func main() {
	// log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
