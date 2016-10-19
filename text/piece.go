package text

import (
	"os"
)

type Piece struct {
	File *os.File
	Off, Len int64
	Prev, Next *Piece
}

func (p *Piece) Split(pos int64) {

	// splitting at divide means job done
	if pos == 0 {
		return
	}

	// splitting work occurs on future piece
	if pos >= p.Len {
		if p.Next == nil {
			return
		}
		p.Next.Split(pos - p.Len)
		return
	}
	
	newPiece := &Piece{
		p.File,
		p.Off + pos,
		p.Len - pos,
		p,
		p.Next,
	}
	p.Next = newPiece
	p.Len = pos
}

func (p *Piece) pieceAt(pos int64) (int64, *Piece) {
	if pos < p.Len {
		return pos, p	
	}
	
	// TODO: double check this is correct: return pos as end of final piece
	if p.Next == nil {
		return p.Len, p
	}

	return p.Next.pieceAt(pos - p.Len)
}

func (p *Piece) bytesSingle() ([]byte, error) {
	buf := make([]byte, p.Len)
	_, err := p.File.ReadAt(buf, p.Off)
	return buf, err
}

func join(a, b *Piece) {
	if a != nil {
		a.Next = b	
	}
	if b != nil {
		b.Prev = a	
	}
}

func patch(a, b, c *Piece) {
	join(a, b)
	join(b, c)
}

func (p *Piece) Bytes() ([]byte, error) {
	var bufFull []byte
	for piece := p; piece != nil; piece = piece.Next {
		buf, err := piece.bytesSingle()
		if err != nil {
			return bufFull, err
		}
		bufFull = append(bufFull, buf...)
	}
	return bufFull, nil
}
