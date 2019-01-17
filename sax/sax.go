// Package sax contains a SAX parser for the KeePass2 XML document.
package sax

import (
	"encoding/xml"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"io"
)

func Parse(r io.Reader, c chan<- types.Entry) error {
	defer close(c)

	key := ""
	entry := types.Entry{}
	inHistory := false
	inEntry := false
	inKey := false
	inValue := false

	decoder := xml.NewDecoder(r)

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "History":
				inHistory = true

			case "Entry":
				inEntry = true

			case "Key":
				inKey = true

			case "Value":
				inValue = true
			}
		case xml.CharData:
			// Only proceed if this is the content of a Key or a Value node inside an
			// Entry node, but outside a History node.
			if (inKey == false && inValue == false) || inEntry == false || inHistory == true {
				break
			}

			if inKey == true {
				key = string(t)
			}

			switch key {
			case "Password":
				entry.Password = string(t)
			case "Title":
				entry.Title = string(t)
			case "URL":
				entry.URL = string(t)
			case "UserName":
				entry.UserName = string(t)
			default:
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "History":
				inHistory = false

			case "Entry":
				if inHistory == false {
					c <- entry
					entry = types.Entry{}
				}
				inEntry = false

			case "Key":
				inKey = false

			case "Value":
				inValue = false
			}
		}
	}

	return nil
}
