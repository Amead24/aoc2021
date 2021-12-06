package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func part01(inputs []string) (int, error) {
	// so far the best idea I have would be to
	// keep some list of tuples [(0, 0), (0, 1), ...]
	// but even that would be O(n*m)

	gamma := 0
	epsilon := 0

	diagnostics := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	// meh - O(n*m), is there a better way?
	for _, input := range inputs {
		for ii, row := range input {
			if row == 48 {
				// rune 48 == '0'
				diagnostics[ii][0]++
			} else {
				// rune 49 == '1'
				diagnostics[ii][1]++
			}
		}
	}

	for i, stat := range diagnostics {
		binary := int(math.Pow(2, float64(len(diagnostics)-i-1)))

		// if the first element is higher,
		// then gamma gets set to zero
		// with epsilon set to one.
		if stat[0] > stat[1] {
			epsilon += binary
		} else {
			gamma += binary
		}
	}

	return gamma * epsilon, nil
}

func part02(inputs []string) (int, error) {
	tmpOxygenOptions := make([]string, 0)
	oxygenOptions := make([]string, len(inputs))
	copy(oxygenOptions, inputs)

	// meh - There has got to be a better way?
	for i := 0; i < len(inputs[0]); i++ {

		// tracking the number of zeros & ones
		bitTracker := []int{0, 0}
		for _, oxygenOption := range oxygenOptions {
			if oxygenOption[i] == '0' {
				bitTracker[0]++
			} else {
				bitTracker[1]++
			}
		}

		// loop back through looking for the most common bit w/ pos 'i'
		for _, oxygenOption := range oxygenOptions {
			if bitTracker[0] > bitTracker[1] {
				if oxygenOption[i] == '0' {
					tmpOxygenOptions = append(tmpOxygenOptions, oxygenOption)
				}
			} else {
				if oxygenOption[i] == '1' {
					tmpOxygenOptions = append(tmpOxygenOptions, oxygenOption)
				}
			}
		}

		if len(oxygenOptions) != 1 {
			oxygenOptions = tmpOxygenOptions
			tmpOxygenOptions = make([]string, 0) // reset
		}
	}

	tmpCarbonDioxideOptions := make([]string, 0)
	carbonDioxideOptions := make([]string, len(inputs))
	copy(carbonDioxideOptions, inputs)

	for i := 0; i < len(inputs[0]); i++ {

		// tracking the number of zeros & ones
		bitTracker := []int{0, 0}
		for _, carbonDioxideOption := range carbonDioxideOptions {
			if carbonDioxideOption[i] == '0' {
				bitTracker[0]++
			} else {
				bitTracker[1]++
			}
		}

		// loop back through looking for the least common bit w/ pos 'i'
		for _, carbonDioxideOption := range carbonDioxideOptions {
			if bitTracker[0] <= bitTracker[1] {
				if carbonDioxideOption[i] == '0' {
					log.Printf("looking for %d == 0", i)
					tmpCarbonDioxideOptions = append(tmpCarbonDioxideOptions, carbonDioxideOption)
				}
			} else {
				if carbonDioxideOption[i] == '1' {
					log.Printf("looking for %d == 0", i)
					tmpCarbonDioxideOptions = append(tmpCarbonDioxideOptions, carbonDioxideOption)
				}
			}
		}

		log.Printf("tmp co2 == %s", tmpCarbonDioxideOptions)

		if len(carbonDioxideOptions) != 1 {
			carbonDioxideOptions = tmpCarbonDioxideOptions
			tmpCarbonDioxideOptions = make([]string, 0) // reset
		}
	}

	log.Printf("final 02 == %s", oxygenOptions)
	log.Printf("final C02 == %s", carbonDioxideOptions)

	oxygenInt := 0
	carbonDioxideInt := 0
	for i := range inputs[0] {
		oxygenInt += int(math.Pow(2, float64(len(inputs[0])-i-1)))
		carbonDioxideInt += int(math.Pow(2, float64(len(inputs[0])-i-1)))
	}

	return oxygenInt * carbonDioxideInt, nil
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

	log.Printf("Part 1 == %d", p1Answer) // 3985686

	p2Answer, err := part02(inputs)
	if err != nil {
		log.Fatalf("part02 failed")
	}

	log.Printf("Part 2 == %d", p2Answer) // 2555739
}
