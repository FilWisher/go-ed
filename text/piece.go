package text

import (
	_ "fmt"
	"os"
	"bytes"
)

type Piece struct {
	Next, Prev *Piece

	Off, Len int64
	
	// specify whether contents is buffer or file
	IsBuffer bool
	// contents
	File *os.File
	Buf *[]byte
}

func PieceFromFile(Filename string) (*Piece, error) {
	File, err := os.Open(Filename) // TODO: ensure readonly
	if err != nil {
		return nil, err	
	}
	
	fi, err := File.Stat()
	if err != nil {
		return nil, err	
	}

	return &Piece{
		Off: 0,
		Len: fi.Size(),
		File: File,
		IsBuffer: false,
	}, nil
}

func NewPiece() *Piece {
	return &Piece{}
}

// split piece at pos and connect reattach links 
// return pieces on either side of the split
func (p *Piece) Split(pos int64) (*Piece, *Piece) {
	
	before, off := p.PieceAt(pos)

	if pos >= p.Len {
		return p, nil
	}
	before.Next = &Piece{
		Next: before.Next,
		Prev: before,
		Off: before.Off + off,
		Len: before.Len - off,
		File: before.File,
		IsBuffer: before.IsBuffer,
	}
	before.Len = pos

	return before, before.Next
}

func (p *Piece) Bytes() ([]byte, error) {
	buf := make([]byte, p.Len)
	n, err := p.File.ReadAt(buf, p.Off)
	return buf[:n], err
}

func (p *Piece) Content() ([]byte, error) {
	var contents []byte
	
	for pp := p; pp != nil; pp = pp.Next {
		buf, err := pp.Bytes()
		if err != nil {
			return contents, err
		}
		contents = bytes.Join([][]byte{contents, buf}, []byte(""))
	}
	return contents, nil
}

func (p *Piece) Insert(pos int64, np *Piece) {
	before, after := p.Split(pos)
	before.Next = np
	np.Next = after
	after.Prev = np
	np.Prev = before
}

func (p *Piece) PieceAt(pos int64) (*Piece, int64) {
	var pp *Piece
	for pp = p; pos > pp.Len; pp = pp.Next {
		pos -= pp.Len
	}
	return pp, pos
}
