package arrays_slices

func Sum(number []int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}
