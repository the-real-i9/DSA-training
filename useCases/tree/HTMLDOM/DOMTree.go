package HTMLDOM

// Note: All block level elements have first and last text nodes of "\n"
func NewDocument(title string) *DOMTree {

	htmlElem := NewElementNode("html", NilENPs)

	headElem := NewElementNode("head", NilENPs)

	bodyElem := NewElementNode("body", NilENPs)

	titleElem := NewElementNode("title", NilENPs)

	titleText := NewTextNode(title)

	titleElem.AppendChild(titleText)
	headElem.AppendChild(titleElem)
	htmlElem.Append(headElem, bodyElem)

	return &DOMTree{Root: htmlElem}
}

type DOMTree struct {
	Root *DOMTreeNode
}
