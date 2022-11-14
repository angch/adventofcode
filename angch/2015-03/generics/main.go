package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/angch/adventofcode/angch/vector"
)

var directions = map[rune]vector.Point[int]{
	'^': vector.New(0, -1),
	'v': vector.New(0, 1),
	'>': vector.New(1, 0),
	'<': vector.New(-1, 0),
}

func advent03(input string) int {
	done := make(map[vector.Point[int]]int)
	vec := vector.Point[int]{}
	done[vec] = 1
	count := 1
	for _, v := range input {
		_ = vec.SelfAdd(directions[v])
		done[vec]++
		if done[vec] == 1 {
			count++
		}
	}
	return count
}

func advent03b(input string) int {
	done := make(map[vector.Point[int]]int)
	vec := [2]vector.Point[int]{}
	done[vec[0]] = 1
	count := 1
	for k, v := range input {
		_ = vec[k%2].SelfAdd(directions[v])
		done[vec[k%2]]++
		if done[vec[k%2]] == 1 {
			count++
		}
	}
	return count
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test1 := advent03(">")
	if test1 != 2 {
		log.Fatal(test1)
	}
	test1 = advent03("^>v<")
	if test1 != 4 {
		log.Fatal(test1)
	}
	test1 = advent03("^v^v^v^v^v")
	if test1 != 2 {
		log.Fatal(test1)
	}

	test1 = advent03b("^v")
	if test1 != 3 {
		log.Fatal(test1)
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
		//log.Println(advent03(t))
		one := advent03(t)
		two := advent03b(t)
		sum += one
		sum2 += two
	}
	fmt.Println(sum, sum2)
}
