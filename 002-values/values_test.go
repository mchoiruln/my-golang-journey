package main

import (
	"testing"
)

func TestSumString(t *testing.T) {
	if "golang" != SumString("go", "lang") {
		t.Fatal("Test Fail")
	}
}

func TestSumInt(t *testing.T) {
	if 2 != SumInt(1, 1) {
		t.Fatal("Test Fail")
	}

	if 10 != SumInt(4, 6) {
		t.Fatal("Test Fail")
	}
}

func TestDivideInt(t *testing.T) {
	if 5 != DivideInt(25, 5) {
		t.Fatal("Test Fail")
	}
}

func TestAnd(t *testing.T) {
	if false != And(true, false) {
		t.Fatal("Test Fail")
	}
}

func TestOr(t *testing.T) {
	if true != Or(true, false) {
		t.Fatal("Test Fail")
	}
}

func TestNot(t *testing.T) {
	if false != Not(true) {
		t.Fatal("Test Fail")
	}
}
