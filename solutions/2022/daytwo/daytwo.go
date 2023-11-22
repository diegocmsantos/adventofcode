package daytwo

import (
	"github.com/diegocmsantos/adventofcode/solutions"
)

// Part 1
// A : X Rock
// B : Y Paper
// C : Z Scissors
// var scores = map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}

// Part 2
// A : X Lose
// B : Y Draw
// C : Z Win
var scores = map[string]int{"B X": 1, "C Y": 6, "A Z": 8, "A X": 3, "B Y": 5, "C Z": 7, "C X": 2, "A Y": 4, "B Z": 9}

func Solution() int {
	input, err := solutions.ReadFile("/Users/macuser/go/src/github.com/diegocmsantos/adventofcode/solutions/2022/daytwo/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, i := range input {
		sum += scores[i]
	}
	return sum
}
