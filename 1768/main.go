package main

func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))

	var minLen int
	if len(word1) > len(word2) {
		minLen = len(word2)
	} else {
		minLen = len(word1)
	}

	for i := 0; i < minLen; i++ {
		result = append(result, word1[i])
		result = append(result, word2[i])
	}

	if len(word2) > len(word1) {
		result = append(result, word2[minLen:]...)
	} else {
		result = append(result, word1[minLen:]...)
	}

	return string(result)
}

func main() {
	println(mergeAlternately("abc", "pqr"))
	println(mergeAlternately("ab", "pqrs"))
	println(mergeAlternately("abcd", "pq"))
}
