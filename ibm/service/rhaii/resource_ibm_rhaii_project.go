// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	rhaiiProjectResourceName = "ibm_rhaii_project"
	rhaiiDefaultTimeout      = 10 * time.Minute
)

var (
	_ resource.Resource                   = &rhaiiProjectResource{}
	_ resource.ResourceWithConfigure      = &rhaiiProjectResource{}
	_ resource.ResourceWithImportState    = &rhaiiProjectResource{}
	_ resource.ResourceWithModifyPlan     = &rhaiiProjectResource{}
	_ resource.ResourceWithValidateConfig = &rhaiiProjectResource{}
)

var (
	// Resource controller rule: 180 characters or less, no special characters
	// other than space - . _ :
	rhaiiNameRegexp = regexp.MustCompile(`^[\p{L}\p{N} ._:-]+$`)
	// User tag rule of the IBM Cloud tagging documentation.
	rhaiiTagRegexp = regexp.MustCompile(`^[A-Za-z0-9:_ .-]+$`)
	// Access tags are key:value.
	rhaiiAccessTagRegexp = regexp.MustCompile(`^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`)
)

// NewRhaiiProjectResource returns the ibm_rhaii_project resource. A project is a
// resource controller instance of the "instructlab" service. The resource talks
// to the Global Catalog, Resource Controller and Global Tagging APIs directly.
func NewRhaiiProjectResource() resource.Resource {
	return &rhaiiProjectResource{}
}

type rhaiiProjectResource struct {
	meta conns.ClientSession
}

