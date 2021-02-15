package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	increment()
	join()
}

func increment() {
	defer timeTrack("increment")()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func join() {
	defer timeTrack("join")()
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func timeTrack(fname string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", fname, time.Since(start))
	}
}
