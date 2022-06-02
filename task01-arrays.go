package homework

func average(input [15]float32) (result float32) {
	sum := 0
	for i := range input {
		sum += i
	}
	result = float32(sum) / float32(len(input))
	return
}
