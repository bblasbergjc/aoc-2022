package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type operation func(old int) int

type test func(x int) int

type Monkey struct {
	Items     []int
	Operation operation
	Test      test
}

// expects an operation string. Ex: old + 7
func parseOperation(operation string) operation {
	//special case
	if operation == "old * old" {
		return func(old int) int {
			return old * old
		}
	}

	parts := strings.Split(operation, " ")
	op := parts[1]
	rawNum := parts[2]
	num, err := strconv.Atoi(rawNum)
	checkErr(err)

	if op == "+" {
		return func(old int) int {
			return old + num
		}
	} else {
		return func(old int) int {
			return old * num
		}
	}
}

func parseLastNumber(line string) int {
	parts := strings.Split(line, " ")
	rawNum := parts[len(parts)-1]

	num, err := strconv.Atoi(rawNum)
	checkErr(err)

	return num
}

func parseMonkeys(lines []string) map[int]*Monkey {
	monkies := make(map[int]*Monkey)
	currentMonkey := 0
	for i := 0; i < len(lines); i += 1 {
		trimmedLine := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmedLine, "Monkey") {
			monkies[currentMonkey] = &Monkey{}
		} else if strings.HasPrefix(trimmedLine, "Starting items") {
			itemsRaw := strings.TrimPrefix(trimmedLine, "Starting items: ")

			items := make([]int, 0)
			for _, item := range strings.Split(itemsRaw, ", ") {
				num, err := strconv.Atoi(item)
				checkErr(err)

				items = append(items, num)
			}

			monkies[currentMonkey].Items = items
		} else if strings.HasPrefix(trimmedLine, "Operation") {
			rawOp := strings.TrimPrefix(trimmedLine, "Operation: new = ")
			op := parseOperation(rawOp)

			monkies[currentMonkey].Operation = op
		} else if strings.HasPrefix(trimmedLine, "Test") {
			divisor := parseLastNumber(trimmedLine)

			i += 1
			trueCase := parseLastNumber(strings.TrimSpace(lines[i]))
			i += 1
			falseCase := parseLastNumber(strings.TrimSpace(lines[i]))

			monkies[currentMonkey].Test = func(x int) int {
				if x%divisor == 0 {
					return trueCase
				} else {
					return falseCase
				}
			}
		} else {
			currentMonkey += 1
		}
	}

	return monkies
}

func partOne(monkies map[int]*Monkey) int {
	inspections := make([]int, len(monkies))
	for round := 0; round < 20; round += 1 {
		for i := 0; i < len(monkies); i += 1 {
			monkey := monkies[i]

			for _, worryLevel := range monkey.Items {
				newWorryLevel := monkey.Operation(worryLevel)
				newWorryLevel = newWorryLevel / 3

				toMonkey := monkey.Test(newWorryLevel)
				monkies[toMonkey].Items = append(monkies[toMonkey].Items, newWorryLevel)

				inspections[i] += 1
			}

			monkey.Items = make([]int, 0)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	// multiply hightest 2
	return inspections[0] * inspections[1]
}

// tried to big.Int this :) apparently math is required :(
func partTwo(monkies map[int]*Monkey) int {
	return 0
}

func main() {
	data, err := os.ReadFile("./day11.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1] // trim empty last line

	monkies := parseMonkeys(lines)

	fmt.Println("Part 1:", partOne(monkies))
	fmt.Println("Part 2:", partTwo(monkies))
}
