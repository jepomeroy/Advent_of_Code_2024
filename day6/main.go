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
	idx := strings.Index(data, "^")
	if idx != -1 {
		return guard{row: idx / width, col: idx % width, direction: 0}, nil
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

func detectLoop(visited map[int][]int, guard guard, width int) bool {
	if visited[guard.row*width+guard.col] == nil {
		return false
	}

	for _, dir := range visited[guard.row*width+guard.col] {
		if dir == guard.direction {
			return true
		}
	}

	return false
}

func trapGuard(data string, guard guard, trap int, height int, width int) error {
	// mark each spot the guard has turned at, once we get a repeat, they're in a loop
	visited := map[int][]int{}

	for {
		if willExit(guard, height, width) {
			return fmt.Errorf("Guard is not in a loop")
		}

		switch guard.direction {
		case 0:
			if data[(guard.row-1)*width+guard.col] == '#' || (guard.row-1)*width+guard.col == trap {
				if detectLoop(visited, guard, width) {
					return nil
				}

				visited[guard.row*width+guard.col] = append(visited[guard.row*width+guard.col], guard.direction)
				guard.direction = 1
			} else {
				guard.row--
			}
		case 1:
			if data[guard.row*width+guard.col+1] == '#' || guard.row*width+guard.col+1 == trap {
				if detectLoop(visited, guard, width) {
					return nil
				}

				visited[guard.row*width+guard.col] = append(visited[guard.row*width+guard.col], guard.direction)
				guard.direction = 2
			} else {
				guard.col++
			}
		case 2:
			if data[(guard.row+1)*width+guard.col] == '#' || (guard.row+1)*width+guard.col == trap {
				if detectLoop(visited, guard, width) {
					return nil
				}

				visited[guard.row*width+guard.col] = append(visited[guard.row*width+guard.col], guard.direction)
				guard.direction = 3
			} else {
				guard.row++
			}
		case 3:
			if data[guard.row*width+guard.col-1] == '#' || guard.row*width+guard.col-1 == trap {
				if detectLoop(visited, guard, width) {
					return nil
				}

				visited[guard.row*width+guard.col] = append(visited[guard.row*width+guard.col], guard.direction)
				guard.direction = 0
			} else {
				guard.col--
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

	// part1(string(buffer), height, width)
	part2(string(buffer), height, width)
}

func part1(data string, height int, width int) {
	value := 0
	guard, err := findGuard(data, width)

	check(err, "No guard found")

	value = trackGuard(data, guard, height, width)

	fmt.Println(value)
}

func part2(data string, height int, width int) {
	value := 0
	guard, err := findGuard(data, width)

	check(err, "No guard found")

	for trap, b := range data {
		if b != '#' && b != '^' {
			fmt.Println("Checking trap at", trap)
			err := trapGuard(data, guard, trap, height, width)

			if err == nil {
				value++
			}
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
