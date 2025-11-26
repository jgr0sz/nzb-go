package src

/*
	Custom structs used to reflect the composition of NZBs.
	Using tags, we map these fields to XML regions (encoding/xml is a godsend).
*/

//The main NZB type. Meta[] stores data under the <meta> tag, and Files[] stores a slice of <file> tag data.
type Nzb struct {
	Meta []Meta `xml:"head>meta"`
	Files []File `xml:"file"`
}

//Data associated with the <file> tag.
type File struct {
	Poster string `xml:"poster,attr"`
	Date int64 `xml:"date,attr"`
	Subject string `xml:"subject,attr"`
	Groups []string `xml:"groups>group"`
	Segments []Segment `xml:"segments>segment"`
}


//Data associated with the <segment> tag. Stored in a slice within a File.
type Segment struct {
	Bytes int `xml:"bytes,attr"`
	Number int `xml:"number,attr"`
	ID string `xml:",chardata"`
}

//Data associated with the <meta> tag within <head>. There are many variants of meta type="..." handled separately.
type Meta struct {
	Type string `xml:"type,attr"`
	Value string `xml:",chardata"`
}
