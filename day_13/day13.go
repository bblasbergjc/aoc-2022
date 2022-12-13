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

const (
	KeepGoing = iota
	InOrder
	OutOfOrder
)

type NumberOrList struct {
	Number *int
	List   *[]NumberOrList
}

func (n *NumberOrList) ToString() string {
	if n.Number != nil {
		return fmt.Sprint(*n.Number)
	}

	print := func(i []NumberOrList) string {
		out := ""

		var recur func([]NumberOrList)
		recur = func(items []NumberOrList) {
			for i, item := range items {
				if item.List != nil {
					out += "["
					recur(*item.List)
					out += "]"
				} else {
					out += fmt.Sprint(*(item.Number))
				}

				if i < len(items)-1 {
					out += ","
				}
			}
		}

		out += "["
		recur(i)
		out += "]"

		return out
	}

	return print(*n.List)
}

type Pair struct {
	Left  []NumberOrList
	Right []NumberOrList
}

// WARNING: this code is disgusting
func parseInput(lines []string) []Pair {
	pairs := make([]Pair, 0)

	var l []NumberOrList
	var r []NumberOrList
	for _, line := range lines {
		if line == "" { // we've parsed both pairs, get ready for the next
			pairs = append(pairs, Pair{l, r})
			l = nil
			r = nil

			continue
		}

		stack := make([]NumberOrList, 0)
		rawNum := ""

		openLists := make([]*[]NumberOrList, 0)
		for _, ch := range line[1 : len(line)-1] { // remove ends because we know that part is a list
			if ch == '[' { // start a new list
				lst := NumberOrList{nil, &[]NumberOrList{}}

				if len(openLists) == 0 { // not in an open list, put it directly on the stack
					stack = append(stack, lst)
				} else { // otherwise put it in the most recently opened list
					*openLists[len(openLists)-1] = append(*openLists[len(openLists)-1], lst)
				}
				openLists = append(openLists, lst.List)
			} else if ch == ']' {
				if rawNum != "" {
					num, err := strconv.Atoi(rawNum)
					checkErr(err)

					packet := NumberOrList{&num, nil}

					if len(openLists) > 0 {
						*openLists[len(openLists)-1] = append(*openLists[len(openLists)-1], packet)
					} else {
						stack = append(stack, packet)
					}
				}
				rawNum = ""

				//close the most recently opened list
				openLists = openLists[0 : len(openLists)-1]
			} else if ch == ',' {
				if rawNum != "" {
					num, err := strconv.Atoi(rawNum)
					checkErr(err)

					packet := NumberOrList{&num, nil}

					if len(openLists) > 0 {
						*openLists[len(openLists)-1] = append(*openLists[len(openLists)-1], packet)
					} else {
						stack = append(stack, packet)
					}
				}

				rawNum = ""
			} else { //number
				rawNum += string(ch)
			}
		}

		if rawNum != "" { //ended on a number
			num, err := strconv.Atoi(rawNum)
			checkErr(err)

			packet := NumberOrList{&num, nil}

			stack = append(stack, packet)
		}

		if l == nil {
			l = stack
		} else {
			r = stack
		}
	}

	return pairs
}

func printCompare(depth int, left string, right string) {
	compareStr := "%" + fmt.Sprint(2*depth) + "v" + "Compare: %s vs %s\n"
	fmt.Printf(compareStr, "", left, right)
}

func compare(p Pair) bool {
	var compareFunc func(pair Pair, depth int) int
	compareFunc = func(pair Pair, depth int) int {
		left := pair.Left
		right := pair.Right
		printCompare(depth, stringAll(left), stringAll(right))

		for i := range left {
			if i == len(right) { // right ran out of inputs
				fmt.Print("Right ran out of items ")
				return OutOfOrder
			}

			l := left[i]
			r := right[i]

			if l.Number != nil && r.Number != nil { //both integers
				printCompare(depth+1, l.ToString(), r.ToString())

				if *l.Number == *r.Number {
					continue // keep going
				}

				if *l.Number < *r.Number {
					fmt.Print(*l.Number, " < ", *r.Number, " ")
					return InOrder
				} else {
					fmt.Print(*l.Number, " > ", *r.Number, " ")
					return OutOfOrder
				}
			}

			if l.List != nil && r.List != nil { //both lists
				res := compareFunc(Pair{*l.List, *r.List}, depth+1)

				if res != KeepGoing {
					return res
				}
			}

			// mismatch, left list
			if l.List != nil && r.List == nil {
				res := compareFunc(Pair{*l.List, []NumberOrList{{r.Number, nil}}}, depth+1)

				if res != KeepGoing {
					return res
				}
			}

			// mismatch, right list
			if l.List == nil && r.List != nil {
				res := compareFunc(Pair{[]NumberOrList{{l.Number, nil}}, *r.List}, depth+1)

				if res != KeepGoing {
					return res
				}
			}
		}

		if len(left) < len(right) {
			fmt.Print("Left ran out of items ")
			return InOrder
		} else {
			return KeepGoing
		}
	}

	res := compareFunc(p, 0)

	if res == InOrder {
		fmt.Println("IN order")
	} else if res == OutOfOrder {
		fmt.Println("OUT of order")
	} else {
		fmt.Println("KEEP GOING - so , yes in order")
	}

	return res != OutOfOrder
}

func partOne(pairs []Pair) int {
	sum := 0

	for i, pair := range pairs {
		if compare(pair) {
			sum += i + 1
		}
	}
	return sum
}

func stringAll(items []NumberOrList) string {
	out := "["
	for i, n := range items {
		out += n.ToString()

		if i < len(items)-1 {
			out += ","
		}
	}
	out += "]"
	return out
}

func main() {
	data, err := os.ReadFile("./day13.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")

	pairs := parseInput(lines)

	// // print our input to make sure we read it right
	// for i := range pairs {
	// 	fmt.Println(stringAll(pairs[i].Left))
	// 	fmt.Println(stringAll(pairs[i].Right))
	// 	fmt.Println()
	// }

	fmt.Println("Part 1:", partOne(pairs))
}
