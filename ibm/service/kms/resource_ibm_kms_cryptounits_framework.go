// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	kpCryptoUnit "github.com/IBM/keyprotect-go-client/dedicated"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &kmsCryptoUnitsResource{}
	_ resource.ResourceWithConfigure   = &kmsCryptoUnitsResource{}
	_ resource.ResourceWithImportState = &kmsCryptoUnitsResource{}
)

// NewKmsCryptoUnitsResource creates a new instance of the resource.
func NewKmsCryptoUnitsResource() resource.Resource {
	return &kmsCryptoUnitsResource{}
}

// kmsCryptoUnitsResource is the resource implementation.
type kmsCryptoUnitsResource struct {
	client conns.ClientSession
}

// kmsCryptoUnitsResourceModel describes the resource data model.
type kmsCryptoUnitsResourceModel struct {
	URL          types.String `tfsdk:"url"`
	UsePrivate   types.Bool   `tfsdk:"use_private_endpoint"`
	ID           types.String `tfsdk:"id"`
	InstanceID   types.String `tfsdk:"instance_id"`
	Region       types.String `tfsdk:"region"`
	CryptoUnits  types.Map    `tfsdk:"cryptounits"`
	SignatureKey types.Set    `tfsdk:"signature_key"`
	MasterKey    types.Set    `tfsdk:"master_key"`
}

// signatureKeyModel describes the signature key nested block.
type signatureKeyModel struct {
	Filepath   types.String `tfsdk:"filepath"`
	Passphrase types.String `tfsdk:"passphrase"`
	Owner      types.String `tfsdk:"owner"`
	Exists     types.Bool   `tfsdk:"exists"`
}

// masterKeyModel describes the master  key nested block.
type masterKeyModel struct {
	KeyShareFiles []keyShareFileModel `tfsdk:"keysharefile"`
	KeyName       types.String        `tfsdk:"keyname"`
	Exists        types.Bool          `tfsdk:"exists"`
}

// keyShareFileModel describes individual key share file entries.
type keyShareFileModel struct {
	Filepath types.String `tfsdk:"filepath"`
	Token    types.String `tfsdk:"token"`
}

// Metadata returns the resource type name.
func (r *kmsCryptoUnitsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_cryptounits"
}

// Schema defines the schema for the resource.
func (r *kmsCryptoUnitsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages IBM Key Protect or HPCS crypto units initialization. This resource initializes crypto units with signature keys and master  keys.",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL to use when targeting the resource",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the resource (same as instance_id).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"instance_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key protect or HPCS instance GUID or CRN.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Area where the key protect dedicated instance resides.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"use_private_endpoint": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to use private endpoint.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"cryptounits": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Map of crypto unit IDs to their current states.",
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"signature_key": schema.SetNestedBlock{
				Description: "Credentials for the user to create sessions with the crypto units.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"filepath": schema.StringAttribute{
							Required:    true,
							Description: "The filepath to store the signature key.",
						},
						"passphrase": schema.StringAttribute{
							Required:    true,
							Sensitive:   true,
							Description: "The passphrase of the signature key.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"owner": schema.StringAttribute{
							Required:    true,
							Description: "The owner of the signature key.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"exists": schema.BoolAttribute{
							Required:    true,
							Description: "True if the file in the filepath, False if it should be generated.",
							PlanModifiers: []planmodifier.Bool{
								forceNewIfFalseModifier{},
							},
						},
					},
				},
			},
			"master_key": schema.SetNestedBlock{
				Description: "Attributes related to the master  key.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"keyname": schema.StringAttribute{
							Required:    true,
							Description: "The name of the master  key shown on the crypto unit (max 8 characters).",
						},
						"exists": schema.BoolAttribute{
							Required:    true,
							Description: "True if the files are present the keysharefile.filepath, False if it should be generated.",
							PlanModifiers: []planmodifier.Bool{
								forceNewIfFalseModifier{},
							},
						},
					},
					Blocks: map[string]schema.Block{
						"keysharefile": schema.SetNestedBlock{
							Description: "Key share file configuration with filepath and token.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"filepath": schema.StringAttribute{
										Required:    true,
										Description: "The filepath to store the key share file.",
									},
									"token": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: "The token associated with the key share file.",
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

type forceNewIfFalseModifier struct{}

func (m forceNewIfFalseModifier) Description(_ context.Context) string {
	return "Forces replacement if the boolean value is set to false."
}

func (m forceNewIfFalseModifier) MarkdownDescription(_ context.Context) string {
	return "Forces replacement if the boolean value is set to false."
}

func (m forceNewIfFalseModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// Do nothing if the value is unknown or null to avoid conflicting with other logic
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}

	// Force replacement if the planned value is false
	if !req.PlanValue.ValueBool() {
		resp.RequiresReplace = true
	}
}

