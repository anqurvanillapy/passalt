package main

import (
	"fmt"
	"testing"
)

func ExampleNewAndCheck() {
	hashStr := New("foo")
	fmt.Println(Check(hashStr, "foo"))
	fmt.Println(Check(hashStr, "bar"))
	// Output:
	// true
	// false
}

func TestNew(t *testing.T) {
	hashStr := New("foo")
	expected := "sha512$31cb48a8eec9aa9c$1411bf8dd5178a201862f638cbab738bb252982ec0a73085c61e611cf45932cfcea9f908bb91a4133061b34d4f4c4e890add289ac5e43380aeaaff93344758e2"
	if hashStr != expected {
		t.Error("Expected ", expected, ", got ", hashStr)
	}
}

func TestCheck(t *testing.T) {
	hashStr := New("foo")

	if Check(hashStr, "foo") != true {
		t.Error("Expected \"foo\" to be valid, got false")
	}

	if Check(hashStr, "bar") != false {
		t.Error("Expected \"bar\" to be invalid, got true")
	}
}
