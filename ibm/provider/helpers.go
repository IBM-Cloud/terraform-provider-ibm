// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GetStringFromEnv checks environment variables for a value if the given StringValue is null
func GetStringFromEnv(value basetypes.StringValue, envVars []string) basetypes.StringValue {
	if !value.IsNull() {
		return value
	}
	for _, envVar := range envVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			return basetypes.NewStringValue(envValue)
		}
	}
	return basetypes.NewStringNull()
}

// GetInt64FromEnv checks environment variables for a value if the given Int64Value is null
func GetInt64FromEnv(value basetypes.Int64Value, envVars []string) basetypes.Int64Value {
	if !value.IsNull() {
		return value
	}
	for _, envVar := range envVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			if intValue, err := strconv.ParseInt(envValue, 10, 64); err == nil {
				return basetypes.NewInt64Value(intValue)
			}
		}
	}
	return basetypes.NewInt64Null()
}

// GetBoolFromEnv checks environment variables for a value if the given BoolValue is null
func GetBoolFromEnv(value basetypes.BoolValue, envVars []string) basetypes.BoolValue {
	if !value.IsNull() {
		return value
	}
	for _, envVar := range envVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			if boolValue, err := strconv.ParseBool(envValue); err == nil {
				return basetypes.NewBoolValue(boolValue)
			}
		}
	}
	return basetypes.NewBoolNull()
}
