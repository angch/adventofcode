package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	k := 0
	for scanner.Scan() {
		t := scanner.Text()
		i := 0
		fmt.Sscanf(t, "%d", &i)
		//log.Println(i)
		k += i
	}
	fmt.Println(k)
}

func main() {

	k := 0
	freq := make(map[int]bool, 0)
	freq[k] = true
	for {
		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			t := scanner.Text()
			i := 0
			fmt.Sscanf(t, "%d", &i)
			k += i
			//log.Println(freq, i, k)
			if freq[k] {
				fmt.Println(k)
				return
			}
			freq[k] = true
		}
		file.Close()
	}
	//fmt.Println(k)
}
