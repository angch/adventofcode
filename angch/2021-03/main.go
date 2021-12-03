package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	// file, _ := os.Open("test.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	oxygen := []string{}
	co2 := []string{}
	counts := make([][2]int, 26)
	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		lines = append(lines, line)
		oxygen = append(oxygen, line)
		co2 = append(co2, line)

		for k, v := range line {
			if v == '0' {
				counts[k][0]++
			} else {
				counts[k][1]++
			}
		}
		length = len(line)
	}
	gamma := 0
	epsilon := 0
	for i := 0; i < length; i++ {
		// fmt.Println("pos ", i, " ", counts[i])
		gamma <<= 1
		epsilon <<= 1
		if counts[i][0] < counts[i][1] {
			gamma |= 1
			// fmt.Print("1")
		} else {
			epsilon |= 1
			// fmt.Print("0")
		}
	}
	fmt.Println(gamma * epsilon)

	sum1, sum2 := 0, 0

	for i := 0; i < length; i++ {
		counts := [2]int{0, 0}
		for _, v := range oxygen {
			if v[i] == '0' {
				counts[0]++
			} else {
				counts[1]++
			}
		}
		mostcommon := '0'
		if counts[0] <= counts[1] {
			mostcommon = '1'
		}
		oxygen2 := make([]string, 0)
		for _, v := range oxygen {
			if v[i] == byte(mostcommon) {
				oxygen2 = append(oxygen2, v)
			}
		}
		oxygen = oxygen2
		if len(oxygen) == 1 {
			fmt.Println("oxy", oxygen[0])
			for _, v := range oxygen[0] {
				sum1 <<= 1
				if v == '1' {
					sum1 |= 1
				}
			}
			break
		}
		fmt.Println(i+1, oxygen, mostcommon)

	}
	for i := 0; i < length; i++ {
		counts := [2]int{0, 0}
		for _, v := range co2 {
			if v[i] == '0' {
				counts[0]++
			} else {
				counts[1]++
			}
		}
		mostcommon := '0'
		if counts[0] > counts[1] {
			mostcommon = '1'
		}
		co22 := make([]string, 0)
		for _, v := range co2 {
			if v[i] == byte(mostcommon) {
				co22 = append(co22, v)
			}
		}
		co2 = co22
		if len(co2) == 1 {
			fmt.Println("co2", co2[0])
			for _, v := range co2[0] {
				sum2 <<= 1
				if v == '1' {
					sum2 |= 1
				}
			}
			break
		}
	}
	fmt.Println(sum1, sum2, sum1*sum2)
}

func part2() {
	file, _ := os.Open("test.txt")
	// file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}

}
func main() {
	part1()
	// part2()
}
