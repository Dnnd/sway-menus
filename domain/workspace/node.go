package workspace

type Node struct {
	Nodes         []*Node `json:"nodes,omitempty"`
	FloatingNodes []*Node `json:"floating_nodes,omitempty"`
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
}

type FlatNode struct {
	Name    string
	Id      int
	Ordinal int
}
