package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	// "strconv"
	"strings"
)

type IntBoolPair struct {
	IntValue  int
	BoolValue bool
}

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

	// part1(string(data))
	part2(string(data))
}

func part1(data string) {
	value := 0

	indices := findSubstrings(data, "mul(")

	for _, index := range indices {
		length, err := validateInstructions(data, index)

		if err != nil {
			continue
		}

		value += processInstruction(data[index : index+length])
	}

	fmt.Println(value)
}

func part2(data string) {
	value := 0
	indices := findSubstrings(data, "mul(")
	enabledList := getEnabledList(data)

	fmt.Println(data)
	fmt.Println(indices)
	fmt.Printf("%+v\n", enabledList)

	for _, index := range indices {
		length, err := validateInstructions(data, index)

		if err != nil {
			continue
		}

		if checkEnabled(index, enabledList) {
			fmt.Printf("Processing index %d\n", index)
			value += processInstruction(data[index : index+length])
		} else {
			fmt.Printf("Skipping index %d\n", index)
		}
	}

	fmt.Println(value)
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}

func checkEnabled(index int, enableList []IntBoolPair) bool {
	enabled := true

	for _, pair := range enableList {
		if pair.IntValue < index {
			enabled = pair.BoolValue
		}

		if pair.IntValue > index {
			break
		}
	}

	return enabled
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

func getEnabledList(data string) []IntBoolPair {
	doList := findSubstrings(data, "do()")
	dontList := findSubstrings(data, "don't()")
	enableList := []IntBoolPair{}

	for i := 0; i < len(data); i++ {
		if slices.Contains(dontList, i) {
			enableList = append(enableList, IntBoolPair{i, false})
		} else if slices.Contains(doList, i) {
			enableList = append(enableList, IntBoolPair{i, true})
		}
	}

	return enableList
}

func processInstruction(data string) int {
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
