// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

const (
	// Arguments
	PIIpSecPolicyAuth        = "pi_policy_authentication"
	PIIpSecPolicyDH          = "pi_policy_dh_group"
	PIIpSecPolicyEncryption  = "pi_policy_encryption"
	PIIpSecPolicyKeyLifetime = "pi_policy_key_lifetime"
	PIIpSecPolicyName        = "pi_policy_name"
	PIIpSecPfs               = "pi_policy_pfs"

	// Attributes
	IPSecPolicyID = "policy_id"
)

func ResourceIBMPIIPSecPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIIPSecPolicyCreate,
		ReadContext:   resourceIBMPIIPSecPolicyRead,
		UpdateContext: resourceIBMPIIPSecPolicyUpdate,
		DeleteContext: resourceIBMPIIPSecPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
			},
			PIIpSecPolicyName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the IPSec Policy",
			},
			PIIpSecPolicyDH: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedIntValues([]int{1, 2, 5, 14, 19, 20, 24}),
				Description:  "DH group of the IPSec Policy",
			},
			PIIpSecPolicyEncryption: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"aes-256-cbc", "aes-192-cbc", "aes-128-cbc", "aes-256-gcm", "aes-128-gcm", "3des-cbc"}),
				Description:  "Encryption of the IPSec Policy",
			},
			PIIpSecPolicyKeyLifetime: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedRangeInt(180, 86400),
				Description:  "Policy key lifetime",
			},
			PIIpSecPfs: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Perfect Forward Secrecy",
			},

			// Optional Attributes
			PIIpSecPolicyAuth: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"hmac-sha-256-128", "hmac-sha1-96", "none"}),
				Description:  "Authentication for the IPSec Policy",
			},

			//Computed Attributes
			IPSecPolicyID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPSec policy ID",
			},
		},
	}
}

func resourceIBMPIIPSecPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	name := d.Get(PIIpSecPolicyName).(string)
	dhGroup := int64(d.Get(PIIpSecPolicyDH).(int))
	encryption := d.Get(PIIpSecPolicyEncryption).(string)
	pfs := d.Get(PIIpSecPfs).(bool)
	keyLifetime := int64(d.Get(PIIpSecPolicyKeyLifetime).(int))
	klt := models.KeyLifetime(keyLifetime)

	body := &models.IPSecPolicyCreate{
		DhGroup:     &dhGroup,
		Encryption:  &encryption,
		KeyLifetime: &klt,
		Name:        &name,
		Pfs:         &pfs,
	}

	if v, ok := d.GetOk(PIIpSecPolicyAuth); ok {
		body.Authentication = models.IPSECPolicyAuthentication(v.(string))
	}

	client := st.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	ipsecPolicy, err := client.CreateIPSecPolicy(body)
	if err != nil {
		log.Printf("[DEBUG] create ipsec policy failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *ipsecPolicy.ID))

	return resourceIBMPIIPSecPolicyRead(ctx, d, meta)
}

func resourceIBMPIIPSecPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	body := &models.IPSecPolicyUpdate{}

	if d.HasChange(PIIpSecPolicyName) {
		name := d.Get(PIIpSecPolicyName).(string)
		body.Name = name
	}
	if d.HasChange(PIIpSecPolicyDH) {
		dhGroup := int64(d.Get(PIIpSecPolicyDH).(int))
		body.DhGroup = dhGroup
	}
	if d.HasChange(PIIpSecPolicyEncryption) {
		encryption := d.Get(PIIpSecPolicyEncryption).(string)
		body.Encryption = encryption
	}
	if d.HasChange(PIIpSecPolicyKeyLifetime) {
		keyLifetime := int64(d.Get(PIIpSecPolicyKeyLifetime).(int))
		body.KeyLifetime = models.KeyLifetime(keyLifetime)
	}
	if d.HasChange(PIIpSecPfs) {
		pfs := d.Get(PIIpSecPfs).(bool)
		body.Pfs = &pfs
	}
	if d.HasChange(PIIpSecPolicyAuth) {
		authentication := d.Get(PIIpSecPolicyAuth).(string)
		body.Authentication = models.IPSECPolicyAuthentication(authentication)
	}

	_, err = client.UpdateIPSecPolicy(policyID, body)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIIPSecPolicyRead(ctx, d, meta)
}

func resourceIBMPIIPSecPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	ipsecPolicy, err := client.GetIPSecPolicy(policyID)
	if err != nil {
		// FIXME: Uncomment when 404 error is available
		// switch err.(type) {
		// case *p_cloud_v_p_n_policies.PcloudIPSecpoliciesGetNotFound:
		// 	log.Printf("[DEBUG] VPN policy does not exist %v", err)
		// 	d.SetId("")
		// 	return nil
		// }
		log.Printf("[DEBUG] get VPN policy failed %v", err)
		return diag.FromErr(err)
	}

	d.Set(IPSecPolicyID, ipsecPolicy.ID)
	d.Set(PIIpSecPolicyName, ipsecPolicy.Name)
	d.Set(PIIpSecPolicyDH, ipsecPolicy.DhGroup)
	d.Set(PIIpSecPolicyEncryption, ipsecPolicy.Encryption)
	d.Set(PIIpSecPolicyKeyLifetime, ipsecPolicy.KeyLifetime)
	d.Set(PIIpSecPfs, ipsecPolicy.Pfs)
	d.Set(PIIpSecPolicyAuth, ipsecPolicy.Authentication)

	return nil
}

func resourceIBMPIIPSecPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)

	err = client.DeleteIPSecPolicy(policyID)
	if err != nil {
		// FIXME: Uncomment when 404 error is available
		// switch err.(type) {
		// case *p_cloud_v_p_n_policies.PcloudIPSecpoliciesDeleteNotFound:
		// 	log.Printf("[DEBUG] VPN policy does not exist %v", err)
		// 	d.SetId("")
		// 	return nil
		// }
		log.Printf("[DEBUG] delete VPN policy failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
