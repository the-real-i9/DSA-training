package HTMLDOMParser

type DOMTreeElement struct {
	ChildElementCount      int
	Children               []*DOMTreeElement
	ClassList              []string
	ClassName              string
	FirstElementChild      *DOMTreeElement
	Id                     string
	InnerHTML              string
	InnerText              string
	LastElementChild       *DOMTreeElement
	NextElementSibling     *DOMTreeElement
	OuterHTML              string
	OuterText              string
	ParentElement          *DOMTreeElement
	PreviousElementSibling *DOMTreeElement
	TagName                string
	TextContent            string
	Title                  string
}
