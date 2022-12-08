package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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
		checkErr(err)

		high, err := strconv.Atoi(vals[1])
		checkErr(err)
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
		if line == "" {
			break
		}

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
		if line == "" {
			break
		}

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
	data, err := os.ReadFile("./day4.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1: ", partOne(lines))
	fmt.Println("Part 2: ", partTwo(lines))
}
