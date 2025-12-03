package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func day3(file string) (part1, part2 int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		_ = t
		// log.Println(t)

		best := 0
		for i := 0; i < len(t); i++ {
			for j := i + 1; j < len(t); j++ {
				a, b, c := t[i]-'0', t[j]-'0', 0
				c = int(a*10 + b)
				if c > best {
					best = c
				}
			}
		}

		// c, _ = strconv.Atoi(string(a) + string(b))
		// log.Println(best)
		part1 += best

		// 12 digits
		digits := []byte(t)
		for len(digits) > 12 {
			// biggest := 0
			biggestIdx := -1
			biggestString := "0"

			for i := 0; i < len(digits); i++ {
				digits2 := make([]byte, len(digits))
				copy(digits2, digits)
				digits2 = append(digits2[:i], digits2[i+1:]...)
				// Cute, of course, it's too big to fit in int
				// d, err := strconv.Atoi(string(digits2))
				// if err != nil {
				// 	log.Fatal(err)
				// }
				if string(digits2) > biggestString {
					// biggest = d
					biggestIdx = i
					biggestString = string(digits2)
				}
			}
			digits = append(digits[:biggestIdx], digits[biggestIdx+1:]...)
		}
		// log.Println("xxx", string(digits))
		d, err := strconv.Atoi(string(digits))
		if err != nil {
			log.Fatal(err)
		}
		part2 += d

	}
	return
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t1 := time.Now()
	part1, part2 := day3("test.txt")
	fmt.Println(part1, part2)
	// 9223372036854775807
	if part1 != 357 || part2 != 3121910778619 {
		log.Fatal("Test failed ", part1, part2)
	}

	part1, part2 = day3("input.txt")
	fmt.Println(part1, part2)
	// 126961588362333
	// 172516781546707
	fmt.Println("Elapsed time:", time.Since(t1))
}