type rhaiiProjectResourceModel struct {
	Name            types.String   `tfsdk:"name"`
	Location        types.String   `tfsdk:"location"`
	PlanName        types.String   `tfsdk:"plan_name"`
	ResourceGroupID types.String   `tfsdk:"resource_group_id"`
	Tags            types.Set      `tfsdk:"tags"`
	AccessTags      types.Set      `tfsdk:"access_tags"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
	rhaiiProjectComputedModel
}

func (r *rhaiiProjectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rhaii_project"
}

func (r *rhaiiProjectResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required:    true,
			Description: "The name of the project.",
			Validators: []validator.String{
				stringvalidator.LengthBetween(1, 180),
				stringvalidator.RegexMatches(rhaiiNameRegexp, "can contain only letters, numbers, spaces and - . _ :"),
			},
		},
		"location": schema.StringAttribute{
			Optional:      true,
			Computed:      true,
			Default:       stringdefault.StaticString(rhaiiDefaultLocation),
			Description:   "The region where the project is created.",
			PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		},
		// The name, not the ID, like SERVICE_PLAN_NAME in the CLI. The ID is the
		// computed plan_id. The catalog marks the service with
		// plan_updateable=false, so a plan change replaces the project.
		"plan_name": schema.StringAttribute{
			Optional:      true,
			Computed:      true,
			Default:       stringdefault.StaticString(rhaiiDefaultPlan),
			Description:   "The name of the pricing plan of the project.",
			PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		},
		"resource_group_id": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The ID of the resource group of the project. The default resource group of the account is used when not set.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		// tags and access_tags are Optional + Computed, as in ibm_is_ssh_key and the
		// other resources of the provider: an unset attribute keeps the tags of the
		// server, and an empty list ([]) detaches all of them.
		"tags": schema.SetAttribute{
			ElementType:   types.StringType,
			Optional:      true,
			Computed:      true,
			Description:   "The user tags of the project.",
			PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			Validators: []validator.Set{setvalidator.ValueStringsAre(
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(rhaiiTagRegexp, "can contain only letters, numbers, spaces and _ . - :"),
			)},
		},
		"access_tags": schema.SetAttribute{
			ElementType:   types.StringType,
			Optional:      true,
			Computed:      true,
			Description:   "The access management tags of the project. The tags must already exist in the account.",
			PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			Validators: []validator.Set{setvalidator.ValueStringsAre(
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(rhaiiAccessTagRegexp, "must be in key:value format"),
			)},
		},
	}
	for k, v := range rhaiiProjectResourceComputedAttributes() {
		attributes[k] = v
	}

	resp.Schema = schema.Schema{
		Description: "Manages a Red Hat AI Inference project.",
		Attributes:  attributes,
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Update: true, Delete: true}),
		},
	}
}

func (r *rhaiiProjectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	meta, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected conns.ClientSession, got %T.", req.ProviderData))
		return
	}
	r.meta = meta
}

// ValidateConfig has no checks today. Value rules live in the schema validators.
func (r *rhaiiProjectResource) ValidateConfig(context.Context, resource.ValidateConfigRequest, *resource.ValidateConfigResponse) {
}

// ModifyPlan does what the SDKv2 resources do in CustomizeDiff:
//   - adds the IC_ENV_TAGS tags of Schematics to the planned tags, because the
//     flex tag helpers always attach them (no endless diff)
//   - fails the plan when an access tag does not exist in the account
func (r *rhaiiProjectResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return // destroy
	}

	var config, plan rhaiiProjectResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if envTags := rhaiiEnvTags(); len(envTags) > 0 && !config.Tags.IsNull() && !config.Tags.IsUnknown() {
		tags, d := rhaiiSetToStrings(ctx, config.Tags)
		resp.Diagnostics.Append(d...)
		merged, d := types.SetValueFrom(ctx, types.StringType, rhaiiMergeTags(tags, envTags))
		resp.Diagnostics.Append(d...)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("tags"), merged)...)
	}

	if r.meta != nil && !config.AccessTags.IsNull() && !config.AccessTags.IsUnknown() {
		accessTags, d := rhaiiSetToStrings(ctx, config.AccessTags)
		resp.Diagnostics.Append(d...)
		missing, err := rhaiiMissingAccessTags(ctx, r.meta, accessTags)
		if err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("access_tags"), "Error checking access tags", err.Error())
		} else if len(missing) > 0 {
			resp.Diagnostics.AddAttributeError(path.Root("access_tags"), "Access tags do not exist",
				fmt.Sprintf("Access tag(s) %s do not exist in the account. Create them first, for example with ibm_iam_access_tag.", strings.Join(missing, ", ")))
		}
	}
}

func (r *rhaiiProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan rhaiiProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	createTimeout, d := plan.Timeouts.Create(ctx, rhaiiDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	rcClient, err := r.meta.ResourceControllerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Resource Controller client", err.Error())
		return
	}
	gcClient, err := r.meta.GlobalCatalogV1API()
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Global Catalog client", err.Error())
		return
	}

	planID, err := rhaiiResolvePlanID(ctx, gcClient, plan.PlanName.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("plan_name"), "Error resolving the plan", err.Error())
		return
	}
	targetCRN, err := rhaiiResolveTargetCRN(ctx, gcClient, planID, plan.Location.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("location"), "Error resolving the deployment", err.Error())
		return
	}

	resourceGroupID := plan.ResourceGroupID.ValueString()
	if plan.ResourceGroupID.IsNull() || plan.ResourceGroupID.IsUnknown() || resourceGroupID == "" {
		resourceGroupID, err = flex.DefaultResourceGroup(r.meta)
		if err != nil {
			resp.Diagnostics.AddError("Error getting the default resource group", err.Error())
			return
		}
	}

	instance, _, err := rcClient.CreateResourceInstanceWithContext(ctx, &rc.CreateResourceInstanceOptions{
		Name:           ptr(plan.Name.ValueString()),
		Target:         &targetCRN,
		ResourceGroup:  &resourceGroupID,
		ResourcePlanID: &planID,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating the Red Hat AI Inference project", fmt.Sprintf("CreateResourceInstanceWithContext failed: %s", err))
		return
	}
	id := *instance.ID

	// Save the ID right away, so that a failed wait leaves a tainted project in
	// the state instead of an orphan in the account.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)

	if err := waitForRhaiiProject(ctx, rcClient, id, rhaiiWaitCreating, createTimeout, rhaiiCreateRefresh); err != nil {
		resp.Diagnostics.AddError("Error waiting for the Red Hat AI Inference project to be active", fmt.Sprintf("project %s: %s", id, err))
		return
	}

	// Tag errors are warnings. An error here would mark the new project as
	// tainted, and the next apply would destroy and recreate it.
	crn := strDeref(instance.CRN)
	tags, d := rhaiiSetToStrings(ctx, plan.Tags)
	resp.Diagnostics.Append(d...)
	if len(tags) > 0 || len(rhaiiEnvTags()) > 0 {
		if err := rhaiiUpdateTags(r.meta, crn, rhaiiUserTagType, nil, tags); err != nil {
			resp.Diagnostics.AddWarning("Error attaching tags to the Red Hat AI Inference project", err.Error())
		}
	}
	accessTags, d := rhaiiSetToStrings(ctx, plan.AccessTags)
	resp.Diagnostics.Append(d...)
	if len(accessTags) > 0 {
		if err := rhaiiUpdateTags(r.meta, crn, rhaiiAccessTagType, nil, accessTags); err != nil {
			resp.Diagnostics.AddWarning("Error attaching access tags to the Red Hat AI Inference project", err.Error())
		}
	}

	r.readInto(ctx, id, &plan, &resp.State, &resp.Diagnostics, true)
}

func (r *rhaiiProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state rhaiiProjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.readInto(ctx, state.ID.ValueString(), &state, &resp.State, &resp.Diagnostics, false)
}

func (r *rhaiiProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state rhaiiProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	updateTimeout, d := plan.Timeouts.Update(ctx, rhaiiDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	rcClient, err := r.meta.ResourceControllerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Resource Controller client", err.Error())
		return
	}
	id := state.ID.ValueString()

	if !plan.Name.Equal(state.Name) {
		if _, _, err := rcClient.UpdateResourceInstanceWithContext(ctx, &rc.UpdateResourceInstanceOptions{
			ID:   &id,
			Name: ptr(plan.Name.ValueString()),
		}); err != nil {
			resp.Diagnostics.AddError("Error renaming the Red Hat AI Inference project", fmt.Sprintf("UpdateResourceInstanceWithContext failed: %s", err))
			return
		}
		if err := waitForRhaiiProject(ctx, rcClient, id, rhaiiWaitUpdating, updateTimeout, rhaiiUpdateRefresh); err != nil {
			resp.Diagnostics.AddError("Error waiting for the Red Hat AI Inference project to be updated", fmt.Sprintf("project %s: %s", id, err))
			return
		}
	}

	crn := state.CRN.ValueString()
	if !plan.Tags.IsUnknown() && !plan.Tags.Equal(state.Tags) {
		oldTags, d := rhaiiSetToStrings(ctx, state.Tags)
		resp.Diagnostics.Append(d...)
		newTags, d := rhaiiSetToStrings(ctx, plan.Tags)
		resp.Diagnostics.Append(d...)
		if err := rhaiiUpdateTags(r.meta, crn, rhaiiUserTagType, oldTags, newTags); err != nil {
			resp.Diagnostics.AddError("Error updating the tags of the Red Hat AI Inference project", err.Error())
			return
		}
	}
	if !plan.AccessTags.IsUnknown() && !plan.AccessTags.Equal(state.AccessTags) {
		oldTags, d := rhaiiSetToStrings(ctx, state.AccessTags)
		resp.Diagnostics.Append(d...)
		newTags, d := rhaiiSetToStrings(ctx, plan.AccessTags)
		resp.Diagnostics.Append(d...)
		if err := rhaiiUpdateTags(r.meta, crn, rhaiiAccessTagType, oldTags, newTags); err != nil {
			resp.Diagnostics.AddError("Error updating the access tags of the Red Hat AI Inference project", err.Error())
			return
		}
	}

	r.readInto(ctx, id, &plan, &resp.State, &resp.Diagnostics, true)
}

func (r *rhaiiProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state rhaiiProjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deleteTimeout, d := state.Timeouts.Delete(ctx, rhaiiDefaultTimeout)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	rcClient, err := r.meta.ResourceControllerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Resource Controller client", err.Error())
		return
	}
	id := state.ID.ValueString()

	httpResp, err := rcClient.DeleteResourceInstanceWithContext(ctx, &rc.DeleteResourceInstanceOptions{
		ID:        &id,
		Recursive: flex.PtrToBool(true),
	})
	if err != nil {
		if httpResp != nil && (httpResp.StatusCode == 404 || httpResp.StatusCode == 410) {
			return
		}
		resp.Diagnostics.AddError("Error deleting the Red Hat AI Inference project", fmt.Sprintf("DeleteResourceInstanceWithContext failed: %s", err))
		return
	}
	if err := waitForRhaiiProject(ctx, rcClient, id, rhaiiWaitDeleting, deleteTimeout, rhaiiDeleteRefresh); err != nil {
		resp.Diagnostics.AddError("Error waiting for the Red Hat AI Inference project to be deleted", fmt.Sprintf("project %s: %s", id, err))
	}
}

// ImportState accepts the CRN or the GUID of the project. Read normalizes the ID to the CRN.
func (r *rhaiiProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// readInto reads the project and writes it to the state. After create and
// update (afterApply), the tags of the plan are kept when they are known, so the
// state matches the plan exactly; the tagging API can return tags in another
// form (for example lower case), which shows as a difference on the next plan.
func (r *rhaiiProjectResource) readInto(ctx context.Context, id string, m *rhaiiProjectResourceModel, state *tfsdk.State, diags *diag.Diagnostics, afterApply bool) {
	rcClient, err := r.meta.ResourceControllerV2API()
	if err != nil {
		diags.AddError("Error initializing the Resource Controller client", err.Error())
		return
	}

	instance, httpResp, err := rcClient.GetResourceInstanceWithContext(ctx, &rc.GetResourceInstanceOptions{ID: &id})
	if err != nil {
		if !afterApply && httpResp != nil && httpResp.StatusCode == 404 {
			state.RemoveResource(ctx)
			return
		}
		diags.AddError("Error reading the Red Hat AI Inference project", fmt.Sprintf("GetResourceInstanceWithContext failed: %s", err))
		return
	}
	if !afterApply && isRhaiiInstanceGone(instance) {
		diags.AddWarning("Red Hat AI Inference project removed",
			fmt.Sprintf("Project %s is in state %q and is removed from the Terraform state.", id, strDeref(instance.State)))
		state.RemoveResource(ctx)
		return
	}
	if !isRhaiiInstance(instance) {
		diags.AddError("Not a Red Hat AI Inference project",
			fmt.Sprintf("Resource instance %s is not a Red Hat AI Inference project (service is %q, expected %q).", id, strDeref(instance.ResourceID), rhaiiServiceName))
		return
	}

	values, d := flattenRhaiiProject(ctx, r.meta, instance, &m.rhaiiProjectComputedModel)
	diags.Append(d...)
	if diags.HasError() {
		return
	}
	m.Name = types.StringValue(values.Name)
	m.Location = types.StringValue(values.Location)
	m.PlanName = types.StringValue(values.PlanName)
	m.ResourceGroupID = types.StringValue(values.ResourceGroupID)

	tags, accessTags, d := rhaiiReadTags(ctx, r.meta, strDeref(instance.CRN))
	diags.Append(d...)
	if !afterApply || m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tags
	}
	if !afterApply || m.AccessTags.IsNull() || m.AccessTags.IsUnknown() {
		m.AccessTags = accessTags
	}
	diags.Append(state.Set(ctx, m)...)
}
