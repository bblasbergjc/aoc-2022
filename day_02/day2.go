package main

import (
	"fmt"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

var pointsMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

/*    X Y Z
 *  A 3 0 6
 *  B 6 3 0
 *  C 0 6 3
 */
var winMatrix = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 0,
		"C": 6,
	},

	"Y": {
		"A": 6,
		"B": 3,
		"C": 0,
	},
	"Z": {
		"A": 0,
		"B": 6,
		"C": 3,
	},
}

var outcomeMatrix = map[string]map[string]string{
	"A": {
		"X": "C", // lose
		"Y": "A", // draw
		"Z": "B", // win
	},

	"B": {
		"X": "A",
		"Y": "B",
		"Z": "C",
	},
	"C": {
		"X": "B",
		"Y": "C",
		"Z": "A",
	},
}

func calcPoints(mine, theirs string) int {
	gamePoints := winMatrix[mine][theirs]
	choicePoints := pointsMap[mine]

	return gamePoints + choicePoints
}

func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		if len(choices) < 2 {
			break
		}
		total += calcPoints(choices[1], choices[0])
	}

	return total
}

func partTwo(lines []string) int {

	outcomeScores := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	total := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		if len(choices) < 2 {
			break
		}

		theirs := choices[0]
		outcome := choices[1]
		choice := outcomeMatrix[theirs][outcome]
		pointsScored := pointsMap[choice] + outcomeScores[outcome]

		total += pointsScored
	}

	return total
}

func main() {
	lines := ParseLines("./day2.txt")

	fmt.Printf("Total Score: %d\n", partOne(lines))
	fmt.Printf("Total Score: %d\n", partTwo(lines))
}
