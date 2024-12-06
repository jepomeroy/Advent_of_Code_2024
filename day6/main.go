package main

import (
	"fmt"
	"os"
	"strings"
)

type guard struct {
	row       int
	col       int
	direction int // 0 = up, 1 = right, 2 = down, 3 = left
}

func findGuard(data string, width int) (guard, error) {
	for i, b := range "^>v<" {
		idx := strings.Index(data, string(b))
		if idx != -1 {
			return guard{row: idx / width, col: idx % width, direction: i}, nil
		}
	}
	return guard{}, fmt.Errorf("No guard found")
}

func willExit(guard guard, height int, width int) bool {
	if guard.row-1 < 0 || guard.row+1 >= height || guard.col-1 < 0 || guard.col+1 >= width {
		return true
	}

	return false
}

func trackGuard(data string, guard guard, height int, width int) int {
	visited := make([]bool, height*width)
	// Start at 1 because we already visited the starting position
	value := 1
	visited[guard.row*width+guard.col] = true

	for {
		if willExit(guard, height, width) {
			return value
		}

		switch guard.direction {
		case 0:
			if data[(guard.row-1)*width+guard.col] == '#' {
				guard.direction = 1
			} else {
				guard.row--

				if !visited[guard.row*width+guard.col] {
					value++
				}
				visited[guard.row*width+guard.col] = true
			}
		case 1:
			if data[guard.row*width+guard.col+1] == '#' {
				guard.direction = 2
			} else {
				guard.col++

				if !visited[guard.row*width+guard.col] {
					value++
				}
				visited[guard.row*width+guard.col] = true
			}
		case 2:
			if data[(guard.row+1)*width+guard.col] == '#' {
				guard.direction = 3
			} else {
				guard.row++

				if !visited[guard.row*width+guard.col] {
					value++
				}
				visited[guard.row*width+guard.col] = true
			}
		case 3:
			if data[guard.row*width+guard.col-1] == '#' {
				guard.direction = 0
			} else {
				guard.col--

				if !visited[guard.row*width+guard.col] {
					value++
				}
				visited[guard.row*width+guard.col] = true
			}
		}
	}
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

	height := strings.Count(string(data), "\n")
	width := len(strings.Split(string(data), "\n")[0])
	buffer := make([]byte, 0)
	for _, b := range data {
		if b != '\n' {
			buffer = append(buffer, b)
		}
	}

	part1(string(buffer), height, width)
	// part2(updates, rules)
}

func part1(data string, height int, width int) {
	value := 0
	guard, err := findGuard(data, width)

	check(err, "No guard found")

	value = trackGuard(data, guard, height, width)

	fmt.Println(value)
}

func part2(data string) {
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
