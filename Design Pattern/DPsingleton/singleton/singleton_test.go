package singleton

import "testing"

func TestSingleton(t *testing.T) {
	s := New("lazyboy", 22, "China")
	s1 := New("asdf", 23, "Nibelungenlied")
	if s != s1 {
		t.Error("singleton is wrong.")
	}
}
