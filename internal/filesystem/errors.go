package filesystem

import (
	"fmt"
)

type DocumentNotOpenErr struct {
	DocumentHandler DocumentHandler
}

func (e *DocumentNotOpenErr) Error() string {
	return fmt.Sprintf("document is not open: %s", e.DocumentHandler.URI())
}

type MetadataAlreadyExistsErr struct {
	DocumentHandler DocumentHandler
}

func (e *MetadataAlreadyExistsErr) Error() string {
	return fmt.Sprintf("document metadata already exists: %s", e.DocumentHandler.URI())
}

type MetadataNotExistsErr struct {
	DocumentHandler DocumentHandler
}

func (e *MetadataNotExistsErr) Error() string {
	return fmt.Sprintf("document metadata does not exist: %s", e.DocumentHandler.URI())
}
