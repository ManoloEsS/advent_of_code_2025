package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const input = "input.txt"

// TODO: modify for using queue approach to check only modifiable positions and its potentially modifiable neighbors
func main() {
	start := time.Now()
	var grid [][]rune
	var reachable int
	removed := -1

	f, err := os.Open(input)
	if err != nil {
		log.Printf("could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	for removed != 0 {
		removed = 0
		for lineIdx, line := range grid {
			for i := 0; i < len(line); i++ {
				adjacent := 0
				if line[i] == '@' {
					if lineIdx > 0 {
						//check spaces above
						if i > 0 {
							if grid[lineIdx-1][i-1] == '@' {
								adjacent += 1
							}
						}

						if grid[lineIdx-1][i] == '@' {
							adjacent += 1
						}

						if i < len(line)-1 {
							if grid[lineIdx-1][i+1] == '@' {
								adjacent += 1
							}
						}
					}

					//check spaces on same line
					if i > 0 {
						if grid[lineIdx][i-1] == '@' {
							adjacent += 1
						}
					}

					if i < len(line)-1 {
						if grid[lineIdx][i+1] == '@' {
							adjacent += 1
						}
					}

					//check spaces on line below
					if lineIdx < len(grid)-1 {
						if i > 0 {
							if grid[lineIdx+1][i-1] == '@' {
								adjacent += 1
							}
						}
						if grid[lineIdx+1][i] == '@' {
							adjacent += 1
						}
						if i < len(line)-1 {
							if grid[lineIdx+1][i+1] == '@' {
								adjacent += 1
							}
						}
					}

					if adjacent < 4 {
						reachable += 1
						removed += 1
						grid[lineIdx][i] = '.'
					}
				}
			}
		}
	}
	fmt.Println(reachable)
	elapsed := time.Since(start)

	fmt.Printf("Program took: %s\n", elapsed)
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
