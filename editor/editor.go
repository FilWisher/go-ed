package editor

import (
	"github.com/filwisher/godit/screen"
)

type Editor struct {
	Cur screen.Cursor
	Screen screen.Screen
}

func NewEditor(h, w int) Editor {
	return Editor{Screen: screen.Screen{h, w}}
}
