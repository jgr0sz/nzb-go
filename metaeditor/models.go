package metaeditor

import "github.com/jgr0sz/nzbgo/parser"

//Main object struct for the meta editor. Contains a Meta splice to hold attributes and their content and a pointer to the original Nzb.
type NzbMetaEditor struct {
	Metadata []parser.Meta
	Nzb *parser.Nzb
}
