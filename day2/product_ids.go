package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input string = "input.txt"

func main() {
	invalidIds := make(map[int]bool)
	f, err := os.Open(input)
	if err != nil {
		log.Printf("could not read file: %s", input)
	}
	defer f.Close()

	fi, _ := f.Stat()
	buf := make([]byte, fi.Size())

	_, err = f.Read(buf)
	rawText := string(buf)
	inputText := strings.TrimRight(rawText, "\n")

	ranges := strings.Split(inputText, ",")

	for _, item := range ranges {
		limits := strings.Split(item, "-")

		lower, err := strconv.Atoi(limits[0])
		if err != nil {
			log.Printf("couldn't convert lower limit to int: %s", err)
		}

		upper, err := strconv.Atoi(limits[1])
		if err != nil {
			log.Printf("couldn't convert upper limit to int: %s", err)
		}

		for num := lower; num <= upper; num++ {
			numString := strconv.Itoa(num)

			for i := 1; i <= (len(numString) / 2); i++ {
				pattern := numString[:i]
				if len(numString)%len(pattern) == 0 {
					repeats := checkNumRepeatsTwice(numString[:i], numString)
					// if len(pattern) == 1 && len(numString)%2 == 0 && repeats == true {
					// 	invalidIds[num] = true
					// }
					if repeats == true {
						invalidIds[num] = true
					}
				}
			}
		}
	}
	fmt.Println(invalidIds)
	sum := 0
	for num, _ := range invalidIds {
		sum += num
	}
	fmt.Println(sum)
}

func checkNumRepeatsTwice(pattern, numString string) bool {
	for i := 0; i+len(pattern) <= len(numString); i += len(pattern) {
		if numString[i:i+len(pattern)] != pattern {
			return false
		}
		if len(numString) == 1 {
			return false
		}

	}
	return true
}
