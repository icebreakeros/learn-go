/*
Package main is the main package of this program.
*/
package main

import (
	"fmt"
	"runtime"
)

const ok = true

func main() {
	var hello = "Hi there!!!"
	var prompt = "This is my first program! "
	if 5 > 2 {
		for i := 0; i < 3; i++ {
			hey()
			fmt.Println("CPU numbers:", runtime.NumCPU())
			fmt.Println("Golang version:", runtime.Version())
		}
	}
	fmt.Println(hello, ok)
	fmt.Println(prompt, ok)
	bye()
}
