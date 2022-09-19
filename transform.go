package xmldom

type TransformFunc func([]byte) ([]byte, error)

func (n *Node) Transform(f TransformFunc) error {
	data, err := f([]byte(n.XML()))
	if err != nil {
		return err
	}
	return n.ParseXML(string(data))
}

func (n *Node) ParseXML(data string) error {
	doc := *n.Document
	if err := doc.ParseXML(data); err != nil {
		return err
	}

	doc.Root.ChangeDocumentTo(n.Document, n.Parent)
	*n = *doc.Root

	return nil
}

func (n *Node) ChangeDocumentTo(doc *Document, parent *Node) {
	n.Document = doc
	n.Parent = parent
	for _, child := range n.Children {
		child.ChangeDocumentTo(doc, n)
	}
}
