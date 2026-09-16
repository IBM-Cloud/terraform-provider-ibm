// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	result, err := buildURL("https://example.cloud.ibm.com/v2/things", nil)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things", result)

	// Query parameters are sorted so the URL is stable between runs.
	result, err = buildURL("https://example.cloud.ibm.com/v2/things", map[string]string{"limit": "10", "account_id": "abc"})
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things?account_id=abc&limit=10", result)

	// Parameters already on the URL survive, and are overridden by the map.
	result, err = buildURL("https://example.cloud.ibm.com/v2/things?limit=1&sort=name", map[string]string{"limit": "10"})
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things?limit=10&sort=name", result)

	_, err = buildURL("", nil)
	assert.Error(t, err)

	_, err = buildURL("ftp://example.com/thing", nil)
	assert.Error(t, err)

	_, err = buildURL("/v2/things", nil)
	assert.Error(t, err)
}

func TestJoinURL(t *testing.T) {
	result, err := joinURL("https://example.cloud.ibm.com/v2/things", "abc-123")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things/abc-123", result)

	// A trailing slash does not produce a doubled separator.
	result, err = joinURL("https://example.cloud.ibm.com/v2/things/", "abc-123")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things/abc-123", result)

	// The query string on the base URL is preserved.
	result, err = joinURL("https://example.cloud.ibm.com/v2/things?version=2", "abc-123")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things/abc-123?version=2", result)

	// Identifiers with reserved characters are escaped.
	result, err = joinURL("https://example.cloud.ibm.com/v2/things", "a b/c")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.cloud.ibm.com/v2/things/a%20b%2Fc", result)
}

func TestExtractID(t *testing.T) {
	id, err := extractID(`{"id": "abc-123"}`, "id")
	assert.NoError(t, err)
	assert.Equal(t, "abc-123", id)

	id, err = extractID(`{"metadata": {"guid": "xyz"}}`, "metadata.guid")
	assert.NoError(t, err)
	assert.Equal(t, "xyz", id)

	// Integral JSON numbers must not gain a decimal point.
	id, err = extractID(`{"id": 4211}`, "id")
	assert.NoError(t, err)
	assert.Equal(t, "4211", id)

	id, err = extractID(`{"id": true}`, "id")
	assert.NoError(t, err)
	assert.Equal(t, "true", id)

	// Failure modes that must be reported rather than silently ignored.
	_, err = extractID("", "id")
	assert.Error(t, err)

	_, err = extractID("not json", "id")
	assert.Error(t, err)

	_, err = extractID(`{"name": "thing"}`, "id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "available keys are [name]")

	_, err = extractID(`{"id": null}`, "id")
	assert.Error(t, err)

	_, err = extractID(`{"id": ""}`, "id")
	assert.Error(t, err)

	_, err = extractID(`{"id": {"nested": "value"}}`, "id")
	assert.Error(t, err)

	_, err = extractID(`{"id": "abc"}`, "id.nested")
	assert.Error(t, err)

	_, err = extractID(`[{"id": "abc"}]`, "id")
	assert.Error(t, err)
}

func TestResponseIsSuccess(t *testing.T) {
	assert.True(t, (&restResponse{StatusCode: 200}).isSuccess(nil))
	assert.True(t, (&restResponse{StatusCode: 204}).isSuccess(nil))
	assert.False(t, (&restResponse{StatusCode: 300}).isSuccess(nil))
	assert.False(t, (&restResponse{StatusCode: 404}).isSuccess(nil))
	assert.True(t, (&restResponse{StatusCode: 404}).isSuccess([]int{404}))
	assert.False(t, (&restResponse{StatusCode: 500}).isSuccess([]int{404}))
}

func TestStringMapAndIntList(t *testing.T) {
	assert.Equal(t, map[string]string{}, stringMap(nil))
	assert.Equal(t, map[string]string{"a": "b"}, stringMap(map[string]interface{}{"a": "b", "": "skipped"}))
	assert.Equal(t, []int{404, 409}, intList([]interface{}{404, 409}))
	assert.Nil(t, intList(nil))
}

func TestTruncateBody(t *testing.T) {
	assert.Equal(t, "short", truncateBody("short"))

	long := make([]byte, maxLoggedBodyLength+10)
	for index := range long {
		long[index] = 'x'
	}
	truncated := truncateBody(string(long))
	assert.Len(t, truncated, maxLoggedBodyLength+len("... (truncated)"))
}

func TestResponseError(t *testing.T) {
	err := responseError("GET", "https://example.cloud.ibm.com/v2/things", &restResponse{StatusCode: 500, Body: "boom"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
	assert.Contains(t, err.Error(), "boom")

	err = responseError("GET", "https://example.cloud.ibm.com/v2/things", &restResponse{StatusCode: 500})
	assert.Contains(t, err.Error(), "empty body")
}
