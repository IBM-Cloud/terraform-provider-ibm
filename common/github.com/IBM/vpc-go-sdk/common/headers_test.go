package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSystemInfo(t *testing.T) {
	var sysinfo = GetSystemInfo()
	assert.NotNil(t, sysinfo)
	assert.True(t, strings.Contains(sysinfo, "arch="))
	assert.True(t, strings.Contains(sysinfo, "os="))
	assert.True(t, strings.Contains(sysinfo, "go.version="))
}

func TestGetSdkHeaders(t *testing.T) {
	var headers = GetSdkHeaders("myService", "v123", "myOperation")
	assert.NotNil(t, headers)

	var foundIt bool

	_, foundIt = headers[HEADER_NAME_USER_AGENT]
	assert.True(t, foundIt)
	_, foundIt = headers[X_REQUEST_ID]
	assert.True(t, foundIt)
}
