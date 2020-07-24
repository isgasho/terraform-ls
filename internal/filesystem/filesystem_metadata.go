package filesystem

import (
	"github.com/hashicorp/hcl/v2"
)

func (fs *fsystem) markDocumentAsOpen(dh DocumentHandler) error {
	if !fs.documentMetadataExists(dh) {
		return &MetadataNotExistsErr{dh}
	}

	fs.docMeta[dh.URI()].setOpen(true)
	return nil
}

func (fs *fsystem) createDocumentMetadata(dh DocumentHandler, text []byte) error {
	if fs.documentMetadataExists(dh) {
		return &MetadataAlreadyExistsErr{dh}
	}

	fs.docMeta[dh.URI()] = NewDocumentMetadata(dh, text)
	return nil
}

func (fs *fsystem) documentMetadataExists(dh DocumentHandler) (bool) {
	_, ok := fs.docMeta[dh.URI()]
	return ok
}

func (fs *fsystem) updateDocumentMetadata(fh VersionedDocumentHandler, changes DocumentChanges) error {
	// TODO
	return nil
}

func (fs *fsystem) getDocumentMetadata(dh DocumentHandler) (*documentMetadata, error) {
	dm, ok := fs.docMeta[dh.URI()]
	if !ok {
		return nil, &MetadataNotExistsErr{dh}
	}

	return dm, nil
}

func (fs *fsystem) applyDocumentChange(change DocumentChange) error {

	// if the range is regarded as nil, we regard it as full content change
	if rangeIsNil(change.Range()) {
		d.change([]byte(change.Text()))
		return nil
	}
	// b := &bytes.Buffer{}
	// b.Grow(len(d.content) + diffLen(change))
	// b.Write(d.content[:change.Range().Start.Byte])
	// b.WriteString(change.Text())
	// b.Write(d.content[change.Range().End.Byte:])

	// f.change(b.Bytes())
	return nil
}

// HCL column and line indexes start from 1, therefore if the any index
// contains 0, we assume it is an undefined range
func rangeIsNil(r hcl.Range) bool {
	return r.End.Column == 0 && r.End.Line == 0
}

func diffLen(change DocumentChange) int {
	rangeLen := change.Range().End.Byte - change.Range().Start.Byte
	return len(change.Text()) - rangeLen
}