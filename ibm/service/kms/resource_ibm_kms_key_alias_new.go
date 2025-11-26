package kms

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	kp "github.com/IBM/keyprotect-go-client"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &KMSKeyAliasResource{}
var _ resource.ResourceWithImportState = &KMSKeyAliasResource{}

type KMSKeyAliasResource struct {
	client interface{}
}

func NewKMSKeyAliasResource(client interface{}) resource.Resource {
	return &KMSKeyAliasResource{
		client: client,
	}
}

func (r *KMSKeyAliasResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_key_alias_new"
}

func (r *KMSKeyAliasResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "KMS key alias resource",

		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs instance GUID or CRN",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"alias": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs key alias name",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key_id": schema.StringAttribute{
				MarkdownDescription: "Key ID",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"existing_alias": schema.StringAttribute{
				MarkdownDescription: "Existing Alias of the Key",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"endpoint_type": schema.StringAttribute{
				MarkdownDescription: "public or private",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key alias identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *KMSKeyAliasResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan KMSKeyAliasResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	instanceID := getInstanceIDFromCRN(plan.InstanceID.ValueString())
	kpAPI, _, err := r.populateKPClient(plan, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	aliasName := plan.Alias.ValueString()
	var id string
	if !plan.KeyID.IsNull() && plan.KeyID.ValueString() != "" {
		id = plan.KeyID.ValueString()
	} else if !plan.ExistingAlias.IsNull() && plan.ExistingAlias.ValueString() != "" {
		id = plan.ExistingAlias.ValueString()
	}

	stkey, err := kpAPI.CreateKeyAlias(ctx, aliasName, id)
	if err != nil {
		resp.Diagnostics.AddError("Error creating key alias", err.Error())
		return
	}

	key, err := kpAPI.GetKey(ctx, stkey.KeyID)
	if err != nil {
		resp.Diagnostics.AddError("Error getting key", err.Error())
		return
	}

	plan.Id = basetypes.NewStringValue(fmt.Sprintf("%s:alias:%s", stkey.Alias, key.CRN))
	plan.KeyID = basetypes.NewStringValue(key.ID)

	if strings.Contains((kpAPI.URL).String(), "private") || strings.Contains(kpAPI.Config.BaseURL, "private") {
		plan.EndpointType = basetypes.NewStringValue("private")
	} else {
		plan.EndpointType = basetypes.NewStringValue("public")
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *KMSKeyAliasResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state KMSKeyAliasResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := strings.Split(state.Id.ValueString(), ":alias:")
	if len(id) < 2 {
		resp.Diagnostics.AddError("Invalid ID", fmt.Sprintf("Incorrect ID %s: Id should be a combination of keyAlias:alias:keyCRN", state.Id.ValueString()))
		return
	}

	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(id[1])
	kpAPI, _, err := r.populateKPClient(state, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	key, err := kpAPI.GetKey(ctx, keyid)
	if err != nil {
		if kpError, ok := err.(*kp.Error); ok {
			if kpError.StatusCode == 404 || kpError.StatusCode == 409 {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError("Error getting key", err.Error())
		return
	}

	if key.State == 5 { // Deleted state
		resp.State.RemoveResource(ctx)
		return
	}

	state.Alias = basetypes.NewStringValue(id[0])
	state.KeyID = basetypes.NewStringValue(key.ID)
	state.InstanceID = basetypes.NewStringValue(instanceID)

	if strings.Contains((kpAPI.URL).String(), "private") || strings.Contains(kpAPI.Config.BaseURL, "private") {
		state.EndpointType = basetypes.NewStringValue("private")
	} else {
		state.EndpointType = basetypes.NewStringValue("public")
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *KMSKeyAliasResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No-op: all fields require replace
	var plan KMSKeyAliasResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *KMSKeyAliasResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state KMSKeyAliasResourceModel
	req.State.Get(ctx, &state)

	id := strings.Split(state.Id.ValueString(), ":alias:")
	if len(id) < 2 {
		resp.Diagnostics.AddError("Invalid ID", fmt.Sprintf("Incorrect ID %s", state.Id.ValueString()))
		return
	}

	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(id[1])
	kpAPI, _, err := r.populateKPClient(state, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	err = kpAPI.DeleteKeyAlias(ctx, id[0], keyid)
	if err != nil {
		if kpError, ok := err.(*kp.Error); ok {
			if kpError.StatusCode == 404 {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError("Error deleting key alias", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *KMSKeyAliasResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *KMSKeyAliasResource) populateKPClient(model KMSKeyAliasResourceModel, instanceID string) (*kp.Client, *string, error) {
	kpAPI, err := r.client.(conns.ClientSession).KeyManagementAPI()
	if err != nil {
		return nil, nil, err
	}

	var endpointType string
	if !model.EndpointType.IsNull() {
		endpointType = model.EndpointType.ValueString()
	}

	rsConClient, err := r.client.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, nil, err
	}

	instanceData, resp, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &instanceID,
	})
	if err != nil || instanceData == nil {
		return nil, nil, flex.FmtErrorf("[ERROR] Error retrieving resource instance: %s with resp code: %s", err, resp)
	}

	extensions := instanceData.Extensions
	kpAPI.URL, err = KmsEndpointURL(kpAPI, endpointType, extensions)
	if err != nil {
		return nil, nil, err
	}

	kpAPI.Config.InstanceID = instanceID
	return kpAPI, instanceData.CRN, nil
}

type KMSKeyAliasResourceModel struct {
	InstanceID    types.String `tfsdk:"instance_id"`
	Alias         types.String `tfsdk:"alias"`
	KeyID         types.String `tfsdk:"key_id"`
	ExistingAlias types.String `tfsdk:"existing_alias"`
	EndpointType  types.String `tfsdk:"endpoint_type"`
	Id            types.String `tfsdk:"id"`
}
