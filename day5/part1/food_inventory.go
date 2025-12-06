package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input = "../input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges [][]int
	var fresh []int
	var reachedIds bool

outer:
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reachedIds = true
			continue
		}

		if !reachedIds {
			nums := []int{}
			stringNums := strings.Split(line, "-")
			for _, s := range stringNums {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalf("could not convert string to int: %s", err)
				}

				nums = append(nums, num)
			}

			ranges = append(ranges, nums)
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not convert id string to int: %s", err)
		}

		for _, r := range ranges {
			if id > r[0] && id < r[1] {
				fresh = append(fresh, id)
				continue outer
			}
		}
	}

	fmt.Println(len(fresh))
}
