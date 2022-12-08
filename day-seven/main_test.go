package main

import "testing"

func Test_p1(t *testing.T) {
	result := p1(input)
	expected := 1077191

	if result != expected {
		t.Errorf("p1() = %v, want %v", result, expected)
	}
}

func Test_p2(t *testing.T) {
	result := p2(input)
	expected := 5649896

	if result != expected {
		t.Errorf("p2() = %v, want %v", result, expected)
	}
}
