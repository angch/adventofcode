package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isupper(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
func islower(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func foo(s string, ignore byte) int {

	if ignore != 0 {
		o := ""
		for _, v := range s {
			if lower(byte(v)) == ignore {
				continue
			}
			o += string(v)
		}
		s = o
	}

a:
	for k, v := range s {
		if k+1 < len(s) {
			//fmt.Println(s, string(lower(byte(v))), string(lower(byte(s[k+1]))))
			if lower(byte(v)) != lower(byte(s[k+1])) {
				continue
			}
			//fmt.Println("a")
			if isupper(byte(v)) != isupper(s[k+1]) {
				if k > 0 {
					s = s[0:k] + s[k+2:len(s)]
				} else {
					s = s[k+2 : len(s)]
				}
				//fmt.Println(s)
				goto a
			}
		}
	}
	return len(s)
}

func foo2(s string) int {
	replaceMe := make([]string, 0)
	for a := 'a'; a <= 'z'; a++ {
		b, c := string(a), strings.ToUpper(string(a))
		replaceMe = append(replaceMe, b+c)
		replaceMe = append(replaceMe, c+b)
	}
	//fmt.Println(replaceMe)
a:
	for oldLen := 0; ; oldLen = len(s) {
		oldLen = len(s)
		for _, v := range replaceMe {
			s = strings.Replace(s, v, "", 1)
			if len(s) != oldLen {
				continue a
			}
		}
		break
	}
	return len(s)
}

// Both parts, one iter
// Returns part1_len, part2_len, part2_char
func foo3(s string) (int, int, string) {
	l := len([]byte(s))
	best := make([][]byte, 27)
	for k := range best {
		best[k] = make([]byte, 0, l)
	}
	//prev := ':'
	//prevIsLower := false
	for _, v := range s {
		lower := byte(v)
		currentIsLower := true

		if v >= 'A' && v <= 'Z' {
			lower += ('a' - 'A')
			currentIsLower = false
		}

		for k2, v2 := range best {
			if lower-'a'+1 == byte(k2) {
				continue
			}
			if len(v2) == 0 {
				best[k2] = append(best[k2], byte(v))
				continue
			}
			last := v2[len(v2)-1]
			lastLower := last
			if lastLower >= 'A' && lastLower <= 'Z' {
				lastLower += ('a' - 'A')
			}
			if lower == lastLower && currentIsLower != islower(last) {
				//fmt.Println("f")
				best[k2] = best[k2][0:(len(v2) - 1)]
			} else {
				//fmt.Println("f")
				best[k2] = append(best[k2], byte(v))
			}
		}
		//bestLower = len(best[0])
		//prevIsLower = currentIsLower
		//prev = v
	}
	least := len(best[0])
	leastk := 0
	for k, v := range best {
		if least > len(v) {
			least = len(v)
			leastk = k
		}
		//fmt.Println(k, len(v))
	}
	//fmt.Println(string(best[0]))
	return len(best[0]), least, string(leastk + 'a' - 1)
}

// Both parts, one iter
// Returns part1_len, part2_len, part2_char
func foo4(s []byte) (int, int, string) {
	l := len([]byte(s))
	best := make([][]byte, 27)
	lengths := [27]int{}
	for k := range best {
		best[k] = make([]byte, l, l)
	}
	for _, v := range s {
		if v < 'A' {
			continue
		}
		lower := v
		currentIsLower := true

		if v >= 'A' && v <= 'Z' {
			lower += ('a' - 'A')
			currentIsLower = false
		}

		for k2, v2 := range best {
			if lower-'a' == byte(k2) {
				continue
			}
			if lengths[k2] == 0 {
				best[k2][0] = v
				lengths[k2] = 1
				continue
			}
			last := v2[lengths[k2]-1]
			lastLower := last

			lastIsLower := true
			if lastLower >= 'A' && lastLower <= 'Z' {
				lastLower += ('a' - 'A')
				lastIsLower = false
			}
			if lower == lastLower && currentIsLower != lastIsLower {
				lengths[k2]--
			} else {
				best[k2][lengths[k2]] = v
				lengths[k2]++
			}
		}
	}
	least := lengths[26]
	leastk := 0
	for k, _ := range best {
		if least > lengths[k] {
			least = lengths[k]
			leastk = k
		}
	}

	return lengths[26], least, string(leastk + 'a')
}

func main() {
	// fmt.Println(foo3("aA"))
	// fmt.Println(foo3("abBA"))
	// fmt.Println(foo3("abAB"))
	// fmt.Println(foo3("aabAAB"))
	// fmt.Println(foo3("dabAcCaCBAcCcaDA"))
	// fmt.Println(foo("aA", 0))
	// fmt.Println(foo("abBA", 0))
	// fmt.Println(foo("abAB", 0))
	// fmt.Println(foo("aabAAB", 0))
	// fmt.Println(foo("dabAcCaCBAcCcaDA", 0))
	input, _ := ioutil.ReadFile("input.txt")

	// Facepalm.
	input2 := strings.TrimSpace(string(input))
	fmt.Println(len(input2), len(input))
	if false {
		fmt.Println(foo(input2, 0))
	}
	if false {
		// 28.8 seconds
		min := 99999
		best := ';'
		for i := 'a'; i <= 'z'; i++ {
			k := foo(input2, byte(i))
			fmt.Println(i, k)
			if k < min {
				min = k
				best = i
			}
		}
		fmt.Println(best, min)
	}
	if true {
		// 0.42 seconds
		fmt.Println(foo3(input2))
	}
}
