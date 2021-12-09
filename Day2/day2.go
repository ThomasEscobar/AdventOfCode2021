package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	var x int
	var y int

	for _, instruction := range lines {
		// Parse instructions
		direction := strings.Fields(instruction)[0]
		value, err := strconv.Atoi(strings.Fields(instruction)[1])
		if err != nil {
			log.Fatalf("Failed to convert string to int")
		}

		// Calculate movement
		switch direction {
		case "up":
			y -= value
		case "down":
			y += value
		case "forward":
			x += value
		}
	}

	fmt.Printf("The final coordonates are %v and %v\n", x, y)
	return x * y
}

func solve_part_2(lines []string) int {
	var x int
	var y int
	var aim int

	for _, instruction := range lines {
		// Parse instructions
		direction := strings.Fields(instruction)[0]
		value, err := strconv.Atoi(strings.Fields(instruction)[1])
		if err != nil {
			log.Fatalf("Failed to convert string to int")
		}

		// Calculate movement
		switch direction {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			x += value
			y += value * aim
		}
	}

	fmt.Printf("The final coordonates are %v and %v\n", x, y)
	return x * y
}
