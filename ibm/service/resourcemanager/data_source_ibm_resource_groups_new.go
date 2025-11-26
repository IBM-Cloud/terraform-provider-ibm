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

var _ datasource.DataSource = &ResourceGroupsDataSource{}

type ResourceGroupsDataSource struct {
	client interface{}
}

func NewResourceGroupsDataSource(client interface{}) datasource.DataSource {
	return &ResourceGroupsDataSource{
		client: client,
	}
}

func (d *ResourceGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_groups_new"
}

func (d *ResourceGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "List IBM Cloud resource groups",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Resource group name",
				Optional:            true,
			},
			"is_default": schema.BoolAttribute{
				MarkdownDescription: "Default Resource group",
				Optional:            true,
			},
			"include_deleted": schema.BoolAttribute{
				MarkdownDescription: "Include deleted resource groups",
				Optional:            true,
			},
			"date": schema.StringAttribute{
				MarkdownDescription: "The date in the format of YYYY-MM which returns resource groups. Deleted resource groups will be excluded before this month",
				Optional:            true,
			},
			"resource_groups": schema.ListNestedAttribute{
				MarkdownDescription: "List of resource groups",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The ID of the resource group",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Resource group name",
							Computed:            true,
						},
						"is_default": schema.BoolAttribute{
							MarkdownDescription: "Default Resource group",
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
						"resource_linkages": schema.ListAttribute{
							MarkdownDescription: "An array of the resources that linked to the resource group",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"account_id": schema.StringAttribute{
							MarkdownDescription: "Account ID",
							Computed:            true,
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Data source ID",
			},
		},
	}
}

func (d *ResourceGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ResourceGroupsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rMgtClient, err := d.client.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Resource Manager client", err.Error())
		return
	}

	userDetails, err := d.client.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		resp.Diagnostics.AddError("Error getting user details", err.Error())
		return
	}
	accountID := userDetails.UserAccount

	resourceGroupListOptions := rg.ListResourceGroupsOptions{
		AccountID: &accountID,
	}

	if !config.Name.IsNull() {
		name := config.Name.ValueString()
		resourceGroupListOptions.Name = &name
	}

	if !config.IsDefault.IsNull() {
		defaultBool := config.IsDefault.ValueBool()
		resourceGroupListOptions.Default = &defaultBool
	}

	if !config.IncludeDeleted.IsNull() {
		includeDeletedBool := config.IncludeDeleted.ValueBool()
		resourceGroupListOptions.IncludeDeleted = &includeDeletedBool
	}

	if !config.Date.IsNull() {
		dateStr := config.Date.ValueString()
		resourceGroupListOptions.Date = &dateStr
	}

	rgList, response, err := rMgtClient.ListResourceGroupsWithContext(ctx, &resourceGroupListOptions)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error listing resource groups",
			fmt.Sprintf("Error: %s, Response: %s", err.Error(), response))
		return
	}

	if rgList == nil || rgList.Resources == nil {
		resp.Diagnostics.AddError(
			"No resource groups found",
			"No resource groups were returned from the API")
		return
	}

	resourceGroups := []ResourceGroupItem{}
	for _, resourceGroup := range rgList.Resources {
		rgItem := ResourceGroupItem{}

		if resourceGroup.ID != nil {
			rgItem.ID = basetypes.NewStringValue(*resourceGroup.ID)
		}
		if resourceGroup.Name != nil {
			rgItem.Name = basetypes.NewStringValue(*resourceGroup.Name)
		}
		if resourceGroup.Default != nil {
			rgItem.IsDefault = basetypes.NewBoolValue(*resourceGroup.Default)
		}
		if resourceGroup.State != nil {
			rgItem.State = basetypes.NewStringValue(*resourceGroup.State)
		}
		if resourceGroup.CRN != nil {
			rgItem.Crn = basetypes.NewStringValue(*resourceGroup.CRN)
		}
		if resourceGroup.CreatedAt != nil {
			rgItem.CreatedAt = basetypes.NewStringValue(resourceGroup.CreatedAt.String())
		}
		if resourceGroup.UpdatedAt != nil {
			rgItem.UpdatedAt = basetypes.NewStringValue(resourceGroup.UpdatedAt.String())
		}
		if resourceGroup.TeamsURL != nil {
			rgItem.TeamsURL = basetypes.NewStringValue(*resourceGroup.TeamsURL)
		}
		if resourceGroup.PaymentMethodsURL != nil {
			rgItem.PaymentMethodsURL = basetypes.NewStringValue(*resourceGroup.PaymentMethodsURL)
		}
		if resourceGroup.QuotaURL != nil {
			rgItem.QuotaURL = basetypes.NewStringValue(*resourceGroup.QuotaURL)
		}
		if resourceGroup.QuotaID != nil {
			rgItem.QuotaID = basetypes.NewStringValue(*resourceGroup.QuotaID)
		}
		if resourceGroup.AccountID != nil {
			rgItem.AccountID = basetypes.NewStringValue(*resourceGroup.AccountID)
		}
		if resourceGroup.ResourceLinkages != nil && len(resourceGroup.ResourceLinkages) > 0 {
			rl := []string{}
			for _, r := range resourceGroup.ResourceLinkages {
				rl = append(rl, r.(string))
			}
			linkages, _ := types.ListValueFrom(ctx, types.StringType, rl)
			rgItem.ResourceLinkages = linkages
		}

		resourceGroups = append(resourceGroups, rgItem)
	}

	config.ResourceGroups = resourceGroups
	config.Id = basetypes.NewStringValue(accountID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

type ResourceGroupsDataSourceModel struct {
	Name           types.String        `tfsdk:"name"`
	IsDefault      types.Bool          `tfsdk:"is_default"`
	IncludeDeleted types.Bool          `tfsdk:"include_deleted"`
	Date           types.String        `tfsdk:"date"`
	ResourceGroups []ResourceGroupItem `tfsdk:"resource_groups"`
	Id             types.String        `tfsdk:"id"`
}

type ResourceGroupItem struct {
	ID                types.String `tfsdk:"id"`
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
	ResourceLinkages  types.List   `tfsdk:"resource_linkages"`
	AccountID         types.String `tfsdk:"account_id"`
}
