package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	Add Operation = iota
	Multiply
	Combine
)

func generateCombinations(n int) [][]Operation {
	if n == 1 {
		return [][]Operation{{}}
	}

	subCombinations := generateCombinations(n - 1)
	combinations := make([][]Operation, 0, len(subCombinations)*3)

	for _, subCombination := range subCombinations {
		for _, op := range []Operation{Add, Multiply, Combine} {
			newCombination := append([]Operation{op}, subCombination...)
			combinations = append(combinations, newCombination)
		}
	}

	return combinations
}

type equation struct {
	answer int
	values []int
}

func (eq equation) solvable() int {
	fmt.Println()
	fmt.Println("answer", eq.answer)
	fmt.Println("values", eq.values)
	fmt.Println()

	opIndex := len(eq.values)

	operations := generateCombinations(opIndex)

	for opset := range int(math.Pow(3, float64(opIndex-1))) {
		if eq.solve(opset, operations[opset]) == eq.answer {
			return eq.answer
		}
	}

	return 0
}

func (eq equation) solve(opset int, operation []Operation) int {
	accumulator := eq.values[0]
	str := ""
	fmt.Println("opset", opset)

	for i := 1; i < len(eq.values); i++ {
		switch operation[i-1] {
		case 0:
			str += "+"
			accumulator += eq.values[i]
		case 1:
			str += "*"
			accumulator *= eq.values[i]
		case 2:
			str += "|"
			tmp := strconv.Itoa(accumulator) + strconv.Itoa(eq.values[i])
			accumulator, _ = strconv.Atoi(tmp)
		}
	}

	fmt.Printf("[op: %d] %s => %d\n\n", opset, str, accumulator)

	return accumulator
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

	equations := []equation{}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			parts := strings.Split(string(line), ":")
			answer, err := strconv.Atoi(parts[0])
			check(err, "Error converting answer to int")

			values := []int{}

			for _, v := range strings.Split(strings.TrimSpace(string(parts[1])), " ") {
				value, err := strconv.Atoi(v)
				check(err, "Error converting value to int")
				values = append(values, value)
			}

			equations = append(equations, equation{answer, values})
		}
	}

	// part1(equations)
	part2(equations)
}

func part1(equations []equation) {
	value := 0

	for _, eq := range equations {
		value += eq.solvable()
	}

	fmt.Println(value)
}

func part2(equations []equation) {
	value := 0

	for _, eq := range equations {
		value += eq.solvable()
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
