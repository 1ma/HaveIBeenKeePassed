package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/sax"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"net/http"
	"os"
	"strings"
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
		password := entry.Password

		jash := sha1.Sum([]byte(password))
		reference := strings.ToUpper(hex.EncodeToString(jash[:]))
		prefics := strings.ToUpper(hex.EncodeToString(jash[:3])[:5])

		url := fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefics)

		client := http.DefaultClient
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "HaveIBeenKeePassed")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)

		compromized := false
		for scanner.Scan() {
			fullHash := prefics + scanner.Text()[:35]
			// fmt.Println(fullHash)

			if reference == fullHash {
				compromized = true
				fmt.Printf("%s: COMPROMiZED (%s)\n", entry.Title, entry.Password)
			}
		}

		if compromized == false {
			fmt.Println("%s: SAFE\n", entry.Title)
		}
	}
}
