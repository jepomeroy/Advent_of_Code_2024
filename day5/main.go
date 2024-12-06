package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type rulesType map[int][]int

type update struct {
	pageNums []int
}

func (u *update) fixUpdate(rules rulesType) {
	sort.Slice(u.pageNums, func(i, j int) bool {
		left := u.pageNums[i]
		right := u.pageNums[j]
		rule, found := rules[left]

		if found && slices.Contains(rule, right) {
			return true
		}

		return false
	})
}

func (u update) validate(rules rulesType) bool {
	for i := 0; i < len(u.pageNums)-1; i++ {
		rules, found := rules[u.pageNums[i]]

		if !found {
			return false
		}

		if !slices.Contains(rules, u.pageNums[i+1]) {
			return false
		}
	}

	return true
}

func (u update) getMiddlePage() int {
	idx := int(math.Ceil(float64(len(u.pageNums) / 2)))

	return u.pageNums[idx]
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

	foundBreak := false
	updates := []update{}
	rules := rulesType{}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			foundBreak = true
			continue
		}

		if !foundBreak {
			parts := strings.Split(line, "|")
			first, err := strconv.Atoi(parts[0])
			check(err, "Error converting first part to int")

			second, err := strconv.Atoi(parts[1])
			check(err, "Error converting second part to int")

			rules[first] = append(rules[first], second)
		} else {
			newUpdate := update{pageNums: []int{}}
			for _, num := range strings.Split(line, ",") {
				num, err := strconv.Atoi(num)
				check(err, "Error converting update to int")
				newUpdate.pageNums = append(newUpdate.pageNums, num)
			}

			updates = append(updates, newUpdate)
		}
	}

	// part1(updates, rules)
	part2(updates, rules)
}

func part1(updates []update, rules rulesType) {
	value := 0
	for _, update := range updates {
		if update.validate(rules) {
			value += update.getMiddlePage()
		}
	}

	fmt.Println(value)
}

func part2(updates []update, rules rulesType) {
	value := 0

	for _, update := range updates {
		if !update.validate(rules) {
			update.fixUpdate(rules)
			fmt.Printf("Fixed update: %v\n", update)

			value += update.getMiddlePage()
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
