// Package hibp encapsulates all code related to communicating
// with the HaveIBeenPwned API and parsing the HTTP responses.
package hibp

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/tobischo/gokeepasslib/v2"
	"net/http"
	"strings"
)

func Check(c <-chan gokeepasslib.Entry) {
	client := http.DefaultClient

	for entry := range c {
		rawSHA1 := sha1.Sum([]byte(entry.GetPassword()))
		hexSHA1 := strings.ToUpper(hex.EncodeToString(rawSHA1[:]))

		req, _ := http.NewRequest("GET", "https://api.pwnedpasswords.com/range/"+hexSHA1[:5], nil)
		req.Header.Set("User-Agent", "HaveIBeenKeePassed/0.1")

		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s: API ERROR (could not check password)\n", entry.GetTitle())
			continue
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
			fmt.Printf("%s: COMPROMISED (%s)\n", entry.GetTitle(), entry.GetPassword())
		} else {
			fmt.Printf("%s: SAFE\n", entry.GetTitle())
		}
	}
}
