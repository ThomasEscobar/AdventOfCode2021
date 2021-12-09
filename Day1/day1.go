package main

import (
	"fmt"
	"log"
	"strconv"

	"toolbox"
)

func main() {

	// Read input file
	lines, err := toolbox.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Part 1
	solution_part_1 := solve_part_1(lines)

	fmt.Printf("Solution for part 1 is %v", solution_part_1)
}

func solve_part_1(lines []string) int {
	var increasedCount int
	var previousNumber int
	for i, line := range lines {
		// Skip first line
		if i == 0 {
			continue
		}

		// Get number with ID i
		currentNumber, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert input to int")
		}

		// Get number with ID i-1
		previousNumber, err = strconv.Atoi(lines[i-1])
		if err != nil {
			log.Fatalf("Failed to convert input to int")
		}

		if currentNumber > previousNumber {
			increasedCount++
		}
	}

	return increasedCount
}
