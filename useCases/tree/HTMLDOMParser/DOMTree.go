package HTMLDOMParser

func NewDocument(title string) *DOMTree {
	root := &DOMTreeElement{TagName: "HTML", OuterHTML: "<html></html>"}

	// insert head
	// insert title in head
	// insert body

	return &DOMTree{Root: root}
}

type DOMTree struct {
	Root *DOMTreeElement
}
