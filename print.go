package xmldom

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func stringifyProcInst(pi *xml.ProcInst) string {
	if pi == nil {
		return ""
	}
	return fmt.Sprintf("<?%s %s?>", pi.Target, string(pi.Inst))
}

func stringifyDirective(directive *xml.Directive) string {
	if directive == nil {
		return ""
	}
	return fmt.Sprintf("<!%s>", string(*directive))
}

type printer struct{}

func (p *printer) printXML(buf *bytes.Buffer, n *Node, level int, indent string) {
	pretty := len(indent) > 0

	if pretty {
		buf.WriteString(strings.Repeat(indent, level))
	}

	if n.IsTextNode() {
		if n.Document.TextSafeMode {
			xml.EscapeText(buf, []byte(n.Text)) // nolint:errcheck
		} else {
			buf.WriteString(n.Text)
		}
		return
	}

	space := n.GetNamespace(n.Name.Space)

	buf.WriteByte('<')
	if space != nil {
		buf.WriteString(space.Name.Local)
		buf.WriteByte(':')
	}
	buf.WriteString(n.Name.Local)

	if level == 0 && n.Parent != nil && space != nil {
		buf.WriteByte(' ')
		buf.WriteString(space.Name.Space)
		buf.WriteByte(':')
		buf.WriteString(space.Name.Local)
		buf.WriteByte('=')
		buf.WriteByte('"')
		xml.Escape(buf, []byte(space.Value))
		buf.WriteByte('"')
	}
	for _, attr := range n.Attributes {
		buf.WriteByte(' ')
		if attr.Name.Space == xmlnsPrefix {
			buf.WriteString(attr.Name.Space)
			buf.WriteByte(':')
		} else if space := n.GetNamespace(attr.Name.Space); space != nil {
			buf.WriteString(space.Name.Local)
			buf.WriteByte(':')
		}
		buf.WriteString(attr.Name.Local)
		buf.WriteByte('=')
		buf.WriteByte('"')
		xmlEscape(buf, []byte(attr.Value))
		buf.WriteByte('"')
	}
	if n.Document.EmptyElementTag {
		if len(n.Children) == 0 {
			buf.WriteString(" />")
			return
		}
	}

	buf.WriteByte('>')

	for _, c := range n.Children {
		if c.IsTextNode() {
			p.printXML(buf, c, level+1, "")
		} else {
			if pretty {
				buf.WriteByte('\n')
			}
			p.printXML(buf, c, level+1, indent)
		}
	}
	if pretty && len(n.Children) > 0 && !(len(n.Children) == 1 && n.Children[0].IsTextNode()) {
		buf.WriteByte('\n')
		buf.WriteString(strings.Repeat(indent, level))
	}
	buf.WriteString("</")
	if space != nil {
		buf.WriteString(space.Name.Local)
		buf.WriteByte(':')
	}
	buf.WriteString(n.Name.Local)
	buf.WriteByte('>')
}

func xmlEscape(w io.Writer, value []byte) {
	var res bytes.Buffer
	xml.Escape(&res, value)

	out := res.String()
	newOut := strings.Replace(out, "&#34;", "&quot;", -1)

	fmt.Fprintf(w, "%s", newOut)
}
