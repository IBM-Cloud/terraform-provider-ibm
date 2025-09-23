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
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEAuthenticationAlg),
				Description:  "Authentication algorithm type",
			},

			isIKEEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEEncryptionAlg),
				Description:  "Encryption alogorithm type",
			},

			isIKEDhGroup: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_ike_policy", isIKEDhGroup),
				Description:  "IKE DH group",
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
	authenticationAlg := d.Get(isIKEAuthenticationAlg).(string)
	encryptionAlg := d.Get(isIKEEncryptionAlg).(string)
	dhGroup := int64(d.Get(isIKEDhGroup).(int))

	diag := ikepCreate(context, d, meta, authenticationAlg, encryptionAlg, name, dhGroup)
	if diag != nil {
		return diag
	}
	return resourceIBMISIKEPolicyRead(context, d, meta)
}

func ikepCreate(context context.Context, d *schema.ResourceData, meta interface{}, authenticationAlg, encryptionAlg, name string, dhGroup int64) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateIkePolicyOptions{
		AuthenticationAlgorithm: &authenticationAlg,
		EncryptionAlgorithm:     &encryptionAlg,
		DhGroup:                 &dhGroup,
		Name:                    &name,
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
	if err = d.Set("encryption_algorithm", ikePolicy.EncryptionAlgorithm); err != nil {
		err = fmt.Errorf("Error setting encryption_algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ike_policy", "read", "set-encryption_algorithm").GetDiag()
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
	if d.HasChange(isIKEName) || d.HasChange(isIKEAuthenticationAlg) || d.HasChange(isIKEEncryptionAlg) || d.HasChange(isIKEDhGroup) || d.HasChange(isIKEVERSION) || d.HasChange(isIKEKeyLifeTime) {
		name := d.Get(isIKEName).(string)
		authenticationAlg := d.Get(isIKEAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIKEEncryptionAlg).(string)
		keyLifetime := int64(d.Get(isIKEKeyLifeTime).(int))
		dhGroup := int64(d.Get(isIKEDhGroup).(int))
		ikeVersion := int64(d.Get(isIKEVERSION).(int))

		ikePolicyPatchModel := &vpcv1.IkePolicyPatch{}
		ikePolicyPatchModel.Name = &name
		ikePolicyPatchModel.AuthenticationAlgorithm = &authenticationAlg
		ikePolicyPatchModel.EncryptionAlgorithm = &encryptionAlg
		ikePolicyPatchModel.KeyLifetime = &keyLifetime
		ikePolicyPatchModel.DhGroup = &dhGroup
		ikePolicyPatchModel.IkeVersion = &ikeVersion
		ikePolicyPatch, err := ikePolicyPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for IkePolicyPatch: %s", err.Error()), "ibm_is_ike_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
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
