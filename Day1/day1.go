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
	fmt.Printf("Solution for part 1 is %v\n", solution_part_1)

	// Part 2
	solution_part_2 := solve_part_2(lines)
	fmt.Printf("Solution for part 2 is %v\n", solution_part_2)
}

func solve_part_1(lines []string) int {
	var increasedCount int

	for i, line := range lines {
		// Skip first line
		if i == 0 {
			continue
		}

		// Get number with ID i
		currentNumber, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert string to int")
		}

		// Get number with ID i-1
		previousNumber, err := strconv.Atoi(lines[i-1])
		if err != nil {
			log.Fatalf("Failed to convert string to int")
		}

		// Check for increase
		if currentNumber > previousNumber {
			increasedCount++
		}
	}

	return increasedCount
}

func solve_part_2(lines []string) int {
	var increasedCount int

	for i := range lines {
		// Skip first 2 lines and last line
		if i < 2 || i == len(lines)-1 {
			continue
		}

		// Get sliding sum for ID i
		currentSum := get_sliding_sum(lines, i)

		// Get sliding sum for ID i-1
		previousSum := get_sliding_sum(lines, i-1)

		// Check for increase
		if currentSum > previousSum {
			increasedCount++
		}
	}

	return increasedCount
}

func get_sliding_sum(lines []string, index int) int {
	var slidingSum int

	// Fixed window size of 3
	for i := index - 1; i <= index+1; i++ {
		number, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatalf("Failed to convert string to int")
		}
		slidingSum += number
	}

	return slidingSum
}
