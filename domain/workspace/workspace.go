package workspace

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/domain/node"
)

type Workspaces []*Workspace
type Leaves []*node.FlatNode
type ErrLeafNotFound struct {
	Ordinal int
}

const WorkspacesCountEstimate = 10
const LeavesCountEstimate = 4

func (e *ErrLeafNotFound) Error() string {
	return fmt.Sprintf("leaf with %d ordinal not found", e.Ordinal)
}

type Workspace struct {
	Root   *node.Node `json:"-"`
	Leaves Leaves     `json:"nodes"`
}

func ExtractWorkspaces(root *node.Node) Workspaces {
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

func extractLeaves(root *node.Node, workspace *Workspace) {
	if len(root.Nodes) == 0 {
		workspace.Leaves = append(workspace.Leaves, &node.FlatNode{
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

func extractWorkspacesImpl(current *node.Node, workspaces *Workspaces) {
	// skip scratchpad
	if current.Name == "__i3" {
		return
	}
	if current.Type == "workspace" {
		workspace := &Workspace{
			Root:   current,
			Leaves: make([]*node.FlatNode, 0, LeavesCountEstimate),
		}
		extractLeaves(current, workspace)
		*workspaces = append(*workspaces, workspace)
		return
	}
	for _, child := range current.Nodes {
		extractWorkspacesImpl(child, workspaces)
	}
}

func (workspaces *Workspaces) Find(ordinal int) (*node.FlatNode, error) {
	for _, ws := range *workspaces {
		for _, leaf := range ws.Leaves {
			if leaf.Ordinal == ordinal {
				return leaf, nil
			}
		}
	}
	return nil, &ErrLeafNotFound{Ordinal: ordinal}
}
