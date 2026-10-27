// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const resourceInstanceParametersDefaultTimeout = 30 * time.Minute

// Poll timing for async updates; variables so unit tests can shorten them.
var (
	resourceInstanceParametersPollDelay    = 10 * time.Second
	resourceInstanceParametersPollInterval = 10 * time.Second
)

var (
	_ resource.Resource                   = &resourceInstanceParametersResource{}
	_ resource.ResourceWithConfigure      = &resourceInstanceParametersResource{}
	_ resource.ResourceWithValidateConfig = &resourceInstanceParametersResource{}
	_ resource.ResourceWithImportState    = &resourceInstanceParametersResource{}
)

// NewResourceInstanceParametersResource PATCHes
// /v2/resource_instances/{id} with service parameters that can only be known after
// the instance exists (trusted profile, COS bucket), so they cannot be set on
// ibm_resource_instance itself without a dependency cycle.
func NewResourceInstanceParametersResource() resource.Resource {
	return &resourceInstanceParametersResource{}
}

type resourceInstanceParametersResource struct {
	client *rc.ResourceControllerV2
}

type resourceInstanceParametersModel struct {
	ID                 types.String   `tfsdk:"id"`
	ResourceInstanceID types.String   `tfsdk:"resource_instance_id"`
	Parameters         types.String   `tfsdk:"parameters"`
	DestroyParameters  types.String   `tfsdk:"destroy_parameters"`
	State              types.String   `tfsdk:"state"`
	Timeouts           timeouts.Value `tfsdk:"timeouts"`
}

func (r *resourceInstanceParametersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_instance_parameters"
}

func (r *resourceInstanceParametersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Applies service-specific parameters to an existing resource instance through the Resource Controller update API. Use it for parameters that depend on resources created after the instance, such as enabling CSPM or CDR on SCC Workload Protection. Drift is checked only on the top-level keys you set, so server-generated fields never show as changes.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Same as resource_instance_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"resource_instance_id": schema.StringAttribute{
				Required:      true,
				Description:   "GUID or CRN of the resource instance to update.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"parameters": schema.StringAttribute{
				Required:    true,
				Description: "JSON object of service parameters to send, for example jsonencode({ enable_cspm = true, target_accounts = [...] }). Create and update wait until an asynchronous update finishes.",
			},
			"destroy_parameters": schema.StringAttribute{
				Optional:    true,
				Description: "JSON object of parameters to send on destroy, for example jsonencode({ enable_cspm = false }). When omitted, destroy only removes the resource from state, which matches the restapi_object behaviour it replaces. If the instance is already gone (404 or 410), destroy succeeds without sending anything.",
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "State of the resource instance after the last update or refresh.",
			},
		},
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Update: true, Delete: true}),
		},
	}
}

func (r *resourceInstanceParametersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	session, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	client, err := session.ResourceControllerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Resource Controller Client", err.Error())
		return
	}
	r.client = client
}

// ValidateConfig checks the JSON of parameters and destroy_parameters at plan time,
// so a typo in destroy_parameters does not first surface during terraform destroy.
func (r *resourceInstanceParametersResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config resourceInstanceParametersModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for name, value := range map[string]types.String{"parameters": config.Parameters, "destroy_parameters": config.DestroyParameters} {
		if value.IsNull() || value.IsUnknown() {
			continue
		}
		if _, err := parseParametersJSON(value.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(path.Root(name), "Invalid JSON", fmt.Sprintf("%s %s", name, err))
		}
	}
}

// ImportState takes the instance GUID or CRN. Only the parameters set in config after
// import are managed; the first apply re-sends them (the update is idempotent).
func (r *resourceInstanceParametersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("resource_instance_id"), req.ID)...)
}

