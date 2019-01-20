package keepass

import (
	"github.com/1ma/HaveIBeenKeePassed/types"
	"github.com/tobischo/gokeepasslib/v2"
)

func Parse(db *gokeepasslib.Database) <-chan types.Entry {
	c := make(chan types.Entry, 1024)
	_ = db.UnlockProtectedEntries()
	parseGroups(db.Content.Root.Groups, c)

	close(c)
	return c
}

func parseGroups(gs []gokeepasslib.Group, c chan<- types.Entry) {
	for i := range gs {
		parseGroup(&gs[i], c)
	}
}

func parseGroup(g *gokeepasslib.Group, c chan<- types.Entry) {
	if g.Groups != nil {
		parseGroups(g.Groups, c)
	}

	es := g.Entries
	if es == nil {
		return
	}

	for i := range es {
		e := types.Entry{Title: es[i].GetTitle(), Password:es[i].GetPassword()}
		c <- e
	}
}
