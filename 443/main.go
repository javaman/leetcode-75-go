package main

import "strconv"

func compress(chars []byte) int {
	encode := func(np int, nc byte, l int) int {
		chars[np] = nc
		np++
		if l > 1 {
			str := strconv.Itoa(l)
			for _, v := range str {
				chars[np] = byte(v)
				np++
			}
		}
		return np
	}
	nextPosition := 0
	nextChar := chars[0]
	l := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] != nextChar {
			nextPosition = encode(nextPosition, nextChar, l)
			l = 1
			nextChar = chars[i]
		} else {
			l++
		}
	}
	nextPosition = encode(nextPosition, nextChar, l)
	return nextPosition
}
