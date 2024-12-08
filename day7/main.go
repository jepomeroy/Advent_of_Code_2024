package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	answer int
	values []int
}

func (eq equation) solvable() int {
	fmt.Println()
	fmt.Println("answer", eq.answer)
	fmt.Println("values", eq.values)

	for opset := range int(math.Pow(2, float64(len(eq.values)-1))) {
		if eq.solve(opset) == eq.answer {
			return eq.answer
		}
	}

	return 0
}

func (eq equation) solve(opset int) int {
	accumulator := eq.values[0]
	str := ""

	for i := 1; i < len(eq.values); i++ {
		if opset&(1<<(i-1)) != 0 {
			// multiplication
			str += "*"
			accumulator *= eq.values[i]
		} else {
			// addition
			str += "+"
			accumulator += eq.values[i]
		}
	}

	fmt.Printf("%s = %d\n", str, accumulator)
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

	part1(equations)
	// part2(equations)
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

	fmt.Println(value)
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}
