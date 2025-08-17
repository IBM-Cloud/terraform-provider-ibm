// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isVPCAddressPrefixPrefixName = "name"
	isVPCAddressPrefixZoneName   = "zone"
	isVPCAddressPrefixCIDR       = "cidr"
	isVPCAddressPrefixVPCID      = "vpc"
	isVPCAddressPrefixHasSubnets = "has_subnets"
	isVPCAddressPrefixDefault    = "is_default"
	isAddressPrefix              = "address_prefix"
)

func ResourceIBMISVpcAddressPrefix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVpcAddressPrefixCreate,
		ReadContext:   resourceIBMISVpcAddressPrefixRead,
		UpdateContext: resourceIBMISVpcAddressPrefixUpdate,
		DeleteContext: resourceIBMISVpcAddressPrefixDelete,
		Exists:        resourceIBMISVpcAddressPrefixExists,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isVPCAddressPrefixPrefixName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_address_prefix", isVPCAddressPrefixPrefixName),
				Description:  "Name",
			},
			isVPCAddressPrefixZoneName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone name",
			},

			isVPCAddressPrefixCIDR: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_address_prefix", isVPCAddressPrefixCIDR),
				Description:  "CIDIR address prefix",
			},
			isVPCAddressPrefixDefault: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Is default prefix for this zone in this VPC",
			},

			isVPCAddressPrefixVPCID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC id",
			},

			isVPCAddressPrefixHasSubnets: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Boolean value, set to true if VPC instance have subnets",
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the VPC resource",
			},

			isAddressPrefix: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the address prefix",
			},
		},
	}
}

func ResourceIBMISAddressPrefixValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPCAddressPrefixPrefixName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPCAddressPrefixCIDR,
			ValidateFunctionIdentifier: validate.ValidateOverlappingAddress,
			Type:                       validate.TypeString,
			ForceNew:                   true,
			Required:                   true})

	ibmISAddressPrefixResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_address_prefix", Schema: validateSchema}
	return &ibmISAddressPrefixResourceValidator
}

func resourceIBMISVpcAddressPrefixCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	isDefault := false
	prefixName := d.Get(isVPCAddressPrefixPrefixName).(string)
	zoneName := d.Get(isVPCAddressPrefixZoneName).(string)
	cidr := d.Get(isVPCAddressPrefixCIDR).(string)
	vpcID := d.Get(isVPCAddressPrefixVPCID).(string)
	if isDefaultPrefix, ok := d.GetOk(isVPCAddressPrefixDefault); ok {
		isDefault = isDefaultPrefix.(bool)
	}

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	conns.IbmMutexKV.Lock(isVPCAddressPrefixKey)
	defer conns.IbmMutexKV.Unlock(isVPCAddressPrefixKey)

	err := vpcAddressPrefixCreate(context, d, meta, prefixName, zoneName, cidr, vpcID, isDefault)
	if err != nil {
		return err
	}
	return resourceIBMISVpcAddressPrefixRead(context, d, meta)
}

func vpcAddressPrefixCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, zone, cidr, vpcID string, isDefault bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateVPCAddressPrefixOptions{
		Name:      &name,
		VPCID:     &vpcID,
		CIDR:      &cidr,
		IsDefault: &isDefault,
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	}
	addrPrefix, _, err := sess.CreateVPCAddressPrefixWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPCAddressPrefixWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	addrPrefixID := *addrPrefix.ID
	d.SetId(fmt.Sprintf("%s/%s", vpcID, addrPrefixID))
	d.Set(isAddressPrefix, addrPrefixID)
	return nil
}

func resourceIBMISVpcAddressPrefixRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "sep-id-parts").GetDiag()
	}

	vpcID := parts[0]
	addrPrefixID := parts[1]
	error := vpcAddressPrefixGet(context, d, meta, vpcID, addrPrefixID)
	if error != nil {
		return error
	}

	return nil
}

