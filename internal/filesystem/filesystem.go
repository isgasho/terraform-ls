package filesystem

import (	
	"io/ioutil"
	"log"
	"sync"

	"github.com/spf13/afero"
)

type fsystem struct {
	afs afero.Fs

	docMeta   map[string]*documentMetadata
	docMetaMu sync.RWMutex

	logger *log.Logger
}

func NewFilesystem() *fsystem {
	return &fsystem{
		afs:    afero.NewMemMapFs(),
		logger: log.New(ioutil.Discard, "", 0),
	}
}

func (fs *fsystem) SetLogger(logger *log.Logger) {
	fs.logger = logger
}

func (fs *fsystem) CreateDocument(dh DocumentHandler, text []byte) error {
	f, err := fs.afs.Create(dh.FullPath())
	if err != nil {
		return err
	}
	_, err = f.Write(text)
	if err != nil {
		return err
	}

	return fs.createDocumentMetadata(dh, text)
}

func (fs *fsystem) CreateAndOpenDocument(dh DocumentHandler, text []byte) error {
	err := fs.CreateDocument(dh, text)
	if err != nil {
		return err
	}

	return fs.markDocumentAsOpen(dh)
}

func (fs *fsystem) ChangeDocument(fh VersionedDocumentHandler, changes DocumentChanges) error {
	// TODO
	return nil
}

func (fs *fsystem) CloseDocument(fh DocumentHandler) error {
	// TODO
	return nil
}

func (fs *fsystem) GetDocument(dh DocumentHandler) (Document, error) {
	dm, err := fs.getDocumentMetadata(dh)
	if err != nil {
		return nil, err
	}

	return &document{
		meta: dm,
		fo: fs.afs,
	}, nil
}
