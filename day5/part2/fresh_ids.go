package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const input = "../input.txt"

func main() {
	start := time.Now()
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var nums []int
		stringNums := strings.Split(line, "-")

		for _, s := range stringNums {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not convert string to int: %s", err)
			}

			nums = append(nums, num)
		}

		if len(ranges) == 0 {
			ranges = append(ranges, nums)
			continue
		}

		extended := false
		var toDeleteQueue []int

		for i, r := range ranges {
			if nums[1] >= r[0] && nums[0] <= r[1] {
				nums[0] = min(nums[0], r[0])
				nums[1] = max(nums[1], r[1])

				toDeleteQueue = prependInt(toDeleteQueue, i)
				extended = true
			}
		}

		if extended {
			for _, toDelete := range toDeleteQueue {
				ranges = slices.Delete(ranges, toDelete, toDelete+1)
			}
		}

		ranges = append(ranges, nums)

	}

	fmt.Println(ranges)
	totalFresh := 0
	for _, r := range ranges {
		totalFresh += ((r[1] - r[0]) + 1)
	}

	fmt.Println(totalFresh)
	elapsed := time.Since(start)
	fmt.Printf("Program took: %s\n", elapsed)
}

func prependInt(x []int, y int) []int {
	result := make([]int, len(x)+1)
	result[0] = y
	copy(result[1:], x)
	return result
}

