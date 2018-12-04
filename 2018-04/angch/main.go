package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	/*[1518-11-01 00:00] Guard #10 begins shift
	  [1518-11-01 00:05] falls asleep
	  [1518-11-01 00:25] wakes up
	  [1518-11-01 00:30] falls asleep
	  [1518-11-01 00:55] wakes up
	*/
	re1 := regexp.MustCompile(`^\[(\d+-\d+-\d+)\s*(\d+):(\d+)\]\s*Guard #(\d+) begins shift`)
	re2 := regexp.MustCompile(`^\[(\d+-\d+-\d+)\s*(\d+):(\d+)\]\s*\w+\s*\w*$`)
	re3 := regexp.MustCompile(`^\[(\d+-\d+-\d+\s*\d+:\d+)\]\s*(.*)$`)

	scanner := bufio.NewScanner(file)
	guardSleeps := make(map[int]int)
	currentGuard := -1
	lines := make(map[string]string)
	times := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		c := re3.FindAllStringSubmatch(t, -1)
		//b := re2.FindAllStringSubmatch(t, -1)
		log.Println(c)
		lines[c[0][1]] = t
		times = append(times, c[0][1])
	}
	sort.Strings(times)
	log.Println(times)
	sleepTime := -1
	matrix := make(map[int]map[int]int, 0)
	for _, cron := range times {
		t := lines[cron]
		a := re1.FindAllStringSubmatch(t, -1)
		b := re2.FindAllStringSubmatch(t, -1)
		log.Println(t)

		if len(a) > 0 {
			currentGuard, _ = strconv.Atoi(a[0][4])
			// h, _ := strconv.Atoi(a[0][2])
			// m, _ := strconv.Atoi(a[0][3])
			// if h >= 12 {
			// 	h -= 12
			// }
			//startTime := h*60 + m
			sleepTime = -1
		} else if len(b) > 0 {
			h, _ := strconv.Atoi(b[0][2])
			m, _ := strconv.Atoi(b[0][3])
			if h >= 12 {
				h -= 12
			}

			if sleepTime < 0 {
				sleepTime = h*60 + m
			} else {
				sleepDuration := h*60 + m - sleepTime
				log.Println("Guard", currentGuard, sleepDuration)
				guardSleeps[currentGuard] += sleepDuration

				if matrix[currentGuard] == nil {
					matrix[currentGuard] = make(map[int]int)
				}
				for ti := 0; ti < sleepDuration; ti++ {
					matrix[currentGuard][ti+sleepTime]++
				}
				sleepTime = -1
			}
		}
	}
	maxSleep := 0
	maxGuard := -1

	for k, v := range guardSleeps {
		if v > maxSleep {
			maxGuard = k
			maxSleep = v
		}
	}
	log.Println(guardSleeps)
	mostTime := -1
	mostMinute := -1
	for k, v := range matrix[maxGuard] {
		if v > mostTime {
			mostTime = v
			mostMinute = k
		}
	}
	fmt.Println(maxGuard, maxSleep)
	log.Println(matrix[maxGuard])
	fmt.Println(mostMinute, mostTime)

	fmt.Println(maxGuard * mostMinute)

	// Part 2
	counts := make([]map[int]int, 1000)
	for guard, times := range matrix {
		for k, v := range times {
			if counts[k] == nil {
				counts[k] = make(map[int]int)
			}
			counts[k][guard] += v
		}
	}
	maxMin2 := -1
	maxGuard2 := -1
	maxTime2 := -1
	for time, v := range counts {
		for guard, times := range v {
			if times > maxMin2 {
				maxMin2 = times
				maxGuard2 = guard
				maxTime2 = time
			}
		}
	}

	log.Print(maxTime2, maxGuard2, maxTime2*maxGuard2)
	//

	// 	//log.Println("a", a)
	// 	log.Println("b", b)
	// }
}
