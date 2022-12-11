package main

import (
	"reflect"
	"testing"
)

var crt = [][]string{
	{"#", "#", "#", "#", ".", ".", "#", "#", ".", ".", "#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", "#", "#", ".", ".", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", "#", "#", "."},
	{"#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", ".", ".", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#", "."},
	{"#", "#", "#", ".", ".", "#", ".", ".", ".", ".", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", "#", "#", "#", ".", ".", "#", ".", ".", ".", ".", "#", "#", "#", ".", ".", ".", ".", ".", "#", "."},
	{"#", ".", ".", ".", ".", "#", ".", "#", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", ".", ".", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#", "."},
	{"#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", "#", ".", "#", ".", ".", "#", ".", "#", ".", ".", ".", ".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
	{"#", "#", "#", "#", ".", ".", "#", "#", "#", ".", "#", "#", "#", "#", ".", "#", ".", ".", "#", ".", "#", "#", "#", ".", ".", "#", "#", "#", "#", ".", "#", ".", ".", ".", ".", ".", "#", "#", ".", "."},
}

func Test_p1(t *testing.T) {
	result := p1(input)
	expected := 13920

	if result != expected {
		t.Errorf("p1() = %v, want %v", result, expected)
	}
}

func Test_p2(t *testing.T) {
	result := p2(input)
	expected := crt

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("p2() = %v, want %v", result, expected)
	}
}
