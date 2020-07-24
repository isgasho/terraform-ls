package filesystem

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-ls/internal/source"
)

type Document interface {
	DocumentHandler
	Text() ([]byte, error)
	Lines() source.Lines
	Version() int
}

type DocumentHandler interface {
	URI() string
	FullPath() string
	Dir() string
	Filename() string
}

type VersionedDocumentHandler interface {
	DocumentHandler
	Version() int
}

type DocumentChange interface {
	Text() string
	Range() hcl.Range
}

type DocumentChanges []DocumentChange

type Filesystem interface {
	// LS-specific methods
	CreateDocument(DocumentHandler, []byte) error
	CreateAndOpenDocument(DocumentHandler, []byte) error
	GetDocument(DocumentHandler) (Document, error)
	CloseDocument(DocumentHandler) error
	ChangeDocument(VersionedDocumentHandler, DocumentChanges) error

	// TODO: standard afero.Fs methods
	// Create(name string) (File, error)
	// Mkdir(name string, perm os.FileMode) error
	// MkdirAll(path string, perm os.FileMode) error
	// Open(name string) (File, error)
	// OpenFile(name string, flag int, perm os.FileMode) (File, error)
	// Remove(name string) error
	// RemoveAll(path string) error
	// Rename(oldname, newname string) error
	// Stat(name string) (os.FileInfo, error)
	// Name() string
	// Chmod(name string, mode os.FileMode) error
	// Chtimes(name string, atime time.Time, mtime time.Time) error
}
