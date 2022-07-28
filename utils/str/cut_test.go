package str

import (
	"testing"
)

func TestCutRight(t *testing.T) {
	if CutRight("aaaacbb", "bb") != "aaaac" {
		t.Error()
	}
}
