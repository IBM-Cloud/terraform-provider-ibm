package mqcloud

import (
	"errors"
	"strings"

	"os"

	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

// Mock function to replace fetchServicePlan in tests
func mockFetchServicePlan(meta interface{}, instanceID string) (string, error) {
	if instanceID == "reserved_deployment_plan_wildcard_id" {
		return "reserved-deployment-chaturanga2", nil
	}
	if instanceID == "reserved_deployment_plan_id" {
		return "reserved-deployment", nil
	}
	if instanceID == "default_plan_id" {
		return "default", nil
	}
	if instanceID == "invalid_plan_id" {
		return "", errors.New("[ERROR] Failed to fetch instance")
	}
	return "default", nil
}

var dummyInterface interface{}

func Test_checkSIPlan(t *testing.T) {
	tests := []struct {
		name                 string
		cachePlan            string
		wantErr              bool
		expectedErrorMessage string
		enforceRDP           bool
	}{
		{
			name:      "Cache Hit: Reserved Deployment Plan",
			cachePlan: "cached-reserved-deployment-chaturanga",
			wantErr:   false,
		},
		{
			name:                 "Cache Hit: Default Plan",
			cachePlan:            "cached-default-plan",
			wantErr:              true,
			expectedErrorMessage: "[ERROR] Terraform is only supported for Reserved Deployment, Reserved Capacity, and Reserved Capacity Subscription Plans. Your Service Plan is: cached-default-plan for the instance cached-default-plan",
		},
		{
			name:      "Reserved Deployment Plan Wildcard",
			cachePlan: "reserved_deployment_plan_wildcard_id",
			wantErr:   false,
		},
		{
			name:      "Reserved Deployment Plan",
			cachePlan: "reserved_deployment_plan_id",
			wantErr:   false,
		},
		{
			name:                 "Default Plan",
			cachePlan:            "default_plan_id",
			wantErr:              true,
			expectedErrorMessage: "[ERROR] Terraform is only supported for Reserved Deployment, Reserved Capacity, and Reserved Capacity Subscription Plans. Your Service Plan is: default for the instance default_plan_id",
		},
		{
			name:                 "fetchServicePlanFunc Error: Invalid id",
			wantErr:              true,
			cachePlan:            "invalid_plan_id",
			expectedErrorMessage: "[ERROR] Failed to fetch instance",
		},
	}
	fetchServicePlanFunc = mockFetchServicePlan
	defer func() {
		fetchServicePlanFunc = fetchServicePlan
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
				"service_instance_guid": {Type: schema.TypeString, Required: true},
			}, map[string]interface{}{
				"service_instance_guid": tt.cachePlan,
			})

			// Save to planCache for testing caching.
			if strings.Contains(tt.cachePlan, "cached") {
				planCache[tt.cachePlan] = tt.cachePlan
			}

			err := checkSIPlan(d, dummyInterface)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkSIPlan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && err.Error() != tt.expectedErrorMessage {
				t.Errorf("checkSIPlan() error = %v, expectedErrorMessage %v", err.Error(), tt.expectedErrorMessage)
			}
		})
	}
}
