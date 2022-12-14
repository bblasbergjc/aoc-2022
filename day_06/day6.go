package main

import (
	"fmt"
	"os"

	. "github.com/bblasbergjc/aoc-2022/util"
)

func partOne(line string) int {
	charCount := make(map[rune]int)
	previousChars := make([]rune, 0)
	for i, ch := range line {
		if i > 3 {
			// decrement count of oldest char
			removeChar := previousChars[0]
			count := charCount[removeChar]
			charCount[removeChar] = count - 1

			// remove oldest char from window
			previousChars = previousChars[1:]
			previousChars = append(previousChars, ch)

			// increment or initialize count for current char
			count, ok := charCount[ch]
			if ok {
				charCount[ch] = count + 1
			} else {
				charCount[ch] = 1
			}

			// check for unique keys. Looping here because it's only 4 values
			hasMultiple := false
			for _, v := range charCount {
				if v > 1 {
					hasMultiple = true
				}
			}

			if !hasMultiple {
				return i + 1
			}
		} else {
			previousChars = append(previousChars, ch)
			count, ok := charCount[ch]
			if ok {
				charCount[ch] = count + 1
			} else {
				charCount[ch] = 1
			}
		}
	}

	return 0
}

func partTwo(line string) int {
	charCount := make(map[rune]int)
	previousChars := make([]rune, 0)
	for i, ch := range line {
		if i > 13 {
			// decrement count of oldest char
			removeChar := previousChars[0]
			count := charCount[removeChar]
			charCount[removeChar] = count - 1

			// remove oldest char from window
			previousChars = previousChars[1:]
			previousChars = append(previousChars, ch)

			// increment or initialize count for current char
			count, ok := charCount[ch]
			if ok {
				charCount[ch] = count + 1
			} else {
				charCount[ch] = 1
			}

			// check for unique keys. Looping here because it's only 4 values
			hasMultiple := false
			for _, v := range charCount {
				if v > 1 {
					hasMultiple = true
				}
			}

			if !hasMultiple {
				return i + 1
			}
		} else {
			previousChars = append(previousChars, ch)
			count, ok := charCount[ch]
			if ok {
				charCount[ch] = count + 1
			} else {
				charCount[ch] = 1
			}
		}
	}

	return 0
}

func main() {
	data, err := os.ReadFile("./day6.txt")
	CheckErr(err)

	line := string(data)

	fmt.Println("Part 1: ", partOne(line))
	fmt.Println("Part 2: ", partTwo(line))
}
