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
	bareMetalWorkerStatusUndeploying     = "undeploying"
	bareMetalWorkerStatusUndeployed      = "undeployed"
	bareMetalWorkerStatusReloadPending   = "reload_pending"
	bareMetalWorkerStatusReloading       = "reloading"
	bareMetalWorkerStatusReloaded        = "reloaded"
	bareMetalWorkerStatusDeploying       = "deploying"
	bareMetalWorkerStatusDeployed        = "deployed"
	bareMetalWorkerStatusDeployFailed    = "deploy_failed"
	bareMetalWorkerStatusReloadingFailed = "reloading_failed"
)

var (
	_ action.Action              = &containerVpcBareMetalWorkerReloadAction{}
	_ action.ActionWithConfigure = &containerVpcBareMetalWorkerReloadAction{}
)

func NewContainerVpcBareMetalWorkerReloadAction() action.Action {
	return &containerVpcBareMetalWorkerReloadAction{}
}

type containerVpcBareMetalWorkerReloadAction struct {
	workerClient    containerv1.Workers
	vpcWorkerClient containerv2.Workers
}

type workerReloadModel struct {
	ClusterNameID     types.String `tfsdk:"cluster_name_id"`
	BareMetalServerID types.String `tfsdk:"bare_metal_server_id"`
	Timeout           types.String `tfsdk:"timeout"`
	NoWait            types.Bool   `tfsdk:"no_wait"`
}

func (a *containerVpcBareMetalWorkerReloadAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "ibm_container_vpc_bare_metal_worker_reload"
}

func (a *containerVpcBareMetalWorkerReloadAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reloads a bare metal worker node in a VPC cluster. This action triggers a reload operation for the specified worker node. Actions do not return output values.",
		Attributes: map[string]schema.Attribute{
			"cluster_name_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID or name of the VPC cluster containing the bare metal worker node.",
			},
			"bare_metal_server_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the bare metal server.",
			},
			"timeout": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum time to wait for the bare metal worker reload to complete, for example `30m` or `1h`. If not specified, defaults to `45m`. Ignored when no_wait is true.",
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

func (a *containerVpcBareMetalWorkerReloadAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *containerVpcBareMetalWorkerReloadAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config workerReloadModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterNameID := config.ClusterNameID.ValueString()
	bareMetalServerId := config.BareMetalServerID.ValueString()

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
		Message: fmt.Sprintf("Reloading bare metal server '%s' in cluster '%s'...", bareMetalServerId, clusterNameID),
	})

	// Send reload request
	params := containerv1.WorkerUpdateParam{
		Action: "reload",
	}
	err := a.workerClient.Update("", bareMetalServerId, params, containerv1.ClusterTargetHeader{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initiate Bare Metal Worker Reload",
			fmt.Sprintf("Failed to reload bare metal server '%s' in cluster '%s': %s", bareMetalServerId, clusterNameID, err.Error()),
		)
		return
	}

	// 	Return immediately if no_wait set to true
	if noWait {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Bare metal server '%s' reload submitted (no-wait mode)", bareMetalServerId),
		})
		return
	}

	// Otherwise, wait for the reload to complete
	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Waiting for bare metal server '%s' reload to complete (timeout: %v)...", bareMetalServerId, timeout),
	})

	_, waitErr := waitForBareMetalWorkerReloadAvailable(ctx, a.vpcWorkerClient, bareMetalServerId, clusterNameID, timeout)
	if waitErr != nil {
		resp.Diagnostics.AddError(
			"Bare Metal Worker Reload Failed",
			fmt.Sprintf("Failed waiting for bare metal server '%s' reload in cluster '%s': %s", bareMetalServerId, clusterNameID, waitErr.Error()),
		)
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Bare metal server '%s' reload completed successfully in cluster '%s'", bareMetalServerId, clusterNameID),
	})
}

func waitForBareMetalWorkerReloadAvailable(ctx context.Context, vpcWorkerClient containerv2.Workers, bareMetalServerId string, clusterNameId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for bare metal worker node to be deployed.")
	stateConf := &retry.StateChangeConf{

		Pending: []string{bareMetalWorkerStatusUndeploying, bareMetalWorkerStatusUndeployed,
			bareMetalWorkerStatusReloadPending, bareMetalWorkerStatusReloading, bareMetalWorkerStatusReloaded, bareMetalWorkerStatusDeploying},
		Target:     []string{bareMetalWorkerStatusDeployed, bareMetalWorkerStatusDeployFailed, bareMetalWorkerStatusReloadingFailed},
		Refresh:    bareMetalWorkerReloadRefreshFunc(vpcWorkerClient, bareMetalServerId, clusterNameId),
		Timeout:    timeout,
		Delay:      20 * time.Second, // 20 seconds delay before the check start
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func bareMetalWorkerReloadRefreshFunc(vpcWorkerClient containerv2.Workers, bareMetalServerId string, clusterNameId string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		worker, err := vpcWorkerClient.Get(clusterNameId, bareMetalServerId, containerv2.ClusterTargetHeader{})
		if err != nil {
			return nil, "", fmt.Errorf("error getting bare metal server worker node: %s", err)
		}
		switch worker.LifeCycle.ActualState {
		case "reloading_failed":
			return worker, worker.LifeCycle.ActualState, fmt.Errorf("bare metal server worker node is in reloading_failed state")

		case "deploy_failed":
			return worker, worker.LifeCycle.ActualState, fmt.Errorf("bare metal server worker node is in deploy_failed state")
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
