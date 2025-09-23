// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isIpSecName              = "name"
	isIpSecAuthenticationAlg = "authentication_algorithm"
	isIpSecEncryptionAlg     = "encryption_algorithm"
	isIpSecPFS               = "pfs"
	isIpSecKeyLifeTime       = "key_lifetime"
	isIPSecResourceGroup     = "resource_group"
	isIPSecEncapsulationMode = "encapsulation_mode"
	isIPSecTransformProtocol = "transform_protocol"
	isIPSecVPNConnections    = "vpn_connections"
	isIPSecVPNConnectionName = "name"
	isIPSecVPNConnectionId   = "id"
	isIPSecVPNConnectionHref = "href"
)

func ResourceIBMISIPSecPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISIPSecPolicyCreate,
		ReadContext:   resourceIBMISIPSecPolicyRead,
		UpdateContext: resourceIBMISIPSecPolicyUpdate,
		DeleteContext: resourceIBMISIPSecPolicyDelete,
		Exists:        resourceIBMISIPSecPolicyExists,
		Importer:      &schema.ResourceImporter{},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceIPSecPolicyValidate(diff)
				}),
		),

		Schema: map[string]*schema.Schema{
			isIpSecName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecName),
				Description:  "IPSEC name",
			},

			isIpSecAuthenticationAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecAuthenticationAlg),
				Description:  "Authentication alorothm",
			},

			isIpSecEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecEncryptionAlg),
				Description:  "Encryption algorithm",
			},

			isIpSecPFS: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecPFS),
				Description:  "PFS info",
			},

			isIPSecResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group info",
			},

			isIpSecKeyLifeTime: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      3600,
				ValidateFunc: validate.ValidateKeyLifeTime,
				Description:  "IPSEC key lifetime",
			},

			isIPSecEncapsulationMode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPSEC encapsulation mode",
			},

			isIPSecTransformProtocol: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPSEC transform protocol",
			},

			isIPSecVPNConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isIPSecVPNConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIPSecVPNConnectionId: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIPSecVPNConnectionHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
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

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMISIPSECValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	authentication_algorithm := "md5, sha1, sha256, sha512, sha384, disabled"
	encryption_algorithm := "triple_des, aes128, aes256, aes128gcm16, aes192gcm16, aes256gcm16, aes192"
	pfs := "disabled, group_2, group_5, group_14, group_19, group_15, group_16, group_17, group_18, group_20, group_21, group_22, group_23, group_24, group_31"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIpSecName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIpSecAuthenticationAlg,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              authentication_algorithm})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIpSecEncryptionAlg,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              encryption_algorithm})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIpSecPFS,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              pfs})

	ibmISIPSECResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_ipsec_policy", Schema: validateSchema}
	return &ibmISIPSECResourceValidator
}

func resourceIBMISIPSecPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] Ip Sec create")
	name := d.Get(isIpSecName).(string)
	authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
	encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
	pfs := d.Get(isIpSecPFS).(string)

	diag := ipsecpCreate(context, d, meta, authenticationAlg, encryptionAlg, name, pfs)
	if diag != nil {
		return diag
	}
	return resourceIBMISIPSecPolicyRead(context, d, meta)
}

func ipsecpCreate(context context.Context, d *schema.ResourceData, meta interface{}, authenticationAlg, encryptionAlg, name, pfs string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateIpsecPolicyOptions{
		AuthenticationAlgorithm: &authenticationAlg,
		EncryptionAlgorithm:     &encryptionAlg,
		Pfs:                     &pfs,
		Name:                    &name,
	}

	if keylt, ok := d.GetOk(isIpSecKeyLifeTime); ok {
		keyLifetime := int64(keylt.(int))
		options.KeyLifetime = &keyLifetime
	} else {
		keyLifetime := int64(3600)
		options.KeyLifetime = &keyLifetime
	}

	if rgrp, ok := d.GetOk(isIPSecResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	ipSec, _, err := sess.CreateIpsecPolicyWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateIpsecPolicyWithContext failed: %s", err.Error()), "ibm_is_ipsec_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*ipSec.ID)
	log.Printf("[INFO] ipSec policy : %s", *ipSec.ID)
	return nil
}

func resourceIBMISIPSecPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	return ipsecpGet(context, d, meta, id)
}

func ipsecpGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	iPsecPolicy, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIpsecPolicyWithContext failed: %s", err.Error()), "ibm_is_ipsec_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(iPsecPolicy.Name) {
		if err = d.Set("name", iPsecPolicy.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-name").GetDiag()
		}
		if err = d.Set(flex.ResourceName, iPsecPolicy.Name); err != nil {
			err = fmt.Errorf("Error setting resource_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-resource_name").GetDiag()
		}
	}
	if err = d.Set("authentication_algorithm", iPsecPolicy.AuthenticationAlgorithm); err != nil {
		err = fmt.Errorf("Error setting authentication_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-authentication_algorithm").GetDiag()
	}
	if err = d.Set("encryption_algorithm", iPsecPolicy.EncryptionAlgorithm); err != nil {
		err = fmt.Errorf("Error setting encryption_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-encryption_algorithm").GetDiag()
	}
	if iPsecPolicy.ResourceGroup != nil {
		if err = d.Set(isIPSecResourceGroup, *iPsecPolicy.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *iPsecPolicy.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-resource_group_name").GetDiag()
		}

	} else {
		d.Set(isIPSecResourceGroup, nil)
	}
	if err = d.Set("pfs", iPsecPolicy.Pfs); err != nil {
		err = fmt.Errorf("Error setting pfs: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-pfs").GetDiag()
	}
	if !core.IsNil(iPsecPolicy.KeyLifetime) {
		if err = d.Set("key_lifetime", flex.IntValue(iPsecPolicy.KeyLifetime)); err != nil {
			err = fmt.Errorf("Error setting key_lifetime: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-key_lifetime").GetDiag()
		}
	}
	if err = d.Set("encapsulation_mode", iPsecPolicy.EncapsulationMode); err != nil {
		err = fmt.Errorf("Error setting encapsulation_mode: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-encapsulation_mode").GetDiag()
	}
	if err = d.Set("transform_protocol", iPsecPolicy.TransformProtocol); err != nil {
		err = fmt.Errorf("Error setting transform_protocol: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-transform_protocol").GetDiag()
	}

	connList := make([]map[string]interface{}, 0)
	if iPsecPolicy.Connections != nil && len(iPsecPolicy.Connections) > 0 {
		for _, connection := range iPsecPolicy.Connections {
			conn := map[string]interface{}{}
			conn[isIPSecVPNConnectionName] = *connection.Name
			conn[isIPSecVPNConnectionId] = *connection.ID
			conn[isIPSecVPNConnectionHref] = *connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIPSecVPNConnections, connList)
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		err = fmt.Errorf("Error featching Base Controller URL: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-resource_controller_url").GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/ipsecpolicies")
	// d.Set(flex.ResourceCRN, *ipSec.Crn)
	return nil
}

func resourceIBMISIPSecPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	err := ipsecpUpdate(context, d, meta, id)
	if err != nil {
		return err
	}

	return resourceIBMISIPSecPolicyRead(context, d, meta)
}

func ipsecpUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.UpdateIpsecPolicyOptions{
		ID: &id,
	}
	if d.HasChange(isIpSecName) || d.HasChange(isIpSecAuthenticationAlg) || d.HasChange(isIpSecEncryptionAlg) || d.HasChange(isIpSecPFS) || d.HasChange(isIpSecKeyLifeTime) {
		name := d.Get(isIpSecName).(string)
		authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
		pfs := d.Get(isIpSecPFS).(string)
		keyLifetime := int64(d.Get(isIpSecKeyLifeTime).(int))

		ipsecPolicyPatchModel := &vpcv1.IPsecPolicyPatch{
			Name:                    &name,
			AuthenticationAlgorithm: &authenticationAlg,
			EncryptionAlgorithm:     &encryptionAlg,
			Pfs:                     &pfs,
			KeyLifetime:             &keyLifetime,
		}
		ipsecPolicyPatch, err := ipsecPolicyPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for IPsecPolicyPatch: %s", err.Error()), "ibm_is_ipsec_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.IPsecPolicyPatch = ipsecPolicyPatch

		_, _, err = sess.UpdateIpsecPolicyWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateIpsecPolicyWithContext failed: %s", err.Error()), "ibm_is_ipsec_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISIPSecPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	return ipsecpDelete(context, d, meta, id)
}

func ipsecpDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicyWithContext(context, getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIpsecPolicyWithContext failed: %s", err.Error()), "ibm_is_ipsec_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	deleteIpsecPolicyOptions := &vpcv1.DeleteIpsecPolicyOptions{
		ID: &id,
	}
	response, err = sess.DeleteIpsecPolicyWithContext(context, deleteIpsecPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteIpsecPolicyWithContext failed: %s", err.Error()), "ibm_is_ipsec_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISIPSecPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()
	return ipsecpExists(d, meta, id)
}

func ipsecpExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "exits", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	options := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicy(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIpsecPolicy failed: %s", err.Error()), "ibm_is_ipsec_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
