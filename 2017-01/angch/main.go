package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetOutput(ioutil.Discard)

	test1 := advent01a("test.txt")
	fmt.Println(test1)
	if test1 != 3 {
		log.Fatal("err")
	}
}

func advent01a(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	count2 := 0
	for scanner.Scan() {
		t := scanner.Text()
		count2++
		words := strings.Split(t, "")
		prev := -1

		for _, w := range words {
			i, _ := strconv.Atoi(w)
			//fmt.Printf("%#v\n", i)
			if i == prev {
				count += prev
				log.Printf("%#v y\n", i)

			} else {
				log.Printf("%#v n\n", i)
			}
			prev = i
		}
		i, _ := strconv.Atoi(words[0])
		//fmt.Printf("%#v\n", i)
		if i == prev {
			count += prev
			log.Printf("%#v y\n", i)
		} else {
			log.Printf("%#v n\n", i)
		}
		prev = i

		//count++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(count, count2)
	return count
}
