// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyprotect_dedicated "github.com/IBM/keyprotect-go-client/dedicated"
)

func ResourceIBMKmsCryptoUnits() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMKmsCryptoUnitsCreate,
		ReadContext:   resourceIBMKmsCryptoUnitsRead,
		UpdateContext: resourceIBMKmsCryptoUnitsUpdate,
		DeleteContext: resourceIBMKmsCryptoUnitsDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URL to use when targeting the resource",
			},
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Key protect or hpcs instance GUID or CRN",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"region": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Description:      "area where the key protect dedicated instance resides",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"use_private_endpoint": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "set to true if the private endpoint should be used, otherwise false",
				Default:     false,
			},
			"should_zeroize": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "set to true if the resource should be zeroized. Zeroizing if harmful if not understood. Set to false as a safe default",
				Default:     false,
			},
			"cryptounits": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"signature_key": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    1,
				Description: "Credentials for the user will use to create sessions with the cryptounits",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filepath": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The filepath to store the signature key",
						},
						"passphrase": {
							Type:             schema.TypeString,
							Required:         true,
							Sensitive:        true,
							Description:      "The passphrase of the signature_key",
							DiffSuppressFunc: suppressTokenDiff,
						},
						"owner": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The owner of the signature_key",
						},
						"exists": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set to true if the signature_key file exists from the filepath, false if it should be generated",
						},
					},
				},
			},
			"master_key": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    1,
				Description: "Attributes related to the master backup key",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keysharefile": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "Key share file configuration with filepath and token",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filepath": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The filepath to store the key share file",
									},
									"token": {
										Type:             schema.TypeString,
										Required:         true,
										Sensitive:        true,
										Description:      "The token associated with the key share file",
										DiffSuppressFunc: suppressTokenDiff,
									},
								},
							},
						},
						"keyname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the master backup key shown on the cryptounit",
						},
						"exists": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set to true if keysharefiles are all existing files, false if it should be generated",
						},
					},
				},
			},
		},
	}
}

func resourceIBMKmsCryptoUnitsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpOpts, err := createKPCryptoOptsFromV2(ctx, d)
	if err != nil {
		return diag.FromErr(err)
	}

	// Initialize KMS crypto unit client
	kmsCryptoUnitClient, err := meta.(conns.ClientSession).KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		return diag.Errorf("failed to initialize KMS crypto unit client: %v", err)
	}

	// Parse and validate master backup key configuration
	masterKeySpec, err := parseMasterKey(d)
	if err != nil {
		return diag.Errorf("failed to parse master_key: %v", err)
	}

	// Parse and validate signature key configuration
	sigKeySpec, err := parseSignatureKey(d)
	if err != nil {
		return diag.Errorf("failed to parse signature_key: %v", err)
	}

	// Initialize crypto units with proper error context
	err = kmsCryptoUnitClient.InitializeCryptoUnits(ctx, sigKeySpec, masterKeySpec, kpOpts.InstanceID)
	if err != nil {
		return diag.Errorf("failed to initialize crypto units for instance %s: %s", kpOpts.InstanceID, err.Error())
	}

	// Set resource ID for state management
	d.SetId(kpOpts.InstanceID)
	kmsCryptoUnitClient.DisconnectAll()

	return resourceIBMKmsCryptoUnitsRead(ctx, d, meta)
}

