package solutions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func DayOne() int {

	file, err := os.Open("solutions/dayoneinput.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var count, previous int
	for scanner.Scan() {
		line := scanner.Text()

		if count == 0 {
			count++
			continue
		}

		lineInt, err := strconv.Atoi(string(line))
		if err != nil {
			panic("err converting number")
		}
		if lineInt > previous {
			count++
		}
		previous = lineInt
	}
	defer file.Close()

	return count - 1
}
