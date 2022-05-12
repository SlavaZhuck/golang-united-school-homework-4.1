package reverse_string

// package main

func ReverseString(input string) (output string) {
	// solution goes here
	temp := []rune(input)
	i := 0
	for _, char := range input {
		temp[len(input)-1-i] = char
		i++
	}
	return string(temp)
}

// func main() {
// 	m := "test_string"

// 	fmt.Println(m)
// 	fmt.Println(ReverseString(m))
// }
