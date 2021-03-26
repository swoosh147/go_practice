// Вывод аргументов командной строки

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ":"
	}
	fmt.Println(len(os.Args))
	// fmt.Println(os.Args)
	fmt.Println(s)
	finish := time.Since(start).Seconds()
	fmt.Println(finish)
}
