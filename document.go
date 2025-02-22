package xmldom

import (
	"bytes"
	"encoding/xml"
)

const (
	DEFAULT_XML_HEADER = `<?xml version="1.0" encoding="UTF-8"?>`
	xmlURL             = "http://www.w3.org/XML/1998/namespace"
	xmlnsPrefix        = "xmlns"
	xmlPrefix          = "xml"
)

func NewDocument(name string) *Document {
	d := &Document{
		ProcInst:        DEFAULT_XML_HEADER,
		EmptyElementTag: true,
		TextSafeMode:    true,
	}
	d.Root = &Node{
		Document: d,
		Name: xml.Name{
			Local: name,
		},
	}
	return d
}

type Document struct {
	ProcInst        string
	Directives      []string
	EmptyElementTag bool
	TextSafeMode    bool
	Root            *Node
}

func (d *Document) XML() string {
	buf := new(bytes.Buffer)
	buf.WriteString(d.ProcInst)
	for _, directive := range d.Directives {
		buf.WriteString(directive)
	}
	buf.WriteString(d.Root.XML())
	return buf.String()
}

func (d *Document) XMLPretty() string {
	buf := new(bytes.Buffer)
	if len(d.ProcInst) > 0 {
		buf.WriteString(d.ProcInst)
		buf.WriteByte('\n')
	}
	for _, directive := range d.Directives {
		buf.WriteString(directive)
		buf.WriteByte('\n')
	}
	buf.WriteString(d.Root.XMLPretty())
	buf.WriteByte('\n')
	return buf.String()
}

func (d *Document) XMLPrettyEx(indent string) string {
	buf := new(bytes.Buffer)
	if len(d.ProcInst) > 0 {
		buf.WriteString(d.ProcInst)
		buf.WriteByte('\n')
	}
	for _, directive := range d.Directives {
		buf.WriteString(directive)
		buf.WriteByte('\n')
	}
	buf.WriteString(d.Root.XMLPrettyEx(indent))
	buf.WriteByte('\n')
	return buf.String()
}