// parseMasterKey extracts and validates master backup key configuration from TypeSet
func parseMasterKey(d *schema.ResourceData) (*keyprotect_dedicated.MasterKeyPartsSpec, error) {
	mbkSet, ok := d.GetOk("master_key")
	if !ok {
		return nil, fmt.Errorf("master_key is required but not provided")
	}

	mbkSetTyped, ok := mbkSet.(*schema.Set)
	if !ok {
		return nil, fmt.Errorf("master_key must be a set")
	}

	mbkList := mbkSetTyped.List()
	if len(mbkList) == 0 {
		return nil, fmt.Errorf("master_key must contain at least one element")
	}

	mbkData, ok := mbkList[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("master_key[0] must be a map")
	}

	// Extract and validate keysharefile blocks
	keyShareFileRaw, ok := mbkData["keysharefile"]
	if !ok {
		return nil, fmt.Errorf("keysharefile is required in master_key")
	}

	keyShareFileSet, ok := keyShareFileRaw.(*schema.Set)
	if !ok {
		return nil, fmt.Errorf("keysharefile must be a set")
	}

	keyShareFileList := keyShareFileSet.List()
	if len(keyShareFileList) == 0 {
		return nil, fmt.Errorf("keysharefile must contain at least one block")
	}

	// Track unique filepaths to ensure no duplicates
	filepathMap := make(map[string]bool)
	keyShareFiles := make([]string, 0, len(keyShareFileList))

	// Process each keysharefile block
	for i, item := range keyShareFileList {
		keyShareFileData, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("keysharefile[%d] must be a map", i)
		}

		// Extract filepath
		filePathRaw, ok := keyShareFileData["filepath"]
		if !ok {
			return nil, fmt.Errorf("filepath is required in keysharefile[%d]", i)
		}

		filePath, ok := filePathRaw.(string)
		if !ok {
			return nil, fmt.Errorf("filepath in keysharefile[%d] must be a string", i)
		}
		if filePath == "" {
			return nil, fmt.Errorf("filepath in keysharefile[%d] cannot be empty", i)
		}

		// Resolve relative path to absolute path
		resolvedPath, err := resolveRelativePath(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve filepath in keysharefile[%d]: %v", i, err)
		}

		// Check for duplicate filepaths
		if filepathMap[resolvedPath] {
			return nil, fmt.Errorf("duplicate filepath detected in keysharefile[%d]: %s", i, filePath)
		}
		filepathMap[resolvedPath] = true

		// Extract token
		tokenRaw, ok := keyShareFileData["token"]
		if !ok {
			return nil, fmt.Errorf("token is required in keysharefile[%d]", i)
		}

		token, ok := tokenRaw.(string)
		if !ok {
			return nil, fmt.Errorf("token in keysharefile[%d] must be a string", i)
		}
		if token == "" {
			return nil, fmt.Errorf("token in keysharefile[%d] cannot be empty", i)
		}

		// Combine filepath and token in the format expected by the API
		keyShareFileEntry := fmt.Sprintf("%s#%s", resolvedPath, token)
		keyShareFiles = append(keyShareFiles, keyShareFileEntry)
	}

	// Extract and validate keyname
	keyNameRaw, ok := mbkData["keyname"]
	if !ok {
		return nil, fmt.Errorf("keyname is required in master_key")
	}

	keyName, ok := keyNameRaw.(string)
	if !ok {
		return nil, fmt.Errorf("keyname must be a string")
	}
	if keyName == "" {
		return nil, fmt.Errorf("keyname cannot be empty")
	}
	if len(keyName) > 8 {
		return nil, fmt.Errorf("keyname must be 8 characters or less")
	}

	// Extract and validate exists
	keysExistsRaw, ok := mbkData["exists"]
	if !ok {
		return nil, fmt.Errorf("exists is required in master_key")
	}

	keyExists, ok := keysExistsRaw.(bool)
	if !ok {
		return nil, fmt.Errorf("keyname must be a bool")
	}

	return keyprotect_dedicated.NewMasterKeyPartsSpec(
		len(keyShareFiles),
		keyName,
		keyShareFiles,
		keyExists,
	)
}

// parseSignatureKey extracts and validates signature key configuration from TypeSet
func parseSignatureKey(d *schema.ResourceData) (*keyprotect_dedicated.SignatureKeyRequest, error) {
	sigKeySet, ok := d.GetOk("signature_key")
	if !ok {
		return nil, fmt.Errorf("signature_key is required but not provided")
	}

	sigKeySetTyped, ok := sigKeySet.(*schema.Set)
	if !ok {
		return nil, fmt.Errorf("signature_key must be a set")
	}

	sigKeyList := sigKeySetTyped.List()
	if len(sigKeyList) == 0 {
		return nil, fmt.Errorf("signature_key must contain at least one element")
	}

	sigKeyData, ok := sigKeyList[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("signature_key[0] must be a map")
	}

	// Extract and validate filepath
	filePathRaw, ok := sigKeyData["filepath"]
	if !ok {
		return nil, fmt.Errorf("filepath is required in signature_key")
	}

	filePath, ok := filePathRaw.(string)
	if !ok {
		return nil, fmt.Errorf("filepath must be a string")
	}
	if filePath == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}

	// Resolve relative path to absolute path based on terraform execution directory
	resolvedFilePath, err := resolveRelativePath(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve signature_key filepath: %v", err)
	}

	// Extract and validate passphrase
	passphraseRaw, ok := sigKeyData["passphrase"]
	if !ok {
		return nil, fmt.Errorf("passphrase is required in signature_key")
	}

	passphrase, ok := passphraseRaw.(string)
	if !ok {
		return nil, fmt.Errorf("passphrase must be a string")
	}

	// Extract and validate owner
	ownerRaw, ok := sigKeyData["owner"]
	if !ok {
		return nil, fmt.Errorf("owner is required in signature_key")
	}

	owner, ok := ownerRaw.(string)
	if !ok {
		return nil, fmt.Errorf("owner must be a string")
	}
	if owner == "" {
		return nil, fmt.Errorf("owner cannot be empty")
	}

	// Extract and validate filepath
	existsRaw, ok := sigKeyData["exists"]
	if !ok {
		return nil, fmt.Errorf("exists is required in signature_key")
	}

	fileExists, ok := existsRaw.(bool)
	if !ok {
		return nil, fmt.Errorf("filepath must be a bool")
	}

	return keyprotect_dedicated.NewSignatureKeyRequest(resolvedFilePath, passphrase, owner, fileExists)
}

// resolveRelativePath converts a path to be relative to the current working directory
// where terraform is executed. If the path is already absolute, it returns it unchanged.
func resolveRelativePath(inputPath string) (string, error) {
	// If path is already absolute, return as-is
	if filepath.IsAbs(inputPath) {
		return inputPath, nil
	}

	// Get current working directory (where terraform is run)
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Join the cwd with the relative path and clean it
	resolvedPath := filepath.Join(cwd, inputPath)
	return filepath.Clean(resolvedPath), nil
}

