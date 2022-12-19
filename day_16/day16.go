package main

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

type RoomSet map[string]*Room

func (s RoomSet) Add(room *Room) {
	if s == nil {
		s = make(map[string]*Room)
	}

	s[room.Name] = room
}

func (s RoomSet) Contains(name string) bool {
	if s == nil {
		return false
	}

	_, ok := s[name]

	return ok
}

func (s RoomSet) CalculateTotalPressure() int {
	if s == nil {
		return 0
	}

	pressure := 0

	for _, room := range s {
		pressure += room.FlowRate
	}

	return pressure
}

func (s RoomSet) Clone() RoomSet {
	new := RoomSet{}

	for k, v := range s {
		new[k] = v
	}

	return new
}

type Room struct {
	Name     string
	FlowRate int
	Edges    []*Room
}

func parseInput(lines []string) map[string]*Room {
	edges := make(map[string][]string)
	rooms := make(map[string]*Room)

	for _, line := range lines {
		room := Room{}

		fmt.Sscanf(line, "Valve %s has flow rate=%d", &room.Name, &room.FlowRate)

		rawValves := strings.Split(line, ";")[1]
		rawValves = strings.TrimPrefix(rawValves, " tunnel leads to valve ")
		rawValves = strings.TrimPrefix(rawValves, " tunnels lead to valves ")

		edgesRaw := strings.Split(rawValves, ", ")
		edges[room.Name] = edgesRaw
		rooms[room.Name] = &room
	}

	for _, room := range rooms {
		room.Edges = make([]*Room, 0)
		for _, name := range edges[room.Name] {
			edgeRoom, ok := rooms[name]
			if !ok {
				CheckErr(errors.New("Edge room doesnt exist: " + name))
			}
			room.Edges = append(room.Edges, edgeRoom)
		}
	}

	return rooms
}

func dfs(root *Room, minutesLeft int, openValves RoomSet, roomsSize int, depth int) int {
	fmt.Printf("[%s] minutes: %d, openValves len: %d, roomsSize: %d\n", root.Name, minutesLeft, len(openValves), roomsSize)

	if depth > 15 {
		return 0
	}
	if minutesLeft == 0 {
		return 0
	}

	if len(openValves) == roomsSize { // all the rooms are open, run out the clock
		pressure := 0
		for m := minutesLeft; m > 0; m -= 1 {
			pressure += openValves.CalculateTotalPressure()
		}

		return pressure
	}

	pressure := openValves.CalculateTotalPressure()
	openValves.Add(root)

	highestPressure := 0
	for _, room := range root.Edges {
		if room.Name == root.Name {
			continue
		}
		p := dfs(room, minutesLeft-1, openValves.Clone(), roomsSize, depth+1)
		if p > highestPressure {
			highestPressure = p
		}
	}

	return pressure + highestPressure
}

func partOne(rooms map[string]*Room) int {
	return dfs(rooms["AA"], 30, RoomSet{}, len(rooms), 0)
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day16.txt")
	rooms := parseInput(lines)

	fmt.Println("Part 1:", partOne(rooms))
}
