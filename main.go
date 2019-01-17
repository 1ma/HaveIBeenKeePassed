package main

import (
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/keepass"
	"os"
)

func main() {
	file, err := os.Open("/home/marcel/workspace/HaveIBeenKeePassed/Sample.kdbx")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	raw, err := keepass.Decode(file, "1234")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw))
}
