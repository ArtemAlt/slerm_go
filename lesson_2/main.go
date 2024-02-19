package main

import "fmt"

func IsContainsDuplicates(arr []int) bool {
	result := false
	mapDuplicate := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if mapDuplicate[arr[i]] {
			result = true
			break
		}
		mapDuplicate[arr[i]] = true
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	result := IsContainsDuplicates(arr)
	fmt.Println(result)
}
