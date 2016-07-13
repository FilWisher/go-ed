package editor

import (
	"testing"
)

var (
	e    Editor
	ch   chan bool = make(chan bool)
	cmds *Trie     = NewCmdSet()
)

func simpleCmd(val string, t *testing.T) func(Editor, chan bool) Editor {
	return func(e Editor, ch chan bool) Editor {
		t.Log(val)
		return e
	}
}

func TestTrie(t *testing.T) {

	pairs := []struct {
		Key   []byte
		Value Cmd
	}{
		{[]byte("wil"), simpleCmd("one", t)},
		{[]byte("will"), simpleCmd("two", t)},
		{[]byte("willy"), simpleCmd("three", t)},
		{[]byte(""), simpleCmd("hello", t)},
	}

	for _, pair := range pairs {
		cmds.Add(pair.Key, pair.Value)
	}

	for _, pair := range pairs {
		got := cmds.Find(pair.Key)
		e = got(e, ch)
	}
}
