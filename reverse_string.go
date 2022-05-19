package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	output_runes := make([]rune, len(runes))
	for i, r := range runes {
		output_runes[len(runes)-i-1] = r
	}
	return string(output_runes)
}
