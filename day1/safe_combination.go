package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const input string = "input.txt"

func main() {
	var (
		position    = 50
		maxPos      = 99
		minPos      = 0
		zeroCounter = 0
	)

	f, err := os.Open(input)
	if err != nil {
		log.Printf("could not read file: %s", input)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Printf("could not read integer from string")
		}

		fullCycles := num / 100
		zeroCounter += fullCycles

		tensNum := num % 100

		if tensNum == 0 {
			continue
		}

		if line[0] == 'R' {
			newPosition := position + tensNum
			if newPosition > maxPos {
				position = newPosition - 100
				zeroCounter++
			} else {
				position = newPosition
			}
		}

		if line[0] == 'L' {
			newPosition := position - tensNum
			if newPosition < minPos {
				if position > 0 {
					zeroCounter++
				}
				position = 100 + newPosition
			} else {
				position = newPosition
			}
			if position == 0 {
				zeroCounter++
			}
		}
	}
	fmt.Println(zeroCounter)
}
