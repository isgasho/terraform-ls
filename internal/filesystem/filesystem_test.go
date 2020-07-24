package filesystem

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-ls/internal/source"
)

func TestFilesystem_Change_notOpen(t *testing.T) {
	fs := NewFilesystem()

	var changes DocumentChanges
	changes = append(changes, &testChange{})
	h := &testHandler{"file:///doesnotexist"}

	err := fs.ChangeDocument(h, changes)

	expectedErr := &DocumentNotOpenErr{h}
	if err == nil {
		t.Fatalf("Expected error: %s", expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error.\nexpected: %#v\ngiven: %#v",
			expectedErr, err)
	}
}

func TestFilesystem_Change_closed(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///doesnotexist"}
	fs.CreateAndOpenDocument(&testDocument{
		testHandler: fh,
	}, []byte{})
	err := fs.CloseDocument(fh)
	if err != nil {
		t.Fatal(err)
	}

	var changes DocumentChanges
	changes = append(changes, &testChange{})
	err = fs.ChangeDocument(fh, changes)

	expectedErr := &DocumentNotOpenErr{fh}
	if err == nil {
		t.Fatalf("Expected error: %s", expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error.\nexpected: %#v\ngiven: %#v",
			expectedErr, err)
	}
}

func TestFilesystem_Close_closed(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///doesnotexist"}
	fs.CreateAndOpenDocument(&testDocument{
		testHandler: fh,
	}, []byte{})
	err := fs.CloseDocument(fh)
	if err != nil {
		t.Fatal(err)
	}

	err = fs.CloseDocument(fh)

	expectedErr := &DocumentNotOpenErr{fh}
	if err == nil {
		t.Fatalf("Expected error: %s", expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error.\nexpected: %#v\ngiven: %#v",
			expectedErr, err)
	}
}

func TestFilesystem_Change_noChanges(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///test.tf"}
	fs.CreateAndOpenDocument(&testDocument{
		testHandler: fh,
	}, []byte{})

	var changes DocumentChanges
	err := fs.ChangeDocument(fh, changes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilesystem_Change_multipleChanges(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///test.tf"}
	fs.CreateAndOpenDocument(&testDocument{
		testHandler: fh,
	}, []byte{})

	var changes DocumentChanges
	changes = append(changes, &testChange{text: "ahoy"})
	changes = append(changes, &testChange{text: ""})
	changes = append(changes, &testChange{text: "quick brown fox jumped over\nthe lazy dog"})
	changes = append(changes, &testChange{text: "bye"})

	err := fs.ChangeDocument(fh, changes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilesystem_GetDocument_success(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///test.tf"}
	err := fs.CreateAndOpenDocument(&testDocument{
		testHandler: fh,
	}, []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}

	f, err := fs.GetDocument(fh)
	if err != nil {
		t.Fatal(err)
	}

	expectedFile := &documentMetadata{
		isOpen:    true,
	}
	opts := []cmp.Option{
		cmp.AllowUnexported(documentMetadata{}),
	}
	if diff := cmp.Diff(expectedFile, f, opts...); diff != "" {
		t.Fatalf("File doesn't match: %s", diff)
	}
}

func TestFilesystem_GetDocument_unopenedFile(t *testing.T) {
	fs := NewFilesystem()

	fh := &testHandler{"file:///test.tf"}
	_, err := fs.GetDocument(fh)

	expectedErr := &DocumentNotOpenErr{fh}
	if err == nil {
		t.Fatalf("Expected error: %s", expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error.\nexpected: %#v\ngiven: %#v",
			expectedErr, err)
	}
}

type testDocument struct {
	*testHandler
	text string
}

func (f *testDocument) Text() []byte {
	return []byte(f.text)
}

func (f *testDocument) Lines() source.Lines {
	return source.Lines{}
}

type testHandler struct {
	uri string
}

func (fh *testHandler) URI() string {
	return fh.uri
}

func (fh *testHandler) FullPath() string {
	return ""
}

func (fh *testHandler) Dir() string {
	return ""
}

func (fh *testHandler) Filename() string {
	return ""
}
func (fh *testHandler) Version() int {
	return 0
}
