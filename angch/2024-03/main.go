package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func day3(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	t2 := ""
	for scanner.Scan() {
		t := scanner.Text()
		t2 += t
	}
	out := reg.FindAllStringSubmatch(t2, -1)
	out2 := reg.FindAllStringSubmatchIndex(t2, -1)
	for _, v := range out {
		if len(v) >= 2 {
			a, b := 0, 0
			fmt.Sscanf(v[1], "%d", &a)
			fmt.Sscanf(v[2], "%d", &b)
			part1 += a * b
		}
	}
	log.Println(out, out2)
	reg2 := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	_ = reg2
	out3 := reg2.FindAllStringSubmatch(t2, -1)
	log.Println(out3)
	on := true
	for _, v := range out3 {
		if len(v) >= 4 {
			if on {
				a, b := 0, 0
				fmt.Sscanf(v[2], "%d", &a)
				fmt.Sscanf(v[3], "%d", &b)
				log.Printf("on %d %d %+v\n", a, b, v)
				part2 += a * b
			}
			if v[0] == "don't()" {
				on = false
			} else if v[0] == "do()" {
				on = true
			}
		}

	}

	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, part2 := day3("test.txt")

	_, part2 = day3("test2.txt")
	fmt.Println(part1, part2)
	if part1 != 161 || part2 != 48 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day3("input.txt"))
}
