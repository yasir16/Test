package main

import (
	"fmt"
	"reflect"
)

func mappingStr(str string) map[string]int {
	mapping := make(map[string]int)
	for _, word := range str {
		mapping[string(word)]++
	}
	return mapping
}
func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	mappingstr1 := mappingStr(str1)
	mappingstr2 := mappingStr(str2)
	if !reflect.DeepEqual(mappingstr1, mappingstr2) {
		return false
	}
	return true

}
func main() {
	fmt.Println(isAnagram("asuu", "usaa"))
}
