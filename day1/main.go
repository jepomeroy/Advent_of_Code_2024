package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
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

	lines := strings.Split(string(data), "\n")

	left, right := []int{}, []int{}
	hash := make(map[int]int)

	for _, line := range lines {
		nums := strings.Fields(line)

		if len(nums) == 2 {
			l, err := strconv.Atoi(nums[0])
			check(err, "Error converting string to int")
			left = append(left, l)

			r, err := strconv.Atoi(nums[1])
			check(err, "Error converting string to int")
			right = append(right, r)
			hash[r]++
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	// part1(left, right)
	part2(left, hash)
}

func part1(left, right []int) {
	value := 0

	for i := 0; i < len(left); i++ {
		v := left[i] - right[i]

		value += int(math.Abs(float64(v)))
	}

	fmt.Println(value)
}

func part2(left []int, hash map[int]int) {
	value := 0

	for _, v := range left {
		value += v * hash[v]
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
