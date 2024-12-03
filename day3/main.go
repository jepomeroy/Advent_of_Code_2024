package main

import (
	"fmt"
	"os"
	"strconv"

	// "strconv"
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

	part1(string(data))
	// part2(data)
}

func part1(data string) {
	value := 0

	indices := findSubstrings(data, "mul(")

	fmt.Println("Indices: ", indices)

	for _, index := range indices {
		length, err := validateInstructions(data, index)

		if err != nil {
			continue
		}

		fmt.Println("Length: ", length)

		value += processInstruction(data[index : index+length])

		fmt.Println("Value: ", value, "\n")
	}

	fmt.Println(value)
}

func part2(data string) {
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}

func findSubstrings(str, substr string) []int {
	var indices []int

	for i := 0; i < len(str); {
		index := strings.Index(str[i:], substr)
		if index == -1 {
			break // No more occurrences found
		}

		indices = append(indices, i+index)
		i += index + len(substr) // Move to next potential position
	}

	return indices
}

func processInstruction(data string) int {
	fmt.Println("Processing: ", data)
	nums := data[len("mul(") : len(data)-1]

	f, err := strconv.Atoi(strings.Split(nums, ",")[0])
	check(err, "Error converting first number")

	s, err := strconv.Atoi(strings.Split(nums, ",")[1])
	check(err, "Error converting second number")

	return f * s
}

func validateInstructions(data string, offset int) (int, error) {
	// check up to 12 character in search; 3 digits for each number,
	// 1 comma and 1 closing bracket, plus the "mul(" prefix
	for i := 4; i < 12; i++ {
		fmt.Println("Checking: ", data[offset+i])
		switch data[offset+i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', ',':
			continue
		case ')':
			fmt.Println("Found closing bracket")
			return i + 1, nil
		default:
			return 0, fmt.Errorf("Invalid character found")
		}
	}

	return 0, fmt.Errorf("No closing bracket found")
}
