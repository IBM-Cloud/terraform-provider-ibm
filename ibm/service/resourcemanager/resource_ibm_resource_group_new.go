// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &ResourceGroupResource{}
var _ resource.ResourceWithImportState = &ResourceGroupResource{}

type ResourceGroupResource struct {
	client interface{}
}

func NewResourceGroupResource(client interface{}) resource.Resource {
	return &ResourceGroupResource{
		client: client,
	}
}

func (r *ResourceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group_new"
}

func (r *ResourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a resource group resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the resource group",
				Required:            true,
			},
			"default": schema.BoolAttribute{
				MarkdownDescription: "Specifies whether its default resource group or not",
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "Tags associated with the resource group",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The full CRN associated with the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "The date when the resource group was initially created",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "The date when the resource group was last updated",
				Computed:            true,
			},
			"teams_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the team details that associated with the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"payment_methods_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the payment methods details that associated with the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"quota_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the quota details that associated with the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"quota_id": schema.StringAttribute{
				MarkdownDescription: "An alpha-numeric value identifying the quota ID associated with the resource group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_linkages": schema.SetAttribute{
				MarkdownDescription: "An array of the resources that linked to the resource group",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource group ID",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ResourceGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceGroupModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := r.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	userDetails, err := r.client.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		resp.Diagnostics.AddError("Error getting user details", err.Error())
		return
	}
	accountID := userDetails.UserAccount

	name := plan.Name.ValueString()
	resourceGroupCreate := rg.CreateResourceGroupOptions{
		Name:      &name,
		AccountID: &accountID,
	}

	resourceGroup, response, err := rMgtClient.CreateResourceGroup(&resourceGroupCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating resource group",
			fmt.Sprintf("Error: %s, Response: %s", err.Error(), response))
		return
	}

	plan.Id = basetypes.NewStringValue(*resourceGroup.ID)

	// Read back the full state
	err = r.readResourceGroup(ctx, &plan, rMgtClient)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource group after create", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ResourceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceGroupModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := r.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	err = r.readResourceGroup(ctx, &state, rMgtClient)
	if err != nil {
		// Check if resource was deleted (404 error)
		if err.Error() == "resource group not found" {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading resource group", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ResourceGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceGroupModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := r.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	resourceGroupID := state.Id.ValueString()
	hasChange := false

	if !plan.Name.Equal(state.Name) {
		name := plan.Name.ValueString()
		resourceGroupUpdate := rg.UpdateResourceGroupOptions{
			ID:   &resourceGroupID,
			Name: &name,
		}

		_, response, err := rMgtClient.UpdateResourceGroup(&resourceGroupUpdate)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating resource group",
				fmt.Sprintf("Error: %s, Response: %s", err.Error(), response))
			return
		}
		hasChange = true
	}

	if hasChange {
		err = r.readResourceGroup(ctx, &plan, rMgtClient)
		if err != nil {
			resp.Diagnostics.AddError("Error reading resource group after update", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ResourceGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceGroupModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := r.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	resourceGroupID := state.Id.ValueString()
	resourceGroupDelete := rg.DeleteResourceGroupOptions{
		ID: &resourceGroupID,
	}

	response, err := rMgtClient.DeleteResourceGroup(&resourceGroupDelete)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("[WARN] Resource Group is not found")
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting resource group",
			fmt.Sprintf("Error: %s, Response: %s", err.Error(), response))
		return
	}
}

func (r *ResourceGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResourceGroupResource) readResourceGroup(ctx context.Context, model *ResourceGroupModel, client *rg.ResourceManagerV2) error {
	resourceGroupID := model.Id.ValueString()
	resourceGroupGet := rg.GetResourceGroupOptions{
		ID: &resourceGroupID,
	}

	resourceGroup, response, err := client.GetResourceGroup(&resourceGroupGet)
	if err != nil || resourceGroup == nil {
		if response != nil && response.StatusCode == 404 {
			return fmt.Errorf("resource group not found")
		}
		return fmt.Errorf("error retrieving resource group: %s. API Response: %s", err, response)
	}

	model.Name = basetypes.NewStringValue(*resourceGroup.Name)
	if resourceGroup.State != nil {
		model.State = basetypes.NewStringValue(*resourceGroup.State)
	}
	if resourceGroup.Default != nil {
		model.Default = basetypes.NewBoolValue(*resourceGroup.Default)
	}
	if resourceGroup.CRN != nil {
		model.Crn = basetypes.NewStringValue(*resourceGroup.CRN)
	}
	if resourceGroup.CreatedAt != nil {
		model.CreatedAt = basetypes.NewStringValue(resourceGroup.CreatedAt.String())
	}
	if resourceGroup.UpdatedAt != nil {
		model.UpdatedAt = basetypes.NewStringValue(resourceGroup.UpdatedAt.String())
	}
	if resourceGroup.TeamsURL != nil {
		model.TeamsURL = basetypes.NewStringValue(*resourceGroup.TeamsURL)
	}
	if resourceGroup.PaymentMethodsURL != nil {
		model.PaymentMethodsURL = basetypes.NewStringValue(*resourceGroup.PaymentMethodsURL)
	}
	if resourceGroup.QuotaURL != nil {
		model.QuotaURL = basetypes.NewStringValue(*resourceGroup.QuotaURL)
	}
	if resourceGroup.QuotaID != nil {
		model.QuotaID = basetypes.NewStringValue(*resourceGroup.QuotaID)
	}
	if resourceGroup.ResourceLinkages != nil && len(resourceGroup.ResourceLinkages) > 0 {
		rl := []string{}
		for _, r := range resourceGroup.ResourceLinkages {
			rl = append(rl, r.(string))
		}
		linkages, _ := types.SetValueFrom(ctx, types.StringType, rl)
		model.ResourceLinkages = linkages
	}

	return nil
}

type ResourceGroupModel struct {
	Name              types.String `tfsdk:"name"`
	Default           types.Bool   `tfsdk:"default"`
	State             types.String `tfsdk:"state"`
	Tags              types.Set    `tfsdk:"tags"`
	Crn               types.String `tfsdk:"crn"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	TeamsURL          types.String `tfsdk:"teams_url"`
	PaymentMethodsURL types.String `tfsdk:"payment_methods_url"`
	QuotaURL          types.String `tfsdk:"quota_url"`
	QuotaID           types.String `tfsdk:"quota_id"`
	ResourceLinkages  types.Set    `tfsdk:"resource_linkages"`
	Id                types.String `tfsdk:"id"`
}
