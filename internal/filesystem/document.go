package filesystem

import (
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/terraform-ls/internal/source"
	"github.com/spf13/afero"
)



type fileOpener interface {
	Open(name string) (afero.File, error)
}

type document struct {
	meta *documentMetadata
	fo fileOpener
}

func (d *document) Text() ([]byte, error) {
	f, err := d.fo.Open(d.meta.dh.FullPath())
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func (d *document) FullPath() string {
	return d.meta.dh.FullPath()
}

func (d *document) Dir() string {
	return filepath.Dir(d.meta.dh.FullPath())
}

func (d *document) Filename() string {
	return filepath.Base(d.meta.dh.FullPath())
}

func (d *document) URI() string {
	return URIFromPath(d.meta.dh.FullPath())
}

func (d *document) Lines() source.Lines {
	return d.meta.Lines()
}

func (d *document) Version() int {
	return d.meta.version
}

type documentMetadata struct {
	dh DocumentHandler

	isOpen     bool
	version int
	lines      source.Lines
}

func NewDocumentMetadata(dh DocumentHandler, content []byte) *documentMetadata {
	return &documentMetadata{
		dh: dh,
		lines: source.MakeSourceLines(dh.Filename(), content),
	}
}

func (d *documentMetadata) setOpen(isOpen bool) {
	d.isOpen = isOpen
}

func (d *documentMetadata) setVersion(version int) {
	d.version = version
}

func (d *documentMetadata) updateLines(content []byte) {
	d.lines = source.MakeSourceLines(d.dh.Filename(), content)
}

func (d *documentMetadata) Lines() source.Lines {
	return d.lines
}
