package main

import (
	"fmt"

	. "github.com/bblasbergjc/aoc-2022/util"
)

var priorities map[byte]int

func init() {
	priorities = make(map[byte]int)
	priority := 1
	for r := 'a'; r <= 'z'; r += 1 {
		priorities[string(r)[0]] = priority
		priority += 1
	}

	for r := 'A'; r <= 'Z'; r += 1 {
		priorities[string(r)[0]] = priority
		priority += 1
	}
}

func split(items string) (string, string) {
	half := len(items) / 2
	return items[:half], items[half:]
}

// find common letter between halves
func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		first, second := split(line)

		firstItems := make(map[byte]struct{})
		secondItems := make(map[byte]struct{})

		// make a set for each half of all their letters
		for i := range first {
			firstItems[first[i]] = struct{}{}
			secondItems[second[i]] = struct{}{}
		}

		// look for common letters and sum their priorities
		for key := range firstItems {
			if _, ok := secondItems[key]; ok {
				total += priorities[key]
			}
		}
	}
	return total
}

// find common letter between each group of three
func partTwo(lines []string) int {
	total := 0

	groupItems := make([]map[byte]struct{}, 0)
	for _, line := range lines {
		items := make(map[byte]struct{})
		// make a set of all their letters
		for i := range line {
			items[line[i]] = struct{}{}
		}

		if len(groupItems) == 2 { // we have all letters from the group, search for common
			for k := range items {
				_, firstHas := groupItems[0][k]
				_, sndHas := groupItems[1][k]

				if firstHas && sndHas {
					total += priorities[k]
					break
				}
			}

			// reset for next group
			groupItems = make([]map[byte]struct{}, 0)
		} else {
			groupItems = append(groupItems, items)
		}
	}
	return total
}
func main() {
	lines := ParseLines("./day3.txt")

	fmt.Println("Part 1: ", partOne(lines))
	fmt.Println("Part 2: ", partTwo(lines))
}
