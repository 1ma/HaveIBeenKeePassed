// Package hibp encapsulates all code related to communicating
// with the HaveIBeenPwned HTTP API and parsing the responses.
package hibp

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"net/http"
	"strings"
)

func Check(c <-chan types.Entry) {
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
