package text_test

import (
	"testing"
	"github.com/filwisher/go-ed/text"	
	"io/ioutil"
)

const (
	FILENAME = "test.txt"
	CONTENTS = "abcdefghijklmnopqrstuvwxyz"
)

func TestNewText(t *testing.T) {

	err := ioutil.WriteFile(FILENAME, []byte(CONTENTS), 0600)
	if err != nil {
		t.Errorf("could not make file: %s", err.Error())	
	}

	txt, err := text.NewText(FILENAME)
	if err != nil {
		t.Errorf("could not open %s: %s", FILENAME)	
	}
	defer txt.Exit()
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != CONTENTS {
		t.Errorf("got %s but expected %s", buf, CONTENTS)	
	}
}

func TestInsert(t *testing.T) {

	err := ioutil.WriteFile(FILENAME, []byte(CONTENTS), 0600)
	if err != nil {
		t.Errorf("could not make file: %s", err.Error())	
	}
	
	txt, err := text.NewText(FILENAME)
	if err != nil {
		t.Errorf("could not open %s: %s", FILENAME)	
	}
	defer txt.Exit()
	
	split := int64(5)
	txt.Insert(split, []byte(CONTENTS))
	expected := CONTENTS[:split] + CONTENTS + CONTENTS[split:]
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}

func TestAppend(t *testing.T) {
	
	err := ioutil.WriteFile(FILENAME, []byte(CONTENTS), 0600)
	if err != nil {
		t.Errorf("could not make file: %s", err.Error())	
	}

	txt, err := text.NewText(FILENAME)
	if err != nil {
		t.Errorf("could not open %s: %s", FILENAME)	
	}
	defer txt.Exit()
		
	txt.Insert(txt.First.Len, []byte(CONTENTS))
	expected := CONTENTS + CONTENTS
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}

	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}

func TestDelete(t *testing.T) {
	
	err := ioutil.WriteFile(FILENAME, []byte(CONTENTS), 0600)
	if err != nil {
		t.Errorf("could not make file: %s", err.Error())	
	}
	
	txt, err := text.NewText(FILENAME)
	if err != nil {
		t.Errorf("could not open %s: %s", FILENAME)	
	}
	defer txt.Exit()

	pos := 2
	len := 24
	txt.Delete(2,24)
	expected := CONTENTS[:pos] + CONTENTS[pos+len:]
	
	buf, err := txt.First.Bytes()
	if err != nil {
		t.Errorf("could not read bytes %s", err.Error())	
	}
	if string(buf) != expected {
		t.Errorf("got %s but expected %s", buf, expected)	
	}
}

func TestSave(t *testing.T) {
	
	err := ioutil.WriteFile(FILENAME, []byte(CONTENTS), 0600)
	if err != nil {
		t.Errorf("could not make file: %s", err.Error())
	}
	
	txt, err := text.NewText(FILENAME)
	if err != nil {
		t.Errorf("could not open %s: %s", FILENAME)	
	}
	defer txt.Exit()
	
	txt.Insert(txt.First.Len, []byte(CONTENTS))
	err = txt.Save()
	if err != nil {
		t.Errorf("could not save: %s", err.Error())	
	}
	
	expected := CONTENTS + CONTENTS
	
	got, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		t.Errorf("could not open file %s", err.Error())
	}
	
	if string(got) != expected {
		t.Errorf("got %s but expected %s", got, expected)	
	}
}
