package utils_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/kevbeltrao/websocket/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetLogger(t *testing.T) {
	var buf bytes.Buffer
	customLogger := log.New(&buf, "customLogger: ", log.LstdFlags)
	utils.SetLogger(customLogger)

	utils.LogInfo("Test Info Message")
	assert.Contains(t, buf.String(), "customLogger: ")
	assert.Contains(t, buf.String(), "INFO: Test Info Message")
}

func TestLogInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "testLogger: ", log.LstdFlags)
	utils.SetLogger(logger)

	utils.LogInfo("Test Info")
	require.Contains(t, buf.String(), "testLogger: ")
	require.Contains(t, buf.String(), "INFO: Test Info")
}

func TestLogError(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "testLogger: ", log.LstdFlags)
	utils.SetLogger(logger)

	utils.LogError("Test Error")
	require.Contains(t, buf.String(), "testLogger: ")
	require.Contains(t, buf.String(), "ERROR: Test Error")
}
