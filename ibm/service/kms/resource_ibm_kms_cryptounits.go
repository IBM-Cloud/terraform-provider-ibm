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
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
				Computed:    true,
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
				Computed:         true,
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
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the crypto unit",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current state of the crypto unit",
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return schema.HashString(m["id"].(string))
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
							Description: "The filepath to store the signature key or find an exisiting signature key",
						},
						"passphrase": {
							Type:             schema.TypeString,
							Required:         true,
							Sensitive:        true,
							Description:      "The passphrase of the signature_key",
							DiffSuppressFunc: suppresspassphraseDiff,
						},
						"owner": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The owner of the signature_key",
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
							Description: "Key share file configuration with filepath and passphrase",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filepath": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The filepath to store the key share file or find an existing one",
									},
									"passphrase": {
										Type:             schema.TypeString,
										Required:         true,
										Sensitive:        true,
										Description:      "The passphrase associated with the key share file",
										DiffSuppressFunc: suppresspassphraseDiff,
									},
								},
							},
						},
						"keyname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the master backup key shown on the cryptounit",
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

	// Inspect current unit states before deciding whether to initialize.
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		return diag.Errorf("failed to list crypto units for instance %s: %v", kpOpts.InstanceID, err)
	}

	allReserved := len(cryptoUnitsResponse.CryptoUnits) > 0
	allKMSInitialized := len(cryptoUnitsResponse.CryptoUnits) > 0
	for _, cu := range cryptoUnitsResponse.CryptoUnits {
		if cu.State != keyprotect_dedicated.CryptoUnitStateReserved {
			allReserved = false
		}
		if cu.State != keyprotect_dedicated.CryptoUnitStateKMSInitialized {
			allKMSInitialized = false
		}
	}

	switch {
	case allKMSInitialized:
		// All units already fully initialized — idempotent re-run, skip to Read.

	case allReserved:
		// Clean slate — proceed with initialization.
		err = kmsCryptoUnitClient.InitializeCryptoUnits(ctx, sigKeySpec, masterKeySpec, kpOpts.InstanceID)
		if err != nil {
			return diag.Errorf("failed to initialize crypto units for instance %s: %s", kpOpts.InstanceID, err.Error())
		}

	default:
		// Units are in a mixed or unexpected state (e.g. partial failure).
		// Collect the offending states to surface a helpful error.
		stateSet := make(map[string]struct{})
		for _, cu := range cryptoUnitsResponse.CryptoUnits {
			stateSet[string(cu.State)] = struct{}{}
		}
		states := make([]string, 0, len(stateSet))
		for s := range stateSet {
			states = append(states, s)
		}
		return diag.Errorf(
			"crypto units for instance %s are in an unexpected state %v — "+
				"expected all %q (clean) or all %q (already initialized). "+
				"Zeroize the instance and retry.",
			kpOpts.InstanceID, states,
			keyprotect_dedicated.CryptoUnitStateReserved,
			keyprotect_dedicated.CryptoUnitStateKMSInitialized,
		)
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

		// Extract passphrase
		passphraseRaw, ok := keyShareFileData["passphrase"]
		if !ok {
			return nil, fmt.Errorf("passphrase is required in keysharefile[%d]", i)
		}

		passphrase, ok := passphraseRaw.(string)
		if !ok {
			return nil, fmt.Errorf("passphrase in keysharefile[%d] must be a string", i)
		}
		if passphrase == "" {
			return nil, fmt.Errorf("passphrase in keysharefile[%d] cannot be empty", i)
		}

		// Combine filepath and passphrase in the format expected by the API
		keyShareFileEntry := fmt.Sprintf("%s#%s", resolvedPath, passphrase)
		keyShareFiles = append(keyShareFiles, keyShareFileEntry)
	}

	// keyExists is true only when every keysharefile path exists on disk
	keyExists := true
	for path := range filepathMap {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			keyExists = false
			break
		}
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
	// Check whether the resolved filepath actually exists on disk
	_, statErr := os.Stat(resolvedFilePath)
	fileExists := !os.IsNotExist(statErr)

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

