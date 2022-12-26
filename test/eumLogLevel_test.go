package test

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEumLogLevel(t *testing.T) {
	assert.Equal(t, "Info", eumLogLevel.NoneLevel.ToString())
	assert.Equal(t, "Info", eumLogLevel.GetName(eumLogLevel.NoneLevel))

	assert.Equal(t, "Error", eumLogLevel.Error.ToString())
	assert.Equal(t, "Error", eumLogLevel.GetName(eumLogLevel.Error))

	assert.Equal(t, "Debug", eumLogLevel.Debug.ToString())
	assert.Equal(t, "Debug", eumLogLevel.GetName(eumLogLevel.Debug))

	assert.Equal(t, "Critical", eumLogLevel.Critical.ToString())
	assert.Equal(t, "Critical", eumLogLevel.GetName(eumLogLevel.Critical))

	assert.Equal(t, "Info", eumLogLevel.Information.ToString())
	assert.Equal(t, "Info", eumLogLevel.GetName(eumLogLevel.Information))

	assert.Equal(t, "Trace", eumLogLevel.Trace.ToString())
	assert.Equal(t, "Trace", eumLogLevel.GetName(eumLogLevel.Trace))

	assert.Equal(t, "Warn", eumLogLevel.Warning.ToString())
	assert.Equal(t, "Warn", eumLogLevel.GetName(eumLogLevel.Warning))

}
