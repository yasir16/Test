package main

import (
	"fmt"
)

func reverseString(str string) string {
	var result string
	for _, singleStr := range str {
		result = string(singleStr) + result
	}
	return result
}

func solution(arr []string) []string {
	var result []string
	for _, strArry := range arr {
		if strArry == reverseString(strArry) {
			result = append(result, strArry)
		}
	}
	return result
}
func main() {
	arr := []string{"232", "moom", "test"}
	fmt.Println(solution(arr))
}
