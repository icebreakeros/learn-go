package gover

import (
	"fmt"
	"runtime"
)

// Gover is a function that returns Golang version
func Gover() {
	fmt.Println("GO version: ", runtime.Version())
}
