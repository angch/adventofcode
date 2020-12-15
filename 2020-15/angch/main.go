package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func do(fileName string) (int, int) {
	ret1, ret2 := 0, 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nums := make(map[int][]int)
	last := 0
	pos := 0
	for scanner.Scan() {
		l := scanner.Text()
		a := strings.Split(l, ",")
		for k, v := range a {
			num, _ := strconv.Atoi(v)
			_, ok := nums[num]
			if !ok {
				nums[num] = make([]int, 0, 2)
			}
			nums[num] = append(nums[num], k+1)
			pos = k + 1
			last = num
		}
	}

	out := 0
	for {
		if len(nums[last]) == 1 {
			out = 0
		} else {
			out = nums[last][len(nums[last])-1] - nums[last][len(nums[last])-2]
		}
		pos++

		_, ok := nums[out]
		if !ok {
			nums[out] = make([]int, 0, 2)
		}
		if len(nums[out]) <= 1 {
			nums[out] = append(nums[out], pos)
		} else {
			nums[out][0] = nums[out][1]
			nums[out][1] = pos
		}
		last = out
		if pos == 2020 {
			ret1 = last
		}
		if pos >= 30000000 {
			// if pos >= 2020 {
			break
		}
	}
	ret2 = last

	return ret1, ret2
}

func main() {
	log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	log.Println(do("input.txt"))
}
