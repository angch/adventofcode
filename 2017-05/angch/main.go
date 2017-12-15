package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//log.SetOutput(ioutil.Discard)

	test1 := advent05b("test1.txt")
	fmt.Println(test1)
	if test1 != 10 {
		log.Fatal("err")
	}
	o := advent05b("input.txt")
	fmt.Println(o)

}

func advent05a(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	program := make([]int, 0)

	for scanner.Scan() {
		t := scanner.Text()
		words := strings.Split(t, " ")
		i := 0
		fmt.Sscanf(words[0], "%d", &i)
		program = append(program, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := 0

	fmt.Println(program)

	ip, nip := 0, 0
	for ip < len(program) {
		nip = ip + program[ip]
		program[ip]++
		ip = nip
		count++
	}

	return count
}

func advent05b(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	program := make([]int, 0)

	for scanner.Scan() {
		t := scanner.Text()
		words := strings.Split(t, " ")
		i := 0
		fmt.Sscanf(words[0], "%d", &i)
		program = append(program, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := 0

	fmt.Println(program)

	ip, nip := 0, 0
	for ip < len(program) {
		nip = ip + program[ip]
		if program[ip] >= 3 {
			program[ip]--
		} else {
			program[ip]++
		}

		ip = nip
		count++
	}

	return count
}
