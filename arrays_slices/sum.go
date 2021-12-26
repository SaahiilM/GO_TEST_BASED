package arraysslices

func Sum(number [5]int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}
