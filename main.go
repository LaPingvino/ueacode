// The package ueacode provides support for creating and checking codes
// as used by e.g. the Esperanto organizations UEA and FEL. The algorithm 
// used here was originally devised by Danny ten Haaf and reconstructed from
// memory from conversations with Ralph with testing on some actual codes by
// Joop Kiefte, motivated by a conversation with Robin van der Vliet.
package ueacode

import "strings"

// Number checks with a very loose check if a letter is uppercase or lowercase
// and returns the number used for the checksum. 
func Number(letter rune) int {
	if letter > 90 {
		return int(letter) - 96
	} else {
		return int(letter) - 64
	}
}

// Letter reverses the process used by Number, and can provide either an upper
// or lower case result as a rune.
func Letter(number int, ucase bool) rune {
	if ucase {
		return rune(number + 64)
	} else {
		return rune(number + 96)
	}
}

// Check verifies if the code given is a correct code. The code is split on the 
// hyphen sign '-'. The code also returns false if this sign is not present.
func Check(code string) bool {
	letters := strings.Split(code, "-")
	if len(letters) < 2 {
		return false
	}
	return Calculate(letters[0]) == Number(rune(letters[1][0]))
}

// Calculate generates a checksum number from 1 to 26 that can be read out as a
// letter with the Letter function.
func Calculate(code string) (r int) {
	for i, c := range code {
		r += Number(c) * (i + 2)
	}
	r %= 26
	r = 26 - r
	return r
}
