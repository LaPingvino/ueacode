package ueacode

import "strings"

func Number(letter rune) int {
	if letter > 90 {
		return int(letter) - 96
	} else {
		return int(letter) - 64
	}
}

func Letter(number int, ucase bool) rune {
	if ucase {
		return rune(number + 64)
	} else {
		return rune(number + 96)
	}
}

func Check(code string) bool {
	letters := strings.Split(code, "-")
	if len(letters) < 2 {
		return false
	}
	return Calculate(letters[0]) == Number(rune(letters[1][0]))
}

func Calculate(code string) (r int) {
	for i, c := range code {
		r += Number(c) * (i + 2)
	}
	r %= 26
	r = 26 - r
	return r
}
