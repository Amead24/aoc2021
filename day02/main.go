package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func part01(inputs []string) (int, error) {
	vertical := 0
	horizontal := 0
	for i, dirAndUnits := range inputs {
		dir, strUnit := func() (string, string) {
			dirAndUnit := strings.Split(dirAndUnits, " ")
			return dirAndUnit[0], dirAndUnit[1]
		}()

		unit, err := strconv.Atoi(strUnit)
		if err != nil {
			return 0, err
		}

		// forward is a simple aggregation while
		// up and down would be a sum
		if dir == "forward" {
			horizontal += unit
		} else if dir == "up" {
			vertical += (unit * -1)
		} else if dir == "down" {
			vertical += unit
		} else {
			log.Fatalf("unknown direction on line - %d", i)
		}
	}

	return horizontal * vertical, nil
}

func part02(inputs []string) (int, error) {
	vertical := 0
	horizontal := 0
	aim := 0
	for i, dirAndUnits := range inputs {
		dir, strUnit := func() (string, string) {
			dirAndUnit := strings.Split(dirAndUnits, " ")
			return dirAndUnit[0], dirAndUnit[1]
		}()

		unit, err := strconv.Atoi(strUnit)
		if err != nil {
			return 0, err
		}

		// forward is still a simple aggregation while
		// down/up require tracking & resetting aim after use
		if dir == "forward" {
			horizontal += unit
			vertical += (unit * aim)
		} else if dir == "up" {
			// vertical += (unit * -1)
			aim += (unit * -1)
		} else if dir == "down" {
			// vertical += unit
			aim += unit
		} else {
			log.Fatalf("unknown direction on line - %d", i)
		}
		log.Printf("horizontal, depth, & aim - %d, %d, %d", horizontal, vertical, aim)
	}

	return horizontal * vertical, nil
}

func main() {
	inputs, err := readLines("inputs.txt")
	if err != nil {
		log.Fatalf("reading input failed")
	}

	p1Answer, err := part01(inputs)
	if err != nil {
		log.Fatalf("part01 failed")
	}

	log.Printf("Part 1 == %d", p1Answer) // 1936494

	p2Answer, err := part02(inputs)
	if err != nil {
		log.Fatalf("part02 failed")
	}

	log.Printf("Part 2 == %d", p2Answer) // 1997106066
}
