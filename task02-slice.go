package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	for i, v := range input {
		j := len(input) - i - 1
		result[j] = v
	}
	return
}
