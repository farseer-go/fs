package test

import (
	"testing"

	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/snc"
	"github.com/stretchr/testify/assert"
)

func TestEumLogLevel(t *testing.T) {
	assert.Equal(t, "None", eumLogLevel.NoneLevel.ToString())
	assert.Equal(t, eumLogLevel.NoneLevel, eumLogLevel.GetEnum("???"))

	assert.Equal(t, "Error", eumLogLevel.Error.ToString())
	assert.Equal(t, eumLogLevel.Error, eumLogLevel.GetEnum("Error"))

	assert.Equal(t, "Debug", eumLogLevel.Debug.ToString())
	assert.Equal(t, eumLogLevel.Debug, eumLogLevel.GetEnum("Debug"))

	assert.Equal(t, "Critical", eumLogLevel.Critical.ToString())
	assert.Equal(t, eumLogLevel.Critical, eumLogLevel.GetEnum("Critical"))

	assert.Equal(t, "Info", eumLogLevel.Information.ToString())
	assert.Equal(t, eumLogLevel.Information, eumLogLevel.GetEnum("Info"))
	assert.Equal(t, eumLogLevel.Information, eumLogLevel.GetEnum("Information"))

	assert.Equal(t, "Trace", eumLogLevel.Trace.ToString())
	assert.Equal(t, eumLogLevel.Trace, eumLogLevel.GetEnum("Trace"))

	assert.Equal(t, "Warn", eumLogLevel.Warning.ToString())
	assert.Equal(t, eumLogLevel.Warning, eumLogLevel.GetEnum("Warning"))
	assert.Equal(t, eumLogLevel.Warning, eumLogLevel.GetEnum("Warn"))

	var e = eumLogLevel.Debug
	b, _ := snc.Marshal(e)
	e = eumLogLevel.Information
	_ = e.UnmarshalJSON(b)
	assert.True(t, e == eumLogLevel.Debug)
}
