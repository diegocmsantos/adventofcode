package dayone

import (
	"strconv"

	"github.com/diegocmsantos/adventofcode/solutions"
)

func Solution() int {
	input, err := solutions.ReadFile("/Users/macuser/go/src/github.com/diegocmsantos/adventofcode/solutions/2022/dayone/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	max := 0
	for _, line := range input {
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}
	return max
}
