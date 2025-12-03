package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const input = "input.txt"

func main() {
	start := time.Now()
	var joltages int

	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var lineJoltage string
		var startIndex, highestIndex int

		for remaining := 11; remaining >= 0; remaining-- {
			highestBatt := '0'

			for index, num := range line[startIndex : len(line)-remaining] {
				if num > highestBatt {
					highestBatt = num
					highestIndex = startIndex + index
				}
			}

			lineJoltage += string(highestBatt)
			startIndex = highestIndex + 1
		}

		intLineJoltage, err := strconv.Atoi(lineJoltage)
		if err != nil {
			log.Fatalf("could not convert lineJoltage to int: %s", err)
		}
		joltages += intLineJoltage
	}

	// totalJoltage := 0
	// for _, j := range joltages {
	// 	totalJoltage += j
	// }

	fmt.Println(joltages)
	elapsed := time.Since(start)
	fmt.Printf("Program took: %s\n", elapsed)
}
