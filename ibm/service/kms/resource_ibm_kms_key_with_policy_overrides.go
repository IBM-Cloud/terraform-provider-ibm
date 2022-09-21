// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMKmsKeyWithPolicyOverrides() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMKmsKeyWithPolicyOverridesCreate,
		ReadContext:   resourceIBMKmsKeyWithPolicyOverridesRead,
		UpdateContext: resourceIBMKmsKeyWithPolicyOverridesUpdate,
		DeleteContext: resourceIBMKmsKeyWithPolicyOverridesDelete,
		Exists:        resourceIBMKmsKeyExists,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Key protect or HPCS instance GUID or CRN",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"key_ring_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "default",
				Description: "Key Ring for the Key",
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key ID",
			},
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Key name",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "type of service hs-crypto or kms",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private"}),
				Description:  "public or private",
				ForceNew:     true,
			},
			"standard_key": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				ForceNew:    true,
				Description: "Standard key type",
			},
			"payload": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"encrypted_nonce": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Only for imported root key",
			},
			"iv_value": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Only for imported root key",
			},
			"force_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "set to true to force delete the key",
				ForceNew:    false,
				Default:     false,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Crn of the key",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The date the key material expires. The date format follows RFC 3339. You can set an expiration date on any key on its creation. A key moves into the Deactivated state within one hour past its expiration date, if one is assigned. If you create a key without specifying an expiration date, the key does not expire",
				ForceNew:    true,
			},
			"instance_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key protect or HPCS instance CRN",
			},
			"rotation": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Specifies the key rotation time interval in months, with a minimum of 1, and a maximum of 12",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "If set to true, Key Protect enables a rotation policy on a single key.",
						},
						"interval_month": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validate.ValidateAllowedRangeInt(1, 12),
							Description:  "Specifies the key rotation time interval in months",
						},
					},
				},
			},
			"dual_auth_delete": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Data associated with the dual authorization delete policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "If set to true, Key Protect enables a dual authorization policy on a single key.",
						},
					},
				},
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
			},
		},
	}
}

func resourceIBMKmsKeyWithPolicyOverridesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceID := getInstanceIDFromCRN(d.Get("instance_id").(string))
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	kpAPI.Config.KeyRing = d.Get("key_ring_id").(string)

	name := d.Get("key_name").(string)
	standardKey := d.Get("standard_key").(bool)
	payload := d.Get("payload").(string)
	encryptedNonce := d.Get("encrypted_nonce").(string)
	iv := d.Get("iv_value").(string)
	policy := getPolicyFromSchema(d)

	var expiration *time.Time
	if es, ok := d.GetOk("expiration_date"); ok {
		expiration_string := es.(string)
		// parse string to required time format
		expiration_time, err := time.Parse(time.RFC3339, expiration_string)
		if err != nil {
			return diag.Errorf("[ERROR] Invalid time format (the date format follows RFC 3339): %s", err)
		}
		expiration = &expiration_time
	} else {
		expiration = nil
	}
	key, err := kpAPI.CreateImportedKeyWithPolicyOverrides(context, name, expiration, payload, encryptedNonce, iv, standardKey, nil, policy)
	if err != nil {
		return diag.Errorf("[ERROR] Error while creating key: %s", err)
	}

	d.SetId(key.CRN)
	return resourceIBMKmsKeyWithPolicyOverridesUpdate(context, d, meta)
}

func resourceIBMKmsKeyWithPolicyOverridesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceCRN, instanceID, keyid := getInstanceAndKeyDataFromCRN(d.Id())

	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	key, err := kpAPI.GetKey(context, keyid)
	if err != nil {
		kpError := err.(*kp.Error)
		if kpError.StatusCode == 404 || kpError.StatusCode == 409 {
			d.SetId("")
			return nil
		}
		return diag.Errorf("[ERROR] Get Key failed with error while reading Key: %s", err)
	} else if key.State == 5 { //Refers to Deleted state of the Key
		d.SetId("")
		return nil
	}

	err = setKeyDetails(d, meta, instanceID, instanceCRN, key, kpAPI)
	if err != nil {
		return diag.FromErr(err)
	}

	policies, err := kpAPI.GetPolicies(context, keyid)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read policies: %s", err)
	}
	if len(policies) == 0 {
		log.Printf("No Policy Configurations read\n")
	} else {
		d.Set("rotation", flex.FlattenKeyPoliciesKey(policies)[0]["rotation"])
		d.Set("dual_auth_delete", flex.FlattenKeyPoliciesKey(policies)[0]["dual_auth_delete"])
	}
	return nil
}

func resourceIBMKmsKeyWithPolicyOverridesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	if d.HasChange("force_delete") {
		d.Set("force_delete", d.Get("force_delete").(bool))
	}
	if d.HasChange("rotation") || d.HasChange("dual_auth_delete") {
		_, rotationOk := d.GetOk("rotation")
		_, dualAuthOk := d.GetOk("dual_auth_delete")
		if !rotationOk || !dualAuthOk {
			log.Println("Warning: Removing Policy details does not delete the policies of the Key. Key Policies get deleted when the associated key resource is destroyed.")
			return resourceIBMKmsKeyWithPolicyOverridesRead(context, d, meta)
		}
		_, instanceID, key_id := getInstanceAndKeyDataFromCRN(d.Id())
		kpAPI, _, err := populateKPClient(d, meta, instanceID)
		if err != nil {
			return diag.FromErr(err)
		}

		err = resourceHandlePolicies(context, d, kpAPI, meta, key_id)
		if err != nil {
			return diag.Errorf("Could not create policies: %s", err)
		}
	}
	return resourceIBMKmsKeyWithPolicyOverridesRead(context, d, meta)
}

func resourceIBMKmsKeyWithPolicyOverridesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(d.Id())
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	force := d.Get("force_delete").(bool)
	f := kp.ForceOpt{
		Force: force,
	}

	_, err1 := kpAPI.DeleteKey(context, keyid, kp.ReturnRepresentation, f)
	if err1 != nil {
		return diag.Errorf("[ERROR] Error while deleting: %s", err1)
	}
	d.SetId("")
	return nil
}
