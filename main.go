package main

import (
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/hibp"
	"github.com/1ma/HaveIBeenKeePassed/keepass2"
	"github.com/tobischo/gokeepasslib/v2"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s database.kdbx\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s: No such file\n", os.Args[1])
		os.Exit(1)
	}

	defer file.Close()

	fmt.Print("Enter Masker Key: ")
	mk, _ := terminal.ReadPassword(0)
	fmt.Println()

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(string(mk))
	err = gokeepasslib.NewDecoder(file).Decode(db)

	if err != nil {
		fmt.Printf("%s: Could not decrypt database\n", os.Args[1])
		os.Exit(1)
	}

	c := make(chan gokeepasslib.Entry, 128)

	go keepass2.Parse(db, c)

	hibp.Check(c)
}