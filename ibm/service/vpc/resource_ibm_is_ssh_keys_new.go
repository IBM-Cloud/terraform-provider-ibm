package vpc

import (
	"context"
	"log"
	"regexp"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &SSHKeyResource{}

type SSHKeyResource struct {
	client interface{}
}

func NewSSHKeyResource(client interface{}) resource.Resource {
	return &SSHKeyResource{
		client: client,
	}
}

func (r *SSHKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_is_ssh_key_new"
}

func (r *SSHKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "ssh key resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name for this key. The name must not be used by another key in the region. If unspecified, the name will be a hyphenated list of randomly-selected words",
				Optional:            true,
				Computed:            true,
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "A unique public SSH key to import, in OpenSSH format (consisting of three space-separated fields: the algorithm name, base64-encoded key, and a comment). The algorithm and comment fields may be omitted, as only the key field is imported. Keys of type rsa may be 2048 or 4096 bits in length, however 4096 is recommended. Keys of type ed25519 are 256 bits in length.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: "Tags for the ssh key",
				Optional:            true,
				ElementType:         types.StringType,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(stringvalidator.RegexMatches(regexp.MustCompile("read|write|read and write"), "Must be either of read|write|read and write")),
				},
			},
			"access_tags": schema.ListAttribute{
				MarkdownDescription: "Access list for this ssh key",
				Required:            true,
				ElementType:         types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The crypto-system used by this key",
				Optional:            true,
				Computed:            true,
			},
			"fingerprint": schema.StringAttribute{
				MarkdownDescription: "SSH key Fingerprint info",
				Computed:            true,
			},
			"length": schema.Int64Attribute{
				MarkdownDescription: "SSH key Length",
				Computed:            true,
			},
			"resource_group": schema.StringAttribute{
				MarkdownDescription: "The resource group to use. If unspecified, the account's default resource group will be used.",
				Optional:            true,
				Computed:            true,
			},
			"resource_controller_url": schema.StringAttribute{
				MarkdownDescription: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
				Optional:            true,
				Computed:            true,
			},
			"resource_name": schema.StringAttribute{
				MarkdownDescription: "The name of the resource",
				Computed:            true,
			},
			"resource_crn": schema.StringAttribute{
				MarkdownDescription: "The crn of the resource",
				Computed:            true,
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The crn of the resource",
				Computed:            true,
			},
			"resource_group_name": schema.StringAttribute{
				MarkdownDescription: "The resource group name in which resource is provisioned",
				Computed:            true,
			},

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key Id",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *SSHKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SSHKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}
	sess, err := r.client.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
	}

	options := &vpcv1.CreateKeyOptions{}
	if plan.ResourceGroup.ValueString() != "" {
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: core.StringPtr(plan.ResourceGroup.ValueString()),
		}
	}
	if plan.Name.ValueString() != "" {
		options.Name = core.StringPtr(plan.Name.ValueString())
	}
	if plan.Type.ValueString() != "" {
		options.Type = core.StringPtr(plan.Type.ValueString())
	}
	key, _, err := sess.CreateKey(options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating key",
			err.Error(),
		)
		return
	}

	plan.Id = basetypes.NewStringValue(*key.ID)
	plan.Fingerprint = basetypes.NewStringValue(*key.Fingerprint)
	plan.Length = basetypes.NewInt64Value(int64(*key.Length))

	if !plan.Tags.IsNull() {
		newList := plan.Tags.Elements()
		err = flex.UpdateGlobalTagsElementsUsingCRN(nil, newList, r.client, *key.CRN, "", isKeyUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc SSH Key (%s) tags: %s", plan.Id.ValueString(), err)
		}
	}

	if !plan.AccessTags.IsNull() {
		newList := plan.AccessTags.Elements()
		err = flex.UpdateGlobalTagsElementsUsingCRN(nil, newList, r.client, *key.CRN, "", isKeyAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc SSH Key (%s) access tags: %s", plan.Id.ValueString(), err)
		}
	}

	resp.State.Set(ctx, plan)
}

