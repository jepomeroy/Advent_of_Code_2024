package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type trailHead struct {
	x int
	y int
}

func findTrailHeads(topoMap []int, width int) []trailHead {
	trailHeads := []trailHead{}

	for i, v := range topoMap {
		if v == 0 {
			trailHeads = append(trailHeads, trailHead{x: i % width, y: i / width})
		}
	}

	return trailHeads
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

	width := len(strings.Split(string(data), "\n")[0])
	height := len(strings.Split(string(data), "\n")) - 1

	fmt.Println(width, height)

	topoMap := []int{}
	for _, r := range string(data) {
		if r == '\n' {
			continue
		}

		num, err := strconv.Atoi(string(r))
		check(err, "Error converting to int")
		topoMap = append(topoMap, num)
	}

	part1(topoMap, width, height)
	// part2(topoMap)
}

func part1(topoMap []int, width int, height int) {
	value := 0

	trailHeads := findTrailHeads(topoMap, width)
	fmt.Println(topoMap)
	fmt.Println(trailHeads)

	fmt.Println(value)
}

func part2(diskMap []int) {
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
