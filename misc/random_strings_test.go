package misc

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	s, err := GenerateRandomString(50)
	if err != nil {
		t.Error(err)
	}
	if (len(s) < 50) {
		t.Error("unexpected length")
	}
}