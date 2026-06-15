// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
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
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"authentication_algorithm", "authentication_algorithms"},
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecAuthenticationAlg),
				Description:  "Authentication alorothm",
				Deprecated:   "`authentication_algorithm` is deprecated in favor of `authentication_algorithms`. The existing `authentication_algorithm` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `authentication_algorithms`. Use `authentication_algorithms` to configure multiple authentication algorithms. This enhancement adds support for multi-algorithm authentication while preserving compatibility with earlier single-algorithm configurations.",
			},
			"authentication_algorithms": &schema.Schema{
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     3,
				ExactlyOneOf: []string{"authentication_algorithm", "authentication_algorithms"},
				Computed:     true,
				Description:  "The authentication algorithms to use for IPsec Negotiation.The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it.",
				Elem: &schema.Schema{Type: schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecAuthenticationAlg),
				},
			},

			isIpSecEncryptionAlg: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"encryption_algorithm", "encryption_algorithms"},
				Deprecated:   "`encryption_algorithm` is deprecated in favor of `encryption_algorithms`. The existing `encryption_algorithm` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `encryption_algorithms`. Use `encryption_algorithms` to configure multiple encryption algorithms. This enhancement adds support for multi-algorithm encryption while preserving compatibility with earlier single-algorithm configurations.",
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecEncryptionAlg),
				Description:  "Encryption algorithm",
			},

			"encryption_algorithms": &schema.Schema{
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     3,
				Computed:     true,
				ExactlyOneOf: []string{"encryption_algorithm", "encryption_algorithms"},
				Description:  "The encryption algorithms to use for IPsec Negotiation.The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecEncryptionAlg),
				},
			},

			isIpSecPFS: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"pfs", "pfs_groups"},
				Deprecated:   "`pfs` is deprecated in favor of `pfs_groups`. The existing `pfs` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `pfs_groups`. Use `pfs_groups` to configure multiple Perfect Forward Secrecy (PFS) groups. This enhancement adds support for multi-group PFS configurations while preserving compatibility with earlier single-group configurations.",
				ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecPFS),
				Description:  "PFS info",
			},

			"pfs_groups": &schema.Schema{
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     12,
				Computed:     true,
				ExactlyOneOf: []string{"pfs", "pfs_groups"},
				Description:  "The Perfect Forward Secrecy groups to use for IPsec negotiation.The order of the Perfect Forward Secrecy groups in this array indicates their priority for negotiation, with each Perfect Forward Secrecy group having priority over the one after it.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_is_ipsec_policy", isIpSecPFS),
				},
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

	diag := ipsecpCreate(context, d, meta, name)
	if diag != nil {
		return diag
	}
	return resourceIBMISIPSecPolicyRead(context, d, meta)
}

