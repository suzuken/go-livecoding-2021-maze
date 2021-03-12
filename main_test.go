package main

import "testing"

func TestInit(t *testing.T) {
	m := Maze{}
	m.init(7, 7)
	if got := m.width; got != 7 {
		t.Fatalf("want %d, got %d", 7, got)
	}
}
