package common

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
}
