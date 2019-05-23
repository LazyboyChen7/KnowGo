package factory

import "testing"

func TestGetNameA(t *testing.T) {
	f := new(Factory)
	p := f.Create("a")
	if p.GetName() != "A" {
		t.Error("Something is wrong.")
	}
}
func TestGetNameB(t *testing.T) {
	f := new(Factory)
	p := f.Create("b")
	if p.GetName() != "B" {
		t.Error("Something is wrong.")
	}
}
