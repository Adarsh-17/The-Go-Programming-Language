package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	fmt.Println(strings.Join(os.Args, ","))
	end1 := time.Now()
	start2 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += s + sep + arg
		sep = " "
	}
	fmt.Println(s)
	end2 := time.Now()
	dur1 := end1.Sub(start1)
	dur2 := end2.Sub(start2)
	fmt.Println("Duration of strings.Join:", dur1)
	fmt.Println("Duration of manual concatenation:", dur2)
}
