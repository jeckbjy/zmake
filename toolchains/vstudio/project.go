package vstudio

import "encoding/xml"

// https://docs.microsoft.com/en-us/cpp/build/reference/project-files?view=vs-2019
type vsProject struct {
	XMLName        xml.Name          `xml:"Project"`
	DefaultTargets string            `xml:"DefaultTargets, attr"`
	ToolsVersion   string            `xml:"ToolsVersion,attr"`
	Xmlns          string            `xml:"xmlns,attr"` // xml namespace
	ItemGroups     []vsItemGroup     `xml:"ItemGroup"`
	Globals        []vsPropertyGroup `xml:"PropertyGroup"`
	Import         vsImport          `xml:"Import"`
	Configuration  []vsPropertyGroup `xml:"PropertyGroup"`
}

type vsItemGroup struct {
	XMLName        xml.Name                 `xml:"ItemGroup"`
	Label          string                   `xml:"Label, attr"`
	Configurations []vsProjectConfiguration `xml:"ProjectConfiguration"`
}

type vsProjectConfiguration struct {
	XMLName       xml.Name `xml:"ProjectConfiguration"`
	Include       string   `xml:"Include,attr"`
	Configuration string   `xml:"Configuration"`
	Platform      string   `xml:"Platform"`
}

type vsPropertyGroup struct {
	XMLName       xml.Name `xml:"PropertyGroup"`
	Label         string   `xml:"Label, attr"`
	Condition     string   `xml:"Condition"`
	ProjectGuid   string   `xml:"ProjectGuid"`
	Keyword       string   `xml:"Keyword"`
	RootNamespace string   `xml:"RootNamespace"`
}

type vsImport struct {
	XMLName xml.Name `xml:"Import"`
	Project string   `xml:"Project,attr"`
}
