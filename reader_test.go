package main

import (
	"testing"
	"os"
)

func TestEmpty(t *testing.T) {
	file, _ := os.Open("./testdata/nodata.bin")

	c := make(chan Result)
	subject := NewReader(file, 820, &c)
	subject.Read()
}
