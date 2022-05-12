package reverse_string

// package main

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// func main() {
// 	m := "абвгдеёжз"

// 	fmt.Println(m)
// 	fmt.Println(ReverseString(m))
// }
