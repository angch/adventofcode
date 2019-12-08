package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if true {
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
		layer := 0
		layerCounts := make([][]int, 0)

		for k, v := range data {
			// x = k % (width)
			// y = (k % (width * height)) / width
			layer = k / (width * height)

			if layer >= len(layerCounts) {
				counts := make([]int, 10)
				layerCounts = append(layerCounts, counts)
				// log.Println(layer)
			}
			c := v - '0'

			layerCounts[layer][c]++
		}

		min := 9999
		minLayer := 999
		for k, l := range layerCounts {
			if l[0] < min {
				min = l[0]
				minLayer = k
			}
			log.Printf("layer %d: %d\n", l[0], k)
		}
		//1548 too low
		log.Println(min, minLayer, layerCounts[minLayer][1]*layerCounts[minLayer][2])

		layers := len(data) / width / height
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				c := 3
				for l := 0; l < layers; l++ {
					if data[(l*width*height)+y*(width)+x]-'0' < 2 {
						c = int(data[(l*width*height)+y*(width)+x] - '0')
						break
					}
				}
				if c == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
				// fmt.Print(c)
			}
			fmt.Println()
		}
	}

}
