package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	combos [][]string
}

func (b Board) isWinner(num string) bool {
	hasBingo := false
	for x, combo := range b.combos {
		inARow := 0

		for y, ele := range combo {
			if ele == num {
				b.combos[x][y] = "X"
				inARow++
			} else if ele == "X" {
				inARow++
			} else {
				inARow = 0
			}
		}

		if inARow == 5 {
			hasBingo = true
		}
	}

	return hasBingo
}

func (b Board) Sum(strNum string) int {
	// i've created  duplicate numbers with my combos
	// so we need to sum only the de-dupped numbers
	set := make(map[string]int)
	for _, combo := range b.combos {
		for _, ele := range combo {
			if ele != "X" {
				value, _ := strconv.Atoi(ele)
				set[ele] = value
			}
		}
	}

	sum := 0
	for _, value := range set {
		sum += value
	}

	num, _ := strconv.Atoi(strNum)
	return sum * num
}

func NewBoard(strBoard string) Board {
	sliceStep := func(slice []string, start int, stop int, step int) []string {
		newSlice := make([]string, 0)
		for i := start; i < stop; i += step {
			newSlice = append(newSlice, slice[i])
		}
		return newSlice
	}

	sliceBoard := strings.Fields(strBoard)
	return Board{combos: [][]string{
		// horizontals
		sliceStep(sliceBoard, 0, 5, 1),
		sliceStep(sliceBoard, 5, 10, 1),
		sliceStep(sliceBoard, 10, 15, 1),
		sliceStep(sliceBoard, 15, 20, 1),
		sliceStep(sliceBoard, 20, 25, 1),
		// verticals
		sliceStep(sliceBoard, 0, 25, 5),
		sliceStep(sliceBoard, 1, 25, 5),
		sliceStep(sliceBoard, 2, 25, 5),
		sliceStep(sliceBoard, 3, 25, 5),
		sliceStep(sliceBoard, 4, 25, 5),
	}}
}

func part01(inputs []string, numbers []string) (int, error) {
	// need to track which board has the
	// least number of drawings before bingo.
	answer := 0
	lowestWinningDraw := len(numbers)
	for i, input := range inputs {
		board := NewBoard(input)

		for ii := 0; ii < len(numbers); ii++ {
			if board.isWinner(numbers[ii]) {
				// log.Printf("Board %d wins in %d moves", i, ii+1)
				if ii < lowestWinningDraw {
					lowestWinningDraw = ii
					answer = board.Sum(numbers[ii])
					log.Printf("P1 - current leading board[%d] on round %d with value of %d", i, ii, answer)
				}
				break
			}
		}
	}

	return answer, nil
}

func part02(inputs []string, numbers []string) (int, error) {
	// need to track which board has the
	// least number of drawings before bingo.
	answer := 0
	highestWinningDraw := 0
	for i, input := range inputs {
		board := NewBoard(input)

		for ii := 0; ii < len(numbers); ii++ {
			if board.isWinner(numbers[ii]) {
				// log.Printf("Board %d wins in %d moves", i, ii+1)
				if ii >= highestWinningDraw {
					highestWinningDraw = ii
					answer = board.Sum(numbers[ii])
					log.Printf("P2 - current losing board[%d] on round %d with value of %d", i, ii, answer)
				}
				break
			}
		}
	}

	return answer, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var boards []string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText != "" { // need double blank lines at the end - idk?
			lines = append(lines, lineText)
		} else {
			boards = append(boards, strings.Join(lines, " "))
			lines = []string{} // reset
		}
	}

	return boards, scanner.Err()
}

func main() {
	strNumbersCalled := "23,91,18,32,73,14,20,4,10,55,40,29,13,25,48,65,2,80,22,16,93,85,66,21,9,36,47,72,88,58,5,42,53,69,52,8,54,63,76,12,6,99,35,95,82,49,41,17,62,34,51,77,94,7,28,71,92,74,46,79,26,19,97,86,87,37,57,64,1,30,11,96,70,44,83,0,56,90,59,78,61,98,89,43,3,84,67,38,68,27,81,39,15,50,60,24,45,75,33,31"
	numbersCalled := strings.Split(strNumbersCalled, ",")

	inputs, err := readLines("inputs.txt")
	if err != nil {
		log.Fatalf("reading input failed")
	}

	p1Answer, err := part01(inputs, numbersCalled)
	if err != nil {
		log.Fatalf("part01 failed")
	}

	log.Printf("Part 1 == %d", p1Answer)

	p2Answer, err := part02(inputs, numbersCalled)
	if err != nil {
		log.Fatalf("part02 failed")
	}

	log.Printf("Part 2 == %d", p2Answer)
}
