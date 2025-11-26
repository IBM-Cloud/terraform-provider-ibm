// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &ResourceGroupDataSource{}

type ResourceGroupDataSource struct {
	client interface{}
}

func NewResourceGroupDataSource(client interface{}) datasource.DataSource {
	return &ResourceGroupDataSource{
		client: client,
	}
}

func (d *ResourceGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group_new"
}

func (d *ResourceGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieve information about an IBM Cloud resource group",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Resource group name",
				Optional:            true,
				Computed:            true,
			},
			"is_default": schema.BoolAttribute{
				MarkdownDescription: "Default Resource group",
				Optional:            true,
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the resource group",
				Computed:            true,
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The full CRN associated with the resource group",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "The date when the resource group was initially created",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "The date when the resource group was last updated",
				Computed:            true,
			},
			"teams_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the team details that associated with the resource group",
				Computed:            true,
			},
			"payment_methods_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the payment methods details that associated with the resource group",
				Computed:            true,
			},
			"quota_url": schema.StringAttribute{
				MarkdownDescription: "The URL to access the quota details that associated with the resource group",
				Computed:            true,
			},
			"quota_id": schema.StringAttribute{
				MarkdownDescription: "An alpha-numeric value identifying the quota ID associated with the resource group",
				Computed:            true,
			},
			"resource_linkages": schema.SetAttribute{
				MarkdownDescription: "An array of the resources that linked to the resource group",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "Account ID",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource group ID",
			},
		},
	}
}

func (d *ResourceGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ResourceGroupDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := d.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	var defaultGrp bool
	if !config.IsDefault.IsNull() {
		defaultGrp = config.IsDefault.ValueBool()
	}

	var name string
	if !config.Name.IsNull() {
		name = config.Name.ValueString()
	}

	if !defaultGrp && name == "" {
		resp.Diagnostics.AddError(
			"Missing required property",
			"Need a resource group name, or the is_default true")
		return
	}

	userDetails, err := d.client.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		resp.Diagnostics.AddError("Error getting user details", err.Error())
		return
	}
	accountID := userDetails.UserAccount

	resourceGroupList := rg.ListResourceGroupsOptions{
		AccountID: &accountID,
	}

	if defaultGrp {
		resourceGroupList.Default = &defaultGrp
	} else if name != "" {
		resourceGroupList.Name = &name
	}

	rgList, response, err := rMgtClient.ListResourceGroups(&resourceGroupList)
	if err != nil || rgList == nil || rgList.Resources == nil {
		resp.Diagnostics.AddError(
			"Error retrieving resource group",
			fmt.Sprintf("Error: %s, Response: %s", err.Error(), response))
		return
	}

	if len(rgList.Resources) < 1 {
		resp.Diagnostics.AddError(
			"Resource group not found",
			"Given Resource Group is not found in the account")
		return
	}

	resourceGroup := rgList.Resources[0]
	config.Id = basetypes.NewStringValue(*resourceGroup.ID)

	if resourceGroup.Name != nil {
		config.Name = basetypes.NewStringValue(*resourceGroup.Name)
	}
	if resourceGroup.Default != nil {
		config.IsDefault = basetypes.NewBoolValue(*resourceGroup.Default)
	}
	if resourceGroup.State != nil {
		config.State = basetypes.NewStringValue(*resourceGroup.State)
	}
	if resourceGroup.CRN != nil {
		config.Crn = basetypes.NewStringValue(*resourceGroup.CRN)
	}
	if resourceGroup.CreatedAt != nil {
		config.CreatedAt = basetypes.NewStringValue(resourceGroup.CreatedAt.String())
	}
	if resourceGroup.UpdatedAt != nil {
		config.UpdatedAt = basetypes.NewStringValue(resourceGroup.UpdatedAt.String())
	}
	if resourceGroup.TeamsURL != nil {
		config.TeamsURL = basetypes.NewStringValue(*resourceGroup.TeamsURL)
	}
	if resourceGroup.PaymentMethodsURL != nil {
		config.PaymentMethodsURL = basetypes.NewStringValue(*resourceGroup.PaymentMethodsURL)
	}
	if resourceGroup.QuotaURL != nil {
		config.QuotaURL = basetypes.NewStringValue(*resourceGroup.QuotaURL)
	}
	if resourceGroup.QuotaID != nil {
		config.QuotaID = basetypes.NewStringValue(*resourceGroup.QuotaID)
	}
	if resourceGroup.AccountID != nil {
		config.AccountID = basetypes.NewStringValue(*resourceGroup.AccountID)
	}
	if resourceGroup.ResourceLinkages != nil && len(resourceGroup.ResourceLinkages) > 0 {
		rl := []string{}
		for _, r := range resourceGroup.ResourceLinkages {
			rl = append(rl, r.(string))
		}
		linkages, _ := types.SetValueFrom(ctx, types.StringType, rl)
		config.ResourceLinkages = linkages
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

type ResourceGroupDataSourceModel struct {
	Name              types.String `tfsdk:"name"`
	IsDefault         types.Bool   `tfsdk:"is_default"`
	State             types.String `tfsdk:"state"`
	Crn               types.String `tfsdk:"crn"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	TeamsURL          types.String `tfsdk:"teams_url"`
	PaymentMethodsURL types.String `tfsdk:"payment_methods_url"`
	QuotaURL          types.String `tfsdk:"quota_url"`
	QuotaID           types.String `tfsdk:"quota_id"`
	ResourceLinkages  types.Set    `tfsdk:"resource_linkages"`
	AccountID         types.String `tfsdk:"account_id"`
	Id                types.String `tfsdk:"id"`
}
