package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

func readSignalStrength(x, strengthCheckCycle, totalSignalStrengths int) (int, int) {
	totalSignalStrengths += (x * strengthCheckCycle)
	strengthCheckCycle += 40

	return totalSignalStrengths, strengthCheckCycle
}

func partOne(lines []string) int {
	x := 1
	cycle := 1
	strengthCheckCycle := 20  // the cycle we will check the signal strength at
	totalSignalStrengths := 0 // the sum of all our signal strength checks
	for _, line := range lines {
		if line == "noop" {
			cycle += 1

			if cycle >= strengthCheckCycle {
				totalSignalStrengths, strengthCheckCycle = readSignalStrength(x, strengthCheckCycle, totalSignalStrengths)
			}
		} else { // it's an add operation
			addVal, err := strconv.Atoi(strings.Split(line, " ")[1])
			CheckErr(err)

			// if this add wont go through before we need to check signal strength, do it now
			if cycle+2 > strengthCheckCycle {
				totalSignalStrengths, strengthCheckCycle = readSignalStrength(x, strengthCheckCycle, totalSignalStrengths)
			}

			x += addVal
			cycle += 2 // takes 2 cycles to add

			if cycle >= strengthCheckCycle {
				totalSignalStrengths, strengthCheckCycle = readSignalStrength(x, strengthCheckCycle, totalSignalStrengths)
			}
		}
	}
	return totalSignalStrengths
}

func draw(x, cycle int) string {
	position := (cycle - 1) % 40

	ret := ""
	if x-1 == position || x == position || x+1 == position {
		ret += "#"
	} else {
		ret += "."
	}

	if cycle > 1 && cycle%40 == 0 { // this was the final position on this row
		ret += "\n"
	}

	return ret
}

func partTwo(lines []string) {
	x := 1
	cycle := 1
	screen := ""
	for _, line := range lines {
		if line == "noop" {
			screen += draw(x, cycle)
			cycle += 1

		} else { // it's an add operation
			addVal, err := strconv.Atoi(strings.Split(line, " ")[1])
			CheckErr(err)

			screen += draw(x, cycle)
			cycle += 1

			screen += draw(x, cycle)
			x += addVal
			cycle += 1
		}
	}

	fmt.Println(screen)
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day10.txt") // trim empty last line

	fmt.Println("Part 1:", partOne(lines))
	partTwo(lines)
}