// Configure adds the provider configured client to the resource.
func (r *kmsCryptoUnitsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected conns.ClientSession, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

// Create creates the resource and sets the initial Terraform state.
func (r *kmsCryptoUnitsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan kmsCryptoUnitsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract region and instance ID
	url := plan.URL.ValueString()
	region := plan.Region.ValueString()
	instanceID := plan.InstanceID.ValueString()
	usePrivate := plan.UsePrivate.ValueBool()

	kpOpts, err := createKPCryptoOpts(url, region, instanceID, usePrivate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client options: %s", err.Error()),
		)
		return
	}
	// Initialize KMS crypto unit client
	kmsCryptoUnitClient, err := r.client.KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client: %s", err.Error()),
		)
		return
	}

	// Parse master  key configuration
	mbkSpec, diags := r.parseMasterKey(ctx, plan.MasterKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Parse signature key configuration
	sigKeySpec, diags := r.parseSignatureKey(ctx, plan.SignatureKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize crypto units
	err = kmsCryptoUnitClient.InitializeCryptoUnits(ctx, sigKeySpec, mbkSpec, instanceID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize Crypto Units",
			fmt.Sprintf("Unable to initialize crypto units for instance %s: %s", instanceID, err.Error()),
		)
		return
	}

	// Set resource ID
	plan.ID = types.StringValue(instanceID)

	// Ensure computed attributes are set from plan values
	plan.InstanceID = types.StringValue(instanceID)
	plan.Region = types.StringValue(region)
	plan.UsePrivate = types.BoolValue(usePrivate)

	// Read crypto units state
	r.readCryptoUnits(ctx, &plan, kmsCryptoUnitClient, instanceID, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *kmsCryptoUnitsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state kmsCryptoUnitsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract region and instance ID
	url := state.URL.ValueString()
	region := state.Region.ValueString()
	instanceID := state.ID.ValueString()
	usePrivate := state.UsePrivate.ValueBool()

	kpOpts, err := createKPCryptoOpts(url, region, instanceID, usePrivate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client options: %s", err.Error()),
		)
		return
	}
	// Initialize KMS crypto unit client
	kmsCryptoUnitClient, err := r.client.KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client: %s", err.Error()),
		)
		return
	}
	// Read crypto units state
	r.readCryptoUnits(ctx, &state, kmsCryptoUnitClient, instanceID, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *kmsCryptoUnitsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan kmsCryptoUnitsResourceModel
	var state kmsCryptoUnitsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read current state for delete operation
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 1: Zeroize existing crypto units (Delete logic)
	url := state.URL.ValueString()
	region := state.Region.ValueString()
	instanceID := state.ID.ValueString()
	usePrivate := state.UsePrivate.ValueBool()

	kpOpts, err := createKPCryptoOpts(url, region, instanceID, usePrivate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client options for zeroization: %s", err.Error()),
		)
		return
	}

	kmsCryptoUnitClient, err := r.client.KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client for zeroization: %s", err.Error()),
		)
		return
	}

	// List and zeroize crypto units
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to List Crypto Units",
			fmt.Sprintf("Unable to list crypto units for instance %s during update: %s", instanceID, err.Error()),
		)
		return
	}

	for _, cryptoUnit := range cryptoUnitsResponse.CryptoUnits {
		err := kmsCryptoUnitClient.ZeroizeCryptoUnitWithContext(ctx, cryptoUnit.ID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failed to Zeroize Crypto Unit",
				fmt.Sprintf("Failed to update cryptounits upon zeroization of unit %s: %s", cryptoUnit.ID, err.Error()),
			)
			return
		}
	}

	// Step 2: Reinitialize crypto units with new configuration (Create logic)
	url = plan.URL.ValueString()
	region = plan.Region.ValueString()
	instanceID = plan.InstanceID.ValueString()
	usePrivate = plan.UsePrivate.ValueBool()

	kpOpts, err = createKPCryptoOpts(url, region, instanceID, usePrivate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client options for reinitialization: %s", err.Error()),
		)
		return
	}

	kmsCryptoUnitClient, err = r.client.KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client for reinitialization: %s", err.Error()),
		)
		return
	}

	// Parse master key configuration
	mbkSpec, diags := r.parseMasterKey(ctx, plan.MasterKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "master key plan",
		map[string]interface{}{
			"master key plan": plan.MasterKey,
		})

	// Parse signature key configuration
	sigKeySpec, diags := r.parseSignatureKey(ctx, plan.SignatureKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "signature key plan",
		map[string]interface{}{
			"signature key plan": plan.SignatureKey,
		})

	// Initialize crypto units
	err = kmsCryptoUnitClient.InitializeCryptoUnits(ctx, sigKeySpec, mbkSpec, instanceID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Reinitialize Crypto Units",
			fmt.Sprintf("Failed to update cryptounits upon reinitialization for instance %s: %s", instanceID, err.Error()),
		)
		return
	}

	// Set resource ID
	plan.ID = types.StringValue(instanceID)

	// Read crypto units state
	r.readCryptoUnits(ctx, &plan, kmsCryptoUnitClient, instanceID, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *kmsCryptoUnitsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state kmsCryptoUnitsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract region and instance ID

	// Initialize KMS crypto unit client
	url := state.URL.ValueString()
	region := state.Region.ValueString()
	instanceID := state.ID.ValueString()
	usePrivate := state.UsePrivate.ValueBool()

	kpOpts, err := createKPCryptoOpts(url, region, instanceID, usePrivate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client options: %s", err.Error()),
		)
		return
	}
	kmsCryptoUnitClient, err := r.client.KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Initialize KMS Crypto Unit Client",
			fmt.Sprintf("Unable to create KMS crypto unit client: %s", err.Error()),
		)
		return
	}

	// List crypto units
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to List Crypto Units",
			fmt.Sprintf("Unable to list crypto units for instance %s: %s", instanceID, err.Error()),
		)
		return
	}

	// Zeroize each crypto unit
	for _, cryptoUnit := range cryptoUnitsResponse.CryptoUnits {
		err := kmsCryptoUnitClient.ZeroizeCryptoUnitWithContext(ctx, cryptoUnit.ID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failed to Zeroize Crypto Unit",
				fmt.Sprintf("Unable to zeroize crypto unit %s: %s", cryptoUnit.ID, err.Error()),
			)
			return
		}
	}
}

