package HTMLDOMParser

type DOMTreeNode struct {
	Name            string // html, b, u, em, strong
	Parent          *DOMTreeNode
	Children        []*DOMTreeNode
	PreviousSibling *DOMTreeNode
	NextSibling     *DOMTreeNode

	// data
	TextContent string
}
