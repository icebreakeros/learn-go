package main

import "fmt"

const ok = true

func main() {
	var hello = "Hi there!!!"
	var prompt = "This is my first program! "

	hey()
	fmt.Println(hello, ok)
	fmt.Println(prompt, ok)
	bye()
}
