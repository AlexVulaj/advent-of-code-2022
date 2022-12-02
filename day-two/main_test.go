package main

import "testing"

func Test_p1(t *testing.T) {
	result := p1(input)
	expected := 12276

	if result != expected {
		t.Errorf("p1() = %v, want %v", result, expected)
	}
}

func Test_p2(t *testing.T) {
	result := p2(input)
	expected := 9975

	if result != expected {
		t.Errorf("p2() = %v, want %v", result, expected)
	}
}
