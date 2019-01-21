// Package keepass2 encapsulates all code related to
// parsing the passwords out of KeePass2 databases.
package keepass2

import "github.com/tobischo/gokeepasslib/v2"

// Parse accepts a pointer to an open gokeepasslib.Database
// and a gokeepass.Entry channel, and will unlock the database's
// protected entries, parse the whole XML document and put in
// the channel all the password entries it finds in the database.
//
// Once it finishes it will close the channel.
func Parse(db *gokeepasslib.Database, c chan<- gokeepasslib.Entry) {
	_ = db.UnlockProtectedEntries()

	parseGroup(&db.Content.Root.Groups[0], c)

	close(c)
}

func parseGroup(g *gokeepasslib.Group, c chan<- gokeepasslib.Entry) {
	for i := range g.Groups {
		parseGroup(&g.Groups[i], c)
	}

	parseEntries(g.Entries, c)
}

func parseEntries(entries []gokeepasslib.Entry, c chan<- gokeepasslib.Entry) {
	for i := range entries {
		c <- entries[i]
	}
}
