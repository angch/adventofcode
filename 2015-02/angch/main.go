package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func advent02(input string) (int, int) {
	fmt.Println("input", input)
	s := strings.Split(input, "x")
	i := make([]int, len(s))
	max := 0
	for k, v := range s {
		i[k], _ = strconv.Atoi(v)
		if i[k] > max {
			max = i[k]
		}
	}
	area := 2 * (i[0]*i[1] + i[1]*i[2] + i[0]*i[2])
	sort.Ints(i)
	length := 2 * (i[0] + i[1])
	length += i[0] * i[1] * i[2]

	fmt.Println(i)
	area += i[0] * i[1]
	return area, length
}

func main() {
	test1, test2 := advent02("2x3x4")
	if test1 != 58 || test2 != 34 {
		log.Fatal(test1, test2)
	}
	test1, test2 = advent02("1x1x10")
	if test1 != 43 || test2 != 14 {
		log.Fatal(test1, test2)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum, sum2 := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		log.Println(advent02(t))
		one, two := advent02(t)
		sum += one
		sum2 += two
	}
	fmt.Println(sum, sum2)
}
