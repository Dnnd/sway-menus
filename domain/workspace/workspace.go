package workspace

import (
	"fmt"
	"github.com/Dnnd/sway-menus/dmenu"
	"strings"
)

type Workspaces []*Workspace
type Leaves []*FlatNode

type ErrLeafNotFound struct {
	Ordinal int
}

const WorkspacesCountEstimate = 10
const LeavesCountEstimate = 4

func (e *ErrLeafNotFound) Error() string {
	return fmt.Sprintf("leaf with %d ordinal not found", e.Ordinal)
}

type Workspace struct {
	Root   *Node  `json:"-"`
	Leaves Leaves `json:"nodes"`
}

func ExtractWorkspaces(root *Node) Workspaces {
	workspaces := make(Workspaces, 0, WorkspacesCountEstimate)
	extractWorkspacesImpl(root, &workspaces)
	setOrdinals(workspaces)
	return workspaces
}

func setOrdinals(workspaces Workspaces) {
	counter := 0
	for _, ws := range workspaces {
		for _, leaf := range ws.Leaves {
			leaf.Ordinal = counter
			counter += 1
		}
	}
}

func extractLeaves(root *Node, workspace *Workspace) {
	if len(root.Nodes) == 0 {
		workspace.Leaves = append(workspace.Leaves, &FlatNode{
			Name: root.Name,
			Id:   root.Id,
		})
		return
	}
	for _, child := range root.Nodes {
		extractLeaves(child, workspace)
	}
	for _, child := range root.FloatingNodes {
		extractLeaves(child, workspace)
	}
}

func extractWorkspacesImpl(current *Node, workspaces *Workspaces) {
	// skip scratchpad
	if current.Name == "__i3" {
		return
	}
	if current.Type == "workspace" {
		workspace := &Workspace{
			Root:   current,
			Leaves: make([]*FlatNode, 0, LeavesCountEstimate),
		}
		extractLeaves(current, workspace)
		*workspaces = append(*workspaces, workspace)
		return
	}
	for _, child := range current.Nodes {
		extractWorkspacesImpl(child, workspaces)
	}
}

func (workspaces *Workspaces) Find(ordinal int) (*FlatNode, error) {
	for _, ws := range *workspaces {
		for _, leaf := range ws.Leaves {
			if leaf.Ordinal == ordinal {
				return leaf, nil
			}
		}
	}
	return nil, &ErrLeafNotFound{Ordinal: ordinal}
}

const MenuEntryBytesEstimate = 16

type WorkspacesMenu struct {
	entries string
	lines   int
}

func (w WorkspacesMenu) TotalLines() int {
	return w.lines
}

func (w WorkspacesMenu) AsString() string {
	return w.entries
}

func (workspaces *Workspaces) ToMenuData() dmenu.MenuData {
	menu := strings.Builder{}
	menu.Grow(len(*workspaces) * MenuEntryBytesEstimate)
	lines := 0
	for _, ws := range *workspaces {
		lines += len(ws.Leaves)
		for _, node := range ws.Leaves {
			menu.WriteString(ws.Root.Name)
			menu.WriteString("\t")
			menu.WriteString(node.Name)
			menu.WriteString("\n")
		}
	}
	return WorkspacesMenu{entries: menu.String(), lines: lines}
}