func ipsecpCreate(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.CreateIpsecPolicyOptions{
		Name: &name,
	}

	if _, ok := d.GetOk("authentication_algorithm"); ok {
		options.SetAuthenticationAlgorithm(d.Get("authentication_algorithm").(string))
	}
	if _, ok := d.GetOk("authentication_algorithms"); ok {
		var authenticationAlgorithms []string
		for _, v := range d.Get("authentication_algorithms").([]interface{}) {
			authenticationAlgorithmsItem := v.(string)
			authenticationAlgorithms = append(authenticationAlgorithms, authenticationAlgorithmsItem)
		}
		options.SetAuthenticationAlgorithms(authenticationAlgorithms)
	}
	if _, ok := d.GetOk("encryption_algorithm"); ok {
		options.SetEncryptionAlgorithm(d.Get("encryption_algorithm").(string))
	}
	if _, ok := d.GetOk("encryption_algorithms"); ok {
		var encryptionAlgorithms []string
		for _, v := range d.Get("encryption_algorithms").([]interface{}) {
			encryptionAlgorithmsItem := v.(string)
			encryptionAlgorithms = append(encryptionAlgorithms, encryptionAlgorithmsItem)
		}
		options.SetEncryptionAlgorithms(encryptionAlgorithms)
	}
	if _, ok := d.GetOk("pfs"); ok {
		options.SetPfs(d.Get("pfs").(string))
	}
	if _, ok := d.GetOk("pfs_groups"); ok {
		var pfsGroups []string
		for _, v := range d.Get("pfs_groups").([]interface{}) {
			pfsGroupsItem := v.(string)
			pfsGroups = append(pfsGroups, pfsGroupsItem)
		}
		options.SetPfsGroups(pfsGroups)
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
	if !core.IsNil(iPsecPolicy.AuthenticationAlgorithms) {
		if err = d.Set("authentication_algorithms", iPsecPolicy.AuthenticationAlgorithms); err != nil {
			err = fmt.Errorf("Error setting authentication_algorithms: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-authentication_algorithms").GetDiag()
		}
	}
	if err = d.Set("encryption_algorithm", iPsecPolicy.EncryptionAlgorithm); err != nil {
		err = fmt.Errorf("Error setting encryption_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-encryption_algorithm").GetDiag()
	}
	if !core.IsNil(iPsecPolicy.EncryptionAlgorithms) {
		if err = d.Set("encryption_algorithms", iPsecPolicy.EncryptionAlgorithms); err != nil {
			err = fmt.Errorf("Error setting encryption_algorithms: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-encryption_algorithms").GetDiag()
		}
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
	if !core.IsNil(iPsecPolicy.PfsGroups) {
		if err = d.Set("pfs_groups", iPsecPolicy.PfsGroups); err != nil {
			err = fmt.Errorf("Error setting pfs_groups: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ipsec_policy", "read", "set-pfs_groups").GetDiag()
		}
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

	isIpsecNameChangeFlag := d.HasChange(isIpSecName)
	isIpSecAuthenticationAlgChangeFlag := d.HasChange(isIpSecAuthenticationAlg)
	isIpSecAuthenticationAlgsChangeFlag := d.HasChange("authentication_algorithms")
	isIpSecEncryptionAlgChangeFlag := d.HasChange(isIpSecEncryptionAlg)
	isIpSecEncryptionAlgsChangeFlag := d.HasChange("encryption_algorithms")
	isIpSecPFSChangeFlag := d.HasChange(isIpSecPFS)
	isIpSecPFSGroupsChangeFlag := d.HasChange("pfs_groups")
	isIpSecKeyLifeTimeChangeFlag := d.HasChange(isIpSecKeyLifeTime)

	if isIpsecNameChangeFlag || isIpSecAuthenticationAlgChangeFlag || isIpSecAuthenticationAlgsChangeFlag || isIpSecEncryptionAlgChangeFlag || isIpSecEncryptionAlgsChangeFlag || isIpSecPFSChangeFlag || isIpSecPFSGroupsChangeFlag || isIpSecKeyLifeTimeChangeFlag {
		ipsecPolicyPatchModel := &vpcv1.IPsecPolicyPatch{}
		if isIpSecAuthenticationAlgChangeFlag {
			authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
			ipsecPolicyPatchModel.AuthenticationAlgorithm = &authenticationAlg

		}
		if isIpSecEncryptionAlgChangeFlag {
			encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
			ipsecPolicyPatchModel.EncryptionAlgorithm = &encryptionAlg
		}
		if isIpSecPFSChangeFlag {
			pfs := d.Get(isIpSecPFS).(string)
			ipsecPolicyPatchModel.Pfs = &pfs
		}
		if isIpSecKeyLifeTimeChangeFlag {
			keyLifetime := int64(d.Get(isIpSecKeyLifeTime).(int))
			ipsecPolicyPatchModel.KeyLifetime = &keyLifetime
		}
		if isIpSecAuthenticationAlgsChangeFlag {
			ipsecPolicyPatchModel.AuthenticationAlgorithms = interfaceSliceToStringSlice(d.Get("authentication_algorithms").([]interface{}))
		}
		if isIpSecEncryptionAlgsChangeFlag {
			ipsecPolicyPatchModel.EncryptionAlgorithms = interfaceSliceToStringSlice(d.Get("encryption_algorithms").([]interface{}))
		}
		if isIpSecPFSGroupsChangeFlag {
			ipsecPolicyPatchModel.PfsGroups = interfaceSliceToStringSlice(d.Get("pfs_groups").([]interface{}))
		}
		if isIpsecNameChangeFlag {
			name := d.Get(isIpSecName).(string)
			ipsecPolicyPatchModel.Name = &name
		}

		// if isIpSecAuthenticationAlgsChangeFlag || isIpSecEncryptionAlgsChangeFlag || isIpSecPFSGroupsChangeFlag {
		// 	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		// 		ID: core.StringPtr(d.Id()),
		// 	}
		// 	_, etagResponse, etagErr := sess.GetIpsecPolicyWithContext(context, getIpsecPolicyOptions)
		// 	if etagErr != nil {
		// 		if etagResponse != nil && etagResponse.StatusCode == 404 {
		// 			d.SetId("")
		// 			return nil
		// 		}
		// 		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIpsecPolicyWithContext failed: %s", etagErr.Error()), "ibm_is_ipsec_policy", "update")
		// 		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// 		return tfErr.GetDiag()
		// 	}
		// 	eTag := etagResponse.Headers.Get("ETag")
		// 	options.IfMatch = &eTag
		// }

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
