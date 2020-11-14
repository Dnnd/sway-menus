package workspace

import "strings"

const MenuEntryBytesEstimate = 16

type Dmenu struct {
	Entries string
	Lines   int
}

func (wss *Workspaces) ToDmenu() Dmenu {
	menu := strings.Builder{}
	menu.Grow(len(*wss) * MenuEntryBytesEstimate)
	lines := 0
	for _, ws := range *wss {
		lines += len(ws.Leaves)
		for _, node := range ws.Leaves {
			menu.WriteString(ws.Root.Name)
			menu.WriteString("\t")
			menu.WriteString(node.Name)
			menu.WriteString("\n")
		}
	}
	return Dmenu{Entries: menu.String(), Lines: lines}
}
