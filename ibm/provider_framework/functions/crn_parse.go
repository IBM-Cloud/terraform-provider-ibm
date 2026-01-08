// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure CRNServiceNameFunction satisfies the function.Function interface.
var _ function.Function = &CRNServiceNameFunction{}

// CRNServiceNameFunction defines the function implementation for extracting service name from IBM Cloud CRNs.
type CRNServiceNameFunction struct{}

// NewCRNServiceNameFunction returns a new instance of the CRN service name extraction function.
func NewCRNServiceNameFunction() function.Function {
	return &CRNServiceNameFunction{}
}

// Metadata sets the function name.
func (f *CRNServiceNameFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "crn_service_name"
}

// Definition defines the function's parameters and return type.
func (f *CRNServiceNameFunction) Definition(_ context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Extract service name from an IBM Cloud CRN",
		MarkdownDescription: "Given an IBM Cloud CRN string, extracts and returns the service name component. " +
			"CRN format: crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "crn",
				MarkdownDescription: "The IBM Cloud CRN string to parse",
			},
		},
		Return: function.StringReturn{},
	}
}

// Run executes the function logic.
func (f *CRNServiceNameFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var crn string

	// Get the CRN argument
	resp.Error = req.Arguments.Get(ctx, &crn)
	if resp.Error != nil {
		return
	}

	// Validate CRN format
	if !strings.HasPrefix(crn, "crn:") {
		resp.Error = function.NewFuncError("Invalid CRN: must start with 'crn:'")
		return
	}

	// Parse CRN components
	parts := strings.Split(crn, ":")
	if len(parts) < 10 {
		resp.Error = function.NewFuncError("Invalid CRN format: expected at least 10 parts separated by colons")
		return
	}

	// Extract service name (5th component, index 4)
	serviceName := parts[4]

	// Set the result
	resp.Error = resp.Result.Set(ctx, serviceName)
}
