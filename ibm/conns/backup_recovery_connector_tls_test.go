// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"os"
	"testing"
)

func TestBackupRecoveryConnectorSkipTLSVerify(t *testing.T) {
	const envVar = "IBMCLOUD_BACKUP_RECOVERY_CONNECTOR_SKIP_TLS_VERIFY"

	cases := []struct {
		name  string
		value string
		set   bool
		want  bool
	}{
		{name: "unset keeps TLS verification enabled", set: false, want: false},
		{name: "empty keeps TLS verification enabled", set: true, value: "", want: false},
		{name: "false keeps TLS verification enabled", set: true, value: "false", want: false},
		{name: "invalid keeps TLS verification enabled", set: true, value: "yes-please", want: false},
		{name: "true disables TLS verification", set: true, value: "true", want: true},
		{name: "1 disables TLS verification", set: true, value: "1", want: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.set {
				t.Setenv(envVar, tc.value)
			} else {
				t.Setenv(envVar, "")
				// t.Setenv restores the original value; unset for this case.
				unsetEnv(t, envVar)
			}
			if got := backupRecoveryConnectorSkipTLSVerify(); got != tc.want {
				t.Fatalf("backupRecoveryConnectorSkipTLSVerify() = %v, want %v", got, tc.want)
			}
		})
	}
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("failed to unset %s: %v", key, err)
	}
}
