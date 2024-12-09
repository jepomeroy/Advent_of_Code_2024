package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeDiskBlocks(length int, block int) []int {
	blocks := []int{}

	// fmt.Printf("length: %d, block: %c\n", length, block)

	for i := 0; i < length; i++ {
		blocks = append(blocks, block)
	}

	return blocks
}

func expandDiskMap(diskMap []int) []int {
	expandedMap := []int{}
	id := 0
	createBlocks := true
	var block int

	for _, v := range diskMap {
		if createBlocks {
			block = id
			id++
		} else {
			block = -1
		}
		expandedMap = append(expandedMap, makeDiskBlocks(v, block)...)
		createBlocks = !createBlocks
	}

	// fmt.Println(string(expandedMap))
	return expandedMap
}

func defragDiskMap(diskMap *[]int) {
	i, j := 0, len(*diskMap)-1

	for i < j {
		if (*diskMap)[i] != -1 {
			i++
		} else if (*diskMap)[j] == -1 {
			j--
		} else {
			// fmt.Println(string(*diskMap))
			(*diskMap)[i], (*diskMap)[j] = (*diskMap)[j], (*diskMap)[i]
			i++
			j--
		}
	}

	// fmt.Println(string(*diskMap))
}

func calculateCheckSum(diskMap []int) int {
	checksum := 0

	for i, v := range diskMap {
		if v == -1 {
			break
		}

		checksum += (v * i)

		// fmt.Printf("i: %d, d: %d, checksum: %d\n", i, d, checksum)
	}

	return checksum
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

	diskMap := []int{}
	for _, v := range data {
		if v == '\n' {
			continue
		}

		d, err := strconv.Atoi(string(v))
		if err != nil {
			check(err, "Error converting to int")
		}
		diskMap = append(diskMap, d)
	}

	part1(diskMap)
	// part2(diskMap)
}

func part1(diskMap []int) {
	value := 0

	expandedMap := expandDiskMap(diskMap)

	defragDiskMap(&expandedMap)

	value = calculateCheckSum(expandedMap)

	fmt.Println(value)
}

func part2(diskMap []int) {
	value := 0

	fmt.Println(diskMap)
	fmt.Println(value)
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
		os.Exit(1)
	}
}
