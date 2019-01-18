package main

import (
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/hibp"
	"github.com/1ma/HaveIBeenKeePassed/sax"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filepath.xml\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s: No such file\n", os.Args[1])
		os.Exit(1)
	}

	defer file.Close()

	c := make(chan types.Entry, 64)

	go sax.Parse(file, c)

	hibp.Check(c)
}
