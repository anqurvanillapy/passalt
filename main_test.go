package passalt

import (
	"testing"
)

func TestNew(t *testing.T) {
	hashStrLen := len(New("foo"))
	if hashStrLen != 152 {
		t.Error("Expected token length as 152, got", hashStrLen)
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
