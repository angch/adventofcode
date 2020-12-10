package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"sync/atomic"
)

var done = map[string]bool{}

func isValid(num []int, k int) int {
	count := int64(0)
	i := 0
	for _, v := range num {
		if v-i > 3 {
			return 0
		}
		i = v
	}
	count++
	// log.Println("isvalid", num)

	// copy(a, num)

	a := make([]int, len(num))
	// wg := sync.WaitGroup{}
	for i := k; i < len(num)-1; i++ {
		if num[i]-num[i-1] >= 3 {
			continue
		}

		subcount := 1
		if num[i+1]-num[i] == 1 {
			subcount *= 2
			if num[i+2]-num[i+1] == 1 {
				subcount *= 2
			}
		}

		copy(a[:i], num[:i])
		// log.Println("before", a, i)
		copy(a[i:], num[i+1:]) // Shift a[i+1:] left one index.
		a = a[:len(num)-1]
		// log.Println("after", a)

		// j, _ := json.Marshal(a)
		// if done[string(j)] {
		// 	continue
		// }
		// done[string(j)] = true
		// wg.Add(1)
		// go func(a []int, i int) {
		ii := isValid(a, i)
		atomic.AddInt64(&count, int64(ii))
		// wg.Done()
		// }(a, i)
	}
	// wg.Wait()

	return int(count)
}

func do(fileName string, window int) (int, int) {
	done = map[string]bool{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)
	max := -1
	for scanner.Scan() {
		l := scanner.Text()
		n := 0
		fmt.Sscanf(l, "%d", &n)
		numbers = append(numbers, n)
		if n > max {
			max = n
		}
	}

	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, max+3)
	// log.Println(numbers)
	jolt := 0
	diff := make(map[int]int)
	for _, v := range numbers {
		if v == 0 {
			continue
		}
		d := v - jolt
		diff[d]++
		jolt = v
		// fmt.Print(d)
	}
	// fmt.Println()
	// diff[3]++
	// log.Printf("%+v\n", diff)

	diffs := make([]int, 0)
	isRun := false
	j := 0
	for i := 1; i < len(numbers); i++ {
		if isRun {
			if numbers[i]-numbers[i-1] > 1 {
				isRun = false
				diffs = append(diffs, i-j)
				j = -1
			}
		} else {
			if numbers[i]-numbers[i-1] == 1 {
				isRun = true
				j = i
				continue
			}
		}
	}
	// log.Println("diffs", diffs)

	count := 0
	if len(numbers) < 2 {
		count = isValid(numbers, 1)
	} else {
		count = 1
		for _, d := range diffs {
			count *= m(d)
		}
	}
	// count := 0

	// jolt = 0
	// for i := len(numbers)-1;
	// jolt = v
	// }

	return diff[1] * diff[3], count
}

var mult = []int{
	0, //0
	1, //1
	2, //2
	4, //3
	7, //4
	13, 24, 44, 81, 149, 274, 504,
}

func m(i int) int {
	return mult[i]
}

func main() {
	// log.Println(do("test.txt", 5))
	// log.Println(do("test2.txt", 5))
	// log.Println(do("test3.txt", 5))
	log.Println(do("input.txt", 5))
	// ((((((4*2-1)*2-1)*2-2)*2-4)*2-7)*2-13)*2-(13+7+4)

}
