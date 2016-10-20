/* 
	Text represents the state of a file in memory
	It is implemented as a doubly linked-list of pieces
	Two files on disk are associated with text: the original file
	and an append-only file that stores changes/additions 
*/

package text

import (
	"os"
	"time"
	"fmt"
)

type Text struct {
	First, Last *Piece
	File, Changes *os.File
	Filename, Changesname string
	lastWrite int64
}

func NewText(filename string) (*Text, error) {

	// open original file (readonly)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return nil, err	
	}
	
	// get file info for length
	fi, err := file.Stat()
	if err != nil {
		return nil, err	
	}
	
	// open tmp file for changes (read and append)
	tmpname := fmt.Sprintf("%d.wed", time.Now().Unix())
	tmp, err := os.OpenFile(tmpname, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return nil, err	
	}
	
	piece := &Piece{
		File: file,
		Off: 0,
		Len: fi.Size(),
	}	
	
	return &Text{
		First: piece,
		Last: piece,
		File: file,
		Changes: tmp,
		Filename: filename,
		Changesname: tmpname,
	}, nil
}

func (t *Text) insertPiece(pos int64, p *Piece) {
	
	pos, target := t.First.pieceAt(pos)
	
	var pre, post *Piece
	
	if pos == 0 {
		pre = target.Prev
		post = target	
		if pre == nil {
			t.First = p
		}
	}	else {
		target.Split(pos)
		pre = target
		post = target.Next
	}
	
	patch(pre, p, post)			
}

// TODO: add caching so text isn't written to disk every time
func (t *Text) Insert(pos int64, data []byte) error {
	n, err := t.Changes.Write(data)
	if err != nil {
		return err
	}
	piece := &Piece{
		File: t.Changes,
		Off: t.lastWrite,
		Len: int64(n),
	}	
	t.lastWrite += int64(n)
	t.insertPiece(pos, piece)
	return nil
}

func (t *Text) Delete(pos, len int64) {
	split1, first := t.First.pieceAt(pos)
	first.Split(split1)
	
	split2, second := first.pieceAt(pos+len)
	second.Split(split2)
	
	pre := first
	post := second.Next

	join(pre, post)
}

// save to disk
func (t *Text) Save() error {

	file, err := os.OpenFile(t.Filename, os.O_WRONLY, 0600)
	if err != nil {
		return err	
	}
	defer file.Close()
	
	buf, err := t.First.Bytes()
	if err != nil {
		return err	
	}

	_, err = file.Write(buf)
	return err
}
