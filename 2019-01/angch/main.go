package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		t := scanner.Text()
		mass := 0
		fmt.Sscanf(t, "%d", &mass)
		fuel := (mass / 3) - 2
		// fmt.Println(mass, fuel)
		sum2 := fuel
		for {
			fuel = fuel/3 - 2
			if fuel <= 0 {
				break
			}
			sum2 += fuel
		}
		fmt.Println(mass, fuel, sum2)
		sum += sum2
	}
	fmt.Println(sum)

}
