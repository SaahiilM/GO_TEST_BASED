package arrays_slices

func Sum(number []int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTrails(numbersToTrail ...[]int) []int {
	var trails []int
	for _, numbers := range numbersToTrail {
		tail := numbers[1:]
		trails = append(trails, Sum(tail))
	}
	return trails
}
