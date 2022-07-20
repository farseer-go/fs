package parse

import (
	"testing"
)

func TestConvertIntToBool(t *testing.T) {
	Convert(1, false)
	Convert(0, false)
}
