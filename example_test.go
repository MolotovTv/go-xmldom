// nolint:govet
package xmldom_test

import (
	"bytes"
	"fmt"

	"github.com/molotovtv/go-xmldom"
)

const (
	ExampleXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<testsuites>
	<testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
		<properties>
			<property name="go.version">go1.8.1</property>
		</properties>
		<testcase classname="go-xmldom" id="ExampleParseXML" time="0.004"></testcase>
		<testcase classname="go-xmldom" id="ExampleParse" time="0.005"></testcase>
	</testsuite>
</testsuites>`

	ExampleNamespaceXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<ns0:Content xmlns:ns0="namespace_0" xmlns:ns1="namespace_1">
			<ns1:Item>item1</ns1:Item>
			<ns1:Item>item2</ns1:Item>
		</ns0:Content>
		<ns2:Other xmlns:ns2="namespace_2" param="test_param_value"/>
	</S:Body>
</S:Envelope>`

	ExampleInheritNamespaceXML = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
			<ds:SignedInfo>
				<ds:DigestValue>KHDHFKSFH2IEFIDHJKSHFKJHSKDJFH==
				</ds:DigestValue>
			</ds:SignedInfo>
		</ds:Signature>
	</S:Body>
</S:Envelope>`

	ExampleTTML = `<?xml version="1.0" encoding="UTF-8"?>
<tt xmlns="http://www.w3.org/ns/ttml" xmlns:ttm="http://www.w3.org/ns/ttml#metadata" xmlns:ttp="http://www.w3.org/ns/ttml#parameter" xmlns:tts="http://www.w3.org/ns/ttml#styling" xmlns:ebuttm="urn:ebu::tt::metadata" xmlns:ebutts="urn:ebu::tt::style" xml:lang="fra" ttp:cellResolution="40 24" ttp:timeBase="media">
  <head>
	<metadata>
	  <ttm:title />
	  <ttm:desc />
	  <ttm:copyright />
	</metadata>
	<styling>
	  <style xml:id="Style_367" backgroundColor="#101010fe" tts:color="#ebeb0afe" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" />
	  <style xml:id="Style_368" backgroundColor="#00000000" tts:color="#00000000" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" />
	  <style xml:id="Style_369" backgroundColor="#101010fe" tts:color="#ebeb0afe" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" />
	</styling>
	<layout>
	  <region xml:id="Region_367" tts:extent="82% 8%" tts:origin="9% 79%" />
	  <region xml:id="Region_368" tts:extent="0% 0%" tts:origin="50% 0%" />
	  <region xml:id="Region_369" tts:extent="82% 8%" tts:origin="9% 79%" />
	</layout>
  </head>
  <body>
	<div>
	  <p begin="00:00:00.000" end="00:00:01.560" region="Region_367">
		<span style="Style_367">
		  dormant avec son chat,
		  <br />
		  Emma de Saint-Ismier dans l&apos;Isère
		</span>
	  </p>
	  <p begin="00:00:01.560" end="00:00:01.920" region="Region_368">
		<span style="Style_368" />
	  </p>
	  <p begin="00:00:01.920" end="00:00:03.840" region="Region_369">
		<span style="Style_369">
		  se détend en regardant
		  <br />
		  le top des comédies française.
		</span>
	  </p>
	  <p begin="00:00:03.840" end="00:00:05.320" region="Region_369">
		<span style="Style_369">
		  se détend en regardant
		  <br />
		  le top des comédies française.
		</span>
	  </p>
	  <p begin="00:00:05.320" end="00:00:05.600" region="Region_368">
		<span style="Style_368" />
	  </p>
	  <p begin="00:00:05.600" end="00:00:05.760" region="Region_369">
		<span style="Style_369">-On est en rose, les 2.</span>
	  </p>
	  <p begin="00:00:05.760" end="00:00:07.240" region="Region_368">
		<span style="Style_368">-On est en rose, les 2.</span>
	  </p>
	  <p begin="00:00:07.240" end="00:00:07.520" region="Region_367">
		<span style="Style_367" />
	  </p>
	</div>
  </body>
</tt>`
)

func ExampleParseXML() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	fmt.Printf("name = %v\n", node.Name.Local)
	fmt.Printf("attributes.len = %v\n", len(node.Attributes))
	fmt.Printf("children.len = %v\n", len(node.Children))
	fmt.Printf("root = %v", node == node.Root())
	// Output:
	// name = testsuites
	// attributes.len = 0
	// children.len = 1
	// root = true
}

func ExampleParseNamespaceszXML() {
	node := xmldom.Must(xmldom.ParseXML(ExampleNamespaceXml)).Root
	fmt.Printf("name.Local = %v\n", node.Name.Local)
	fmt.Printf("name.Space = %v\n", node.Name.Space)
	fmt.Printf("attributes.len = %v\n", len(node.Attributes))
	fmt.Printf("children.len = %v\n", len(node.Children))
	fmt.Printf("root = %v", node == node.Root())
	// Output:
	// name.Local = Envelope
	// name.Space = http://schemas.xmlsoap.org/soap/envelope/
	// attributes.len = 1
	// children.len = 1
	// root = true
}

func ExampleEmptyElementTag() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleNamespaceXml))
	doc.EmptyElementTag = true
	fmt.Println(doc.Root.XML())
	// Output:
	// <S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/"><S:Body><ns0:Content xmlns:ns0="namespace_0" xmlns:ns1="namespace_1"><ns1:Item>item1</ns1:Item><ns1:Item>item2</ns1:Item></ns0:Content><ns2:Other xmlns:ns2="namespace_2" param="test_param_value" /></S:Body></S:Envelope>
}

func ExampleNode_GetAttribute() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	attr := node.FirstChild().GetAttribute("name")
	fmt.Printf("%v = %v\n", attr.Name.Local, attr.Value)
	// Output:
	// name = github.com/subchen/go-xmldom
}

func ExampleNode_GetChildren() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	children := node.FirstChild().GetChildren("testcase")
	for _, c := range children {
		fmt.Printf("%v: id = %v\n", c.Name.Local, c.GetAttributeValue("id"))
	}
	// Output:
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
}

func ExampleNode_FindByID() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindByID("ExampleParseXML")
	fmt.Println(node.XML())
	// Output:
	// <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
}

func ExampleNode_FindOneByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindOneByName("property")
	fmt.Println(node.XML())
	// Output:
	// <property name="go.version">go1.8.1</property>
}

func ExampleNode_FindByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	nodes := root.FindByName("testcase")
	for _, node := range nodes {
		fmt.Println(node.XML())
	}
	// Output:
	// <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
	// <testcase classname="go-xmldom" id="ExampleParse" time="0.005" />
}

func ExampleNode_Query() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find all children
	fmt.Printf("children = %v\n", len(node.Query("//*")))

	// find node matched tag name
	nodeList := node.Query("//testcase")
	for _, c := range nodeList {
		fmt.Printf("%v: id = %v\n", c.Name.Local, c.GetAttributeValue("id"))
	}
	// Output:
	// children = 6
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
}

func ExampleNode_QueryOne() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find node matched attr name
	c := node.QueryOne("//testcase[@id='ExampleParseXML']")
	fmt.Printf("%v: id = %v\n", c.Name.Local, c.GetAttributeValue("id"))
	// Output:
	// testcase: id = ExampleParseXML
}

func ExampleDocument_XML() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XML())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><!DOCTYPE junit SYSTEM "junit-result.dtd"><testsuites><testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom"><properties><property name="go.version">go1.8.1</property></properties><testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" /><testcase classname="go-xmldom" id="ExampleParse" time="0.005" /></testsuite></testsuites>
}

func ExampleDocument_XMLPretty() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <!DOCTYPE junit SYSTEM "junit-result.dtd">
	// <testsuites>
	//   <testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
	//     <properties>
	//       <property name="go.version">go1.8.1</property>
	//     </properties>
	//     <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
	//     <testcase classname="go-xmldom" id="ExampleParse" time="0.005" />
	//   </testsuite>
	// </testsuites>
}

func ExampleNewDocument() {
	doc := xmldom.NewDocument("testsuites")

	testsuiteNode := doc.Root.CreateNode("testsuite").SetAttributeValue("name", "github.com/subchen/go-xmldom")
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 1").CreateTextNode("PASS")
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 2").CreateTextNode("FAIL")

	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <testsuites>
	//   <testsuite name="github.com/subchen/go-xmldom">
	//     <testcase name="case 1">PASS</testcase>
	//     <testcase name="case 2">FAIL</testcase>
	//   </testsuite>
	// </testsuites>
}

func ExampleInheritNamespace() {
	doc := xmldom.NewDocument("")
	doc.EmptyElementTag = false
	doc.TextSafeMode = false
	err := doc.Parse(
		bytes.NewReader(
			[]byte(ExampleInheritNamespaceXML),
		),
	)
	fmt.Println(err)
	fmt.Println(doc.Root.XML())

	node := doc.Root.QueryOne("//Body/Signature/SignedInfo")
	fmt.Println(node.XML())
	// Output:
	// <nil>
	// <S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/"><S:Body><ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:SignedInfo><ds:DigestValue>KHDHFKSFH2IEFIDHJKSHFKJHSKDJFH==
	// 				</ds:DigestValue></ds:SignedInfo></ds:Signature></S:Body></S:Envelope>
	// <ds:SignedInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:DigestValue>KHDHFKSFH2IEFIDHJKSHFKJHSKDJFH==
	// 				</ds:DigestValue></ds:SignedInfo>
}

func ExampleParseTTML() {
	doc, err := xmldom.Parse(bytes.NewReader([]byte(ExampleTTML)))
	fmt.Println(err)
	fmt.Println(doc.XML())

	// Output:
	// <nil>
	// <?xml version="1.0" encoding="UTF-8"?><tt xmlns="http://www.w3.org/ns/ttml" xmlns:ttm="http://www.w3.org/ns/ttml#metadata" xmlns:ttp="http://www.w3.org/ns/ttml#parameter" xmlns:tts="http://www.w3.org/ns/ttml#styling" xmlns:ebuttm="urn:ebu::tt::metadata" xmlns:ebutts="urn:ebu::tt::style" xml:lang="fra" ttp:cellResolution="40 24" ttp:timeBase="media"><head><metadata><ttm:title /><ttm:desc /><ttm:copyright /></metadata><styling><style xml:id="Style_367" backgroundColor="#101010fe" tts:color="#ebeb0afe" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" /><style xml:id="Style_368" backgroundColor="#00000000" tts:color="#00000000" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" /><style xml:id="Style_369" backgroundColor="#101010fe" tts:color="#ebeb0afe" tts:fontSize="80%" tts:lineHeight="100%" tts:textAlign="left" ebutts:linePadding="0c" ebutts:multiRowAlign="auto" /></styling><layout><region xml:id="Region_367" tts:extent="82% 8%" tts:origin="9% 79%" /><region xml:id="Region_368" tts:extent="0% 0%" tts:origin="50% 0%" /><region xml:id="Region_369" tts:extent="82% 8%" tts:origin="9% 79%" /></layout></head><body><div><p begin="00:00:00.000" end="00:00:01.560" region="Region_367"><span style="Style_367">dormant avec son chat,<br />Emma de Saint-Ismier dans l&#39;Isère</span></p><p begin="00:00:01.560" end="00:00:01.920" region="Region_368"><span style="Style_368" /></p><p begin="00:00:01.920" end="00:00:03.840" region="Region_369"><span style="Style_369">se détend en regardant<br />le top des comédies française.</span></p><p begin="00:00:03.840" end="00:00:05.320" region="Region_369"><span style="Style_369">se détend en regardant<br />le top des comédies française.</span></p><p begin="00:00:05.320" end="00:00:05.600" region="Region_368"><span style="Style_368" /></p><p begin="00:00:05.600" end="00:00:05.760" region="Region_369"><span style="Style_369">-On est en rose, les 2.</span></p><p begin="00:00:05.760" end="00:00:07.240" region="Region_368"><span style="Style_368">-On est en rose, les 2.</span></p><p begin="00:00:07.240" end="00:00:07.520" region="Region_367"><span style="Style_367" /></p></div></body></tt>
}
