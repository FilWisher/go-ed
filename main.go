package main

import (
	"fmt"
	"github.com/filwisher/go-ed/editor"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

var (
	e        editor.Editor = editor.NewEditor(10, 10)
	bindings *editor.Trie  = loadCommands(commands)
)

func redrawScreen() {
	fmt.Printf("%d,%d", e.Cur.X, e.Cur.Y)
}

/* read keypress from terminal */
func getKeypress() ([]byte, error) {
	data := make([]byte, 4)
	count, err := os.Stdin.Read(data)

	if err != nil {
		return nil, err
	}

	return data[:count], nil
}

/* loop: read keypress and lookup apply corresponding command */
func commandLoop(end chan bool) {
	for {
		key, err := getKeypress()
		if err != nil {
			log.Fatal(err.Error())
		}

		cmd := bindings.Find(key)
		e = cmd(e, end)

		redrawScreen()
		fmt.Printf("\t(%d)\n", key)
	}
}

func main() {

	end := make(chan bool)

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err.Error())
	}
	defer terminal.Restore(0, oldState)

	go commandLoop(end)
	<-end
}
