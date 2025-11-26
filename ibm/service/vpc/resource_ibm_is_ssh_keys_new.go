package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
				CustomType:          SSHPublicKeyType{},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: "Tags for the ssh key",
				Optional:            true,
				ElementType:         types.StringType,
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
	plan.Name = basetypes.NewStringValue(*key.Name)
	plan.PublicKey = NewSSHPublicKeyValue(*key.PublicKey)
	plan.Type = basetypes.NewStringValue(*key.Type)
	plan.Fingerprint = basetypes.NewStringValue(*key.Fingerprint)
	plan.Length = basetypes.NewInt64Value(int64(*key.Length))
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
	plan.ResourceName = basetypes.NewStringValue(*key.Name)
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

	key, response, err := sess.GetKey(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error getting key",
			err.Error(),
		)
		return
	}

	state.Name = basetypes.NewStringValue(*key.Name)
	state.PublicKey = NewSSHPublicKeyValue(*key.PublicKey)
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
	state.ResourceName = basetypes.NewStringValue(*key.Name)
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

	id := state.Id.ValueString()
	keyUpdated := false

	// Update name if changed
	if !plan.Name.Equal(state.Name) {
		input := &vpcv1.UpdateKeyOptions{}
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
			resp.Diagnostics.AddError(
				fmt.Sprintf("Error updating key %v", res),
				err.Error(),
			)
			return
		}
		keyUpdated = true
	}

	// Get current CRN for tag operations
	getOptions := &vpcv1.GetKeyOptions{
		ID: core.StringPtr(id),
	}
	key, _, err := sess.GetKey(getOptions)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting key after update",
			err.Error(),
		)
		return
	}

	// Update tags if changed
	if !plan.Tags.Equal(state.Tags) {
		oldList := state.Tags.Elements()
		newList := plan.Tags.Elements()
		err = flex.UpdateGlobalTagsElementsUsingCRN(oldList, newList, r.client, *key.CRN, "", isKeyUserTagType)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating tags",
				err.Error(),
			)
			return
		}
		keyUpdated = true
	}

	// Update access tags if changed
	if !plan.AccessTags.Equal(state.AccessTags) {
		oldList := state.AccessTags.Elements()
		newList := plan.AccessTags.Elements()
		err = flex.UpdateGlobalTagsElementsUsingCRN(oldList, newList, r.client, *key.CRN, "", isKeyAccessTagType)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating access tags",
				err.Error(),
			)
			return
		}
		keyUpdated = true
	}

	// Refresh state from API if anything was updated
	if keyUpdated {
		key, response, err := sess.GetKey(getOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Error refreshing key state after update",
				err.Error(),
			)
			return
		}

		plan.Name = basetypes.NewStringValue(*key.Name)
		plan.PublicKey = NewSSHPublicKeyValue(*key.PublicKey)
		plan.Type = basetypes.NewStringValue(*key.Type)
		plan.Fingerprint = basetypes.NewStringValue(*key.Fingerprint)
		plan.Length = basetypes.NewInt64Value(int64(*key.Length))
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
		plan.ResourceName = basetypes.NewStringValue(*key.Name)
		plan.ResourceCRN = basetypes.NewStringValue(*key.CRN)
		plan.Crn = basetypes.NewStringValue(*key.CRN)
		plan.ResourceGroupName = basetypes.NewStringValue(*key.ResourceGroup.Name)

		tags, _ := flex.GetGlobalTagsElementsUsingCRN(r.client, *key.CRN, "", isKeyUserTagType)
		access, _ := flex.GetGlobalTagsElementsUsingCRN(r.client, *key.CRN, "", isKeyAccessTagType)
		if len(tags) > 0 {
			plan.Tags, _ = basetypes.NewListValue(convertStringSliceToListValue(tags))
		}
		if len(access) > 0 {
			plan.AccessTags, _ = basetypes.NewListValue(convertStringSliceToListValue(access))
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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

	response, err := sess.DeleteKey(options)
	if err != nil {
		// Ignore 404 errors - resource already deleted
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting key",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SSHKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type SSHKeyResourceModel struct {
	Name                  types.String      `tfsdk:"name"`
	PublicKey             SSHPublicKeyValue `tfsdk:"public_key"`
	Id                    types.String      `tfsdk:"id"`
	Type                  types.String      `tfsdk:"type"`
	Length                types.Int64       `tfsdk:"length"`
	Fingerprint           types.String      `tfsdk:"fingerprint"`
	ResourceGroup         types.String      `tfsdk:"resource_group"`
	ResourceControllerURL types.String      `tfsdk:"resource_controller_url"`
	ResourceName          types.String      `tfsdk:"resource_name"`
	ResourceCRN           types.String      `tfsdk:"resource_crn"`
	Crn                   types.String      `tfsdk:"crn"`
	ResourceGroupName     types.String      `tfsdk:"resource_group_name"`
	Tags                  types.List        `tfsdk:"tags"`
	AccessTags            types.List        `tfsdk:"access_tags"`
}

func convertStringSliceToListValue(stringSlice []string) (attr.Type, []attr.Value) {
	valueType := basetypes.StringType{}
	values := make([]attr.Value, len(stringSlice))

	for i, s := range stringSlice {
		values[i] = basetypes.NewStringValue(s)
	}

	return valueType, values
}

func BeautifyResponse(response interface{}) string {
	output, err := json.MarshalIndent(response, "", "    ")
	if err == nil {
		return fmt.Sprintf("%+v\n", string(output))
	}
	return fmt.Sprintf("Error : %#v", response)
}

// to suppress any change shown when keys are same
func suppressSshKeyPublicKeyDiff(old, new string) bool {
	// if there are extra spaces or new lines, suppress that change
	if strings.Compare(strings.TrimSpace(old), strings.TrimSpace(new)) != 0 {
		// if old is empty
		if old != "" {
			//create a new piblickey object from the string
			usePK, error := parseKey(new)
			if error != nil {
				return false
			}
			// returns the key in byte format with an extra added new line at the end
			newkey := strings.TrimRight(string(ssh.MarshalAuthorizedKey(usePK)), "\n")
			// check if both keys are same, if yes suppress the change
			return strings.TrimSpace(strings.TrimPrefix(newkey, old)) == ""
		} else {
			return strings.TrimSpace(strings.TrimPrefix(new, old)) == ""
		}
	} else {
		return true
	}
}

// SSHPublicKeyValue is a custom string value type that implements semantic equality
type SSHPublicKeyValue struct {
	basetypes.StringValue
}

// Type returns the type of the value
func (v SSHPublicKeyValue) Type(_ context.Context) attr.Type {
	return SSHPublicKeyType{}
}

// Equal returns true if the other value is equal
func (v SSHPublicKeyValue) Equal(other attr.Value) bool {
	o, ok := other.(SSHPublicKeyValue)
	if !ok {
		return false
	}

	return v.StringValue.Equal(o.StringValue)
}

// StringSemanticEquals implements semantic equality for SSH public keys
func (v SSHPublicKeyValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(SSHPublicKeyValue)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)
		return false, diags
	}

	return suppressSshKeyPublicKeyDiff(v.StringValue.ValueString(), newValue.StringValue.ValueString()), diags
}

