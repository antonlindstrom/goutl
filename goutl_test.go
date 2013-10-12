package goutl

import "testing"

type ByteSlice [][]byte

func TestBytesToStrings(t *testing.T) {
	bytes := ByteSlice{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
	}

	if BytesToStrings(bytes)[0] != "a" {
		t.Error("Did not get string a from bytes[0]")
	}
	if BytesToStrings(bytes)[1] != "b" {
		t.Error("Did not get string b from bytes[1]")
	}
	if BytesToStrings(bytes)[2] != "c" {
		t.Error("Did not get string c from bytes[2]")
	}
}
