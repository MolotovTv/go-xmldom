// XML DOM processing for Golang, supports xpath query
package xmldom

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func Must(doc *Document, err error) *Document {
	if err != nil {
		panic(err)
	}
	return doc
}

func ParseObject(v interface{}) (*Document, error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return Parse(bytes.NewReader(data))
}

func ParseXML(s string) (*Document, error) {
	return Parse(strings.NewReader(s))
}

func ParseFile(filename string) (*Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Parse(file)
}

func Parse(r io.Reader) (*Document, error) {
	doc := NewDocument("empty")
	err := doc.Parse(r)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (doc *Document) ParseObject(v interface{}) error {
	data, err := xml.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal object error: %v", err)
	}
	return doc.Parse(bytes.NewReader(data))
}

func (doc *Document) ParseXML(s string) error {
	return doc.Parse(strings.NewReader(s))
}

func (doc *Document) Parse(r io.Reader) error {
	p := xml.NewDecoder(r)
	t, err := p.Token()
	if err != nil {
		return err
	}
	doc.Root = nil

	var e *Node
	for t != nil {
		switch token := t.(type) {
		case xml.StartElement:
			// a new node
			el := new(Node)
			el.Document = doc
			el.Parent = e
			el.Name = token.Name

			for _, attr := range token.Attr {
				attribute := attr
				el.Attributes = append(el.Attributes, &attribute)
			}
			if e != nil {
				e.Children = append(e.Children, el)
			}
			e = el

			if doc.Root == nil {
				doc.Root = e
			}
		case xml.EndElement:
			e = e.Parent
		case xml.CharData:
			// text node
			if e != nil {
				// a new node
				el := new(Node)
				el.Document = doc
				el.Parent = e

				if strings.TrimSpace(string(token)) != "" {
					if doc.TextSafeMode {
						el.Text = string(bytes.TrimSpace(token))
					} else {
						el.Text = string(token)
					}
				}
				if el.Text != "" {
					e.Children = append(e.Children, el)
				}
			}
		case xml.ProcInst:
			doc.ProcInst = stringifyProcInst(&token)
		case xml.Directive:
			doc.Directives = append(doc.Directives, stringifyDirective(&token))
		}

		// get the next token
		t, err = p.Token()
	}

	// Make sure that reading stopped on EOF
	if err != io.EOF {
		return err
	}

	// All is good, return the document
	return nil
}
