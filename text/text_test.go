package text_test

import (
	"testing"
	"github.com/filwisher/go-ed/text"	
	"os"
)

const (
	filename = "test.txt"
	contents = "abcdefghijklmnopqrstuvwxyz"
)

func cleanup(txt *text.Text) {
	txt.File.Close()
	txt.Changes.Close()
	os.Remove(txt.Changesname)
}

func TestNewText(t *testing.T) {
	
	txt, err := text.NewText(filename)
	if err != nil {
		t.Errorf("could not open %s: %s", filename)	
	}
	defer cleanup(txt)
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != contents {
		t.Errorf("got %s but expected %s", buf, contents)	
	}
}

func TestInsert(t *testing.T) {

	txt, err := text.NewText(filename)
	if err != nil {
		t.Errorf("could not open %s: %s", filename)	
	}
	defer cleanup(txt)
	
	split := int64(5)
	txt.Insert(split, []byte(contents))
	expected := contents[:split] + contents + contents[split:]
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}

func TestAppend(t *testing.T) {

	txt, err := text.NewText(filename)
	if err != nil {
		t.Errorf("could not open %s: %s", filename)	
	}
	defer cleanup(txt)
		
	txt.Insert(txt.First.Len, []byte(contents))
	expected := contents + contents
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}

func TestDelete(t *testing.T) {
	
	txt, err := text.NewText(filename)
	if err != nil {
		t.Errorf("could not open %s: %s", filename)	
	}
	defer cleanup(txt)

	pos := 2
	len := 24
	txt.Delete(2,24)
	expected := contents[:pos] + contents[pos+len:]
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}
	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}
