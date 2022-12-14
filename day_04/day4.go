package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

type Range struct {
	Low  int
	High int
}

func getRanges(line string) (Range, Range) {
	ranges := strings.Split(line, ",")

	lVals := strings.Split(ranges[0], "-")
	rVals := strings.Split(ranges[1], "-")

	createRange := func(vals []string) Range {
		low, err := strconv.Atoi(vals[0])
		CheckErr(err)

		high, err := strconv.Atoi(vals[1])
		CheckErr(err)
		return Range{
			Low:  low,
			High: high,
		}
	}

	return createRange(lVals), createRange(rVals)
}

// ranges must completely contain other range
func partOne(lines []string) int {
	totalOverlapping := 0
	for _, line := range lines {

		left, right := getRanges(line)

		if left.Low <= right.Low && left.High >= right.High {
			totalOverlapping += 1
		} else if right.Low <= left.Low && right.High >= left.High {
			totalOverlapping += 1
		}
	}
	return totalOverlapping
}

// ranges have any overlap
func partTwo(lines []string) int {
	totalOverlapping := 0
	for _, line := range lines {

		left, right := getRanges(line)
		if left.High >= right.Low && left.High <= right.High {
			totalOverlapping += 1
		} else if right.Low >= left.Low && right.Low <= left.High {
			totalOverlapping += 1
		} else if left.Low >= right.Low && left.Low <= right.High {
			totalOverlapping += 1
		} else if right.High <= left.High && right.High >= left.Low {
			totalOverlapping += 1
		}
	}
	return totalOverlapping
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day4.txt")

	fmt.Println("Part 1: ", partOne(lines))
	fmt.Println("Part 2: ", partTwo(lines))
}
