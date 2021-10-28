package generic_menu

import (
	"strings"
)

type EntryTag string

type Entry struct {
	Tag   EntryTag
	Label string
}

type GenericMenu []Entry

func (g GenericMenu) TotalLines() int {
	return len(g)
}

func (g GenericMenu) AsString() string {
	builder := strings.Builder{}
	entriesInMenu := len(g)
	builder.Grow(entriesInMenu * 10) // estimate
	for _, entry := range g {
		builder.WriteString(entry.Label)
		builder.WriteString("\n")
	}
	return builder.String()
}
