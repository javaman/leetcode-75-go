package main

func main() {
	println(gcdOfStrings("ABCABC", "ABC"))
	println(gcdOfStrings("ABABAB", "ABAB"))
	println(gcdOfStrings("LEET", "CODE"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) != (str2 + str1) {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(i1, i2 int) int {
	for i1 != i2 {
		if i1 > i2 {
			i1 -= i2
		} else {
			i2 -= i1
		}
	}
	return i1
}
