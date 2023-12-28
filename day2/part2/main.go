package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day2/input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	sum := 0

	for _, line := range lines {
		split := strings.Split(line, ":")
		game := strings.Split(split[1], ";")
		maxes := make(map[string]int)
		maxes["red"] = 0
		maxes["green"] = 0
		maxes["blue"] = 0

		for _, iteration := range game {
			cubes := strings.Split(iteration, ",")

			for _, cube := range cubes {
				fields := strings.Fields(cube)
				count, err := strconv.Atoi(fields[0])
				check(err)
				colour := fields[1]
				maxes[colour] = max(maxes[colour], count)
			}
		}

		power := 1
		for _, max := range maxes {
			power *= max
		}

		sum += power
	}

	print("sum of powers: ", sum)
}
