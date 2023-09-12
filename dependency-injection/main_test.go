package main

import (
	"bytes"
	"testing"
)

func TestFileWrite(t *testing.T) {

	var b bytes.Buffer
	database := NewDatabase(&b)
	_ = database
}
