package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := os.Args[1]
	if filepath == "" {
		fmt.Println("no path provided")
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
}
