package main

import (
	"bufio"
	"fmt"
	"os"
)

var filepath = "input.txt"

// var filepath = "test.txt"
var debug = false

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

func part1() {
	file, _ := os.Open(filepath)
	// file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counts := make([][2]int, 26)
	length := 0
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
		length = len(line)
	}
	gamma := 0
	epsilon := 0
	for i := 0; i < length; i++ {
		gamma <<= 1
		epsilon <<= 1
		if counts[i][0] < counts[i][1] {
			gamma |= 1
			// fmt.Print("1")
		} else {
			epsilon |= 1
		}
	}
	fmt.Println(gamma * epsilon)

}

func part2() {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	oxygen := []string{}
	co2 := []string{}
	sum1, sum2 := 0, 0

	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		oxygen = append(oxygen, line)
		co2 = append(co2, line)
		length = len(line)
	}

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
			if debug {
				fmt.Println("oxy", oxygen[0])
			}
			sum1 = str2bin(oxygen[0])
			break
		}
		if debug {
			fmt.Println(i+1, oxygen, mostcommon)
		}
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
		leastcommon := '0'
		if counts[0] > counts[1] {
			leastcommon = '1'
		}
		co22 := make([]string, 0)
		for _, v := range co2 {
			if v[i] == byte(leastcommon) {
				co22 = append(co22, v)
			}
		}
		co2 = co22
		if len(co2) == 1 {
			if debug {
				fmt.Println("co2", co2[0])
			}
			sum2 = str2bin(co2[0])
			break
		}
	}
	fmt.Println(sum1 * sum2)
}

func main() {
	part1()
	part2()
}
