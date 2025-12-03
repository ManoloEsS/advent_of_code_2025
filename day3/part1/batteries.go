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
	f, err := os.Open(input)
	if err != nil {
		log.Printf("could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var voltages []int

	for scanner.Scan() {
		line := scanner.Text()
		highest := -1
		secondHighest := -1

		for i := 0; i < len(line)-1; i++ {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Printf("could not convert string to int: %s", err)
			}

			if num > highest {
				secondHighest = -1
				highest = num
				continue
			}
			if num > secondHighest || num == highest {
				secondHighest = num
				continue
			}

		}

		lastPos, _ := strconv.Atoi(string(line[len(line)-1]))

		if lastPos > secondHighest {
			secondHighest = lastPos
		}

		voltage := highest*10 + secondHighest
		voltages = append(voltages, voltage)
	}

	fmt.Println(voltages)
	sum := 0
	for _, num := range voltages {
		sum += num
	}
	fmt.Println(sum)
}