// suppressTokenDiff always suppresses token differences to prevent storing tokens in state.
// Since tokens are only needed during resource creation and the API doesn't return them,
// we suppress all diffs to avoid storing the token value in the state file.
func suppressTokenDiff(k, old, new string, d *schema.ResourceData) bool {
	// Always suppress token diffs - tokens are write-only and should not be stored in state
	return true
}

func resourceIBMKmsCryptoUnitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpOpts, err := createKPCryptoOptsFromV2(ctx, d)
	if err != nil {
		return diag.FromErr(err)
	}

	// Initialize KMS crypto unit client
	kmsCryptoUnitClient, err := meta.(conns.ClientSession).KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		return diag.Errorf("failed to initialize KMS crypto unit client: %v", err)
	}

	// List crypto units using the context-aware function
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		return diag.Errorf("failed to list crypto units for instance %s: %v", kpOpts.InstanceID, err)
	}

	// Transform response into map format for schema (ID -> State mapping)
	cryptoUnitsMap := make(map[string]interface{})
	if cryptoUnitsResponse.CryptoUnits != nil {
		for _, cu := range cryptoUnitsResponse.CryptoUnits {
			if cu.ID != "" && cu.State != "" {
				cryptoUnitsMap[cu.ID] = string(cu.State)
			}
		}
	}

	// Set the cryptounits field in resource data
	if err := d.Set("cryptounits", cryptoUnitsMap); err != nil {
		return diag.Errorf("failed to set cryptounits: %v", err)
	}

	// Set other fields to maintain state consistency
	d.Set("instance_id", kpOpts.InstanceID)
	d.Set("region", kpOpts.Region)

	return nil
}

func resourceIBMKmsCryptoUnitsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if deleteErr := resourceIBMKmsCryptoUnitsDelete(ctx, d, meta); deleteErr != nil {
		return diag.Errorf("failed to update cryptounits upon zerolization: %v", deleteErr)
	}

	if createErr := resourceIBMKmsCryptoUnitsCreate(ctx, d, meta); createErr != nil {
		return diag.Errorf("failed to update cryptounits upon reinitialization: %v", createErr)
	}

	return nil
}

func resourceIBMKmsCryptoUnitsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpOpts, err := createKPCryptoOptsFromV2(ctx, d)
	if err != nil {
		return diag.FromErr(err)
	}

	// Initialize KMS crypto unit client
	kmsCryptoUnitClient, err := meta.(conns.ClientSession).KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		return diag.Errorf("failed to initialize KMS crypto unit client: %v", err)
	}

	// List crypto units using the context-aware function
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		return diag.Errorf("failed to list crypto units for instance %s: %v", kpOpts.InstanceID, err)
	}
	// Zeroize each crypto unit
	for _, cryptoUnit := range cryptoUnitsResponse.CryptoUnits {
		err := kmsCryptoUnitClient.ZeroizeCryptoUnitWithContext(ctx, cryptoUnit.ID)
		if err != nil {
			return diag.Errorf("failed to zeroize crypto unit %s: %v", cryptoUnit.ID, err)
		}
	}
	// No-op for delete - resource is read-only
	d.SetId("")
	return nil
}

func createKPCryptoOptsFromV2(ctx context.Context, d *schema.ResourceData) (*keyprotect_dedicated.KeyProtectCryptoUnitAPIOptions, error) {
	// Extract and validate required parameters
	var url string
	var region string
	var instanceID string
	var isPrivate bool

	urlStr, urlChk := d.GetOk("url")
	if urlChk {
		url = urlStr.(string)
	}

	regionStr, regionChk := d.GetOk("region")
	if regionChk {
		region = regionStr.(string)
	}

	instanceIDStr, instChk := d.GetOk("instance_id")
	if instChk {
		instanceID = instanceIDStr.(string)
	}

	isPrivateChk, ok := d.GetOk("use_private_endpoint")
	if ok {
		isPrivate = isPrivateChk.(bool)
	}
	return createKPCryptoOpts(ctx, url, region, instanceID, isPrivate)
}

func createKPCryptoOpts(ctx context.Context, url, region, instanceID string, isPrivate bool) (*keyprotect_dedicated.KeyProtectCryptoUnitAPIOptions, error) {
	var isGood bool

	regionChk := len(region) > 0
	instChk := len(instanceID) > 0
	urlChk := len(url) > 0

	isGood = regionChk && instChk || urlChk

	if !isGood {
		return nil, fmt.Errorf("The following requirements must be met: instance_id and region must be set, url must be set")
	}
	if !urlChk {
		if urlStr, err := keyprotect_dedicated.GetServiceURLForRegion(instanceID, region, isPrivate); err != nil {
			return nil, fmt.Errorf("failed to create KeyProtectCryptoUnitAPIOptions using instance_id and region: %w", err)
		} else {
			url = urlStr
		}
	}
	kpOpts, err := keyprotect_dedicated.NewKeyProtectCryptoUnitAPIOptions(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create KeyProtectCryptoUnitAPIOptions using url: %w", err)
	}

	return kpOpts, nil
}
