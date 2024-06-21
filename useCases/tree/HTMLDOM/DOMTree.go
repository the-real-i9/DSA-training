package HTMLDOM

// Note: All block level elements have first and last text nodes of "\n"

func NewDocument(title string) *DOMTree {
	htmlElem := &DOMTreeElementNode{
		TagName:  "HTML",
		NodeName: "HTML",
		NodeType: 1,
	}

	headElem := &DOMTreeElementNode{
		TagName:  "HEAD",
		NodeName: "HEAD",
		NodeType: 1,
	}

	bodyElem := &DOMTreeElementNode{
		TagName:  "BODY",
		NodeName: "BODY",
		NodeType: 1,
	}

	titleElem := &DOMTreeElementNode{
		TagName:  "TITLE",
		NodeName: "TITLE",
		NodeType: 1,
	}

	titleText := &DOMTreeElementNode{
		NodeName:  "#text",
		NodeValue: title,
		Data:      title,
		NodeType:  3,
	}

	newLineText := &DOMTreeElementNode{
		NodeName:  "#text",
		NodeValue: "\n",
		Data:      "\n",
		NodeType:  3,
	}

	// insert head
	// insert title in head
	// insert body

	return &DOMTree{Root: htmlElem}
}

type DOMTree struct {
	Root *DOMTreeElementNode
}