// ImportState imports the resource state.
func (r *kmsCryptoUnitsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// readCryptoUnits reads the current state of crypto units and updates the model.
func (r *kmsCryptoUnitsResource) readCryptoUnits(ctx context.Context, model *kmsCryptoUnitsResourceModel, client interface{}, instanceID string, diagnostics *diag.Diagnostics) {
	// Type assert the client to the actual KeyProtectCryptoUnitAPI type
	kmsCryptoUnitClient, ok := client.(*kpCryptoUnit.KeyProtectCryptoUnitAPI)
	if !ok {
		diagnostics.AddError(
			"Invalid Client Type",
			fmt.Sprintf("Unable to cast client to *kpCryptoUnit.KeyProtectCryptoUnitAPI, got: %T", client),
		)
		return
	}

	// List crypto units
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		diagnostics.AddError(
			"Failed to List Crypto Units",
			fmt.Sprintf("Unable to list crypto units for instance %s: %s", instanceID, err.Error()),
		)
		return
	}

	// Transform response into map format
	cryptoUnitsMap := make(map[string]attr.Value)
	if cryptoUnitsResponse.CryptoUnits != nil {
		for _, cu := range cryptoUnitsResponse.CryptoUnits {
			if cu.ID != "" && cu.State != "" {
				cryptoUnitsMap[cu.ID] = types.StringValue(string(cu.State))
			}
		}
	}

	// Set the cryptounits field
	mapValue, diags := types.MapValue(types.StringType, cryptoUnitsMap)
	diagnostics.Append(diags...)
	if diagnostics.HasError() {
		return
	}

	model.CryptoUnits = mapValue
}

