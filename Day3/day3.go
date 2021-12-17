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
	solutionPart1 := solvePart1(lines)
	fmt.Printf("Solution for part 1 is %v\n", solutionPart1)

	// Part 2
	solutionPart2 := solvePart2(lines)
	fmt.Printf("Solution for part 2 is %v\n", solutionPart2)
}

func solvePart1(lines []string) int {
	sumPerColumn := getSumPerColumn(lines)

	fmt.Printf("sumPerColumn=%v\n", sumPerColumn)

	gammaRateBinary := calculateGammaRate(sumPerColumn, len(lines)/2)
	epsilonRateBinary := getInverseBinary(gammaRateBinary)

	fmt.Printf("gammaBinary=%v, epsilonBinary=%v\n", gammaRateBinary, epsilonRateBinary)

	gammaRateDecimal := convertToDecimal(gammaRateBinary)
	epsilonRateDecimal := convertToDecimal(epsilonRateBinary)

	fmt.Printf("gammaDecimal=%v, epsilonDecimal=%v\n", gammaRateDecimal, epsilonRateDecimal)

	return gammaRateDecimal * epsilonRateDecimal
}

func solvePart2(lines []string) int {

	// Find oxygen generator rating
	oxygenGeneratorRatingBinary := getFirstMatchingBinary(lines, false)
	co2ScrubberRatingBinary := getFirstMatchingBinary(lines, true)

	fmt.Printf("oxygenGeneratorRatingBinary=%v, co2ScrubberRatingBinary=%v\n", oxygenGeneratorRatingBinary, co2ScrubberRatingBinary)

	oxygenGeneratorRatingDecimal := convertToDecimal(oxygenGeneratorRatingBinary)
	co2ScrubberRatingDecimal := convertToDecimal(co2ScrubberRatingBinary)

	fmt.Printf("oxygenGeneratorRatingDecimal=%v, co2ScrubberRatingDecimal=%v\n", oxygenGeneratorRatingDecimal, co2ScrubberRatingDecimal)

	return oxygenGeneratorRatingDecimal * co2ScrubberRatingDecimal
}

func getSumPerColumn(lines []string) []int {
	// Initialise array with 0s
	sumPerColumn := make([]int, len(lines[0]))

	// For each line, add each bit value to the respective sum
	for _, binary := range lines {
		// Parse bits and add to correct sum
		for i, bit := range binary {
			if bit == '1' {
				sumPerColumn[i]++
			}
		}
	}

	return sumPerColumn
}

func calculateGammaRate(sumPerColumn []int, average int) string {
	result := ""

	// For each sum, compare with average to see if it's most common or not
	for _, sum := range sumPerColumn {
		if sum > average {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result
}

func getInverseBinary(binary string) string {
	result := ""

	// For each bit, write opposite bit to new binary
	for _, char := range binary {
		if char == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}

func convertToDecimal(binary string) int {
	if x, err := strconv.ParseInt(binary, 2, 64); err != nil {
		fmt.Println(err)
		return -1
	} else {
		return int(x)
	}
}

func getFirstMatchingBinary(binaries []string, reverseLogic bool) string {
	lengthOfBinaries := len(binaries[0])
	validBinaries := binaries
	i := 0

	for len(validBinaries) > 1 && i <= lengthOfBinaries {
		var tmpSlice []string

		// Find bit to use for filtering binaries
		bitToKeep := getMostCommonBit(i, validBinaries, reverseLogic)
		// Filter binaries to only keep ones containing the bit to keep in the current column index
		for _, binary := range validBinaries {
			if rune(binary[i]) == bitToKeep {
				tmpSlice = append(tmpSlice, binary)
			}
		}

		i++
		validBinaries = tmpSlice
	}

	return validBinaries[0]
}

func getMostCommonBit(indexToCompare int, binaries []string, reverseLogic bool) rune {
	balance := 0
	for _, binary := range binaries {
		if binary[indexToCompare] == '1' {
			balance++
		} else {
			balance--
		}
	}

	if !reverseLogic {
		if balance >= 0 {
			return '1'
		} else {
			return '0'
		}
	} else {
		if balance >= 0 {
			return '0'
		} else {
			return '1'
		}
	}
}
