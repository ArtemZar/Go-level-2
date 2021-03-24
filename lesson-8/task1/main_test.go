package main

import (
	"testing"
)

func TestListDirByReadDir(t *testing.T)  {
	for i:=0;i<10;i++ {
		ListDirByReadDir("../")
		if len(FindFiles)==0 {
			t.Fatal("e r r o r")
		}
	}
}
