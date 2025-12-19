// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ action.Action              = &codeEngineBuildRunAction{}
	_ action.ActionWithConfigure = &codeEngineBuildRunAction{}
)

func NewCodeEngineBuildRunAction() action.Action {
	return &codeEngineBuildRunAction{}
}

type codeEngineBuildRunAction struct {
	client *codeenginev2.CodeEngineV2
}

type buildRunModel struct {
	ProjectID types.String `tfsdk:"project_id"`
	BuildName types.String `tfsdk:"build_name"`
	Timeout   types.Int64  `tfsdk:"timeout"`
}

func (a *codeEngineBuildRunAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "ibm_code_engine_build_run"
}

func (a *codeEngineBuildRunAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Triggers a Code Engine build run and waits for completion. This action initiates a build run for an existing Code Engine build configuration, polls for completion status, and reports success or failure via diagnostics. Actions do not return output values.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the Code Engine project containing the build configuration.",
			},
			"build_name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the Code Engine build configuration to execute. The build must be in 'ready' state.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "Build run timeout in minutes. If not specified, defaults to 30 minutes. Minimum: 5 minutes",
				Validators: []validator.Int64{
					int64validator.AtLeast(5),
				},
			},
		},
	}
}

func (a *codeEngineBuildRunAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	// Cast provider data to ClientSession (same pattern as SDKv2 resources)
	session, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	// Get the Code Engine client from the session
	client, err := session.CodeEngineV2()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Code Engine Client",
			"An unexpected error occurred when creating the Code Engine client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Code Engine Client Error: "+err.Error(),
		)
		return
	}

	a.client = client
}

func (a *codeEngineBuildRunAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config buildRunModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout := 30 * time.Minute
	if !config.Timeout.IsNull() {
		timeout = time.Duration(config.Timeout.ValueInt64()) * time.Minute
	}

	projectID := config.ProjectID.ValueString()
	buildName := config.BuildName.ValueString()

	createOptions := &codeenginev2.CreateBuildRunOptions{
		ProjectID: core.StringPtr(projectID),
		BuildName: core.StringPtr(buildName),
	}

	buildRun, response, err := a.client.CreateBuildRunWithContext(ctx, createOptions)
	if err != nil {
		a.handleCreateError(response, err, resp)
		return
	}

	buildRunName := *buildRun.Name

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Build run '%s' created, waiting for completion...", buildRunName),
	})

	finalBuildRun, err := a.waitForCompletion(ctx, projectID, buildRunName, timeout, resp.SendProgress)
	if err != nil {
		resp.Diagnostics.AddError(
			"Build Run Failed",
			fmt.Sprintf("Build run '%s' did not complete successfully: %s", buildRunName, err.Error()),
		)
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Build run '%s' completed successfully", buildRunName),
	})

	if finalBuildRun.OutputImage != nil {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Container image: %s", *finalBuildRun.OutputImage),
		})
	}
}

func (a *codeEngineBuildRunAction) waitForCompletion(ctx context.Context, projectID, buildRunName string, timeout time.Duration, sendProgress func(action.InvokeProgressEvent)) (*codeenginev2.BuildRun, error) {
	deadline := time.Now().Add(timeout)
	pollInterval := 10 * time.Second
	maxInterval := 30 * time.Second
	backoffMultiplier := 1.5
	lastStatus := ""

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("operation cancelled: %w", ctx.Err())
		default:
		}

		getOptions := &codeenginev2.GetBuildRunOptions{
			ProjectID: core.StringPtr(projectID),
			Name:      core.StringPtr(buildRunName),
		}

		buildRun, response, err := a.client.GetBuildRunWithContext(ctx, getOptions)
		if err != nil {
			if a.isRetryableError(response, err) {
				time.Sleep(pollInterval)
				continue
			}
			return nil, fmt.Errorf("failed to get build run status: %w", err)
		}

		if buildRun.Status != nil {
			currentStatus := *buildRun.Status
			if currentStatus != lastStatus {
				sendProgress(action.InvokeProgressEvent{
					Message: fmt.Sprintf("Build run status: %s", currentStatus),
				})
				lastStatus = currentStatus
			}

			switch currentStatus {
			case "succeeded":
				return buildRun, nil
			case "failed":
				reason := "unknown error"
				if buildRun.StatusDetails != nil && buildRun.StatusDetails.Reason != nil {
					reason = *buildRun.StatusDetails.Reason
				}
				return nil, fmt.Errorf("build failed: %s", reason)
			case "pending", "running":
			default:
				return nil, fmt.Errorf("unknown build run status: %s", currentStatus)
			}
		}

		time.Sleep(pollInterval)
		pollInterval = time.Duration(float64(pollInterval) * backoffMultiplier)
		if pollInterval > maxInterval {
			pollInterval = maxInterval
		}
	}

	return nil, fmt.Errorf("timeout after %v waiting for build run completion", timeout)
}

func (a *codeEngineBuildRunAction) isRetryableError(response *core.DetailedResponse, err error) bool {
	if response == nil {
		return true
	}

	statusCode := response.StatusCode
	return statusCode == 429 ||
		statusCode == 500 ||
		statusCode == 502 ||
		statusCode == 503 ||
		statusCode == 504
}

func (a *codeEngineBuildRunAction) handleCreateError(response *core.DetailedResponse, err error, resp *action.InvokeResponse) {
	if response == nil {
		resp.Diagnostics.AddError(
			"Network Error",
			fmt.Sprintf("Failed to connect to Code Engine API: %s", err.Error()),
		)
		return
	}

	statusCode := response.StatusCode
	switch statusCode {
	case 400:
		resp.Diagnostics.AddError(
			"Invalid Request",
			fmt.Sprintf("The request to create a build run was invalid. Please verify the project_id and build_name are correct. Error: %s", err.Error()),
		)
	case 401:
		resp.Diagnostics.AddError(
			"Authentication Failed",
			fmt.Sprintf("Authentication with IBM Cloud failed. Please verify your API key or credentials are valid and not expired. Error: %s", err.Error()),
		)
	case 403:
		resp.Diagnostics.AddError(
			"Authorization Failed",
			fmt.Sprintf("You do not have permission to create build runs in this project. Please verify you have the 'Editor' role or higher in the resource group. Error: %s", err.Error()),
		)
	case 404:
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("The specified project or build configuration was not found. Please verify the project_id and build_name are correct. Error: %s", err.Error()),
		)
	case 409:
		resp.Diagnostics.AddError(
			"Conflict",
			fmt.Sprintf("Unable to create build run due to a conflict. This may occur if the build is not in 'ready' state or if maximum concurrent builds are reached. Error: %s", err.Error()),
		)
	case 429:
		resp.Diagnostics.AddError(
			"Rate Limit Exceeded",
			fmt.Sprintf("Too many requests to Code Engine API. Please wait a moment and try again. Error: %s", err.Error()),
		)
	case 500, 502, 503, 504:
		resp.Diagnostics.AddError(
			"Service Error",
			fmt.Sprintf("Code Engine service is temporarily unavailable (HTTP %d). Please try again in a few moments. Error: %s", statusCode, err.Error()),
		)
	default:
		resp.Diagnostics.AddError(
			"Build Run Creation Failed",
			fmt.Sprintf("Failed to create build run (HTTP %d): %s", statusCode, err.Error()),
		)
	}
}
