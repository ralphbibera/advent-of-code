package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	reports, err := readInput()
	if err != nil {
		t.Fatal(err)
	}

	safe := 0

	for _, report := range reports {
		if check(report) {
			safe++
		}
	}

	log.Println(safe)
}

func TestDayTwoPartTwo(t *testing.T) {
	reports, err := readInput()
	if err != nil {
		t.Fatal(err)
	}

	safe := 0

	for _, report := range reports {
		for index := 0; index < len(report); index++ {
			arr := append([]int(nil), report...)
			arr = append(arr[:index], arr[index+1:]...)
			if check(arr) {
				safe++
				break
			}
		}
	}

	log.Println(safe)
}

func check(report []int) bool {
	isSafe := true

	var increasing *bool

	for idx := 0; idx < len(report)-1; idx++ {
		prev := report[idx]
		next := report[idx+1]
		difference := float64(prev - next)

		if increasing == nil {
			incr := math.Signbit(difference)
			increasing = &incr
		}

		if 3 < math.Abs(difference) {

			isSafe = false
			break
		}

		if prev == next {

			isSafe = false
			break
		}

		if *increasing {
			if !math.Signbit(difference) {

				isSafe = false
				break
			}
		} else {
			if math.Signbit(difference) {
				isSafe = false
				break
			}
		}
	}

	return isSafe
}

func readInput() ([][]int, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	reports := make([][]int, 0)
	for _, line := range lines {
		reportSplit := strings.Split(line, " ")
		reportNums := make([]int, len(reportSplit))
		for idx, numString := range reportSplit {
			num, err := strconv.Atoi(numString)
			if err != nil {
				return nil, err
			}
			reportNums[idx] = num
		}
		reports = append(reports, reportNums)
	}

	return reports, nil
}