// SSHPublicKeyType is the type for SSHPublicKeyValue
type SSHPublicKeyType struct {
	basetypes.StringType
}

func (t SSHPublicKeyType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	return SSHPublicKeyValue{
		StringValue: in,
	}, diags
}

func (t SSHPublicKeyType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	val, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringVal, ok := val.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", val)
	}

	return SSHPublicKeyValue{
		StringValue: stringVal,
	}, nil
}

// String returns a human readable string
func (t SSHPublicKeyType) String() string {
	return "SSHPublicKeyType"
}

// Equal returns true if the other type is the same
func (t SSHPublicKeyType) Equal(o attr.Type) bool {
	other, ok := o.(SSHPublicKeyType)
	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// Type returns the underlying type
func (t SSHPublicKeyType) Type() attr.Type {
	return t
}

// ValueType returns the value type
func (t SSHPublicKeyType) ValueType(_ context.Context) attr.Value {
	return SSHPublicKeyValue{}
}

// Validate implements type validation
func (t SSHPublicKeyType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if value.IsNull() || !value.IsKnown() {
		return diags
	}

	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"Invalid SSH Public Key Value",
			"String value expected, received: "+value.String(),
		)
		return diags
	}

	return diags
}

var (
	_ attr.Type                = SSHPublicKeyType{}
	_ basetypes.StringValuable = SSHPublicKeyValue{}
)

// Helper functions to create new values
func NewSSHPublicKeyNull() SSHPublicKeyValue {
	return SSHPublicKeyValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewSSHPublicKeyUnknown() SSHPublicKeyValue {
	return SSHPublicKeyValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewSSHPublicKeyValue(value string) SSHPublicKeyValue {
	return SSHPublicKeyValue{
		StringValue: basetypes.NewStringValue(value),
	}
}