func (r *SSHKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SSHKeyResourceModel

	req.State.Get(ctx, &state)

	sess, err := ctx.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
		return
	}

	options := &vpcv1.GetKeyOptions{
		ID: core.StringPtr(state.Id.ValueString()),
	}

	key, _, err := sess.GetKey(options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting key",
			err.Error(),
		)
		return
	}

	state.Name = basetypes.NewStringValue(*key.Name)
	state.PublicKey = basetypes.NewStringValue(*key.PublicKey)
	state.Type = basetypes.NewStringValue(*key.Type)
	state.Fingerprint = basetypes.NewStringValue(*key.Fingerprint)
	state.Length = basetypes.NewInt64Value(int64(*key.Length))
	state.Id = basetypes.NewStringValue(*key.ID)
	state.ResourceGroup = basetypes.NewStringValue(*key.ResourceGroup.ID)
	controller, err := flex.GetBaseController(r.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting base controller",
			err.Error(),
		)
		return
	}
	state.ResourceControllerURL = basetypes.NewStringValue(controller + "/vpc-ext/compute/sshKeys")
	state.ResourceControllerURL = basetypes.NewStringValue(*key.ResourceGroup.ID)
	state.ResourceName = basetypes.NewStringValue(*key.ResourceGroup.Name)
	state.ResourceCRN = basetypes.NewStringValue(*key.CRN)
	state.Crn = basetypes.NewStringValue(*key.CRN)
	state.ResourceGroupName = basetypes.NewStringValue(*key.ResourceGroup.Name)
	// state.Tags, _ = basetypes.NewListValue()
	// state.AccessTags            =

	resp.State.Set(ctx, state)
}

func (r *SSHKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SSHKeyResourceModel
	req.Plan.Get(ctx, &plan)

	var state SSHKeyResourceModel
	req.State.Get(ctx, &state)

	sess, err := ctx.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
		return
	}

	if plan.Name != state.Name {
		options := &vpcv1.UpdateKeyOptions{
			ID: core.StringPtr(state.Id.ValueString()),
		}
		keyPatchModel := &vpcv1.KeyPatch{
			Name: core.StringPtr(plan.Name.ValueString()),
		}
		keyPatch, err := keyPatchModel.AsPatch()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating key patch",
				err.Error(),
			)
			return
		}
		options.KeyPatch = keyPatch
		_, _, err = sess.UpdateKey(options)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating key",
				err.Error(),
			)
			return
		}
	}

	resp.State.Set(ctx, plan)
}

func (r *SSHKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SSHKeyResourceModel
	req.State.Get(ctx, &state)

	sess, err := ctx.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
		return
	}

	options := &vpcv1.DeleteKeyOptions{
		ID: core.StringPtr(state.Id.ValueString()),
	}

	_, err = sess.DeleteKey(options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting key",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

type SSHKeyResourceModel struct {
	Name                  types.String `tfsdk:"name"`
	PublicKey             types.String `tfsdk:"public_key"`
	Id                    types.String `tfsdk:"id"`
	Type                  types.String `tfsdk:"type"`
	Length                types.Int64  `tfsdk:"length"`
	Fingerprint           types.String `tfsdk:"fingerprint"`
	ResourceGroup         types.String `tfsdk:"resource_group"`
	ResourceControllerURL types.String `tfsdk:"resource_controller_url"`
	ResourceName          types.String `tfsdk:"resource_name"`
	ResourceCRN           types.String `tfsdk:"resource_crn"`
	Crn                   types.String `tfsdk:"crn"`
	ResourceGroupName     types.String `tfsdk:"resource_group_name"`
	Tags                  types.List   `tfsdk:"tags"`
	AccessTags            types.List   `tfsdk:"access_tags"`
}
