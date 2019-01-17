// Package keepass encapsulates code related to decrypting
// KeePass2 databases and retrieving their inner contents.
package keepass

import (
	"github.com/tobischo/gokeepasslib/v2"
	"io"
)

// Given an open io.Reader pointing at an encrypted KeePass2 database
// and a candidate password, Decode returns either the database's underlying
// XML document as a byte slice or an error.
func Decode(file io.Reader, pwd string) (raw []byte, err error) {
	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(pwd)
	err = gokeepasslib.NewDecoder(file).Decode(db)

	if err != nil {
		return nil, err
	}

	return db.Content.RawData, nil
}
