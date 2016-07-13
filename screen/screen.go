/*
	Utility wrappers around VT100 escape codes
*/

package screen

import "fmt"

const (
	ESC_CLEAR_SCREEN 	string = "\x1b[2J"
	ESC_CURSOR_SET 		string = "\x1b[%d;%dH"
	ESC_CURSOR_HOME	  	string = "\x1b[H"
	ESC_CURSOR_HIDE	  	string = "\x1b[?25l"
	ESC_CURSOR_SHOW		string = "\x1b[?25h"
)

type Screen struct {
	Height, Width int
}

func (s Screen) Clear() { fmt.Printf(ESC_CLEAR_SCREEN) }

type Cursor struct {
	X, Y int
}

func (c Cursor) Hide() { fmt.Printf(ESC_CURSOR_HIDE) }
func (c Cursor) Show() { fmt.Printf(ESC_CURSOR_SHOW) }
func (c Cursor) Move() { fmt.Printf(ESC_CURSOR_SET, c.Y, c.X) }
func (c Cursor) Set(x, y int) { 
	c.X, c.Y = x, y
	c.Move()
}
