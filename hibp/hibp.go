// Package hibp encapsulates all code related to communicating
// with the HaveIBeenPwned API and parsing its HTTP responses.
package hibp

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/tobischo/gokeepasslib/v2"
	"io"
	"net/http"
	"strings"
)

func Check(c <-chan gokeepasslib.Entry) {
	client := http.DefaultClient

	for entry := range c {
		rawSHA1 := sha1.Sum([]byte(entry.GetPassword()))
		hexSHA1 := strings.ToUpper(hex.EncodeToString(rawSHA1[:]))

		body, err := attackHIBPApi(hexSHA1[:5], client)
		if err != nil {
			fmt.Printf("%s: API ERROR (could not check password)\n", entry.GetTitle())
			continue
		}

		if isCompromised(hexSHA1[5:], body) {
			fmt.Printf("%s: COMPROMISED (%s)\n", entry.GetTitle(), entry.GetPassword())
		} else {
			fmt.Printf("%s: SAFE\n", entry.GetTitle())
		}
	}
}

func attackHIBPApi(prefix string, client *http.Client) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", "https://api.pwnedpasswords.com/range/"+prefix, nil)
	req.Header.Set("User-Agent", "HaveIBeenKeePassed/0.1")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func isCompromised(suffix string, content io.ReadCloser) bool {
	defer content.Close()

	scn := bufio.NewScanner(content)
	for scn.Scan() {
		if suffix == scn.Text()[:35] {
			return true
		}
	}

	return false
}
