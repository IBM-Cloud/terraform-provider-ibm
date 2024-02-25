// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PIPolicyId = "policy_id"
)

func ResourceIBMPIIKEPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIIKEPolicyCreate,
		ReadContext:   resourceIBMPIIKEPolicyRead,
		UpdateContext: resourceIBMPIIKEPolicyUpdate,
		DeleteContext: resourceIBMPIIKEPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			Arg_CloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
			},
			Arg_VPNPolicyName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the IKE Policy",
			},
			Arg_VPNPolicyDhGroup: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedIntValues([]int{1, 2, 5, 14, 19, 20, 24}),
				Description:  "DH group of the IKE Policy",
			},
			Arg_VPNPolicyEncryption: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"aes-256-cbc", "aes-192-cbc", "aes-128-cbc", "aes-256-gcm", "aes-128-gcm", "3des-cbc"}),
				Description:  "Encryption of the IKE Policy",
			},
			Arg_VPNPolicyKeyLifetime: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedRangeInt(180, 86400),
				Description:  "Policy key lifetime",
			},
			Arg_VPNPolicyVersion: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedRangeInt(1, 2),
				Description:  "Version of the IKE Policy",
			},
			Arg_VPNPolicyPresharedKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Preshared key used in this IKE Policy (length of preshared key must be even)",
			},

			// Optional Attributes
			Arg_VPNPolicyAuthentication: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"sha-256", "sha-384", "sha1", "none"}),
				Description:  "Authentication for the IKE Policy",
			},

			//Computed Attributes
			Attr_PolicyId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IKE Policy ID",
			},
		},
	}
}

func resourceIBMPIIKEPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	name := d.Get(Arg_VPNPolicyName).(string)
	dhGroup := int64(d.Get(Arg_VPNPolicyDhGroup).(int))
	encryption := d.Get(Arg_VPNPolicyEncryption).(string)
	presharedKey := d.Get(Arg_VPNPolicyPresharedKey).(string)
	version := int64(d.Get(Arg_VPNPolicyVersion).(int))
	keyLifetime := int64(d.Get(Arg_VPNPolicyKeyLifetime).(int))
	klt := models.KeyLifetime(keyLifetime)

	body := &models.IKEPolicyCreate{
		DhGroup:      &dhGroup,
		Encryption:   &encryption,
		KeyLifetime:  &klt,
		Name:         &name,
		PresharedKey: &presharedKey,
		Version:      &version,
	}

	if v, ok := d.GetOk(Arg_VPNPolicyAuthentication); ok {
		body.Authentication = models.IKEPolicyAuthentication(v.(string))
	}

	client := instance.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	ikePolicy, err := client.CreateIKEPolicy(body)
	if err != nil {
		log.Printf("[DEBUG] create ike policy failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *ikePolicy.ID))

	return resourceIBMPIIKEPolicyRead(ctx, d, meta)
}

func resourceIBMPIIKEPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := instance.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	body := &models.IKEPolicyUpdate{}

	if d.HasChange(Arg_VPNPolicyName) {
		name := d.Get(Arg_VPNPolicyName).(string)
		body.Name = name
	}
	if d.HasChange(Arg_VPNPolicyDhGroup) {
		dhGroup := int64(d.Get(Arg_VPNPolicyDhGroup).(int))
		body.DhGroup = dhGroup
	}
	if d.HasChange(Arg_VPNPolicyEncryption) {
		encryption := d.Get(Arg_VPNPolicyEncryption).(string)
		body.Encryption = encryption
	}
	if d.HasChange(Arg_VPNPolicyKeyLifetime) {
		keyLifetime := int64(d.Get(Arg_VPNPolicyKeyLifetime).(int))
		body.KeyLifetime = models.KeyLifetime(keyLifetime)
	}
	if d.HasChange(Arg_VPNPolicyPresharedKey) {
		presharedKey := d.Get(Arg_VPNPolicyPresharedKey).(string)
		body.PresharedKey = presharedKey
	}
	if d.HasChange(Arg_VPNPolicyVersion) {
		version := int64(d.Get(Arg_VPNPolicyVersion).(int))
		body.Version = version
	}
	if d.HasChange(Arg_VPNPolicyAuthentication) {
		authentication := d.Get(Arg_VPNPolicyAuthentication).(string)
		body.Authentication = models.IKEPolicyAuthentication(authentication)
	}

	_, err = client.UpdateIKEPolicy(policyID, body)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIIKEPolicyRead(ctx, d, meta)
}

func resourceIBMPIIKEPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := instance.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)
	ikePolicy, err := client.GetIKEPolicy(policyID)
	if err != nil {
		// FIXME: Uncomment when 404 error is available
		// switch err.(type) {
		// case *p_cloud_v_p_n_policies.PcloudIkepoliciesGetNotFound:
		// 	log.Printf("[DEBUG] VPN policy does not exist %v", err)
		// 	d.SetId("")
		// 	return nil
		// }
		log.Printf("[DEBUG] get VPN policy failed %v", err)
		return diag.FromErr(err)
	}

	d.Set(PIPolicyId, ikePolicy.ID)
	d.Set(Arg_VPNPolicyName, ikePolicy.Name)
	d.Set(Arg_VPNPolicyDhGroup, ikePolicy.DhGroup)
	d.Set(Arg_VPNPolicyEncryption, ikePolicy.Encryption)
	d.Set(Arg_VPNPolicyKeyLifetime, ikePolicy.KeyLifetime)
	d.Set(Arg_VPNPolicyVersion, ikePolicy.Version)
	d.Set(Arg_VPNPolicyAuthentication, ikePolicy.Authentication)

	return nil
}

func resourceIBMPIIKEPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, policyID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := instance.NewIBMPIVpnPolicyClient(ctx, sess, cloudInstanceID)

	err = client.DeleteIKEPolicy(policyID)
	if err != nil {
		// FIXME: Uncomment when 404 error is available
		// switch err.(type) {
		// case *p_cloud_v_p_n_policies.PcloudIkepoliciesDeleteNotFound:
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
