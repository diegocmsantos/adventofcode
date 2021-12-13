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

	var counter int
	var buff []int
	c := make(chan int)
	for scanner.Scan() {
		line := scanner.Text()
		buff = append(buff, convertStringToInt(line))

		if counter%6 == 0 {
			go diff(c, buff[counter-6:counter])
		}

		for {
			res, ok := <-c
			if !ok {
				break
			}
			counter += res
		}
	}
	defer file.Close()

	return counter - 1
}

func diff(c chan int, ar []int) {
	if len(ar) == 0 {
		close(c)
	}
	var times, sumA, sumB, sumC, sumD int
	if len(ar) == 4 {
		sumA = ar[0] + ar[1] + ar[2]
		sumB = ar[1] + ar[2] + ar[3]
		sumC = ar[2] + ar[3] + ar[4]
		sumD = ar[3] + ar[4] + ar[5]
	}

	if sumB > sumA {
		times++
	}
	if sumD > sumC {
		times++
	}
	c <- times
}

func convertStringToInt(ns string) int {
	n, err := strconv.Atoi(ns)
	if err != nil {
		panic("err converting number")
	}
	return n
}

func shiftArray(ar [3]int, lastNumber int) [3]int {
	// fmt.Printf("shifArray: %+v, %d\n", ar, lastNumber)
	for i := 1; i < len(ar); i++ {
		ar[i-1] = ar[i]
	}
	ar[len(ar)-1] = lastNumber
	// fmt.Printf("shifArray end: %+v\n", ar)
	return ar
}
