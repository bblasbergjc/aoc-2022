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
	XRanges        map[int]Range // map of y level and the x ranges of their search
}

func calcXRanges(sensor Sensor) map[int]Range {
	xranges := make(map[int]Range)

	// start at the top of the sensor range and work our way down to the bottom
	for y := sensor.Point.Y - sensor.BeaconDistance; y <= sensor.Point.Y+sensor.BeaconDistance; y += 1 {
		width := sensor.BeaconDistance - Abs(y, sensor.Point.Y)

		xranges[y] = Range{
			Low:  sensor.Point.X - width,
			High: sensor.Point.X + width,
		}
	}

	return xranges
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

		sensorItem := Sensor{
			Point:          sensor,
			BeaconDistance: sensor.Distance(beacon),
		}

		sensorItem.XRanges = calcXRanges(sensorItem)

		sensors = append(sensors, sensorItem)
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

// Had to read this answer: https://www.reddit.com/r/adventofcode/comments/zmcn64/comment/j0d6du6/?utm_source=share&utm_medium=web2x&context=3
// to understand how to do this
// It's still slow, but possible without a government level super computer :)
func partTwo(data *Data, max int) int {
	for y := 0; y < max; y += 1 {
		for _, sensor := range data.Sensors {
			// check the square just left of the range
			leftCol := sensor.XRanges[y].Low - 1
			if leftCol > 0 && leftCol < max && isFreeSpace(data, leftCol, y, false) {
				fmt.Println("x:", leftCol, "y:", y)
				return (leftCol * tuningMultiplier) + y
			}

			// check the square just right of the range
			rightCol := sensor.XRanges[y].High + 1
			if rightCol > 0 && rightCol < max && isFreeSpace(data, rightCol, y, false) {
				fmt.Println("x:", rightCol, "y:", y)
				return (rightCol * tuningMultiplier) + y
			}
		}

	}

	return -1
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
	fmt.Println("Part 2:", partTwo(data, 4_000_000))
}
