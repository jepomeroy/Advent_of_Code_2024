package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		check(err, "Error reading file")
	}

	lines := strings.Split(string(data), "\n")

	reports := [][]int{}

	for _, line := range lines {
		nums := strings.Fields(line)
		report := []int{}

		if len(nums) > 0 {
			for _, num := range nums {
				n, err := strconv.Atoi(num)

				if err != nil {
					check(err, "Error converting string to int")
				}

				report = append(report, n)
			}

			reports = append(reports, report)
		}
	}

	// part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	safeCount := 0

	for _, report := range reports {
		safe := testArray(report)

		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part2(reports [][]int) {
	safeCount := 0

	for _, fullReport := range reports {
		safe := testArray(fullReport)

		if !safe {
			for i := 0; i < len(fullReport); i++ {
				report := removeIndex(fullReport, i)

				safe = testArray(report)

				if safe {
					break
				}
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}

func removeIndex(s []int, index int) []int {
	newSlice := []int{}

	for i, num := range s {
		if i != index {
			newSlice = append(newSlice, num)
		}
	}

	return newSlice
}

func testArray(report []int) bool {
	safe := true
	increasing := true

	for i := 1; i < len(report); i++ {
		level1, level2 := report[i-1], report[i]

		if i == 1 {
			increasing = level1 < level2
		}

		if increasing && (level1 >= level2 || (level2-level1) > 3) {
			safe = false
			break
		}

		if !increasing && (level1 <= level2 || (level1-level2) > 3) {
			safe = false
			break
		}
	}

	return safe
}
