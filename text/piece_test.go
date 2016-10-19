package text_test

import (
	"testing"
	"github.com/filwisher/go-ed/text"
	"os"
)

func TestSplit(t *testing.T) {

	totalLen := int64(50)
	chain := text.Piece{
		File: os.Stdin,
		Off: 0,
		Len: totalLen,
		Prev: nil,
		Next: nil,
	}
	
	for i := int64(40); i >= 1; i /= 2 {
		chain.Split(i)	
	}	

	count := int64(0)
	for p := &chain; p != nil; p = p.Next {
		count += p.Len	
	}
	
	if count != totalLen {
		t.Errorf("total length of chain expected to be %d but was %d", totalLen)	
	}
}

func atomize(p *text.Piece) {
	length := p.Len
	for i := int64(0); i < length; i++ {
		p.Split(i)	
	}
}

func TestBytes(t *testing.T) {
	// TODO: ensure contents in file before tests
	contents := "abcdefghijklmnopqrstuvwxyz"
	filename := "test.txt"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		t.Errorf("could not open file %s", filename)	
	}
	
	chain := text.Piece{
		File: file,
		Off: 0,
		Len: 26,
		Prev: nil,
		Next: nil,
	}
	
	buf, err := chain.Bytes()
	if err != nil {
		t.Errorf("could not get bytes from piece %s", err.Error())	
	}
	if string(buf) != contents {
		t.Errorf("expected\t%s\ngot%s", buf, contents)
	}
	
	atomize(&chain)
	
	buf, err = chain.Bytes()
	if err != nil {
		t.Errorf("could not get bytes from piece %s", err.Error())
	}
	
	if string(buf) != contents {
		t.Errorf("expected\t%s\ngot%s", buf, contents)
	}
}
