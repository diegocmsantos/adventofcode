package solutions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func DayTwo() int {

	file, err := os.Open("solutions/dayoneinput.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var counter, buffCounter, previous int
	var buff [3]int
	for scanner.Scan() {
		line := scanner.Text()

		if buffCounter < 3 {
			n := convertStringToInt(line)
			buff[buffCounter] = n
			buffCounter++
			continue
		}
		buffCounter = 0

		lineInt := buff[0] + buff[1] + buff[2]
		if lineInt > previous {
			counter++
		}
		previous = lineInt
	}
	defer file.Close()

	return count - 1
}

func convertStringToInt(ns string) int {
	n, err := strconv.Atoi(ns)
	if err != nil {
		panic("err converting number")
	}
	return n
}

func shiftArray(ar []int, lastNumber int) {
	for i := 1; i < len(ar)-1; i++ {
		ar[i-1] = ar[i]
	}
	ar[len(ar)-1] = lastNumber
}
