// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry

import (
	"context"
	"fmt"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	crDefaultStorageMegabytes int64 = 500
	crDefaultTrafficMegabytes int64 = 5120
)

var (
	_ resource.Resource              = &crQuotaResource{}
	_ resource.ResourceWithConfigure = &crQuotaResource{}
)

// NewCrQuotaResource manages only the quotas set in config and
// resets them to the service defaults (500 MB storage, 5120 MB traffic) on destroy.
func NewCrQuotaResource() resource.Resource {
	return &crQuotaResource{}
}

type crQuotaResource struct {
	client *containerregistryv1.ContainerRegistryV1
}

type crQuotaModel struct {
	ID               types.String `tfsdk:"id"`
	Endpoint         types.String `tfsdk:"endpoint"`
	StorageMegabytes types.Int64  `tfsdk:"storage_megabytes"`
	TrafficMegabytes types.Int64  `tfsdk:"traffic_megabytes"`
	ResetOnDestroy   types.Bool   `tfsdk:"reset_on_destroy"`
	StorageUsage     types.Int64  `tfsdk:"storage_usage_bytes"`
	TrafficUsage     types.Int64  `tfsdk:"traffic_usage_bytes"`
}

func (r *crQuotaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cr_quota"
}

func (r *crQuotaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages the account-level storage and pull traffic quotas of IBM Cloud Container Registry in one registry region. Only the quotas you set are managed; the other one is left alone.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The account ID and registry endpoint the quota belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"endpoint": schema.StringAttribute{
				Optional:      true,
				Description:   "Regional Container Registry API endpoint, for example 'https://us.icr.io'. Defaults to the endpoint the provider is configured with.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"storage_megabytes": schema.Int64Attribute{
				Optional:    true,
				Description: "Storage quota in megabytes. -1 means unlimited. Not managed when omitted.",
			},
			"traffic_megabytes": schema.Int64Attribute{
				Optional:    true,
				Description: "Pull traffic quota in megabytes per month. -1 means unlimited. Not managed when omitted.",
			},
			"reset_on_destroy": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "When true, destroy sets the managed quotas back to the service defaults (500 MB storage, 5120 MB traffic). Default: true.",
			},
			"storage_usage_bytes": schema.Int64Attribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				Description:   "Current storage usage in bytes, as reported at the last refresh.",
			},
			"traffic_usage_bytes": schema.Int64Attribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				Description:   "Current pull traffic usage in bytes, as reported at the last refresh.",
			},
		},
	}
}

func (r *crQuotaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, err := crClientFromProviderData(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Container Registry Client", err.Error())
		return
	}
	r.client = client
}

