package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func part1(depths []int) (int, error) {
	increases := 0
	nextDepth := depths[1]
	for i, depth := range depths[:len(depths)-2] {
		if depth < nextDepth {
			increases++
		}

		nextDepth = depths[i+2]
	}

	return increases, nil
}

func part02(depths []int) (int, error) {
	slidingDepths := make([]int, len(depths))
	for i, _ := range depths[:len(depths)-2] {
		slidingDepths[i] = depths[i]
		slidingDepths[i] += depths[i+1]
		slidingDepths[i] += depths[i+2]
	}

	increases := 0
	nextDepth := slidingDepths[1]
	for i, depth := range slidingDepths[:len(slidingDepths)-2] {
		if depth < nextDepth {
			increases++
		}
		nextDepth = slidingDepths[i+2]
	}

	return increases, nil
}

func main() {
	strDepths, err := readLines("inputs.txt")
	if err != nil {
		log.Fatalf("reading input failed")
	}

	depths := make([]int, len(strDepths))
	for i, depth := range strDepths {
		depths[i], err = strconv.Atoi(depth)
		if err != nil {
			log.Fatalf("cant convert %s", err)
		}
	}

	p1Increases, err := part1(depths)
	if err != nil {
		log.Fatalf("part01 failed")
	}

	log.Printf("Part 1 == %d", p1Increases)

	p2Increases, err := part02(depths)
	if err != nil {
		log.Fatalf("part02 failed")
	}

	log.Printf("Part 2 == %d", p2Increases)
}
