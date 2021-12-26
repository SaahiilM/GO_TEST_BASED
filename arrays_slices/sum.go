package arrays_slices

func Sum(number []int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	lengthOfNumbers := len(numbersToSum)
	endsum := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		endsum[i] = Sum(numbers)
	}

	return endsum
}
