package main

import (
	"github.com/filwisher/go-ed/editor"
)

type CmdConfig struct {
	Name string
	Key []byte
	Fn  editor.Cmd	
}

var commands []CmdConfig = []CmdConfig{
	{ "C-X",   []byte{ 24  },  Exit }, 
	{ "ENTER", []byte{ 13  },  Exit }, 
	{ "UP",    []byte{ 'k' }, Up },
	{ "DOWN",  []byte{ 'j' }, Down },
	{ "RIGHT", []byte{ 'l' }, Right },
	{ "LEFT",  []byte{ 'h' }, Left },
}

func Exit(e editor.Editor, end chan bool) editor.Editor { 
	end <- true 
	return e
}

func Up(e editor.Editor, end chan bool) editor.Editor { 
	if e.Cur.Y - 1 >= 0 {
		e.Cur.Y -= 1
	}
	return e
}

func Down(e editor.Editor, end chan bool) editor.Editor { 
	if e.Cur.Y + 1 <= e.Screen.Height {
		e.Cur.Y += 1
	}
	return e
}

func Right(e editor.Editor, end chan bool) editor.Editor { 
	if e.Cur.X + 1 <= e.Screen.Width {
		e.Cur.X += 1
	}
	return e
}

func Left(e editor.Editor, end chan bool) editor.Editor { 
	if e.Cur.X - 1 >= 0 {
		e.Cur.X -= 1
	}
	return e
}

func loadCommands(configs []CmdConfig) *editor.Trie {
	cs := editor.NewCmdSet()
	for _, conf := range configs {
		cs.Add(conf.Key, conf.Fn)	
	}	
	return cs
}
