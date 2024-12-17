package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestDayThreePartOne(t *testing.T) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := string(file)

	pattern := `mul\(\d+,\d+\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(input, -1)

	functions := make([]string, 0)
	for _, match := range matches {
		functions = append(functions, match)
	}

	sum := 0

	for _, function := range functions {
		split := strings.Split(function, ",")
		firstNumStr := cleanString(split[0])
		secondNumStr := cleanString(split[1])

		firstNum, err := strconv.Atoi(string(firstNumStr))
		if err != nil {
			t.Fatal(err)
		}
		secondNum, err := strconv.Atoi(string(secondNumStr))
		if err != nil {
			t.Fatal(err)
		}

		sum += firstNum * secondNum
	}

	log.Println(sum)
}

func cleanString(input string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(input, "")
}

func TestDayThreePartTwo(t *testing.T) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := string(file)

	cleanInput := ""

	split := strings.SplitAfter(input, "do()")

	for idx, s := range split {
		log.Println(idx, "===")
		dontIndex := strings.Index(s, "don't()")
		if dontIndex != -1 {
			cleanInput += s[:dontIndex]
		} else {
			cleanInput += s
		}
	}

	pattern := `mul\(\d+,\d+\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(cleanInput, -1)

	functions := make([]string, 0)
	for _, match := range matches {
		functions = append(functions, match)
	}

	sum := 0

	for _, function := range functions {
		split := strings.Split(function, ",")
		firstNumStr := cleanString(split[0])
		secondNumStr := cleanString(split[1])

		firstNum, err := strconv.Atoi(string(firstNumStr))
		if err != nil {
			t.Fatal(err)
		}
		secondNum, err := strconv.Atoi(string(secondNumStr))
		if err != nil {
			t.Fatal(err)
		}

		sum += firstNum * secondNum
	}

	log.Println(sum)
}
