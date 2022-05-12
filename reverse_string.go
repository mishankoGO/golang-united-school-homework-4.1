package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	for _, elem := range input {
		output = string(elem) + output
	}
	return output
}
