package fileutil

import (
	"testing"
)

func TestFileutil(t *testing.T) {
	result := GetMessage()
	if result != "Message1" {
		t.Errorf("error1")
	}
}
