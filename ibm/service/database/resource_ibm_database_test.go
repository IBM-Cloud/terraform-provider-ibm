// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"gotest.tools/assert"
	"testing"
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
		err := tc.user.Validate()
		if tc.expectedError == "" {
			if err != nil {
				t.Errorf("TestValidateUserPassword: %q, %q unexpected error: %q", tc.user.Username, tc.user.Password, err.Error())
			}
		} else {
			assert.Equal(t, tc.expectedError, err.Error())
		}
	}
}
