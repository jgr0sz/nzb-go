package src

/*
	Custom structs used to reflect the composition of NZBs.
	Using tags, we map these fields to XML regions (encoding/xml is a godsend).
*/

//The main NZB type. Meta[] stores data under the <meta> tag, and Files[] stores a slice of <file> tag data.
type Nzb struct {
	Head  Head `xml:"head" json:"head"`
	Files []File `xml:"file" json:"files"`
}

//Data associated with the <file> tag.
type File struct {
	Poster string `xml:"poster,attr" json:"poster"`
	Date int64 `xml:"date,attr" json:"date"`
	Subject string `xml:"subject,attr" json:"subject"`
	Groups []string `xml:"groups>group" json:"groups"`
	Segments []Segment `xml:"segments>segment" json:"segments"`
}

//Data associated with the <head> tag.
type Head struct {
	Meta []Meta `xml:"meta" json:"meta"`
}

//Data associated with the <segment> tag. Stored in a slice within a File.
type Segment struct {
	Bytes int `xml:"bytes,attr" json:"bytes"`
	Number int `xml:"number,attr" json:"number"`
	ID string `xml:",chardata" json:"id"`
}

//Data associated with the <meta> tag within <head>. There are many variants of meta type="..." handled separately.
type Meta struct {
	Type string `xml:"type,attr" json:"type"`
	Value string `xml:",chardata" json:"value"`
}
