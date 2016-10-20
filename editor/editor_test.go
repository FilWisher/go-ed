package editor_test

import (
	"github.com/filwisher/go-ed/editor"
	"testing"
	"os"
)

func removeAll(filenames []string) {
	for _, filename := range filenames {
		os.Remove(filename)	
	}
}

func TestEditor(t *testing.T) {

	filenames := []string {
		"hello.txt",
		"cool.txt",
	}

	e := editor.NewEditor()
	
	defer func () {
		e.Exit()
		removeAll(filenames)
	}()
	
	for _, filename := range filenames {
		e.AddFile(filename)	
	}

	addedNames := make(map[string]struct{})
	for _, name := range e.ListFiles() {
		addedNames[name] = struct{}{}
	}
	for _, filename := range filenames {
		if _, ok := addedNames[filename]; !ok {
			t.Errorf("filename %s does not occur in editor", filename)
		}
	}
}

func TestRemove(t *testing.T) {
	
	toRemove := "cool.txt"
	filenames := []string {
		"hello.txt",
		toRemove,
	}

	e := editor.NewEditor()
	
	defer func () {
		e.Exit()
		removeAll(filenames)
	}()
	
	for _, filename := range filenames {
		e.AddFile(filename)	
	}
	
	e.Remove(toRemove)
	
	addedNames := make(map[string]struct{})
	for _, name := range e.ListFiles() {
		addedNames[name] = struct{}{}
	}
	if _, ok := addedNames[toRemove]; ok {
		t.Errorf("filename %s should not occur", toRemove)
	}
}
