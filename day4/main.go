package main

import (
	"fmt"
	"os"
	"strings"
)

type finder func(string, int, int, int) int

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
	count := getCount(data, 'X', findWord)
	fmt.Println(count)
}

func part2(data string) {
	count := getCount(data, 'A', findXWord)
	fmt.Println(count)
}

func getCount(data string, char rune, fn finder) int {
	count := 0

	height := strings.Count(data, "\n")
	width := len(strings.Split(data, "\n")[0]) + 1

	for i, c := range data {
		if c == char {
			count += fn(data, i, width, height)
		}
	}

	return count
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}

// Check if the character at the given index is the same as the given character
func checkDirection(data string, idx int, width int, height int, dir int) bool {
	column := idx % width
	row := idx / width

	switch dir {
	case 0:
		// need room for 3 characters above
		if row < 3 {
			return false
		}

		return data[idx-width] == 'M' && data[idx-(2*width)] == 'A' && data[idx-(3*width)] == 'S'
	case 1:
		// need room for 3 characters above and 3 characters to the right
		if row < 3 || column > width-4 {
			return false
		}

		return data[idx-width+1] == 'M' && data[idx-(2*width)+2] == 'A' && data[idx-(3*width)+3] == 'S'
	case 2:
		// need room for 3 characters to the right
		if column > width-4 {
			return false
		}

		return data[idx+1] == 'M' && data[idx+2] == 'A' && data[idx+3] == 'S'
	case 3:
		// need room for 3 characters below and 3 characters to the right
		if row > height-4 || column > width-4 {
			return false
		}

		return data[idx+width+1] == 'M' && data[idx+(2*width)+2] == 'A' && data[idx+(3*width)+3] == 'S'
	case 4:
		// need room for 3 characters below
		if row > height-4 {
			return false
		}

		return data[idx+width] == 'M' && data[idx+(2*width)] == 'A' && data[idx+(3*width)] == 'S'
	case 5:
		// need room for 3 characters below and 3 characters to the left
		if row > height-4 || column < 3 {
			return false
		}

		return data[idx+width-1] == 'M' && data[idx+(2*width)-2] == 'A' && data[idx+(3*width)-3] == 'S'
	case 6:
		// need room for 3 characters to the left
		if column < 3 {
			return false
		}

		return data[idx-1] == 'M' && data[idx-2] == 'A' && data[idx-3] == 'S'
	case 7:
		// need room for 3 characters above and 3 characters to the left
		if row < 3 || column < 3 {
			return false
		}

		return data[idx-width-1] == 'M' && data[idx-(2*width)-2] == 'A' && data[idx-(3*width)-3] == 'S'
	default:
		// Invalid direction
		return false
	}
}

func findWord(data string, idx int, width int, height int) int {
	wordsFound := 0
	for d := range 8 {
		if checkDirection(data, idx, width, height, d) {
			wordsFound++
		}
	}

	return wordsFound
}

// Check if the character at the given index is the same as the given character
func findXWord(data string, idx int, width int, height int) int {
	column := idx % width
	row := idx / width

	// not the first row, not the last row, not the first column, not the last column
	if row < 1 || row > height-2 || column < 1 || column > width-1 {
		return 0
	}

	if (data[idx-width-1] == 'M' && data[idx+width+1] == 'S') || (data[idx-width-1] == 'S' && data[idx+width+1] == 'M') {
		if (data[idx-width+1] == 'M' && data[idx+width-1] == 'S') || (data[idx-width+1] == 'S' && data[idx+width-1] == 'M') {
			return 1
		}
	}

	return 0
}
