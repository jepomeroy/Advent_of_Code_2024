package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type antenna struct {
	row int
	col int
}

func (a antenna) makeAntiNode(rowDiff, colDiff int, first bool) antenna {
	row, col := 0, 0
	if first {
		row = a.row + rowDiff
		col = a.col + colDiff
	} else {
		row = a.row - rowDiff
		col = a.col - colDiff
	}

	return antenna{row: row, col: col}
}

func (a antenna) isAntiNodeValid(width int, height int) bool {
	if a.row < 0 || a.row >= height {
		return false
	}

	if a.col < 0 || a.col >= width {
		return false
	}

	return true
}

func (a antenna) findNextAntiNode(rowDiff, colDiff, width, height int, first bool) []antenna {
	antiNodes := []antenna{}
	nextNode := a.makeAntiNode(rowDiff, colDiff, first)

	if nextNode.isAntiNodeValid(width, height) {
		antiNodes = append(antiNodes, nextNode)
		antiNodes = append(antiNodes, nextNode.findNextAntiNode(rowDiff, colDiff, width, height, first)...)
	} else {
		return antiNodes
	}

	return antiNodes
}

func (a antenna) findAntiNodes(b antenna, width int, height int) []antenna {
	validAntiNodes := []antenna{}
	colDiff := a.col - b.col
	rowDiff := a.row - b.row

	validAntiNodes = append(validAntiNodes, a)

	antiNode1 := a.makeAntiNode(rowDiff, colDiff, true)
	if antiNode1.isAntiNodeValid(width, height) {
		fmt.Printf("antinode1: %v\n", antiNode1)
		validAntiNodes = append(validAntiNodes, antiNode1)
		validAntiNodes = append(validAntiNodes, antiNode1.findNextAntiNode(rowDiff, colDiff, width, height, true)...)
	}

	antiNode2 := b.makeAntiNode(rowDiff, colDiff, false)
	if antiNode2.isAntiNodeValid(width, height) {
		fmt.Printf("antinode2: %v\n", antiNode2)
		validAntiNodes = append(validAntiNodes, antiNode2)
		validAntiNodes = append(validAntiNodes, antiNode2.findNextAntiNode(rowDiff, colDiff, width, height, false)...)
	}

	return validAntiNodes
}

type antennaData struct {
	width            int
	height           int
	antennaLocations map[rune][]int
}

func findNodes(antData antennaData) int {
	antiNodes := []antenna{}

	// Add all the antennas to the antiNodes list
	for _, value := range antData.antennaLocations {
		for _, position := range value {
			antenna := caculateRowCol(position, antData.width)
			antiNodes = append(antiNodes, antenna)
		}
	}

	for _, value := range antData.antennaLocations {
		for i := 0; i < len(value)-1; i++ {
			antenna1 := caculateRowCol(value[i], antData.width)
			for j := i + 1; j < len(value); j++ {
				antenna2 := caculateRowCol(value[j], antData.width)

				nodes := antenna1.findAntiNodes(antenna2, antData.width, antData.height)

				for _, node := range nodes {
					found := false
					for _, antiNode := range antiNodes {
						if node.row == antiNode.row && node.col == antiNode.col {
							found = true
							break
						}
					}

					if !found {
						antiNodes = append(antiNodes, node)
					}
				}
			}
		}

	}

	return len(antiNodes)
}

func caculateRowCol(position, width int) antenna {
	return antenna{row: position / width, col: position % width}
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

	antData := antennaData{
		width:            len(strings.Split(string(data), "\n")[0]),
		height:           len(strings.Split(string(data), "\n")) - 1,
		antennaLocations: make(map[rune][]int),
	}

	cleanData := ""
	for i, line := range strings.Split(string(data), "\n") {
		for j, char := range line {
			if char == '\n' {
				continue
			}
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				antData.antennaLocations[char] = append(antData.antennaLocations[char], i*antData.width+j)
			}
		}
		cleanData += line
	}

	part1(string(data), antData)
	// part2(data)
}

func part1(data string, antData antennaData) {
	value := 0

	value = findNodes(antData)

	fmt.Println(value)
}

func part2(left []int, hash map[int]int) {
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
