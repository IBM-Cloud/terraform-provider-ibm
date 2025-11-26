package kms

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	kp "github.com/IBM/keyprotect-go-client"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &KMSKeyResource{}
var _ resource.ResourceWithImportState = &KMSKeyResource{}

type KMSKeyResource struct {
	client interface{}
}

func NewKMSKeyResource(client interface{}) resource.Resource {
	return &KMSKeyResource{
		client: client,
	}
}

func (r *KMSKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_key_new"
}

func (r *KMSKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "KMS key resource",

		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs instance GUID or CRN",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key_ring_id": schema.StringAttribute{
				MarkdownDescription: "Key Ring for the Key",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key_id": schema.StringAttribute{
				MarkdownDescription: "Key ID",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"key_name": schema.StringAttribute{
				MarkdownDescription: "Key name",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "type of service hs-crypto or kms",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_type": schema.StringAttribute{
				MarkdownDescription: "public or private",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "description of the key",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"standard_key": schema.BoolAttribute{
				MarkdownDescription: "Standard key type",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"payload": schema.StringAttribute{
				MarkdownDescription: "Key payload",
				Sensitive:           true,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"encrypted_nonce": schema.StringAttribute{
				MarkdownDescription: "Only for imported root key",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"iv_value": schema.StringAttribute{
				MarkdownDescription: "Only for imported root key",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"force_delete": schema.BoolAttribute{
				MarkdownDescription: "set to true to force delete the key",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"crn": schema.StringAttribute{
				MarkdownDescription: "Crn of the key",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration_date": schema.StringAttribute{
				MarkdownDescription: "The date and time that the key expires in the system, in RFC 3339 format",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"instance_crn": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs instance CRN",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"registrations": schema.ListAttribute{
				MarkdownDescription: "Registrations of the key across different services",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key_id":               types.StringType,
						"resource_crn":         types.StringType,
						"prevent_key_deletion": types.BoolType,
					},
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
			"resource_status": schema.StringAttribute{
				MarkdownDescription: "The status of the resource",
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
			"resource_controller_url": schema.StringAttribute{
				MarkdownDescription: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *KMSKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan KMSKeyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	keyData, instanceID, err := r.extractKeyDataFromPlan(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error extracting key data", err.Error())
		return
	}

	kpAPI, _, err := r.populateKPClientForKey(plan, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	kpAPI.Config.KeyRing = plan.KeyRingID.ValueString()

	key, err := kpAPI.CreateKeyWithOptions(ctx, keyData.Name, keyData.Extractable,
		kp.WithExpiration(keyData.Expiration),
		kp.WithPayload(keyData.Payload, &keyData.EncryptedNonce, &keyData.IV, false),
		kp.WithDescription(keyData.Description))
	if err != nil {
		resp.Diagnostics.AddError("Error creating key", err.Error())
		return
	}

	plan.Id = basetypes.NewStringValue(key.CRN)

	// Read the full state
	err = r.readKeyToState(ctx, &plan, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading key after create", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *KMSKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state KMSKeyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	instanceCRN, instanceID, keyid := getInstanceAndKeyDataFromCRN(state.Id.ValueString())
	kpAPI, _, err := r.populateKPClientForKey(state, instanceID)
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

	err = r.setKeyDetails(ctx, &state, instanceID, instanceCRN, key, kpAPI)
	if err != nil {
		resp.Diagnostics.AddError("Error setting key details", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *KMSKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state KMSKeyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Keep force_delete from state (it's not updatable via API)
	plan.ForceDelete = state.ForceDelete

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *KMSKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state KMSKeyResourceModel
	req.State.Get(ctx, &state)

	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(state.Id.ValueString())
	kpAPI, _, err := r.populateKPClientForKey(state, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	force := state.ForceDelete.ValueBool()
	f := kp.ForceOpt{
		Force: force,
	}

	_, err = kpAPI.DeleteKey(ctx, keyid, kp.ReturnRepresentation, f)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting key", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *KMSKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *KMSKeyResource) populateKPClientForKey(model KMSKeyResourceModel, instanceID string) (*kp.Client, *string, error) {
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

	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	instanceData, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
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

func (r *KMSKeyResource) extractKeyDataFromPlan(plan KMSKeyResourceModel) (kp.Key, string, error) {
	instanceID := getInstanceIDFromCRN(plan.InstanceID.ValueString())
	var expiration *time.Time

	if !plan.ExpirationDate.IsNull() && plan.ExpirationDate.ValueString() != "" {
		expiration_string := plan.ExpirationDate.ValueString()
		expiration_time, err := time.Parse(time.RFC3339, expiration_string)
		if err != nil {
			return kp.Key{}, "", flex.FmtErrorf("[ERROR] Invalid time format (the date format follows RFC 3339): %s", err)
		}
		expiration = &expiration_time
	}

	key := kp.Key{
		Name:           plan.KeyName.ValueString(),
		Extractable:    plan.StandardKey.ValueBool(),
		Expiration:     expiration,
		Payload:        plan.Payload.ValueString(),
		Description:    plan.Description.ValueString(),
		EncryptedNonce: plan.EncryptedNonce.ValueString(),
		IV:             plan.IVValue.ValueString(),
	}

	return key, instanceID, nil
}

func (r *KMSKeyResource) readKeyToState(ctx context.Context, state *KMSKeyResourceModel, instanceID string) error {
	instanceCRN, _, keyid := getInstanceAndKeyDataFromCRN(state.Id.ValueString())

	kpAPI, _, err := r.populateKPClientForKey(*state, instanceID)
	if err != nil {
		return err
	}

	key, err := kpAPI.GetKey(ctx, keyid)
	if err != nil {
		return err
	}

	return r.setKeyDetails(ctx, state, instanceID, instanceCRN, key, kpAPI)
}

func (r *KMSKeyResource) setKeyDetails(ctx context.Context, state *KMSKeyResourceModel, instanceID string, instanceCRN string, key *kp.Key, kpAPI *kp.Client) error {
	state.InstanceID = basetypes.NewStringValue(instanceID)
	state.InstanceCRN = basetypes.NewStringValue(instanceCRN)
	state.KeyID = basetypes.NewStringValue(key.ID)
	state.StandardKey = basetypes.NewBoolValue(key.Extractable)
	state.Description = basetypes.NewStringValue(key.Description)
	state.EncryptedNonce = basetypes.NewStringValue(key.EncryptedNonce)
	state.IVValue = basetypes.NewStringValue(key.IV)
	state.KeyName = basetypes.NewStringValue(key.Name)
	state.Crn = basetypes.NewStringValue(key.CRN)

	if strings.Contains((kpAPI.URL).String(), "private") || strings.Contains(kpAPI.Config.BaseURL, "private") {
		state.EndpointType = basetypes.NewStringValue("private")
	} else {
		state.EndpointType = basetypes.NewStringValue("public")
	}

	state.Type = basetypes.NewStringValue(strings.Split(state.Id.ValueString(), ":")[4])
	state.KeyRingID = basetypes.NewStringValue(key.KeyRingID)

	if key.Expiration != nil {
		state.ExpirationDate = basetypes.NewStringValue(key.Expiration.Format(time.RFC3339))
	} else {
		state.ExpirationDate = basetypes.NewStringValue("")
	}

	state.ResourceName = basetypes.NewStringValue(key.Name)
	state.ResourceCRN = basetypes.NewStringValue(key.CRN)
	state.ResourceStatus = basetypes.NewStringValue(strconv.Itoa(key.State))

	rcontroller, err := flex.GetBaseController(r.client)
	if err != nil {
		return err
	}

	crn1 := strings.TrimSuffix(key.CRN, ":key:"+key.ID)
	state.ResourceControllerURL = basetypes.NewStringValue(rcontroller + "/services/kms/" + url.QueryEscape(crn1) + "%3A%3A")

	// Get registrations
	registrations, err := kpAPI.ListRegistrations(ctx, key.ID, "")
	if err != nil {
		return err
	}

	regList := make([]attr.Value, 0)
	for _, r := range registrations.Registrations {
		regObj, _ := types.ObjectValue(
			map[string]attr.Type{
				"key_id":               types.StringType,
				"resource_crn":         types.StringType,
				"prevent_key_deletion": types.BoolType,
			},
			map[string]attr.Value{
				"key_id":               basetypes.NewStringValue(r.KeyID),
				"resource_crn":         basetypes.NewStringValue(r.ResourceCrn),
				"prevent_key_deletion": basetypes.NewBoolValue(r.PreventKeyDeletion),
			},
		)
		regList = append(regList, regObj)
	}

	registrationsList, _ := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key_id":               types.StringType,
				"resource_crn":         types.StringType,
				"prevent_key_deletion": types.BoolType,
			},
		},
		regList,
	)
	state.Registrations = registrationsList

	return nil
}

type KMSKeyResourceModel struct {
	InstanceID            types.String `tfsdk:"instance_id"`
	KeyRingID             types.String `tfsdk:"key_ring_id"`
	KeyID                 types.String `tfsdk:"key_id"`
	KeyName               types.String `tfsdk:"key_name"`
	Type                  types.String `tfsdk:"type"`
	EndpointType          types.String `tfsdk:"endpoint_type"`
	Description           types.String `tfsdk:"description"`
	StandardKey           types.Bool   `tfsdk:"standard_key"`
	Payload               types.String `tfsdk:"payload"`
	EncryptedNonce        types.String `tfsdk:"encrypted_nonce"`
	IVValue               types.String `tfsdk:"iv_value"`
	ForceDelete           types.Bool   `tfsdk:"force_delete"`
	Crn                   types.String `tfsdk:"crn"`
	ExpirationDate        types.String `tfsdk:"expiration_date"`
	InstanceCRN           types.String `tfsdk:"instance_crn"`
	Registrations         types.List   `tfsdk:"registrations"`
	ResourceName          types.String `tfsdk:"resource_name"`
	ResourceCRN           types.String `tfsdk:"resource_crn"`
	ResourceStatus        types.String `tfsdk:"resource_status"`
	ResourceGroupName     types.String `tfsdk:"resource_group_name"`
	ResourceControllerURL types.String `tfsdk:"resource_controller_url"`
	Id                    types.String `tfsdk:"id"`
}
