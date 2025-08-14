package mqcloud

import (
	"os"

	"testing"
)

func TestLoadWaitForQmStatusEnvVar(t *testing.T) {
	tests := []struct {
		name           string
		envValue       string
		expectedOutput bool
	}{
		{
			name:           "EnvVar Set To True",
			envValue:       "true",
			expectedOutput: true,
		},
		{
			name:           "EnvVar Set To False",
			envValue:       "false",
			expectedOutput: false,
		},
		{
			name:           "EnvVar Not Set",
			envValue:       "",
			expectedOutput: true,
		},
		{
			name:           "EnvVar Set To Random String",
			envValue:       "random",
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Setenv("IBMCLOUD_MQCLOUD_WAIT_FOR_QM_STATUS", tt.envValue)

			result := loadWaitForQmStatusEnvVar()
			if result != tt.expectedOutput {
				t.Errorf("loadWaitForQmStatusEnvVar() with env value '%s', got %v, want %v", tt.envValue, result, tt.expectedOutput)
			}

			_ = os.Unsetenv("IBMCLOUD_MQCLOUD_WAIT_FOR_QM_STATUS")
		})
	}
}

func TestIsVersionDowngrade(t *testing.T) {
	tests := []struct {
		oldVersion, newVersion string
		want                   bool
	}{
		// Basic Version Comparison
		{"1.2.3", "1.2.2", true},
		{"1.2.3", "1.2.4", false},
		{"1.2.3", "1.3.0", false},

		// Equal Versions
		{"1.2.3", "1.2.3", false},

		// Edge Cases
		{"", "1.0.0", false},
		{"1.0.0", "", true},
		{"abc.def.ghi", "1.2.3", false},
	}

	for _, tt := range tests {
		t.Run(tt.oldVersion+" to "+tt.newVersion, func(t *testing.T) {
			if got := IsVersionDowngrade(tt.oldVersion, tt.newVersion); got != tt.want {
				t.Errorf("IsVersionDowngrade(%q, %q) = %v, want %v", tt.oldVersion, tt.newVersion, got, tt.want)
			}
		})
	}
}

func TestHandlePlanCheck(t *testing.T) {
	tests := []struct {
		plan       string
		instanceID string
		wantErr    bool
		errMessage string
	}{
		// Test with matching plan
		{"reserved-deployment", "123", false, ""},

		// Test with non-matching plan
		{"Basic_Plan", "123", true, "[ERROR] Terraform is only supported for Reserved Deployment, Reserved Capacity, and Reserved Capacity Subscription Plans. Your Service Plan is: Basic_Plan for the instance 123"},
	}

	for _, tt := range tests {
		t.Run(tt.plan, func(t *testing.T) {
			err := handlePlanCheck(tt.plan, tt.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("handlePlanCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Errorf("handlePlanCheck() error = %v, wantErrMessage %v", err.Error(), tt.errMessage)
			}
		})
	}
}
