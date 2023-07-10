package string

func Concat(input []string) string {
	result := ""
	for _, v := range input {
		result += v + " "
	}

	return result
}
