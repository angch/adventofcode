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
			k2, ok := nums[num]
			if !ok {
				k2 = make([]int, 0, 2)
			}
			nums[num] = append(k2, k+1)
			pos, last = k+1, num
		}
	}

	out := 0
	for {
		k := nums[last]
		if len(k) == 1 {
			out = 0
		} else {
			out = k[1] - k[0]
		}
		pos++

		k, ok := nums[out]
		if !ok {
			k = make([]int, 0, 2)
		}
		if len(k) <= 1 {
			nums[out] = append(k, pos)
		} else {
			k[0], k[1] = k[1], pos
		}
		last = out
		if pos == 2020 {
			ret1 = last
		}
		if pos >= 30000000 {
			break
		}
	}
	ret2 = last

	return ret1, ret2
}

type Vec struct {
	One int32
	Two int32
}

func do2(input []int) (ret1 int, ret2 int) {
	nums := make([]Vec, 30000000)
	last, pos := int32(0), int32(0)
	for _, v := range input {
		pos++
		last = int32(v)
		nums[last] = Vec{One: int32(pos)}
	}

	k := nums[last]
	for {
		if k.Two == 0 {
			last = 0
		} else {
			last = k.One
		}
		pos++
		if pos == 2020 {
			ret1 = int(last)
		}
		if pos >= 30000000 {
			break
		}

		k = nums[last]
		if k.One == 0 {
			k = Vec{
				One: pos,
			}
		} else if k.Two == 0 {
			k = Vec{
				One: pos - k.One,
				Two: pos,
			}
		} else {
			k = Vec{
				One: pos - k.Two,
				Two: pos,
			}
		}
		nums[last] = k
	}
	ret2 = int(last)
	return
}

func do3(input []int) (ret1 int, ret2 int) {
	nums := make([]int32, 30000000)
	last, pos := int32(0), int32(0)
	for _, v := range input {
		pos++
		last = int32(v)
		nums[last] = pos
	}

	for prev := int32(0); pos < 30000000; pos++ {
		if pos == 2020 {
			ret1 = int(last)
		}
		prev, nums[last] = nums[last], pos
		if prev == 0 {
			last = 0
		} else {
			last = pos - prev
		}
	}
	return ret1, int(last)
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do("input.txt"))
	log.Println(do3([]int{12, 20, 0, 6, 1, 17, 7}))
}
