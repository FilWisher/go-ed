/*
	Editor exposes core functionality of multi-file editor
	It does not implement any rendering and therefore any cursor logic
*/

package editor

import (
	"github.com/filwisher/go-ed/text"
)

type Range struct {
	Start, End int64
}

type textState struct {
	Dot Range
	Text *text.Text
}

type Editor struct {
	Texts map[string]*textState
	current string
}

func NewEditor() *Editor {
	return &Editor{
		Texts: make(map[string]*textState),
	}
}

func (e *Editor) ListFiles() []string {
	var filenames []string
	for filename := range e.Texts {
		filenames = append(filenames, filename)
	}
	return filenames
}

func (e *Editor) AddFile(filename string) error {
	t, err := text.NewText(filename)
	if err != nil {
		return err
	}
	e.Texts[filename] = &textState{
		Dot: Range{0,0},
		Text: t,
	}
	e.current = filename
	return nil
}

func (e *Editor) Remove(filename string) {
	ts, ok := e.Texts[filename]
	if !ok {
		return
	}
	ts.Text.Exit()
	delete(e.Texts, filename)
}

func (e *Editor) Exit() {
	// TODO: handle errors from close or ok to leave?
	for _, ts := range e.Texts {
		ts.Text.Exit()	
	}	
}

// TODO: 
// edit current
//	delete portion
//	replace portion
//	insert data
// Save current
// Change current
// 	by name
// 	next
// 	prev
// remove named
// remove current
