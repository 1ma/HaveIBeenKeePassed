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

	checkEntries(c)
}

func checkEntries(c <-chan types.Entry) {
	for entry := range c {
		rawSHA1 := sha1.Sum([]byte(entry.Password))
		hexSHA1 := strings.ToUpper(hex.EncodeToString(rawSHA1[:]))

		client := http.DefaultClient
		req, _ := http.NewRequest("GET", "https://api.pwnedpasswords.com/range/"+hexSHA1[:5], nil)
		req.Header.Set("User-Agent", "HaveIBeenKeePassed/0.1")

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		compromised := false
		scn := bufio.NewScanner(res.Body)
		for scn.Scan() {
			if hexSHA1[5:] == scn.Text()[:35] {
				compromised = true
				break
			}
		}

		_ = res.Body.Close()

		if compromised == true {
			fmt.Printf("%s: COMPROMISED (%s)\n", entry.Title, entry.Password)
		} else {
			fmt.Printf("%s: SAFE\n", entry.Title)
		}
	}
}
