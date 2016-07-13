/*
	Store editor commands in a byte addressed trie.
	Return noop if command not found in trie.
*/

package editor

type Cmd func(Editor, chan bool) Editor

var noop Cmd = func(e Editor, end chan bool) Editor {
	return e
}

type Trie struct {
	Fn       Cmd
	Children map[byte]*Trie
}

func (t *Trie) Add(name []byte, f Cmd) {

	if len(name) == 0 {
		t.Fn = f
		return
	}

	child, ok := t.Children[name[0]]
	if !ok {
		child = NewCmdSet()
		t.Children[name[0]] = child
	}
	child.Add(name[1:], f)
}

func (t *Trie) Find(name []byte) Cmd {

	if len(name) == 0 {
		return t.Fn
	}

	child, ok := t.Children[name[0]]
	if !ok {
		return noop
	}
	return child.Find(name[1:])
}

func NewCmdSet() *Trie {
	return &Trie{noop, make(map[byte]*Trie)}
}
