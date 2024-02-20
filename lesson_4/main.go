package main

import "fmt"

func CountChars(str string) map[rune]int {
	elements := make(map[rune]int)
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		if elements[chars[i]] != 0 {
			elements[chars[i]]++
			continue
		}
		elements[chars[i]] = 1
	}
	return elements
}

func main() {
	s1 := "съешь ещё этих мягких французских булок, да выпей чаю"
	result := CountChars(s1)
	for key, value := range result {
		fmt.Printf("Char '%c' %d \n", key, value)
	}
}
