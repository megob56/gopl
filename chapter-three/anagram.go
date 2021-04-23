package main

import "fmt"

func main() {
	fmt.Println(anagram("hiy", "ih"))
}

func anagram(str1, str2 string) string {
	count := 0
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				count++
			}
		}
	}
	if count == len(str1) && count == len(str2) {
		return "these words are anagrams!"
	}
	return "not anagrams"
}
