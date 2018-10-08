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
