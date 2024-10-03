// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"gotest.tools/assert"
)

func TestValidateUserPassword(t *testing.T) {
	testcases := []struct {
		user          DatabaseUser
		expectedError string
	}{
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "pizzapizzapizza",
				Type:     "database",
			},
			expectedError: "database user (testy) validation error:\npassword must contain at least one number",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "-_pizzapizzapizza",
				Type:     "database",
			},
			expectedError: "database user (testy) validation error:\npassword must not begin with a special character (_-)\npassword must contain at least one number",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "111111111111111",
				Type:     "database",
			},
			expectedError: "database user (testy) validation error:\npassword must contain at least one letter",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "$$$$$$$$$$$$$$a1",
				Type:     "database",
			},
			expectedError: "database user (testy) validation error:\npassword must not contain invalid characters",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "$",
				Type:     "database",
			},
			expectedError: "database user (testy) validation error:\npassword must contain at least one letter\npassword must contain at least one number\npassword must not contain invalid characters",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "aaaaa11111aaaa",
				Type:     "ops_manager",
			},
			expectedError: "database user (testy) validation error:\npassword must contain at least one special character (~!@#$%^&*()=+[]{}|;:,.<>/?_-)",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "password12345678$password",
				Type:     "ops_manager",
			},
			expectedError: "",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "~!@#$%^&*()=+[]{}|;:,.<>/?_-",
				Type:     "ops_manager",
			},
			expectedError: "database user (testy) validation error:\npassword must contain at least one letter\npassword must contain at least one number",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "~!@#$%^&*()=+[]{}|;:,.<>/?_-a1",
				Type:     "ops_manager",
			},
			expectedError: "",
		},
		{
			user: DatabaseUser{
				Username: "testy",
				Password: "pizza1pizzapizza1",
				Type:     "database",
			},
			expectedError: "",
		},
	}
	for _, tc := range testcases {
		err := tc.user.ValidatePassword()
		if tc.expectedError == "" {
			if err != nil {
				t.Errorf("TestValidateUserPassword: %q, %q unexpected error: %q", tc.user.Username, tc.user.Password, err.Error())
			}
		} else {
			assert.Equal(t, tc.expectedError, err.Error())
		}
	}
}

func TestValidateRBACRole(t *testing.T) {
	testcases := []struct {
		user          DatabaseUser
		expectedError string
	}{
		{
			user: DatabaseUser{
				Username: "invalid_format",
				Password: "",
				Type:     "database",
				Role:     core.StringPtr("+admin -all"),
			},
			expectedError: "database user (invalid_format) validation error:\nrole must be in the format +@category or -@category",
		},
		{
			user: DatabaseUser{
				Username: "invalid_operation",
				Password: "",
				Type:     "database",
				Role:     core.StringPtr("~@admin"),
			},
			expectedError: "database user (invalid_operation) validation error:\nrole must be in the format +@category or -@category",
		},
		{
			user: DatabaseUser{
				Username: "invalid_category",
				Password: "",
				Type:     "database",
				Role:     core.StringPtr("+@catfood -@dogfood"),
			},
			expectedError: "database user (invalid_category) validation error:\nrole must contain only allowed categories: all,admin,read,write",
		},
		{
			user: DatabaseUser{
				Username: "one_bad_apple",
				Password: "",
				Type:     "database",
				Role:     core.StringPtr("-@jazz +@read"),
			},
			expectedError: "database user (one_bad_apple) validation error:\nrole must contain only allowed categories: all,admin,read,write",
		},
		{
			user: DatabaseUser{
				Username: "invalid_user_type",
				Password: "",
				Type:     "ops_manager",
				Role:     core.StringPtr("+@all"),
			},
			expectedError: "database user (invalid_user_type) validation error:\nrole is only allowed for the database user",
		},
		{
			user: DatabaseUser{
				Username: "valid",
				Password: "",
				Type:     "database",
				Role:     core.StringPtr("-@all +@read"),
			},
			expectedError: "",
		},
		{
			user: DatabaseUser{
				Username: "blank_role",
				Password: "-@all +@read",
				Type:     "database",
				Role:     core.StringPtr(""),
			},
			expectedError: "",
		},
	}
	for _, tc := range testcases {
		err := tc.user.ValidateRBACRole()
		if tc.expectedError == "" {
			if err != nil {
				t.Errorf("TestValidateRBACRole: %q, %q unexpected error: %q", tc.user.Username, *tc.user.Role, err.Error())
			}
		} else {
			var errMsg string

			if err != nil {
				errMsg = err.Error()
			}

			assert.Equal(t, tc.expectedError, errMsg)
		}
	}
}

func TestAppendSwitchoverWarning(t *testing.T) {
	diags := appendSwitchoverWarning()
	warningNote := "Note: IBM Cloud Databases released new Hosting Models on May 1. All existing multi-tenant instances will have their resources adjusted to Shared Compute allocations during August 2024. To monitor your current resource needs, and learn about how the transition to Shared Compute will impact your instance, see our documentation https://cloud.ibm.com/docs/cloud-databases?topic=cloud-databases-hosting-models"

	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}

	if diags[0].Severity != diag.Warning {
		t.Errorf("expected severity %v, got %v", diag.Warning, diags[0].Severity)
	}

	if diags[0].Summary != warningNote {
		t.Errorf("expected summary %v, got %v", warningNote, diags[0].Summary)
	}
}

func TestPublicServiceEndpointsWarning(t *testing.T) {
	diags := publicServiceEndpointsWarning()
	warningNote := "IBM recommends using private endpoints only to improve security by restricting access to your database to the IBM Cloud private network. For more information, please refer to our security best practices, https://cloud.ibm.com/docs/cloud-databases?topic=cloud-databases-manage-security-compliance."

	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}

	if diags[0].Severity != diag.Warning {
		t.Errorf("expected severity %v, got %v", diag.Warning, diags[0].Severity)
	}

	if diags[0].Summary != warningNote {
		t.Errorf("expected summary %v, got %v", warningNote, diags[0].Summary)
	}
}
