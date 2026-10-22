// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &resourceInstanceParametersResource{}
	_ resource.ResourceWithConfigure = &resourceInstanceParametersResource{}
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
	ID                 types.String `tfsdk:"id"`
	ResourceInstanceID types.String `tfsdk:"resource_instance_id"`
	Parameters         types.String `tfsdk:"parameters"`
	DestroyParameters  types.String `tfsdk:"destroy_parameters"`
	State              types.String `tfsdk:"state"`
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
				Description: "JSON object of service parameters to send, for example jsonencode({ enable_cspm = true, target_accounts = [...] }).",
			},
			"destroy_parameters": schema.StringAttribute{
				Optional:    true,
				Description: "JSON object of parameters to send on destroy, for example jsonencode({ enable_cspm = false }). When omitted, destroy only removes the resource from state, which matches the restapi_object behaviour it replaces.",
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "State of the resource instance after the last update or refresh.",
			},
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

func (r *resourceInstanceParametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceInstanceParametersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.patch(ctx, &plan, plan.Parameters.ValueString()); err != nil {
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
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to Read Resource Instance", err.Error())
		return
	}
	if instance.State != nil && (*instance.State == "removed" || *instance.State == "pending_reclamation") {
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
	if err := r.patch(ctx, &plan, plan.Parameters.ValueString()); err != nil {
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
	if err := r.patch(ctx, &state, state.DestroyParameters.ValueString()); err != nil {
		resp.Diagnostics.AddError("Unable to Reset Resource Instance Parameters", err.Error())
	}
}

func (r *resourceInstanceParametersResource) patch(ctx context.Context, model *resourceInstanceParametersModel, raw string) error {
	var params map[string]any
	if err := json.Unmarshal([]byte(raw), &params); err != nil {
		return fmt.Errorf("parameters must be a JSON object: %w", err)
	}
	opts := r.client.NewUpdateResourceInstanceOptions(model.ResourceInstanceID.ValueString())
	opts.Parameters = params
	instance, _, err := r.client.UpdateResourceInstanceWithContext(ctx, opts)
	if err != nil {
		return err
	}
	model.State = types.StringNull()
	if instance != nil && instance.State != nil {
		model.State = types.StringValue(*instance.State)
	}
	return nil
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
