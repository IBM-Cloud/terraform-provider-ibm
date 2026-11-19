// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const (
	vpcWorkerStatusUndeploying     = "undeploying"
	vpcWorkerStatusUndeployed      = "undeployed"
	vpcWorkerStatusReloadPending   = "reload_pending"
	vpcWorkerStatusReloading       = "reloading"
	vpcWorkerStatusReloaded        = "reloaded"
	vpcWorkerStatusDeploying       = "deploying"
	vpcWorkerStatusDeployed        = "deployed"
	vpcWorkerStatusDeployFailed    = "deploy_failed"
	vpcWorkerStatusReloadingFailed = "reloading_failed"
)

var (
	_ action.Action              = &containerVpcWorkerReloadAction{}
	_ action.ActionWithConfigure = &containerVpcWorkerReloadAction{}
)

func NewContainerVpcWorkerReloadAction() action.Action {
	return &containerVpcWorkerReloadAction{}
}

type containerVpcWorkerReloadAction struct {
	workerClient    containerv1.Workers
	vpcWorkerClient containerv2.Workers
}

type vpcWorkerReloadModel struct {
	ClusterNameID types.String `tfsdk:"cluster_name_id"`
	WorkerID      types.String `tfsdk:"worker_id"`
	Timeout       types.String `tfsdk:"timeout"`
	NoWait        types.Bool   `tfsdk:"no_wait"`
}

func (a *containerVpcWorkerReloadAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "ibm_container_vpc_worker_reload"
}

func (a *containerVpcWorkerReloadAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reloads a worker node (virtual server or bare metal) in a VPC cluster. This action triggers a reload operation for the specified worker node. Actions do not return output values.",
		Attributes: map[string]schema.Attribute{
			"cluster_name_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID or name of the VPC cluster containing the worker node.",
			},
			"worker_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the worker node to reload.",
			},
			"timeout": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum time to wait for the worker reload to complete, for example `30m` or `1h`. If not specified, defaults to `45m`. Ignored when no_wait is true.",
				Validators: []validator.String{
					ValidDuration(),
				},
			},
			"no_wait": schema.BoolAttribute{
				Optional:    true,
				Description: "If true, the action returns immediately after creating the reload request without waiting for completion. Default: false",
			},
		},
	}
}

func (a *containerVpcWorkerReloadAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	session, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. The provider client session could not be established.", req.ProviderData),
		)
		return
	}

	client, err := session.ContainerAPI()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Container Client",
			"An unexpected error occurred when creating Container client.\n\n"+
				"Container Client Error: "+err.Error(),
		)
		return
	}

	a.workerClient = client.Workers()

	vpcClient, err := session.VpcContainerAPI()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create VPC Container Client",
			"An unexpected error occurred when creating the VPC Container client.\n\n"+
				"VPC Container Client Error: "+err.Error(),
		)
		return
	}

	a.vpcWorkerClient = vpcClient.Workers()
}

func (a *containerVpcWorkerReloadAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config vpcWorkerReloadModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterNameID := config.ClusterNameID.ValueString()
	workerID := config.WorkerID.ValueString()

	// Parse timeout duration, default to 45 minutes
	timeout := 45 * time.Minute
	if !config.Timeout.IsNull() {
		timeout, _ = time.ParseDuration(config.Timeout.ValueString())
	}

	noWait := false
	if !config.NoWait.IsNull() {
		noWait = config.NoWait.ValueBool()
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Reloading worker '%s' in cluster '%s'...", workerID, clusterNameID),
	})

	// Send reload request
	params := containerv1.WorkerUpdateParam{
		Action: "reload",
	}
	err := a.workerClient.Update("", workerID, params, containerv1.ClusterTargetHeader{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initiate Worker Reload",
			fmt.Sprintf("Failed to reload worker '%s' in cluster '%s': %s", workerID, clusterNameID, err.Error()),
		)
		return
	}

	// Return immediately if no_wait set to true
	if noWait {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Worker '%s' reload submitted (no-wait mode)", workerID),
		})
		return
	}

	// Otherwise, wait for the reload to complete
	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Waiting for worker '%s' reload to complete (timeout: %v)...", workerID, timeout),
	})

	_, waitErr := waitForVpcWorkerReloadAvailable(ctx, a.vpcWorkerClient, workerID, clusterNameID, timeout)
	if waitErr != nil {
		resp.Diagnostics.AddError(
			"Worker Reload Failed",
			fmt.Sprintf("Failed waiting for worker '%s' reload in cluster '%s': %s", workerID, clusterNameID, waitErr.Error()),
		)
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Worker '%s' reload completed successfully in cluster '%s'", workerID, clusterNameID),
	})
}

func waitForVpcWorkerReloadAvailable(ctx context.Context, vpcWorkerClient containerv2.Workers, workerID string, clusterNameID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for worker node to be deployed.")
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			vpcWorkerStatusUndeploying,
			vpcWorkerStatusUndeployed,
			vpcWorkerStatusReloadPending,
			vpcWorkerStatusReloading,
			vpcWorkerStatusReloaded,
			vpcWorkerStatusDeploying,
		},
		Target:     []string{vpcWorkerStatusDeployed, vpcWorkerStatusDeployFailed, vpcWorkerStatusReloadingFailed},
		Refresh:    vpcWorkerReloadRefreshFunc(vpcWorkerClient, workerID, clusterNameID),
		Timeout:    timeout,
		Delay:      20 * time.Second, // 20 seconds delay before the check start
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func vpcWorkerReloadRefreshFunc(vpcWorkerClient containerv2.Workers, workerID string, clusterNameID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		worker, err := vpcWorkerClient.Get(clusterNameID, workerID, containerv2.ClusterTargetHeader{})
		if err != nil {
			return nil, "", fmt.Errorf("error getting worker node: %s", err)
		}
		switch worker.LifeCycle.ActualState {
		case "reloading_failed":
			return worker, worker.LifeCycle.ActualState, fmt.Errorf("worker node is in reloading_failed state")

		case "deploy_failed":
			return worker, worker.LifeCycle.ActualState, fmt.Errorf("worker node is in deploy_failed state")
		}

		return worker, worker.LifeCycle.ActualState, nil
	}
}

// Duration Validator
var _ validator.String = durationValidator{}

type durationValidator struct{}

func (v durationValidator) Description(ctx context.Context) string {
	return "string must be a valid duration format (e.g., '30m', '1h')"
}

func (v durationValidator) MarkdownDescription(ctx context.Context) string {
	return "string must be a valid duration format (e.g., `30m`, `1h`)"
}

func (v durationValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() {
		return
	}

	val := req.ConfigValue.ValueString()
	if _, err := time.ParseDuration(val); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Timeout Format",
			fmt.Sprintf("Failed to parse timeout '%s': %s. Expected format like '30m' or '1h'.", val, err.Error()),
		)
	}
}

func ValidDuration() validator.String {
	return durationValidator{}
}
