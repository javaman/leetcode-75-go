package main

import "fmt"

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	vowels := map[rune]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	result := []rune(s)

	for left < right {
		_, vowel := vowels[result[left]]
		for left < right && !vowel {
			left++
			_, vowel = vowels[result[left]]
		}
		_, vowel = vowels[result[right]]
		for left < right && !vowel {
			right--
			_, vowel = vowels[result[right]]
		}
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return string(result)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
