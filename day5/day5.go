package day5

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Run() int {

	currentIndex := 0
	// puzzle := []int{
	// 	0, 3, 0, 1, -3,
	// }

	puzzle, err := loadPuzzleInput()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	steps := 0
	for {
		steps++
		newIndex := currentIndex + puzzle[currentIndex]
		if puzzle[currentIndex] >= 3 {
			puzzle[currentIndex]--
		} else {
			puzzle[currentIndex]++
		}
		currentIndex = newIndex

		if currentIndex >= len(puzzle) {
			break
		}
	}

	return steps
}

func loadPuzzleInput() ([]int, error) {

	dat, err := ioutil.ReadFile("puzzleinput.txt")
	if err != nil {
		return nil, err
	}

	spl := strings.Split(string(dat), "\n")
	var puzzle = []int{}

	for _, i := range spl {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		puzzle = append(puzzle, j)
	}
	return puzzle, nil
}
