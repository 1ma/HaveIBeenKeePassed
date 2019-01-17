package main


import (
	"fmt"
	"github.com/tobischo/gokeepasslib/v2"
	"os"
)

func main() {
	file, _ := os.Open("/home/marcel/workspace/HaveIBeenKeePassed/Sample.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("1234")
	err := gokeepasslib.NewDecoder(file).Decode(db)

	if err != nil {
		panic(err)
	}

	db.UnlockProtectedEntries()

	entry := db.Content.Root.Groups[0].Entries[0]
	fmt.Println(entry.GetTitle())
	fmt.Println(entry.GetPassword())
}
