// Package sax contains a SAX parser for the KeePass2 XML document.
package sax

import (
	"encoding/xml"
	"github.com/1ma/HaveIBeenKeePassed/types"
	"io"
)

// This function consumes an io.Reader pointing to a KeePass2 XML
// document, and queues in the channel all the active password
// entries it finds as it parses the document.
//
// Active passwords are found inside Entry elements that are
// not nested in a History element.
//
// Sample layout:
//
// <Entry>
//   <String>
//     <Key>Password</Key>
//     <Value ProtectInMemory="True">12345</Value>
//   </String>
//   <String>
//     <Key>Title</Key>
//     <Value>Sample Entry #2</Value>
//   </String>
// </Entry>
func Parse(r io.Reader, c chan<- types.Entry) {
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
			panic(err)
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
}
