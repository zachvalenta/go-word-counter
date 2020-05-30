package main

import (
	"bytes"
	"testing"
)

func TestCounter(t *testing.T){
	b := bytes.NewBufferString("word1 word2 word3 word4")
	expect := 4
	actual := count(b)
	if expect != actual {
		t.Errorf("expected %d actual %d", expect, actual)
	}
}