// suppresspassphraseDiff suppresses passphrase/passphrase diffs after the initial creation.
// On first create, old is "" and new has the real value — we must NOT suppress this
// or the SDK will strip the value from the diff and d.Get("passphrase") returns "" in Create.
// After creation, the API never returns these values, so old will always be "" in state
// and we suppress to avoid a permanent diff.
func suppresspassphraseDiff(k, old, new string, d *schema.ResourceData) bool {
	// Only suppress when the resource already exists (has an ID) and the stored
	// value is empty (because the API doesn't return it). This prevents a
	// perpetual diff without breaking the initial create.
	return d.Id() != "" && old == ""
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

	// Transform response into a slice of {id, state} objects for TypeSet schema
	cryptoUnitsList := make([]interface{}, 0)
	if cryptoUnitsResponse.CryptoUnits != nil {
		for _, cu := range cryptoUnitsResponse.CryptoUnits {
			if cu.ID != "" {
				cryptoUnitsList = append(cryptoUnitsList, map[string]interface{}{
					"id":    cu.ID,
					"state": string(cu.State),
				})
			}
		}
	}

	// Set the cryptounits field in resource data
	if err := d.Set("cryptounits", cryptoUnitsList); err != nil {
		return diag.Errorf("failed to set cryptounits: %v", err)
	}

	// Set other fields to maintain state consistency
	d.Set("instance_id", kpOpts.InstanceID)
	d.Set("region", kpOpts.Region)

	// Preserve write-only fields that the API does not return.
	// master_key and signature_key contain sensitive passphrases/passphrases that
	// are never returned by the API. If we do not explicitly re-set them here,
	// the TypeSet hash machinery will see empty strings for Sensitive fields
	// and corrupt state, causing "passphrase cannot be empty" on subsequent plans.
	if v, ok := d.GetOk("master_key"); ok {
		d.Set("master_key", v)
	}
	if v, ok := d.GetOk("signature_key"); ok {
		d.Set("signature_key", v)
	}

	return nil
}

func resourceIBMKmsCryptoUnitsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpOpts, err := createKPCryptoOptsFromV2(ctx, d)
	if err != nil {
		return diag.FromErr(err)
	}

	kmsCryptoUnitClient, err := meta.(conns.ClientSession).KeyProtectCryptoUnitAPI(ctx, kpOpts)
	if err != nil {
		return diag.Errorf("failed to initialize KMS crypto unit client: %v", err)
	}

	// Inspect current unit states to decide whether zeroization is needed.
	cryptoUnitsResponse, _, err := kmsCryptoUnitClient.ListCryptoUnitsWithContext(ctx)
	if err != nil {
		return diag.Errorf("failed to list crypto units for instance %s: %v", kpOpts.InstanceID, err)
	}

	allKMSInitialized := len(cryptoUnitsResponse.CryptoUnits) > 0
	for _, cu := range cryptoUnitsResponse.CryptoUnits {
		if cu.State != keyprotect_dedicated.CryptoUnitStateKMSInitialized {
			allKMSInitialized = false
			break
		}
	}

	if allKMSInitialized {
		// All units already fully initialized — nothing to do, skip to Read.
		return resourceIBMKmsCryptoUnitsRead(ctx, d, meta)
	}
	if v, ok := d.GetOk("should_zeroize"); ok {
		if sz, ok := v.(bool); ok && sz {
			tflog.Warn(ctx, "should_zeroize is true — zeroizing all crypto units before re-initialization",
				map[string]interface{}{"instance_id": kpOpts.InstanceID})
			for _, cu := range cryptoUnitsResponse.CryptoUnits {
				if zerr := kmsCryptoUnitClient.ZeroizeCryptoUnitWithContext(ctx, cu.ID); zerr != nil {
					tflog.Warn(ctx, "zeroize failed — keys may linger; delete keys and wait for the key purge",
						map[string]interface{}{"instance_id": kpOpts.InstanceID, "cryptounit": cu.ID})
					return diag.Errorf("failed to zeroize crypto unit %s: %v", cu.ID, zerr)
				}
				tflog.Info(ctx, "successfully zeroized crypto unit",
					map[string]interface{}{"instance_id": kpOpts.InstanceID, "cryptounit": cu.ID})
			}
		} else {
			tflog.Warn(ctx, "should_zeroize is false — all crypto units will keep their state",
				map[string]interface{}{"instance_id": kpOpts.InstanceID})
		}
	}

	return resourceIBMKmsCryptoUnitsCreate(ctx, d, meta)
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
	if v, ok := d.GetOk("should_zeroize"); ok {
		if sz, ok := v.(bool); ok && sz {
			for _, cryptoUnit := range cryptoUnitsResponse.CryptoUnits {
				err := kmsCryptoUnitClient.ZeroizeCryptoUnitWithContext(ctx, cryptoUnit.ID)
				if err != nil {
					tflog.Warn(ctx, "zeroize failed — keys may linger; delete keys and wait for purge",
						map[string]interface{}{"instance_id": kpOpts.InstanceID})
					return diag.Errorf("failed to zeroize crypto unit %s: %v", cryptoUnit.ID, err)
				}
			}
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

	// Ensure Region and InstanceID are always populated on the options struct.
	// NewKeyProtectCryptoUnitAPIOptions parses them from a standard URL format,
	// but if the URL has a non-standard format (e.g. from extensions["endpoints.public"]),
	// the parse may yield empty strings. Use the explicitly-provided values as a fallback.
	if kpOpts.Region == "" && region != "" {
		kpOpts.Region = region
	}
	if kpOpts.InstanceID == "" && instanceID != "" {
		kpOpts.InstanceID = instanceID
	}

	return kpOpts, nil
}
