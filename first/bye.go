package main

import (
	"fmt"
	f "fmt"
	format "fmt"
	"time"
)

func bye() {
	fmt.Println("Bye!")
	start := time.Now()
	t := time.Now()
	elapsed := t.Sub(start)
	format.Println("Time:", elapsed.String())
	f.Println(t.Format("2006-01-02 15:04:05"))
}
