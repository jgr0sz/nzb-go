package metaeditor

import "jgr0sz/nzbgo/parser"

//Main object struct for the meta editor. Contains a Meta splice to hold attributes and their content.
type NzbMetaEditor struct {
	Metadata []parser.Meta
	Nzb *parser.Nzb
}
