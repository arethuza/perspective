package misc

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	s, b, err := GenerateRandomString(50)
	if err != nil {
		t.Error(err)
	}
	if len(s) < 50 {
		t.Error("unexpected s length")
	}
	if len(b) < 50 {
		t.Error("unexpected b length")
	}
}
