package text_test

import (
	"github.com/filwisher/go-ed/text"
	"testing"
	//"os"
	"fmt"
)

func TestSplit(t *testing.T) {
	p, err := text.PieceFromFile("./testfile.txt")
	if err != nil {
		t.Fatalf("could not open File: %s", err.Error())	
	}
	
	len := p.Len
	text, err := p.Bytes()
	if err != nil {
		t.Fatalf("could not read: %s", err.Error())	
	}	
	
	before, after := p.Split(7)
	if before.Len + after.Len != len {
		t.Fatalf("piece sizes not maintained after split")		
	}
	
	fullText, err := p.Content()
	if err != nil {
		t.Fatalf("could not read: %s", err.Error())	
	}	
	if string(text) != string(fullText) {
		t.Fatalf("2.split deformed text: %s != %s", text, string(fullText))	
	}
}

func TestInsert(t *testing.T) {
	p, err := text.PieceFromFile("./testfile.txt")
	if err != nil {
		t.Fatalf("could not open File: %s", err.Error())	
	}
	
	p.Insert(1, &text.Piece{
		File: p.File,
		Off: p.Off + 1,
		Len: 1,
	})
	
	fullText, err := p.Content()
	if err != nil {
		t.Fatalf("could not read: %s", err.Error())	
	}	
	fmt.Printf("got %s\n", fullText)
}

func TestRun(t *testing.T) {

	p, err := text.PieceFromFile("./testfile.txt")
	if err != nil {
		t.Fatalf("could not open File: %s", err.Error())	
	}

	s, err := p.Bytes()
	if err != nil {
		t.Errorf("could not read File: %s", err.Error())	
	}
	
	fmt.Printf("read: %s", s)
}
