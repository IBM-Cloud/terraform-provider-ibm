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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isIKEName              = "name"
	isIKEAuthenticationAlg = "authentication_algorithm"
	isIKEEncryptionAlg     = "encryption_algorithm"
	isIKEDhGroup           = "dh_group"
	isIKEVERSION           = "ike_version"
	isIKEKeyLifeTime       = "key_lifetime"
	isIKEResourceGroup     = "resource_group"
	isIKENegotiationMode   = "negotiation_mode"
	isIKEVPNConnections    = "vpn_connections"
	isIKEVPNConnectionName = "name"
	isIKEVPNConnectionId   = "id"
	isIKEVPNConnectionHref = "href"
	isIKEHref              = "href"
)

func ResourceIBMISIKEPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISIKEPolicyCreate,
		ReadContext:   resourceIBMISIKEPolicyRead,
		UpdateContext: resourceIBMISIKEPolicyUpdate,
		DeleteContext: resourceIBMISIKEPolicyDelete,
		Exists:        resourceIBMISIKEPolicyExists,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isIKEName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEName),
				Description:  "IKE name",
			},

			isIKEAuthenticationAlg: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"authentication_algorithm", "authentication_algorithms"},
				Deprecated:   "`authentication_algorithm` is deprecated in favor of `authentication_algorithms`. The existing `authentication_algorithm` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `authentication_algorithms`. Use `authentication_algorithms` to configure multiple authentication algorithms. This enhancement adds support for multi-algorithm authentication while preserving compatibility with earlier single-algorithm configurations.",
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEAuthenticationAlg),
				Description:  "Authentication algorithm type",
			},
			"authentication_algorithms": &schema.Schema{
				Type:         schema.TypeList,
				MaxItems:     3,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"authentication_algorithm", "authentication_algorithms"},
				Description:  "The authentication algorithms to use for IKE Negotiation.The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEAuthenticationAlg),
				},
			},

			isIKEEncryptionAlg: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"encryption_algorithm", "encryption_algorithms"},
				Deprecated:   "`encryption_algorithm` is deprecated in favor of `encryption_algorithms`. The existing `encryption_algorithm` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `encryption_algorithms`. Use `encryption_algorithms` to configure multiple encryption algorithms. This enhancement adds support for multi-algorithm encryption while preserving compatibility with earlier single-algorithm configurations.",
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEEncryptionAlg),
				Description:  "Encryption alogorithm type",
			},
			"encryption_algorithms": &schema.Schema{
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     3,
				Computed:     true,
				ExactlyOneOf: []string{"encryption_algorithm", "encryption_algorithms"},
				Description:  "The encryption algorithms to use for IKE Negotiation.The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEEncryptionAlg),
				},
			},
			isIKEDhGroup: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"dh_group", "dh_groups"},
				Deprecated:   "`dh_group` is deprecated in favor of `dh_groups`. The existing `dh_group` field will continue to function without any behavior changes to maintain backward compatibility. No migration is required for existing configurations, for newer use `dh_groups`. Use `dh_groups` to configure multiple Diffie-Hellman groups. This enhancement adds support for multi-group DH configurations while preserving compatibility with earlier single-group configurations.",
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEDhGroup),
				Description:  "IKE DH group",
			},
			"dh_groups": &schema.Schema{
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     12,
				Computed:     true,
				ExactlyOneOf: []string{"dh_group", "dh_groups"},
				Description:  "The Diffie-Hellman groups to use for IKE negotiation.The order of the Diffie-Hellman groups in this array indicates their priority for negotiation, with each Diffie-Hellman group having priority over the one after it.",
				Elem: &schema.Schema{
					Type:         schema.TypeInt,
					ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEDhGroup),
				},
			},

			isIKEResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "IKE resource group ID",
			},

			isIKEKeyLifeTime: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      28800,
				ValidateFunc: validate.ValidateKeyLifeTime,
				Description:  "IKE Key lifetime",
			},

			isIKEVERSION: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEVERSION),
				Description:  "IKE version",
			},

			isIKENegotiationMode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IKE negotiation mode",
			},

			isIKEHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IKE href value",
			},

			isIKEVPNConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isIKEVPNConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIKEVPNConnectionId: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIKEVPNConnectionHref: {
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

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMISIKEValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	authentication_algorithm := "md5, sha1, sha256, sha512, sha384"
	encryption_algorithm := "triple_des, aes128, aes192, aes256"
	dh_group := "2, 5, 14, 19, 15, 16, 17, 18, 20, 21, 22, 23, 24, 31"
	ike_version := "1, 2"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIKEName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIKEAuthenticationAlg,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              authentication_algorithm})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIKEEncryptionAlg,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              encryption_algorithm})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIKEDhGroup,
			ValidateFunctionIdentifier: validate.ValidateAllowedIntValue,
			Type:                       validate.TypeInt,
			Required:                   true,
			AllowedValues:              dh_group})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isIKEVERSION,
			ValidateFunctionIdentifier: validate.ValidateAllowedIntValue,
			Type:                       validate.TypeInt,
			Optional:                   true,
			AllowedValues:              ike_version})

	ibmISIKEResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_ike_policy", Schema: validateSchema}
	return &ibmISIKEResourceValidator
}

func resourceIBMISIKEPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] IKE Policy create")
	name := d.Get(isIKEName).(string)

	diag := ikepCreate(context, d, meta, name)
	if diag != nil {
		return diag
	}
	return resourceIBMISIKEPolicyRead(context, d, meta)
}

func ikepCreate(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateIkePolicyOptions{
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
	if _, ok := d.GetOk("dh_group"); ok {
		options.SetDhGroup(int64(d.Get("dh_group").(int)))
	}
	if _, ok := d.GetOk("dh_groups"); ok {
		var dhGroups []int64
		for _, v := range d.Get("dh_groups").([]interface{}) {
			dhGroupsItem := int64(v.(int))
			dhGroups = append(dhGroups, dhGroupsItem)
		}
		options.SetDhGroups(dhGroups)
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

	if keylt, ok := d.GetOk(isIKEKeyLifeTime); ok {
		keyLifetime := int64(keylt.(int))
		options.KeyLifetime = &keyLifetime
	} else {
		keyLifetime := int64(28800)
		options.KeyLifetime = &keyLifetime
	}

	if ikev, ok := d.GetOk(isIKEVERSION); ok {
		ikeVersion := int64(ikev.(int))
		options.IkeVersion = &ikeVersion
	}

	if rgrp, ok := d.GetOk(isIKEResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	ike, _, err := sess.CreateIkePolicyWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateIkePolicyWithContext failed: %s", err.Error()), "ibm_is_ike_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*ike.ID)
	log.Printf("[INFO] ike policy : %s", *ike.ID)
	return nil
}

func resourceIBMISIKEPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	return ikepGet(context, d, meta, id)
}

func ikepGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getikepoptions := &vpcv1.GetIkePolicyOptions{
		ID: &id,
	}
	ikePolicy, response, err := sess.GetIkePolicyWithContext(context, getikepoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIkePolicyWithContext failed: %s", err.Error()), "ibm_is_ike_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("authentication_algorithm", ikePolicy.AuthenticationAlgorithm); err != nil {
		err = fmt.Errorf("Error setting authentication_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-authentication_algorithm").GetDiag()
	}
	if !core.IsNil(ikePolicy.AuthenticationAlgorithms) {
		if err = d.Set("authentication_algorithms", ikePolicy.AuthenticationAlgorithms); err != nil {
			err = fmt.Errorf("Error setting authentication_algorithms: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-authentication_algorithms").GetDiag()
		}
	}
	if err = d.Set("encryption_algorithm", ikePolicy.EncryptionAlgorithm); err != nil {
		err = fmt.Errorf("Error setting encryption_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-encryption_algorithm").GetDiag()
	}
	if !core.IsNil(ikePolicy.EncryptionAlgorithms) {
		if err = d.Set("encryption_algorithms", ikePolicy.EncryptionAlgorithms); err != nil {
			err = fmt.Errorf("Error setting encryption_algorithms: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-encryption_algorithms").GetDiag()
		}
	}
	if ikePolicy.ResourceGroup != nil {
		d.Set(isIKEResourceGroup, *ikePolicy.ResourceGroup.ID)
		d.Set(flex.ResourceGroupName, *ikePolicy.ResourceGroup.Name)
	} else {
		d.Set(isIKEResourceGroup, nil)
	}
	if !core.IsNil(ikePolicy.KeyLifetime) {
		if err = d.Set("key_lifetime", flex.IntValue(ikePolicy.KeyLifetime)); err != nil {
			err = fmt.Errorf("Error setting key_lifetime: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-key_lifetime").GetDiag()
		}
	}
	if err = d.Set("href", ikePolicy.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-href").GetDiag()
	}
	if err = d.Set("negotiation_mode", ikePolicy.NegotiationMode); err != nil {
		err = fmt.Errorf("Error setting negotiation_mode: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-negotiation_mode").GetDiag()
	}
	if err = d.Set("ike_version", flex.IntValue(ikePolicy.IkeVersion)); err != nil {
		err = fmt.Errorf("Error setting ike_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-ike_version").GetDiag()
	}
	if err = d.Set("dh_group", flex.IntValue(ikePolicy.DhGroup)); err != nil {
		err = fmt.Errorf("Error setting dh_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-dh_group").GetDiag()
	}
	if !core.IsNil(ikePolicy.DhGroups) {
		dhGroups := []interface{}{}
		for _, dhGroupsItem := range ikePolicy.DhGroups {
			dhGroups = append(dhGroups, int64(dhGroupsItem))
		}
		if err = d.Set("dh_groups", dhGroups); err != nil {
			err = fmt.Errorf("Error setting dh_groups: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-dh_groups").GetDiag()
		}
	}
	connList := make([]map[string]interface{}, 0)
	if ikePolicy.Connections != nil && len(ikePolicy.Connections) > 0 {
		for _, connection := range ikePolicy.Connections {
			conn := map[string]interface{}{}
			conn[isIKEVPNConnectionName] = *connection.Name
			conn[isIKEVPNConnectionId] = *connection.ID
			conn[isIKEVPNConnectionHref] = *connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIKEVPNConnections, connList)
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		err = fmt.Errorf("Error featching Base Controller URL: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-resource_controller_url").GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/ikepolicies")
	if !core.IsNil(ikePolicy.Name) {
		if err = d.Set("name", ikePolicy.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-name").GetDiag()
		}
		if err = d.Set(flex.ResourceName, ikePolicy.Name); err != nil {
			err = fmt.Errorf("Error setting resource_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-resource_name").GetDiag()
		}
	}
	return nil
}

func resourceIBMISIKEPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	diag := ikepUpdate(context, d, meta, id)
	if diag != nil {
		return diag
	}
	return resourceIBMISIKEPolicyRead(context, d, meta)
}

func ikepUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.UpdateIkePolicyOptions{
		ID: &id,
	}

	nameChangeFlag := d.HasChange(isIKEName)
	authenticationAlgChangeFlag := d.HasChange(isIKEAuthenticationAlg)
	authenticationAlgsChangeFlag := d.HasChange("authentication_algorithms")
	encryptionAlgChangeFlag := d.HasChange(isIKEEncryptionAlg)
	encryptionAlgsChangeFlag := d.HasChange("encryption_algorithms")
	keyLifetimeChangeFlag := d.HasChange(isIKEKeyLifeTime)
	dhGroupChangeFlag := d.HasChange(isIKEDhGroup)
	dhGroupsChangeFlag := d.HasChange("dh_groups")
	ikeVersionChangeFlag := d.HasChange(isIKEVERSION)

	if nameChangeFlag || authenticationAlgChangeFlag || authenticationAlgsChangeFlag || encryptionAlgChangeFlag || encryptionAlgsChangeFlag || dhGroupChangeFlag || dhGroupsChangeFlag || ikeVersionChangeFlag || keyLifetimeChangeFlag {

		ikePolicyPatchModel := &vpcv1.IkePolicyPatch{}
		if nameChangeFlag {
			name := d.Get(isIKEName).(string)
			ikePolicyPatchModel.Name = &name
		}
		if authenticationAlgChangeFlag {
			authenticationAlg := d.Get(isIKEAuthenticationAlg).(string)
			ikePolicyPatchModel.AuthenticationAlgorithm = &authenticationAlg
		}
		if authenticationAlgsChangeFlag {
			authenticationAlgs := d.Get("authentication_algorithms").([]interface{})
			ikePolicyPatchModel.AuthenticationAlgorithms = interfaceSliceToStringSlice(authenticationAlgs)
		}
		if encryptionAlgChangeFlag {
			encryptionAlg := d.Get(isIKEEncryptionAlg).(string)
			ikePolicyPatchModel.EncryptionAlgorithm = &encryptionAlg
		}
		if encryptionAlgsChangeFlag {
			encryptionAlgs := d.Get("encryption_algorithms").([]interface{})
			ikePolicyPatchModel.EncryptionAlgorithms = interfaceSliceToStringSlice(encryptionAlgs)
		}
		if keyLifetimeChangeFlag {
			keyLifetime := int64(d.Get(isIKEKeyLifeTime).(int))
			ikePolicyPatchModel.KeyLifetime = &keyLifetime
		}
		if dhGroupChangeFlag {
			dhGroup := int64(d.Get(isIKEDhGroup).(int))
			ikePolicyPatchModel.DhGroup = &dhGroup
		}
		if dhGroupsChangeFlag {
			dhGroups := d.Get("dh_groups").([]interface{})
			ikePolicyPatchModel.DhGroups = interfaceSliceToInt64Slice(dhGroups)
		}
		if ikeVersionChangeFlag {
			ikeVersion := int64(d.Get(isIKEVERSION).(int))
			ikePolicyPatchModel.IkeVersion = &ikeVersion
		}
		ikePolicyPatch, err := ikePolicyPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for IkePolicyPatch: %s", err.Error()), "ibm_is_ike_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		if !encryptionAlgsChangeFlag {
			ikePolicyPatch["encryption_algorithms"] = nil
		}
		if !dhGroupsChangeFlag {
			ikePolicyPatch["dh_groups"] = nil
		}
		if !authenticationAlgsChangeFlag {
			ikePolicyPatch["authentication_algorithms"] = nil
		}

		// if encryptionAlgsChangeFlag || dhGroupsChangeFlag || authenticationAlgsChangeFlag {
		// 	getIkePolicyOptions := &vpcv1.GetIkePolicyOptions{
		// 		ID: core.StringPtr(d.Id()),
		// 	}
		// 	_, etagResponse, etagErr := sess.GetIkePolicyWithContext(context, getIkePolicyOptions)
		// 	if etagErr != nil {
		// 		if etagResponse != nil && etagResponse.StatusCode == 404 {
		// 			d.SetId("")
		// 			return nil
		// 		}
		// 		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIkePolicyWithContext failed: %s", etagErr.Error()), "ibm_is_ike_policy", "update")
		// 		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// 		return tfErr.GetDiag()
		// 	}
		// 	eTag := etagResponse.Headers.Get("ETag")
		// 	options.IfMatch = &eTag
		// }

		options.IkePolicyPatch = ikePolicyPatch

		_, _, err = sess.UpdateIkePolicyWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateIkePolicyWithContext failed: %s", err.Error()), "ibm_is_ike_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISIKEPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	return ikepDelete(context, d, meta, id)
}

func ikepDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getikepoptions := &vpcv1.GetIkePolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIkePolicyWithContext(context, getikepoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIkePolicyWithContext failed: %s", err.Error()), "ibm_is_ike_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteIkePolicyOptions := &vpcv1.DeleteIkePolicyOptions{
		ID: &id,
	}
	response, err = sess.DeleteIkePolicy(deleteIkePolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteIkePolicyWithContext failed: %s", err.Error()), "ibm_is_ike_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISIKEPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()
	return ikepExists(d, meta, id)
}

func ikepExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "ibm_is_ike_policy", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	options := &vpcv1.GetIkePolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIkePolicy(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIkePolicy failed: %s", err.Error()), "ibm_is_ike_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	return true, nil
}

func interfaceSliceToStringSlice(values []interface{}) []string {
	result := make([]string, 0, len(values))

	for _, v := range values {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}

	return result
}
func interfaceSliceToInt64Slice(values []interface{}) []int64 {
	result := make([]int64, 0, len(values))

	for _, v := range values {
		if s, ok := v.(int); ok {
			result = append(result, int64(s))
		}
	}

	return result
}
