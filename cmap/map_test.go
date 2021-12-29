package cmap

import "testing"

func TestNew(t *testing.T) {
	m := New[string, int]()
	m.Clear()
}
