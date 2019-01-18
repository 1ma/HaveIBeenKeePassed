package main

import (
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/sax"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"os"
)

func main() {
	file, err := os.Open("/home/marcel/workspace/HaveIBeenKeePassed/Sample.xml")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	c := make(chan types.Entry, 128)
	go sax.Parse(file, c)

	for entry := range c {
		fmt.Printf("%+v\n", entry)
	}
}
