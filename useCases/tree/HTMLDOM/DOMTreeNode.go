package HTMLDOM

import "strings"

type ElementNodeProps struct {
	ClassName string
	Title     string
	Id        string
}

var NilENPs = ElementNodeProps{}

func NewElementNode(tagName string, props ElementNodeProps) *DOMTreeNode {
	return &DOMTreeNode{
		NodeType:  1,
		NodeName:  strings.ToUpper(tagName),
		Id:        props.Id,
		ClassList: strings.Split(props.ClassName, " "),
		ClassName: props.ClassName,
		TagName:   tagName,
		Title:     props.Title,
	}
}

func NewTextNode(text string) *DOMTreeNode {
	return &DOMTreeNode{
		NodeType:  3,
		NodeName:  "#text",
		NodeValue: text,
	}
}

type DOMTreeNode struct {
	ChildElementCount      int            // element
	ChildNodes             []*DOMTreeNode // node
	Children               []*DOMTreeNode // element
	ClassList              []string       // element
	ClassName              string         // element
	FirstChild             *DOMTreeNode   // node
	FirstElementChild      *DOMTreeNode   // element
	Id                     string         // element
	LastChild              *DOMTreeNode   // node
	LastElementChild       *DOMTreeNode   // element
	NextElementSibling     *DOMTreeNode   // node | element
	NextSibling            *DOMTreeNode   // node | element
	NodeName               string         // node
	NodeType               int            // node | element
	NodeValue              string         // node | element
	ParentElement          *DOMTreeNode   // node | element
	ParentNode             *DOMTreeNode   // node | element
	PreviousElementSibling *DOMTreeNode   // node | element
	PreviousSibling        *DOMTreeNode   // node | element
	TagName                string         // element
	Title                  string         // element
}

// Insert newNodes after the reference node
func (eln *DOMTreeNode) After(newNodes ...*DOMTreeNode) {

}

// Append newNodes (children) to the reference node
func (eln *DOMTreeNode) Append(newNodes ...*DOMTreeNode) {

}

// Append newChild to the reference node
func (eln *DOMTreeNode) AppendChild(newChild *DOMTreeNode) {

}

// Insert newNodes before the reference node
func (eln *DOMTreeNode) Before(newNodes ...*DOMTreeNode) {

}

// Does reference node contain this node?
func (eln *DOMTreeNode) Contains(node *DOMTreeNode) bool {
	return false
}

// Remove nodeToRemove from reference node's children
func (eln *DOMTreeNode) RemoveChild(nodeToRemove *DOMTreeNode) {

}

// Replace reference node's targetNode child with replacementChild
func (eln *DOMTreeNode) ReplaceChild(targetNode *DOMTreeNode, replacementChild *DOMTreeNode) {

}

// Set reference node's children to the provided replacementChildren
func (eln *DOMTreeNode) ReplaceChildren(replacementChildren ...*DOMTreeNode) {

}

// Among reference node's children insert newChild before targetNode
func (eln *DOMTreeNode) InsertBefore(targetNode *DOMTreeNode, newChild *DOMTreeNode) {

}

// Insert the provided element at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeNode) InsertAdjacentElement(where string, element *DOMTreeNode) {

}

// Insert the provided html (after parsing to DOMTreeNode) at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeNode) InsertAdjacentHTML(where string, html string) {

}

// Insert the provided text (as DOMTreeNode) at the position specified around the reference node
//
// where (position): "afterend" | "beforeend" | "afterbegin" | "beforebegin"
func (eln *DOMTreeNode) InsertAdjacentText(where string, text string) {

}

// Is the childNodes property greater than zero?
func (eln *DOMTreeNode) HasChildNodes() bool {
	return len(eln.ChildNodes) > 0
}
