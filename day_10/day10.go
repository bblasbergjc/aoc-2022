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

func readSignalStrength(x, strengthCheckCycle, totalSignalStrengths int) (int, int) {
	totalSignalStrengths += (x * strengthCheckCycle)
	strengthCheckCycle += 40

	return totalSignalStrengths, strengthCheckCycle
}

func partOne(lines []string) int {
	x := 1
	cycle := 1
	strengthCheckCycle := 20
	totalSignalStrengths := 0
	for _, line := range lines {
		if line == "noop" {
			cycle += 1

			if cycle >= strengthCheckCycle {
				totalSignalStrengths, strengthCheckCycle = readSignalStrength(x, strengthCheckCycle, totalSignalStrengths)
			}
		} else { // it's an add operation
			addVal, err := strconv.Atoi(strings.Split(line, " ")[1])
			checkErr(err)

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

func main() {
	data, err := os.ReadFile("./day10.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1] // trim empty last line

	fmt.Println("Part 1:", partOne(lines))
}
