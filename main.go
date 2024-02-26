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

	f_stats, _ := file.Stat()
	b := make([]byte, f_stats.Size())
	file.Read(b)
	fmt.Println(string(b))
}
