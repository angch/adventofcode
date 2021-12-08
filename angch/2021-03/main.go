package main

import (
	"bufio"
	"fmt"
	"os"
)

func str2bin(s string) int {
	bin := 0
	for _, v := range s {
		bin <<= 1
		if v == '1' {
			bin |= 1
		}
	}
	return bin
}

func day3(filepath string) {
	file, _ := os.Open(filepath)
	// file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counts := make([][2]int, 26)
	length := 0

	o2 := []string{}
	co2 := []string{}

	sum1, sum2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		for k, v := range line {
			if v == '0' {
				counts[k][0]++
			} else {
				counts[k][1]++
			}
		}
		o2 = append(o2, line)
		co2 = append(co2, line)
		length = len(line)
	}
	gamma := 0
	for i := 0; i < length; i++ {
		gamma <<= 1
		if counts[i][0] < counts[i][1] {
			gamma |= 1
		}

		if len(o2) > 1 {
			o2count := [2]int{0, 0}
			for _, v := range o2 {
				if v[i] == '0' {
					o2count[0]++
				} else {
					o2count[1]++
				}
			}
			o2common := '0'
			if o2count[0] <= o2count[1] {
				o2common = '1'
			}
			for j := 0; j < len(o2); j++ {
				if o2[j][i] != byte(o2common) {
					o2 = append(o2[:j], o2[j+1:]...)
					j--
				}
			}
		}

		if len(co2) > 1 {
			co2count := [2]int{0, 0}
			for _, v := range co2 {
				if v[i] == '0' {
					co2count[0]++
				} else {
					co2count[1]++
				}
			}
			co2leastcommon := '0'
			if co2count[0] > co2count[1] {
				co2leastcommon = '1'
			}

			for j := 0; j < len(co2); j++ {
				if co2[j][i] != byte(co2leastcommon) {
					co2 = append(co2[:j], co2[j+1:]...)
					j--
				}
			}
		}
	}
	epsilon := ((1 << length) - 1) & ^gamma
	fmt.Println("Part 1", gamma*epsilon)

	sum1, sum2 = str2bin(o2[0]), str2bin(co2[0])
	fmt.Println("Part 2", sum1, sum2, sum1*sum2)
}

func main() {
	day3("test.txt")
	day3("input.txt")
}
