package main

import (
	"github.com/1ma/HaveIBeenKeePassed/hibp"
	"github.com/1ma/HaveIBeenKeePassed/sax"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("need at least one argument")
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		panic(err)
	}

	c := make(chan types.Entry, 64)

	go sax.Parse(file, c)

	hibp.Check(c)
}
