package xcode

import "encoding/xml"

type xcFileRef struct {
	XMLName  xml.Name `xml:"FileRef"`
	Location string   `xml:"location,attr"` // prefix with [group:|self:]
}

type xcGroup struct {
	XMLName  xml.Name     `xml:"Group"`
	Location string       `xml:"location,attr"`
	Name     string       `xml:"name,attr"`
	Groups   []*xcGroup   `xml:"Group"`
	FileRefs []*xcFileRef `xml:"FileRef"`
}

type xcWorkspace struct {
	XMLName  xml.Name     `xml:"Workspace"`
	Version  string       `xml:"version,attr"`
	Groups   []*xcGroup   `xml:"Group"`
	FileRefs []*xcFileRef `xml:"FileRef"`
}
