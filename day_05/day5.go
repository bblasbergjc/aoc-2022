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

func getInitialStacks() []stack {
	return []stack{
		{"S", "C", "V", "N"},                     //1
		{"Z", "M", "J", "H", "N", "S"},           //2
		{"M", "C", "T", "G", "J", "N", "D"},      //3
		{"T", "D", "F", "J", "W", "R", "M"},      //4
		{"P", "F", "H"},                          //5
		{"C", "T", "Z", "H", "J"},                //6
		{"D", "P", "R", "Q", "F", "S", "L", "Z"}, //7
		{"C", "S", "L", "H", "D", "F", "P", "W"}, //8
		{"D", "S", "M", "P", "F", "N", "G", "Z"}, //9
	}
}

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// returns numberToMove, from, to
func parseLine(line string) (int, int, int) {
	tokens := strings.Split(line, " ")

	numberToMove, err := strconv.Atoi(tokens[1])
	checkErr(err)

	from, err := strconv.Atoi(tokens[3])
	checkErr(err)

	to, err := strconv.Atoi(tokens[5])
	checkErr(err)

	return numberToMove, from - 1, to - 1
}

func partOne(lines []string) string {
	stacks := getInitialStacks()

	for _, line := range lines {
		if line == "" {
			break
		}

		numberToMove, from, to := parseLine(line)

		for i := 0; i < numberToMove; i += 1 {
			crate := stacks[from].Pop()
			stacks[to].Push(crate)
		}
	}

	//grab first item on each stack
	ret := ""
	for i := range stacks {
		ret += stacks[i].Pop()
	}
	return ret
}

func partTwo(lines []string) string {
	stacks := getInitialStacks()

	for _, line := range lines {
		if line == "" {
			break
		}

		numberToMove, from, to := parseLine(line)

		movingStack := new(stack)
		for i := 0; i < numberToMove; i += 1 {
			crate := stacks[from].Pop()
			movingStack.Push(crate)
		}

		for i := 0; i < numberToMove; i += 1 {
			stacks[to].Push(movingStack.Pop())
		}
	}

	//grab first item on each stack
	ret := ""
	for i := range stacks {
		ret += stacks[i].Pop()
	}
	return ret
}

func main() {
	data, err := os.ReadFile("./day5.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1: ", partOne(lines))
	fmt.Println("Part 2: ", partTwo(lines))
}
