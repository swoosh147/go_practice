package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = ":"
		fmt.Println(i, arg)
	}
	// fmt.Println(s)
	finish := time.Since(start).Seconds()
	fmt.Println(finish)

}
