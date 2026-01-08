// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ datasource.DataSourceWithConfigure = &DataSourceIsSshKey{}
	_ datasource.DataSource              = &DataSourceIsSshKey{}
)

type DataSourceIsSshKey struct {
	client interface{}
}

type DataSourceIsSshKeyModel struct {
	ID                    types.String `tfsdk:"id"`
	Href                  types.String `tfsdk:"href"`
	ResourceGroup         types.String `tfsdk:"resource_group"`
	Tags                  types.List   `tfsdk:"tags"`
	Name                  types.String `tfsdk:"name"`
	Type                  types.String `tfsdk:"type"`
	Fingerprint           types.String `tfsdk:"fingerprint"`
	PublicKey             types.String `tfsdk:"public_key"`
	Length                types.Int64  `tfsdk:"length"`
	ResourceControllerUrl types.String `tfsdk:"resource_controller_url"`
	ResourceName          types.String `tfsdk:"resource_name"`
	ResourceCrn           types.String `tfsdk:"resource_crn"`
	Crn                   types.String `tfsdk:"crn"`
	ResourceGroupName     types.String `tfsdk:"resource_group_name"`
	AccessTags            types.List   `tfsdk:"access_tags"`
}

func NewIsSshKeyDataSource() datasource.DataSource {
	return &DataSourceIsSshKey{}
}

func (d *DataSourceIsSshKey) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_is_ssh_key_new"
}

func (d *DataSourceIsSshKey) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	session, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = session
}

func (d *DataSourceIsSshKey) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The ssh key data source retrieves the given ssh key.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier of the SSH key",
				Computed:    true,
			},
			"href": schema.StringAttribute{
				Description: "The URL for this SSH key",
				Computed:    true,
			},
			"resource_group": schema.StringAttribute{
				Description: "Resource group ID",
				Computed:    true,
			},
			"tags": schema.ListAttribute{
				Description: "User Tags for the ssh",
				Computed:    true,
				ElementType: types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ssh key",
			},
			"type": schema.StringAttribute{
				Description: "The ssh key type",
				Computed:    true,
			},
			"fingerprint": schema.StringAttribute{
				Description: "The ssh key Fingerprint",
				Computed:    true,
			},
			"public_key": schema.StringAttribute{
				Description: "SSH Public key data",
				Computed:    true,
			},
			"length": schema.Int64Attribute{
				Description: "The ssh key length",
				Computed:    true,
			},
			"resource_controller_url": schema.StringAttribute{
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
				Computed:    true,
			},
			"resource_name": schema.StringAttribute{
				Description: "The name of the resource",
				Computed:    true,
			},
			"resource_crn": schema.StringAttribute{
				Description: "The crn of the resource",
				Computed:    true,
			},
			"crn": schema.StringAttribute{
				Description: "The crn of the resource",
				Computed:    true,
			},
			"resource_group_name": schema.StringAttribute{
				Description: "The resource group name in which resource is provisioned",
				Computed:    true,
			},
			"access_tags": schema.ListAttribute{
				Description: "List of access tags",
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}

func (d *DataSourceIsSshKey) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DataSourceIsSshKeyModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	name := data.Name.ValueString()
	client, err := d.client.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
	}
	options := &vpcv1.ListKeysOptions{}

	pager, err := client.NewKeysPager(options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching keys",
			err.Error(),
		)
		return
	}
	keys, err := pager.GetAll()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching keys",
			err.Error(),
		)
		return
	}
	var keyfound *vpcv1.Key
	for _, key := range keys {
		if *key.Name == name {
			keyfound = &key
		}
	}
	if keyfound != nil {
		data.ID = types.StringValue(*keyfound.ID)
		data.Href = types.StringValue(*keyfound.Href)
		data.Name = types.StringValue(*keyfound.Name)
		data.ResourceName = types.StringValue(*keyfound.Name)
		data.ResourceCrn = types.StringValue(*keyfound.CRN)
		data.Crn = types.StringValue(*keyfound.CRN)
		data.ResourceGroup = types.StringValue(*keyfound.ResourceGroup.ID)
		data.ResourceGroupName = types.StringValue(*keyfound.ResourceGroup.Name)
		data.Fingerprint = types.StringValue(*keyfound.Fingerprint)
		data.Length = types.Int64Value(*keyfound.Length)
		data.PublicKey = types.StringValue(*keyfound.PublicKey)
		data.Type = types.StringValue(*keyfound.Type)
		controller, err := flex.GetBaseController(d.client)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("RC error %s", name),
				err.Error(),
			)
		}
		data.ResourceControllerUrl = types.StringValue(controller + "/vpc/compute/sshKeys")
		tagsSet, _ := flex.GetGlobalTagsUsingCRN(d.client, *keyfound.CRN, "", isKeyUserTagType)
		accessSet, _ := flex.GetGlobalTagsUsingCRN(d.client, *keyfound.CRN, "", isKeyAccessTagType)
		if tagsSet != nil && tagsSet.Len() > 0 {
			tags := make([]string, 0, tagsSet.Len())
			for _, v := range tagsSet.List() {
				tags = append(tags, v.(string))
			}
			data.Tags, _ = basetypes.NewListValue(convertStringSliceToListValue(tags))
		}
		if accessSet != nil && accessSet.Len() > 0 {
			access := make([]string, 0, accessSet.Len())
			for _, v := range accessSet.List() {
				access = append(access, v.(string))
			}
			data.AccessTags, _ = basetypes.NewListValue(convertStringSliceToListValue(access))
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	} else {
		resp.Diagnostics.AddError(
			"SSH Key Not Found",
			fmt.Sprintf("No key found with the name %s", name),
		)
		return
	}
}
