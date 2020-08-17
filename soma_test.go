package main

import "testing"

func TestSoma(t *testing.T){
	if (Soma(3,5)!=8) {
		t.Errorf("Soma(%q, %q) == %q, want %q", 3, 5, Soma(3,5), 8)		
	}
}