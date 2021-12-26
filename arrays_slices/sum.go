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
