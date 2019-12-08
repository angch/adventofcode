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
	scanner.Scan()
	data := scanner.Text()

	width, height := 25, 6
	layerCounts := make([][]int, 0)

	for k, v := range data {
		layer := k / (width * height)

		if layer >= len(layerCounts) {
			counts := make([]int, 10)
			layerCounts = append(layerCounts, counts)
		}
		c := v - '0'
		layerCounts[layer][c]++
	}

	min, minLayer := 9999, 999
	for k, l := range layerCounts {
		if l[0] < min {
			min, minLayer = l[0], k
		}
	}
	log.Println("Part 1:", min, minLayer, layerCounts[minLayer][1]*layerCounts[minLayer][2])

	log.Println("Part 2:")
	layers := len(data) / width / height
	for y := 0; y < height; y++ {
	a:
		for x := 0; x < width; x++ {
			for l := 0; l < layers; l++ {
				switch data[(l*width*height)+y*(width)+x] - '0' {
				case 0:
					fmt.Print("  ")
					continue a
				case 1:
					fmt.Print("**")
					continue a
				}
			}
			fmt.Println("  ")
		}
		fmt.Println()
	}
}
