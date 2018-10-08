// The check command gives a Unix return code -1 when no code is provided
// as well as a short message, 0 and no text when successful and 1 to 26
// plus printing one letter when a code is given but it doesn't check out.
// With a correct first part of a code, this should be a correct check digit.
package main

import "github.com/lapingvino/ueacode"
import "os"

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
		println("mankas kodo")
	}
	if ueacode.Check(os.Args[1]) {
		os.Exit(0)
	}
	code := ueacode.Calculate(os.Args[1])
	println(string(ueacode.Letter(code, false)))
	os.Exit(code)
}
