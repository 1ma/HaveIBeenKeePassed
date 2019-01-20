package keepass

import (
	"github.com/1ma/HaveIBeenKeePassed/types"
	"github.com/tobischo/gokeepasslib/v2"
)

func Parse(db *gokeepasslib.Database) <-chan types.Entry {
	c := make(chan types.Entry, 4)

	go start(db, c)

	return c
}

func start(db *gokeepasslib.Database, c chan<- types.Entry) {
	_ = db.UnlockProtectedEntries()

	parseGroup(&db.Content.Root.Groups[0], c)

	close(c)
}

func parseGroup(g *gokeepasslib.Group, c chan<- types.Entry) {
	for i := range g.Groups {
		parseGroup(&g.Groups[i], c)
	}

	parseEntries(g.Entries, c)
}

func parseEntries(entries []gokeepasslib.Entry, c chan<- types.Entry) {
	for i := range entries {
		c <- types.Entry{
			Title: entries[i].GetTitle(),
			Password:entries[i].GetPassword(),
		}
	}
}
