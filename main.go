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

	fmt.Println(string(db.Content.RawData))
}