func (r *resourceInstanceParametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceInstanceParametersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, d := plan.Timeouts.Create(ctx, resourceInstanceParametersDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.patch(ctx, &plan, plan.Parameters.ValueString(), timeout); err != nil {
		resp.Diagnostics.AddError("Unable to Update Resource Instance Parameters", err.Error())
		return
	}
	plan.ID = plan.ResourceInstanceID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceInstanceParametersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resourceInstanceParametersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	instance, response, err := r.client.GetResourceInstanceWithContext(ctx,
		r.client.NewGetResourceInstanceOptions(state.ResourceInstanceID.ValueString()))
	if err != nil {
		if isInstanceGone(response) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to Read Resource Instance", err.Error())
		return
	}
	if instance.State != nil && (*instance.State == RsInstanceRemovedStatus || *instance.State == RsInstanceReclamation) {
		resp.State.RemoveResource(ctx)
		return
	}
	if instance.State != nil {
		state.State = types.StringValue(*instance.State)
	}

	// Compare only the keys the user manages. If the server reports a different
	// value for one of them, surface it so Terraform plans a corrective update.
	var desired map[string]any
	if err := json.Unmarshal([]byte(state.Parameters.ValueString()), &desired); err == nil && instance.Parameters != nil {
		observed := make(map[string]any, len(desired))
		drift := false
		for k, v := range desired {
			got, ok := instance.Parameters[k]
			if !ok {
				// Some services do not echo every parameter back; absence is not drift.
				observed[k] = v
				continue
			}
			if !jsonEqual(v, got) {
				drift = true
			}
			observed[k] = got
		}
		if drift {
			if b, err := json.Marshal(observed); err == nil {
				state.Parameters = types.StringValue(string(b))
			}
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *resourceInstanceParametersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceInstanceParametersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, d := plan.Timeouts.Update(ctx, resourceInstanceParametersDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.patch(ctx, &plan, plan.Parameters.ValueString(), timeout); err != nil {
		resp.Diagnostics.AddError("Unable to Update Resource Instance Parameters", err.Error())
		return
	}
	plan.ID = plan.ResourceInstanceID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceInstanceParametersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceInstanceParametersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.DestroyParameters.IsNull() || state.DestroyParameters.ValueString() == "" {
		return
	}
	timeout, d := state.Timeouts.Delete(ctx, resourceInstanceParametersDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	// The instance may be deleted in the same destroy; there is nothing left to reset.
	if err := r.patch(ctx, &state, state.DestroyParameters.ValueString(), timeout); err != nil && !errors.Is(err, errInstanceGone) {
		resp.Diagnostics.AddError("Unable to Reset Resource Instance Parameters", err.Error())
	}
}

var errInstanceGone = errors.New("resource instance not found")

// isInstanceGone reports a deleted instance: 404 Not Found or 410 Gone.
func isInstanceGone(response *core.DetailedResponse) bool {
	return response != nil && (response.StatusCode == http.StatusNotFound || response.StatusCode == http.StatusGone)
}

func parseParametersJSON(raw string) (map[string]any, error) {
	var params map[string]any
	if err := json.Unmarshal([]byte(raw), &params); err != nil || params == nil {
		if err == nil {
			err = errors.New("got null")
		}
		return nil, fmt.Errorf("must be a JSON object: %w", err)
	}
	return params, nil
}

// patch sends the parameters and, for an asynchronous update, waits until it finishes.
func (r *resourceInstanceParametersResource) patch(ctx context.Context, model *resourceInstanceParametersModel, raw string, timeout time.Duration) error {
	params, err := parseParametersJSON(raw)
	if err != nil {
		return fmt.Errorf("parameters %w", err)
	}
	id := model.ResourceInstanceID.ValueString()
	opts := r.client.NewUpdateResourceInstanceOptions(id)
	opts.Parameters = params
	instance, response, err := r.client.UpdateResourceInstanceWithContext(ctx, opts)
	if err != nil {
		if isInstanceGone(response) {
			return fmt.Errorf("%w: %s: %v", errInstanceGone, id, err)
		}
		return err
	}
	if instance != nil && instance.LastOperation != nil && instance.LastOperation.Async != nil && *instance.LastOperation.Async {
		if instance, err = r.waitForUpdate(ctx, id, timeout); err != nil {
			return err
		}
	}
	model.State = types.StringNull()
	if instance != nil && instance.State != nil {
		model.State = types.StringValue(*instance.State)
	}
	return nil
}

// waitForUpdate polls the instance until its asynchronous last operation leaves
// "in progress" and the instance is active again, like ibm_resource_instance does.
func (r *resourceInstanceParametersResource) waitForUpdate(ctx context.Context, id string, timeout time.Duration) (*rc.ResourceInstance, error) {
	conf := &retry.StateChangeConf{
		Pending:      []string{RsInstanceProgressStatus, RsInstanceInactiveStatus},
		Target:       []string{RsInstanceSuccessStatus, RsInstanceUpdateSuccessStatus},
		Timeout:      timeout,
		Delay:        resourceInstanceParametersPollDelay,
		PollInterval: resourceInstanceParametersPollInterval,
		Refresh: func() (interface{}, string, error) {
			instance, response, err := r.client.GetResourceInstanceWithContext(ctx, r.client.NewGetResourceInstanceOptions(id))
			if err != nil {
				if isInstanceGone(response) {
					return nil, "", fmt.Errorf("%w: %s: %v", errInstanceGone, id, err)
				}
				return nil, "", err
			}
			if op := instance.LastOperation; op != nil && op.State != nil && op.Async != nil && *op.Async {
				if *op.State == RsInstanceFailStatus {
					return instance, *op.State, fmt.Errorf("update of resource instance %s failed: %s", id, flex.StringValue(op.Description))
				}
				// A finished operation still waits for the instance itself to be active.
				if *op.State == RsInstanceProgressStatus || instance.State == nil {
					return instance, *op.State, nil
				}
			}
			if instance.State == nil {
				return instance, RsInstanceProgressStatus, nil
			}
			if *instance.State == RsInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("resource instance %s is in state failed", id)
			}
			return instance, *instance.State, nil
		},
	}
	result, err := conf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}
	return result.(*rc.ResourceInstance), nil
}

// jsonEqual compares two decoded JSON values after a normalising round trip, so
// numeric types and key order do not cause false drift.
func jsonEqual(a, b any) bool {
	na, errA := json.Marshal(a)
	nb, errB := json.Marshal(b)
	if errA != nil || errB != nil {
		return false
	}
	var va, vb any
	_ = json.Unmarshal(na, &va)
	_ = json.Unmarshal(nb, &vb)
	return reflect.DeepEqual(va, vb)
}
