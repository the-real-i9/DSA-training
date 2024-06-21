package HTMLDOM

type DOMTreeElementNode struct {
	ChildElementCount      int                   // element
	ChildNodes             []*DOMTreeElementNode // node
	Children               []*DOMTreeElementNode // element
	ClassList              []string              // element
	ClassName              string                // element
	Data                   string                // node
	FirstChild             *DOMTreeElementNode   // node
	FirstElementChild      *DOMTreeElementNode   // element
	Id                     string                // element
	LastChild              *DOMTreeElementNode   // node
	LastElementChild       *DOMTreeElementNode   // element
	NextElementSibling     *DOMTreeElementNode   // node | element
	NextSibling            *DOMTreeElementNode   // node | element
	NodeName               string                // node
	NodeType               int                   // node | element
	NodeValue              string                // node | element
	ParentElement          *DOMTreeElementNode   // node | element
	ParentNode             *DOMTreeElementNode   // node | element
	PreviousElementSibling *DOMTreeElementNode   // node | element
	PreviousSibling        *DOMTreeElementNode   // node | element
	TagName                string                // element
	Title                  string                // element
}

// Insert newNodes after the reference node
func (eln *DOMTreeElementNode) After(newNodes ...DOMTreeElementNode) {

}

// Append newNodes (children) to the reference node
func (eln *DOMTreeElementNode) Append(newNodes ...DOMTreeElementNode) {

}

// Append newChild to the reference node
func (eln *DOMTreeElementNode) AppendChild(newChild DOMTreeElementNode) {

}

// Insert newNodes before the reference node
func (eln *DOMTreeElementNode) Before(newNodes ...DOMTreeElementNode) {

}

// Does reference node contain this node?
func (eln *DOMTreeElementNode) Contains(node DOMTreeElementNode) bool {
	return false
}

// Remove nodeToRemove from reference node's children
func (eln *DOMTreeElementNode) RemoveChild(nodeToRemove DOMTreeElementNode) {

}

// Replace reference node's targetNode child with replacementChild
func (eln *DOMTreeElementNode) ReplaceChild(targetNode DOMTreeElementNode, replacementChild DOMTreeElementNode) {

}

// Set reference node's children to the provided replacementChildren
func (eln *DOMTreeElementNode) ReplaceChildren(replacementChildren ...DOMTreeElementNode) {

}

// Among reference node's children insert newChild before targetNode
func (eln *DOMTreeElementNode) InsertBefore(targetNode DOMTreeElementNode, newChild DOMTreeElementNode) {

}

// Insert the provided element at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeElementNode) InsertAdjacentElement(where string, element DOMTreeElementNode) {

}

// Insert the provided html (after parsing to DOMTreeElementNode) at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeElementNode) InsertAdjacentHTML(where string, html string) {

}

// Insert the provided text (as DOMTreeElementNode) at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeElementNode) InsertAdjacentText(where string, text string) {

}

// Is the childNodes property greater than zero?
func (eln *DOMTreeElementNode) HasChildNodes() bool {
	return len(eln.ChildNodes) > 0
}
