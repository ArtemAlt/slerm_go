package main

import "fmt"

func isSorted(arr []string) bool {
	var result bool
	for i := 0; i < len(arr)-1; i++ {
		if !(arr[i] > arr[i+1]) {
			result = true
			continue
		}
		result = false
	}
	return result
}
func main() {
	arr := []string{"aa", "bb", "cc", "aa"}
	result := isSorted(arr)
	fmt.Println(result)
}