func (r *crQuotaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan crQuotaModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.StorageMegabytes.IsNull() && plan.TrafficMegabytes.IsNull() {
		resp.Diagnostics.AddError("Missing Quota", "Set at least one of storage_megabytes or traffic_megabytes.")
		return
	}

	client, err := crClientForEndpoint(r.client, plan.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Container Registry Endpoint", err.Error())
		return
	}
	if err := r.apply(ctx, client, plan.StorageMegabytes, plan.TrafficMegabytes); err != nil {
		resp.Diagnostics.AddError("Unable to Update Container Registry Quota", err.Error())
		return
	}

	account := ""
	if client.Account != nil {
		account = *client.Account
	}
	plan.ID = types.StringValue(fmt.Sprintf("%s/%s", account, client.Service.GetServiceURL()))
	r.refreshUsage(ctx, client, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *crQuotaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state crQuotaModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, err := crClientForEndpoint(r.client, state.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Container Registry Endpoint", err.Error())
		return
	}

	quota, _, err := client.GetQuotaWithContext(ctx, client.NewGetQuotaOptions())
	if err != nil {
		resp.Diagnostics.AddError("Unable to Read Container Registry Quota", err.Error())
		return
	}
	// Only managed quotas are refreshed, so a quota set outside Terraform for the
	// other dimension never shows as drift.
	if quota != nil && quota.Limit != nil {
		if !state.StorageMegabytes.IsNull() && quota.Limit.StorageBytes != nil {
			state.StorageMegabytes = types.Int64Value(bytesToMegabytes(*quota.Limit.StorageBytes))
		}
		if !state.TrafficMegabytes.IsNull() && quota.Limit.TrafficBytes != nil {
			state.TrafficMegabytes = types.Int64Value(bytesToMegabytes(*quota.Limit.TrafficBytes))
		}
	}
	setUsage(quota, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *crQuotaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state crQuotaModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, err := crClientForEndpoint(r.client, plan.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Container Registry Endpoint", err.Error())
		return
	}

	storage, traffic := plan.StorageMegabytes, plan.TrafficMegabytes
	// A quota that was managed and is now removed from config goes back to its default.
	if storage.IsNull() && !state.StorageMegabytes.IsNull() && plan.ResetOnDestroy.ValueBool() {
		storage = types.Int64Value(crDefaultStorageMegabytes)
	}
	if traffic.IsNull() && !state.TrafficMegabytes.IsNull() && plan.ResetOnDestroy.ValueBool() {
		traffic = types.Int64Value(crDefaultTrafficMegabytes)
	}
	if err := r.apply(ctx, client, storage, traffic); err != nil {
		resp.Diagnostics.AddError("Unable to Update Container Registry Quota", err.Error())
		return
	}

	plan.ID = state.ID
	r.refreshUsage(ctx, client, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *crQuotaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state crQuotaModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !state.ResetOnDestroy.ValueBool() {
		return
	}
	client, err := crClientForEndpoint(r.client, state.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Container Registry Endpoint", err.Error())
		return
	}

	storage, traffic := types.Int64Null(), types.Int64Null()
	if !state.StorageMegabytes.IsNull() {
		storage = types.Int64Value(crDefaultStorageMegabytes)
	}
	if !state.TrafficMegabytes.IsNull() {
		traffic = types.Int64Value(crDefaultTrafficMegabytes)
	}
	if err := r.apply(ctx, client, storage, traffic); err != nil {
		resp.Diagnostics.AddError("Unable to Reset Container Registry Quota", err.Error())
	}
}

func (r *crQuotaResource) apply(ctx context.Context, client *containerregistryv1.ContainerRegistryV1, storage, traffic types.Int64) error {
	if storage.IsNull() && traffic.IsNull() {
		return nil
	}
	opts := client.NewUpdateQuotaOptions()
	if !storage.IsNull() {
		opts.SetStorageMegabytes(storage.ValueInt64())
	}
	if !traffic.IsNull() {
		opts.SetTrafficMegabytes(traffic.ValueInt64())
	}
	_, err := client.UpdateQuotaWithContext(ctx, opts)
	return err
}

func (r *crQuotaResource) refreshUsage(ctx context.Context, client *containerregistryv1.ContainerRegistryV1, model *crQuotaModel) {
	quota, _, err := client.GetQuotaWithContext(ctx, client.NewGetQuotaOptions())
	if err != nil {
		setUsage(nil, model)
		// Usage is informational only; a failed read must not fail the apply.
		return
	}
	setUsage(quota, model)
}

func setUsage(quota *containerregistryv1.Quota, model *crQuotaModel) {
	// Computed values must never be left unknown after apply.
	if model.StorageUsage.IsUnknown() {
		model.StorageUsage = types.Int64Value(0)
	}
	if model.TrafficUsage.IsUnknown() {
		model.TrafficUsage = types.Int64Value(0)
	}
	if quota == nil || quota.Usage == nil {
		return
	}
	if quota.Usage.StorageBytes != nil {
		model.StorageUsage = types.Int64Value(*quota.Usage.StorageBytes)
	}
	if quota.Usage.TrafficBytes != nil {
		model.TrafficUsage = types.Int64Value(*quota.Usage.TrafficBytes)
	}
}

// bytesToMegabytes converts an API limit back to the unit users configure.
// The API reports -1 for unlimited, which is kept as is.
func bytesToMegabytes(b int64) int64 {
	if b < 0 {
		return b
	}
	return b / (1024 * 1024)
}
