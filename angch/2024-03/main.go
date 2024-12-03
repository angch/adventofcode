package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func day3(file string) (part1, part2 int) {
	f, _ := os.Open(file)
	defer f.Close()

	t, _ := io.ReadAll(f)
	t2 := string(t)
	reg := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	out := reg.FindAllStringSubmatch(t2, -1)
	on := 1
	for _, v := range out {
		if len(v) >= 4 {
			switch v[0] {
			case "don't()":
				on = 0
			case "do()":
				on = 1
			default:
				a, b := 0, 0
				fmt.Sscanf(v[2], "%d", &a)
				fmt.Sscanf(v[3], "%d", &b)
				part1 += a * b
				part2 += a * b * on
			}
		}
	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	part1, _ := day3("test.txt")
	_, part2 := day3("test2.txt")
	fmt.Println(part1, part2)
	if part1 != 161 || part2 != 48 {
		log.Fatal("Test failed ", part1, part2)
	}
	fmt.Println(day3("input.txt"))
}
