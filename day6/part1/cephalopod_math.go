package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const input = "../input.txt"

func main() {
	start := time.Now()
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grandTotal := 0
	operationMap := make(map[int]*operationLine)

	for lineIdx := len(lines) - 1; lineIdx >= 0; lineIdx-- {
		symbols := strings.Fields(lines[lineIdx])
		for i, symbol := range symbols {
			if _, ok := operationMap[i]; !ok {
				if symbol == "+" {
					newOp := newOperationLineSum(symbol)
					operationMap[i] = &newOp
				}
				if symbol == "*" {
					newOp := newOperationLineProd(symbol)
					operationMap[i] = &newOp
				}
			} else {
				if operationMap[i].operation == "+" {
					num, err := strconv.Atoi(symbol)
					if err != nil {
						log.Fatalf("could not convert symbol string to int: %s", err)
					}
					operationMap[i].result += num
				}
				if operationMap[i].operation == "*" {
					num, err := strconv.Atoi(symbol)
					if err != nil {
						log.Fatalf("could not convert symbol string to int: %s", err)
					}
					operationMap[i].result *= num
				}
			}
		}
	}
	for _, val := range operationMap {
		grandTotal += val.result
	}

	fmt.Println(grandTotal)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

type operationLine struct {
	operation string
	result    int
}

func newOperationLineSum(symbol string) operationLine {
	return operationLine{
		operation: symbol,
		result:    0,
	}
}
func newOperationLineProd(symbol string) operationLine {
	return operationLine{
		operation: symbol,
		result:    1,
	}
}
