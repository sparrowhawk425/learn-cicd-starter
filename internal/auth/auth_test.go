package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const AUTH_HEADER = "Authorization"
const API_KEY = "ApiKey"

func TestApiKey(t *testing.T) {

	// Valid API Key Auth header
	testKey := "OWEIMOSDIJEISDFOEjoesjf"
	headers := http.Header{}
	headers.Add(AUTH_HEADER, API_KEY+" "+testKey)
	apiKey, err := GetAPIKey(headers)
	require.NoError(t, err)
	assert.Equal(t, testKey, apiKey)

	// Invalid API Key Format
	headers = http.Header{}
	headers.Add(AUTH_HEADER, "Bearer"+" "+testKey)
	apiKey, err = GetAPIKey(headers)
	require.Error(t, err)
	require.Empty(t, apiKey)

	// Invalid Missing Auth Header
	headers = http.Header{}
	apiKey, err = GetAPIKey(headers)
	require.Error(t, err)
	require.Empty(t, apiKey)
}
