package kata

func Solution(str string) []string {
	var result []string

	// if string has odd length, add _ to make it even
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}

	return result
}
