package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestDayOnePartOne(t *testing.T) {
	leftNums, rightNums, err := readInput()
	if err != nil {
		t.Fatal(err)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	sum := 0.0
	for idx := 0; idx < len(leftNums); idx++ {
		sum += math.Abs(float64(leftNums[idx] - rightNums[idx]))
	}

	log.Println(int(sum))
}

func TestDayOnePartTwo(t *testing.T) {
	leftNums, rightNums, err := readInput()
	if err != nil {
		t.Fatal(err)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	lookup := map[int]int{}

	similarityScore := 0

	for idx := 0; idx < len(leftNums); idx++ {
		leftNum := leftNums[idx]
		if _, ok := lookup[leftNum]; ok {
			break
		} else {
			lookup[leftNum] = 0
			for _, rightNum := range rightNums {
				if rightNum == leftNum {
					lookup[leftNum]++
				}
			}
		}
		similarityScore += leftNum * lookup[leftNum]
	}

	log.Println(similarityScore)
}

func readInput() ([]int, []int, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(string(file), "\n")

	leftNums := make([]int, 0)
	rightNums := make([]int, 0)

	for _, line := range lines {
		split := strings.Split(line, "   ")
		leftNum, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		rightNum, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		leftNums = append(leftNums, int(leftNum))
		rightNums = append(rightNums, int(rightNum))
	}

	return leftNums, rightNums, nil
}