func vpcAddressPrefixGet(context context.Context, d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
		VPCID: &vpcID,
		ID:    &addrPrefixID,
	}
	addrPrefix, response, err := sess.GetVPCAddressPrefixWithContext(context, getvpcAddressPrefixOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCAddressPrefixWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(isVPCAddressPrefixVPCID, vpcID); err != nil {
		err = fmt.Errorf("Error setting vpc: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-vpc").GetDiag()
	}
	if !core.IsNil(addrPrefix.IsDefault) {
		if err = d.Set("is_default", addrPrefix.IsDefault); err != nil {
			err = fmt.Errorf("Error setting is_default: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-is_default").GetDiag()
		}
	}
	if !core.IsNil(addrPrefix.Name) {
		if err = d.Set("name", addrPrefix.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-name").GetDiag()
		}
	}
	if addrPrefix.Zone != nil {
		if err = d.Set(isVPCAddressPrefixZoneName, *addrPrefix.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-zone").GetDiag()
		}
	}
	if err = d.Set("cidr", addrPrefix.CIDR); err != nil {
		err = fmt.Errorf("Error setting cidr: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-cidr").GetDiag()
	}
	if err = d.Set("has_subnets", addrPrefix.HasSubnets); err != nil {
		err = fmt.Errorf("Error setting has_subnets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-has_subnets").GetDiag()
	}
	if err = d.Set(isAddressPrefix, addrPrefixID); err != nil {
		err = fmt.Errorf("Error setting address_prefix: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-address_prefix").GetDiag()
	}
	getVPCOptions := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}
	vpc, response, err := sess.GetVPCWithContext(context, getVPCOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.RelatedCRN, *vpc.CRN); err != nil {
		err = fmt.Errorf("Error setting related_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "read", "set-related_crn").GetDiag()
	}

	return nil
}

func resourceIBMISVpcAddressPrefixUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := ""
	isDefault := false
	hasNameChanged := false
	hasIsDefaultChanged := false

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "update", "sep-id-parts").GetDiag()
	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	conns.IbmMutexKV.Lock(isVPCAddressPrefixKey)
	defer conns.IbmMutexKV.Unlock(isVPCAddressPrefixKey)

	if d.HasChange(isVPCAddressPrefixPrefixName) {
		name = d.Get(isVPCAddressPrefixPrefixName).(string)
		hasNameChanged = true
	}
	if d.HasChange(isVPCAddressPrefixDefault) {
		isDefault = d.Get(isVPCAddressPrefixDefault).(bool)
		hasIsDefaultChanged = true
	}
	error := vpcAddressPrefixUpdate(context, d, meta, vpcID, addrPrefixID, name, isDefault, hasNameChanged, hasIsDefaultChanged)
	if error != nil {
		return error
	}

	return resourceIBMISVpcAddressPrefixRead(context, d, meta)
}

func vpcAddressPrefixUpdate(context context.Context, d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID, name string, isDefault, hasNameChanged, hasIsDefaultChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if hasNameChanged || hasIsDefaultChanged {
		updatevpcAddressPrefixoptions := &vpcv1.UpdateVPCAddressPrefixOptions{
			VPCID: &vpcID,
			ID:    &addrPrefixID,
		}

		addressPrefixPatchModel := &vpcv1.AddressPrefixPatch{}
		if hasNameChanged {
			addressPrefixPatchModel.Name = &name
		}
		if hasIsDefaultChanged {
			addressPrefixPatchModel.IsDefault = &isDefault
		}
		addressPrefixPatch, err := addressPrefixPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("addressPrefixPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updatevpcAddressPrefixoptions.AddressPrefixPatch = addressPrefixPatch
		_, _, err = sess.UpdateVPCAddressPrefixWithContext(context, updatevpcAddressPrefixoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPCAddressPrefixWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISVpcAddressPrefixDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "delete", "sep-id-parts").GetDiag()

	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	conns.IbmMutexKV.Lock(isVPCAddressPrefixKey)
	defer conns.IbmMutexKV.Unlock(isVPCAddressPrefixKey)

	error := vpcAddressPrefixDelete(context, d, meta, vpcID, addrPrefixID)
	if error != nil {
		return error
	}

	d.SetId("")
	return nil
}

func vpcAddressPrefixDelete(context context.Context, d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
		VPCID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVPCAddressPrefixWithContext(context, getvpcAddressPrefixOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCAddressPrefixWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletevpcAddressPrefixOptions := &vpcv1.DeleteVPCAddressPrefixOptions{
		VPCID: &vpcID,
		ID:    &addrPrefixID,
	}
	response, err = sess.DeleteVPCAddressPrefixWithContext(context, deletevpcAddressPrefixOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCAddressPrefixWithContext failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISVpcAddressPrefixExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	parts, err := flex.IdParts(d.Id())
	if len(parts) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of vpcID/addrPrefixID", d.Id())
	}
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "exists", "sep-id-parts")
	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	exists, err := vpcAddressPrefixExists(d, meta, vpcID, addrPrefixID)
	return exists, err
}

func vpcAddressPrefixExists(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_address_prefix", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
		VPCID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCAddressPrefix failed: %s", err.Error()), "ibm_is_vpc_address_prefix", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
