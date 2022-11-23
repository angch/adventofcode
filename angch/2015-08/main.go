package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLit(a string) (int, int) {
	count := 0
	esc := false
	for k, v := range a {
		if k == 0 || k == len(a)-1 {
			continue
		}
		if esc {
			esc = false
			if v == 'x' {
				count -= 1
				continue
			}
			count++
			continue
		}
		if v == '\\' {
			esc = true
			continue
		}
		count++
	}
	return len(a), count
}

func countEnc(a string) (int, int) {
	count := 2
	for _, v := range a {
		switch v {
		case '"':
			count += 2
		case '\\':
			count += 2
		default:
			count++
		}
	}
	return len(a), count
}

func day8(file string) (int, int) {
	count1, count2 := 0, 0
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		c1, c2 := countLit(t)
		count1 += c1 - c2
		c1, c2 = countEnc(t)
		count2 += c2 - c1
	}
	return count1, count2
}

func main() {
	fmt.Println(day8("test.txt"))
	fmt.Println(day8("input.txt"))
}
