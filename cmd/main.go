package main

import (
	"fmt"
	"log"
	"os"
	fileexplorer "vivekdubey/golang-practice/internal/file_explorer"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("Expects at least one argument")
	}
	files := os.Args[1:]
	m := make(map[string]bool)
	for _, v := range files {
		err := fileexplorer.CreateFile(v)
		if err != nil {
			m[v] = true
		} else {
			m[v] = false
		}
	}
	for k, v := range m {
		fmt.Printf("File: %v, errored: %v \n", k, v)

	}
}
