package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

var symbols map[rune]bool = map[rune]bool{
	'*': true, '+': true, '-': true, '_': true, '/': true, '\\': true, '%': true, '^': true,
	'=': true, '!': true, '<': true, '>': true, '&': true, '|': true, '~': true, '`': true,
	'?': true, ':': true, ';': true, ',': true, '(': true, ')': true, '[': true, ']': true,
	'{': true, '}': true, '#': true, '$': true, '@': true,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkNumAtIndex(i int, visited map[int]bool, chars []rune) (string, bool) {
	_, ok := visited[i]
	if ok {
		return "", false
	}

	number := ""

	if unicode.IsDigit(chars[i]) {
		number = string(chars[i])
	} else {
		return "", false
	}

	rightIndex := i + 1
	leftIndex := i - 1

	_, ok = visited[leftIndex]
	for leftIndex >= 0 && unicode.IsDigit(chars[leftIndex]) && !ok {
		visited[leftIndex] = true
		number = string(chars[leftIndex]) + number
		leftIndex--
	}

	_, ok = visited[rightIndex]
	for rightIndex < len(chars) && unicode.IsDigit(chars[rightIndex]) && !ok {
		visited[rightIndex] = true
		number = number + string(chars[rightIndex])
		rightIndex++
	}

	return number, true
}

func isSymbol(char rune) bool {
	_, ok := symbols[char]
	return ok
}

func main() {
	file, err := os.Open("day3/input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	sum := 0

	for i, line := range lines {
		aboveChars := []rune{}
		belowChars := []rune{}
		if i > 0 {
			aboveChars = []rune(lines[i-1])
		}
		if i < len(lines)-1 {
			belowChars = []rune(lines[i+1])
		}
		chars := []rune(line)
		linesToParse := [][]rune{aboveChars, chars, belowChars}

		for j, char := range line {
			visited := make(map[int]bool)
			if char != '.' && isSymbol(char) {
				for _, chars := range linesToParse {
					if len(chars) > 0 {
						for k := j - 1; k <= j+1 && j < len(chars); k++ {
							number, ok := checkNumAtIndex(k, visited, chars)
							if ok {
								number, err := strconv.Atoi(number)
								check(err)
								sum += number
							}
						}
						visited = make(map[int]bool)
					}
				}
			}
		}
	}

	println("sum: ", sum)
}
