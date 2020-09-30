package main

import (
	"fmt"
	"reflect"
)

func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	hash := make(map[string]int)
	hash2 := make(map[string]int)

	for _, r := range str1 {
		j := hash[string(r)]
		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}
	for _, r := range str2 {
		j := hash2[string(r)]
		if j == 0 {
			hash2[string(r)] = 1
		} else {
			hash2[string(r)] = j + 1
		}
	}
	isOK := false
	for _, value := range hash {
		if value%2 != 0 {
			isOK = false
		}
	}
	if !reflect.DeepEqual(hash, hash2) && isOK == false {
		return false
	}

	return true
}
func main() {
	fmt.Println(isAnagram("asuu", "usaa"))
}
