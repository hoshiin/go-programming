package main

import (
	"bufio"
	"fmt"
	"os"
)

type dupFiles struct {
	counts int
	files  map[string]struct{}
}

type dupLines map[string]dupFiles

func main() {
	dupLines := make(dupLines)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, dupLines)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, dupLines)
			f.Close()
		}
	}
	for line, dupFiles := range dupLines {
		if dupFiles.counts > 1 {
			fmt.Printf("%d\t%s\t", dupFiles.counts, line)
			for file := range dupFiles.files {
				fmt.Printf("%s\t", file)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(fileName string, f *os.File, dupLines dupLines) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dup, ok := dupLines[input.Text()]
		if !ok {
			files := make(map[string]struct{})
			files[fileName] = struct{}{}
			dupLines[input.Text()] = dupFiles{
				counts: 1,
				files:  files,
			}
			continue
		}
		counts := dup.counts + 1
		if _, ok := dup.files[fileName]; !ok {
			dup.files[fileName] = struct{}{}
		}
		dupLines[input.Text()] = dupFiles{counts: counts, files: dup.files}
	}
	// 注意： input.Err()からのエラーの可能性を無視している
}
