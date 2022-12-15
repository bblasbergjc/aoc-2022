package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	. "github.com/bblasbergjc/aoc-2022/util"
)

type Data struct {
	Sensors []Sensor
	Beacons []Point
	XRange  Range
	YRange  Range
}

func Abs(x, y int) int {
	diff := x - y

	if diff < 0 {
		return diff * -1
	}

	return diff
}

type Range struct {
	Low  int
	High int
}

func (r *Range) MinMax(x int) {
	if x > r.High {
		r.High = x
	} else if x < r.Low {
		r.Low = x
	}
}

func NewRange() Range {
	return Range{
		Low:  math.MaxInt,
		High: math.MinInt,
	}
}

type Point struct {
	Display string
	X       int
	Y       int
}

func (p Point) Distance(o Point) int {
	return Abs(p.X, o.X) + Abs(p.Y, o.Y)
}

type Sensor struct {
	Point          Point
	BeaconDistance int
}

func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	CheckErr(err)

	return num
}

func parseLine(line string) (Point, Point) {
	regex, err := regexp.Compile(`(-?\d+)`)
	CheckErr(err)

	rawNums := regex.FindAllString(line, 4)

	return Point{
			X:       parseInt(rawNums[0]),
			Y:       parseInt(rawNums[1]),
			Display: "S",
		}, Point{
			X:       parseInt(rawNums[2]),
			Y:       parseInt(rawNums[3]),
			Display: "B",
		}
}

func parseInput(lines []string) *Data {
	xRange := NewRange()
	yRange := NewRange()

	var sensors []Sensor
	var beacons []Point
	greatestDistance := 0
	for _, line := range lines {
		sensor, beacon := parseLine(line)

		// set ranges for the board
		xRange.MinMax(sensor.X)
		xRange.MinMax(beacon.X)
		yRange.MinMax(sensor.Y)
		yRange.MinMax(beacon.Y)

		sensors = append(sensors, Sensor{
			Point:          sensor,
			BeaconDistance: sensor.Distance(beacon),
		})
		beacons = append(beacons, beacon)

		if sensor.Distance(beacon) > greatestDistance {
			greatestDistance = sensor.Distance(beacon)
		}
	}

	xRange.Low -= greatestDistance
	xRange.High += greatestDistance
	yRange.Low -= greatestDistance
	yRange.High += greatestDistance

	return &Data{
		Sensors: sensors,
		Beacons: beacons,
		XRange:  xRange,
		YRange:  yRange,
	}
}

func isSensorPoint(sensors []Sensor, x int, y int) bool {
	for _, sensor := range sensors {
		if sensor.Point.X == x && sensor.Point.Y == y {
			return true
		}
	}

	return false
}

func isBeaconPoint(beacons []Point, x int, y int) bool {
	for _, beacon := range beacons {
		if beacon.X == x && beacon.Y == y {
			return true
		}
	}

	return false
}

func isFreeSpace(data *Data, x int, y int, ignoreBeacon bool) bool {
	for _, sensor := range data.Sensors {
		if sensor.Point.Distance(Point{X: x, Y: y}) <= sensor.BeaconDistance {
			if ignoreBeacon {
				return isBeaconPoint(data.Beacons, x, y)
			} else {
				return false
			}
		}
	}

	return true
}

func partOne(data *Data, y int) int {
	count := 0
	for x := data.XRange.Low; x <= data.XRange.High; x += 1 {
		if !isFreeSpace(data, x, y, true) {
			count += 1
		}
	}

	return count
}

const tuningMultiplier = 4_000_000

func searchForFreeSpace(id int, data *Data, xRange Range, yRange Range, found chan int) {
	fmt.Printf("[%d] Starting search for x[%d,%d] and y[%d,%d]\n", id, xRange.Low, xRange.High, yRange.Low, yRange.High)

	count := 1
	tenths := 0
	for x := xRange.Low; x <= xRange.High; x += 1 {
		for y := yRange.Low; y <= yRange.High; y += 1 {

			if count%5_000_000_000 == 0 {
				tenths += 1
				fmt.Printf("[%d] has processed %d/10th of its workload\n", id, tenths)
			}

			if isFreeSpace(data, x, y, false) {
				print("x:", x, "y:", y)
				found <- (tuningMultiplier * x) + y
			}

			count += 1
		}
	}
}

func partTwo(data *Data, max int) int {
	step := max / 320

	id := 0
	found := make(chan int)
	for i := 0; i < max; i += step {
		go searchForFreeSpace(id, data, Range{Low: i, High: i + step}, Range{Low: 0, High: max}, found)
		id += 1
	}

	total := int64(max) * int64(max)
	fmt.Println("Total to process:", total)

	answer := <-found

	return answer
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day15.txt")
	data := parseInput(lines)

	// print the board
	// for y := yRange.Low; y <= yRange.High; y += 1 {
	// 	for x := xRange.Low; x <= xRange.High; x += 1 {
	// 		if isBeaconPoint(beacons, x, y) {
	// 			fmt.Print("B")
	// 		} else if isSensorPoint(sensors, x, y) {
	// 			fmt.Print("S")
	// 		} else if !isFreeSpace(sensors, beacons, x, y, true) {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	//fmt.Println("Part 1:", partOne(sensors, beacons, xRange, yRange, 2000000))
	fmt.Println("Part 2:", partTwo(data, 4000000))
}