// parseMasterKey extracts and validates master  key configuration.
func (r *kmsCryptoUnitsResource) parseMasterKey(ctx context.Context, set types.Set) (*kpCryptoUnit.MasterKeyPartsSpec, diag.Diagnostics) {
	var diags diag.Diagnostics
	var mbkList []masterKeyModel

	diags.Append(set.ElementsAs(ctx, &mbkList, false)...)
	if diags.HasError() {
		return nil, diags
	}

	if len(mbkList) == 0 {
		diags.AddError(
			"Invalid Master  Key Configuration",
			"Master  key block is required",
		)
		return nil, diags
	}

	mbk := mbkList[0]

	if len(mbk.KeyShareFiles) == 0 {
		diags.AddError(
			"Invalid Master  Key Configuration",
			"At least one key share file must be provided",
		)
		return nil, diags
	}

	// Track unique filepaths
	filepathMap := make(map[string]bool)
	keyShareFileEntries := make([]string, 0, len(mbk.KeyShareFiles))

	// Process each key share file
	for i, ksf := range mbk.KeyShareFiles {
		filePath := ksf.Filepath.ValueString()
		token := ksf.Token.ValueString()

		if filePath == "" {
			diags.AddError(
				"Invalid Key Share File Configuration",
				fmt.Sprintf("Filepath in keysharefile[%d] cannot be empty", i),
			)
			return nil, diags
		}

		if token == "" {
			diags.AddError(
				"Invalid Key Share File Configuration",
				fmt.Sprintf("Token in keysharefile[%d] cannot be empty", i),
			)
			return nil, diags
		}

		// Resolve relative path
		resolvedPath, err := resolveRelativePathFramework(filePath)
		if err != nil {
			diags.AddError(
				"Failed to Resolve Filepath",
				fmt.Sprintf("Unable to resolve filepath in keysharefile[%d]: %s", i, err.Error()),
			)
			return nil, diags
		}

		// Check for duplicates
		if filepathMap[resolvedPath] {
			diags.AddError(
				"Duplicate Filepath Detected",
				fmt.Sprintf("Duplicate filepath in keysharefile[%d]: %s", i, filePath),
			)
			return nil, diags
		}
		filepathMap[resolvedPath] = true

		// Combine filepath and token
		keyShareFileEntry := fmt.Sprintf("%s#%s", resolvedPath, token)
		keyShareFileEntries = append(keyShareFileEntries, keyShareFileEntry)
	}

	// Extract and validate keyname
	keyName := mbk.KeyName.ValueString()
	if keyName == "" {
		diags.AddError(
			"Invalid Master Key Configuration",
			"Keyname cannot be empty",
		)
		return nil, diags
	}

	if len(keyName) > 8 {
		diags.AddError(
			"Invalid Master Key Configuration",
			"Keyname must be 8 characters or less",
		)
		return nil, diags
	}

	// Extract and validate existing
	in := mbk.Exists.IsNull()
	if in {
		diags.AddError(
			"Invalid Master Key Configuration",
			"Existing must be specified",
		)
		return nil, diags
	}
	exists := mbk.Exists.ValueBool()

	return &kpCryptoUnit.MasterKeyPartsSpec{
		KeyShareFiles: keyShareFileEntries,
		SlotNo:        3,
		K:             uint8(len(keyShareFileEntries)),
		KeyName:       keyName,
		Exists:        exists,
	}, diags
}

// parseSignatureKey extracts and validates signature key configuration.
func (r *kmsCryptoUnitsResource) parseSignatureKey(ctx context.Context, set types.Set) (*kpCryptoUnit.SignatureKeyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	var sigKeyList []signatureKeyModel

	diags.Append(set.ElementsAs(ctx, &sigKeyList, false)...)
	if diags.HasError() {
		return nil, diags
	}

	if len(sigKeyList) == 0 {
		diags.AddError(
			"Invalid Signature Key Configuration",
			"Signature key block is required",
		)
		return nil, diags
	}

	sigKey := sigKeyList[0]

	// Extract and validate filepath
	filePath := sigKey.Filepath.ValueString()
	if filePath == "" {
		diags.AddError(
			"Invalid Signature Key Configuration",
			"Filepath cannot be empty",
		)
		return nil, diags
	}

	// Resolve relative path
	resolvedFilePath, err := resolveRelativePathFramework(filePath)
	if err != nil {
		diags.AddError(
			"Failed to Resolve Filepath",
			fmt.Sprintf("Unable to resolve signature key filepath: %s", err.Error()),
		)
		return nil, diags
	}
	tflog.Info(
		ctx, "Using signature key filepath",
		map[string]interface{}{
			"filepath": resolvedFilePath,
		},
	)

	// Extract passphrase
	passphrase := sigKey.Passphrase.ValueString()

	// Extract and validate owner
	owner := sigKey.Owner.ValueString()
	if owner == "" {
		diags.AddError(
			"Invalid Signature Key Configuration",
			"Owner cannot be empty",
		)
		return nil, diags
	}

	in := sigKey.Exists.IsNull()
	if in {
		diags.AddError(
			"Invalid Signature Key Configuration",
			"Existing must be specified",
		)
		return nil, diags
	}
	exists := sigKey.Exists.ValueBool()

	return &kpCryptoUnit.SignatureKeyRequest{
		FilePath:   resolvedFilePath,
		Passphrase: passphrase,
		Algorithm:  kpCryptoUnit.SigKeyAlgorithmRSA2048,
		Owner:      owner,
		Exists:     exists,
	}, diags
}

// resolveRelativePathFramework converts a path to be relative to the current working directory.
func resolveRelativePathFramework(inputPath string) (string, error) {
	// If path is already absolute, return as-is
	if filepath.IsAbs(inputPath) {
		return inputPath, nil
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Join and clean the path
	resolvedPath := filepath.Join(cwd, inputPath)
	return filepath.Clean(resolvedPath), nil
}

// Made with Bob
