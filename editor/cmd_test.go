package editor

import (
	"testing"
)

func TestTrie(t *testing.T) {
  
		pairs := []struct{
			Key []byte	
			Value Cmd
		}{
			{ []byte("wil"),   func() { t.Log("one") }  },
			{ []byte("will"),  func() { t.Log("two") }  },
			{ []byte("willy"), func() { t.Log("three") }},
		}	
	
    cmds := NewCmdSet()
		
		for _, pair := range pairs {
			cmds.Add(pair.Key, pair.Value)	
		}
		
		for _, pair := range pairs {
			got := contacts.Find(pair.Key)
			got()
		}
}
