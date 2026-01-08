// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// Ensure IsSshKeyEphemeralResource satisfies the expected interfaces.
var (
	_ ephemeral.EphemeralResource              = &IsSshKeyEphemeralResource{}
	_ ephemeral.EphemeralResourceWithConfigure = &IsSshKeyEphemeralResource{}
)

// IsSshKeyEphemeralResource defines the ephemeral resource implementation for SSH key lookup.
// This ephemeral resource allows you to lookup an existing SSH key by name without storing
// the key data in state, making it ideal for temporary references during resource creation.
type IsSshKeyEphemeralResource struct {
	// client holds the IBM Cloud session for API calls
	client interface{}
}

// IsSshKeyEphemeralModel describes the ephemeral resource data model.
type IsSshKeyEphemeralModel struct {
	Name        types.String `tfsdk:"name"`
	ID          types.String `tfsdk:"id"`
	Fingerprint types.String `tfsdk:"fingerprint"`
	PublicKey   types.String `tfsdk:"public_key"`
	Type        types.String `tfsdk:"type"`
	Length      types.Int64  `tfsdk:"length"`
	CRN         types.String `tfsdk:"crn"`
	Href        types.String `tfsdk:"href"`
}

// NewIsSshKeyEphemeralResource returns a new instance of the SSH key ephemeral resource.
func NewIsSshKeyEphemeralResource() ephemeral.EphemeralResource {
	return &IsSshKeyEphemeralResource{}
}

// Metadata sets the ephemeral resource type name.
func (r *IsSshKeyEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_is_ssh_key"
}

// Configure initializes the ephemeral resource with the provider's client session.
func (r *IsSshKeyEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	session, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = session
}

// Schema defines the schema for the ephemeral resource.
func (r *IsSshKeyEphemeralResource) Schema(_ context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Ephemeral resource to lookup an existing IBM Cloud VPC SSH key by name. " +
			"The key data is available during the Terraform apply operation but is not stored in state, " +
			"making it ideal for temporary references to existing SSH keys.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The name of the SSH key to lookup. This must match an existing SSH key " +
					"in your IBM Cloud VPC infrastructure.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier of the SSH key.",
			},
			"fingerprint": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The SHA256 fingerprint of the SSH key.",
			},
			"public_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public SSH key data.",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The cryptographic algorithm type of the SSH key (e.g., rsa, ed25519).",
			},
			"length": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The length of the SSH key in bits.",
			},
			"crn": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The Cloud Resource Name (CRN) of the SSH key.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The URL for this SSH key.",
			},
		},
	}
}

// Open performs the SSH key lookup operation.
// This is the main method that executes when the ephemeral resource is referenced.
func (r *IsSshKeyEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data IsSshKeyEphemeralModel

	// Read configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()

	// Get VPC client
	client, err := r.client.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating VPC Client",
			fmt.Sprintf("Unable to create VPC API client: %s", err.Error()),
		)
		return
	}

	// List all SSH keys
	options := &vpcv1.ListKeysOptions{}
	pager, err := client.NewKeysPager(options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Listing SSH Keys",
			fmt.Sprintf("Unable to create SSH keys pager: %s", err.Error()),
		)
		return
	}

	keys, err := pager.GetAll()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Fetching SSH Keys",
			fmt.Sprintf("Unable to fetch SSH keys: %s", err.Error()),
		)
		return
	}

	// Find the key by name
	var found *vpcv1.Key
	for _, key := range keys {
		if *key.Name == name {
			found = &key
			break
		}
	}

	if found == nil {
		resp.Diagnostics.AddError(
			"SSH Key Not Found",
			fmt.Sprintf("No SSH key found with name '%s'. Please verify the key exists in your VPC.", name),
		)
		return
	}

	// Populate ephemeral data
	data.ID = types.StringValue(*found.ID)
	data.Fingerprint = types.StringValue(*found.Fingerprint)
	data.PublicKey = types.StringValue(*found.PublicKey)
	data.Type = types.StringValue(*found.Type)
	data.Length = types.Int64Value(*found.Length)
	data.CRN = types.StringValue(*found.CRN)
	data.Href = types.StringValue(*found.Href)

	// Set the result
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
