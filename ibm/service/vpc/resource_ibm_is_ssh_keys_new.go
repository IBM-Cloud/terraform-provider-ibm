package vpc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"golang.org/x/crypto/ssh"
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Optional:            true,
				ElementType:         types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The crypto-system used by this key",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fingerprint": schema.StringAttribute{
				MarkdownDescription: "SSH key Fingerprint info",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"length": schema.Int64Attribute{
				MarkdownDescription: "SSH key Length",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"resource_group": schema.StringAttribute{
				MarkdownDescription: "The resource group to use. If unspecified, the account's default resource group will be used.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_controller_url": schema.StringAttribute{
				MarkdownDescription: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_name": schema.StringAttribute{
				MarkdownDescription: "The name of the resource",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_crn": schema.StringAttribute{
				MarkdownDescription: "The crn of the resource",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "The crn of the resource",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_group_name": schema.StringAttribute{
				MarkdownDescription: "The resource group name in which resource is provisioned",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
	if plan.PublicKey.ValueString() != "" {
		options.PublicKey = core.StringPtr(plan.PublicKey.ValueString())
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

	plan.Name = basetypes.NewStringValue(*key.Name)
	plan.PublicKey = basetypes.NewStringValue(*key.PublicKey)
	plan.Type = basetypes.NewStringValue(*key.Type)
	plan.Fingerprint = basetypes.NewStringValue(*key.Fingerprint)
	plan.Length = basetypes.NewInt64Value(int64(*key.Length))
	plan.Id = basetypes.NewStringValue(*key.ID)
	plan.ResourceGroup = basetypes.NewStringValue(*key.ResourceGroup.ID)
	controller, err := flex.GetBaseController(r.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting base controller",
			err.Error(),
		)
		return
	}
	plan.ResourceControllerURL = basetypes.NewStringValue(controller + "/vpc-ext/compute/sshKeys")
	plan.ResourceGroup = basetypes.NewStringValue(*key.ResourceGroup.ID)
	plan.ResourceName = basetypes.NewStringValue(*key.ResourceGroup.Name)
	plan.ResourceCRN = basetypes.NewStringValue(*key.CRN)
	plan.Crn = basetypes.NewStringValue(*key.CRN)
	plan.ResourceGroupName = basetypes.NewStringValue(*key.ResourceGroup.Name)
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
	// resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SSHKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SSHKeyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	sess, err := r.client.(conns.ClientSession).VpcV1API()
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
	tags, _ := flex.GetGlobalTagsElementsUsingCRN(r.client, *key.CRN, "", isKeyUserTagType)
	access, _ := flex.GetGlobalTagsElementsUsingCRN(r.client, *key.CRN, "", isKeyAccessTagType)
	if len(tags) > 0 {
		state.Tags, _ = basetypes.NewListValue(convertStringSliceToListValue(tags))
	}
	if len(access) > 0 {
		state.AccessTags, _ = basetypes.NewListValue(convertStringSliceToListValue(access))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SSHKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	sess, err := r.client.(conns.ClientSession).VpcV1API()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating session",
			err.Error(),
		)
		return
	}
	var plan, state SSHKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Name.Equal(state.Name) {
		input := &vpcv1.UpdateKeyOptions{}
		id := state.Id.ValueString()
		input.ID = &id
		keyPatchModel := &vpcv1.KeyPatch{}
		keyPatchModel.Name = core.StringPtr(plan.Name.ValueString())
		keyPatchAsPatch, err := keyPatchModel.AsPatch()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error applying asPatch of keyPatchModel",
				err.Error(),
			)
			return
		}
		input.KeyPatch = keyPatchAsPatch

		_, res, err := sess.UpdateKey(input)

		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf(
				"Error creating key %v", res),
				err.Error(),
			)
			return
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	}
	return
}

func (r *SSHKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SSHKeyResourceModel
	req.State.Get(ctx, &state)

	sess, err := r.client.(conns.ClientSession).VpcV1API()
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

func convertStringSliceToListValue(stringSlice []string) (attr.Type, []attr.Value) {
	valueType := basetypes.StringType{}
	values := make([]attr.Value, len(stringSlice))

	for i, s := range stringSlice {
		values[i] = basetypes.NewStringValue(s)
	}

	return valueType, values
}

func parseKeySshNew(s string) (ssh.PublicKey, error) {
	keyBytes := []byte(s)

	// Accepts formats of PublicKey:
	// - <base64 key>
	// - ssh-rsa/ssh-ed25519 <base64 key>
	// - ssh-rsa/ssh-ed25519 <base64 key> <comment>
	// if PublicKey provides other than just base64 key, then first part must be "ssh-rsa" or "ssh-ed25519"
	if subStrs := strings.Split(s, " "); len(subStrs) > 1 && subStrs[0] != "ssh-rsa" && subStrs[0] != "ssh-ed25519" {
		return nil, errors.New("not an RSA key OR ED25519 key")
	}

	pk, _, _, _, e := ssh.ParseAuthorizedKey(keyBytes)
	if e == nil {
		return pk, nil
	}

	decodedKey := make([]byte, base64.StdEncoding.DecodedLen(len(keyBytes)))
	n, e := base64.StdEncoding.Decode(decodedKey, keyBytes)
	if e != nil {
		return nil, e
	}
	decodedKey = decodedKey[:n]

	pk, e = ssh.ParsePublicKey(decodedKey)
	if e == nil {
		return pk, nil
	}
	return nil, e
}

func BeautifyResponse(response interface{}) string {
	output, err := json.MarshalIndent(response, "", "    ")
	if err == nil {
		return fmt.Sprintf("%+v\n", string(output))
	}
	return fmt.Sprintf("Error : %#v", response)
}
