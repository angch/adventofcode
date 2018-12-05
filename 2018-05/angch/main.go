package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isupper(c byte) int {
	if c >= 'A' && c <= 'Z' {
		return 1
	}
	return 0
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
			if isupper(byte(v))^isupper(s[k+1]) > 0 {
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

func main() {
	// fmt.Println(foo("aA"))
	// fmt.Println(foo("abBA"))
	// fmt.Println(foo("abAB"))
	// fmt.Println(foo("aabAAB"))
	// fmt.Println(foo("dabAcCaCBAcCcaDA"))
	input, _ := ioutil.ReadFile("input.txt")

	// Facepalm.
	input2 := strings.TrimSpace(string(input))
	fmt.Println(len(input2), len(input))
	if true {
		fmt.Println(foo(input2, 0))
	}
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
