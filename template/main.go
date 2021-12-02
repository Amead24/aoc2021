package main

import (
	"bufio"
	"log"
	"os"
)

func part01(inputs []string) (int, error) {
	return 0, nil
}

func part02(inputs []string) (int, error) {
	return 0, nil
}

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

func main() {
	inputs, err := readLines("inputs.txt")
	if err != nil {
		log.Fatalf("reading input failed")
	}

	p1Answer, err := part01(inputs)
	if err != nil {
		log.Fatalf("part01 failed")
	}

	log.Printf("Part 1 == %d", p1Answer)

	p2Answer, err := part02(inputs)
	if err != nil {
		log.Fatalf("part02 failed")
	}

	log.Printf("Part 2 == %d", p2Answer)
}
