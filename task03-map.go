package homework

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, len(input))
	for _, v := range input {
		result = append(result, v)
	}
	return
}
