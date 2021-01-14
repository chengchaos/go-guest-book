package learning

import "testing"

func TestTryDefer(t *testing.T) {
	TryDefer()
}

func TestWriteFile(t *testing.T) {
	filename := "test.txt"
	WriteFile(filename)
}
