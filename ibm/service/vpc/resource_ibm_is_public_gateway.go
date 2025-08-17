// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isPublicGatewayName              = "name"
	isPublicGatewayFloatingIP        = "floating_ip"
	isPublicGatewayStatus            = "status"
	isPublicGatewayVPC               = "vpc"
	isPublicGatewayZone              = "zone"
	isPublicGatewayFloatingIPAddress = "address"
	isPublicGatewayTags              = "tags"
	isPublicGatewayAccessTags        = "access_tags"

	isPublicGatewayProvisioning     = "provisioning"
	isPublicGatewayProvisioningDone = "available"
	isPublicGatewayDeleting         = "deleting"
	isPublicGatewayDeleted          = "done"
	isPublicGatewayCRN              = "crn"
	isPublicGatewayResourceGroup    = "resource_group"
)

func ResourceIBMISPublicGateway() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISPublicGatewayCreate,
		ReadContext:   resourceIBMISPublicGatewayRead,
		UpdateContext: resourceIBMISPublicGatewayUpdate,
		DeleteContext: resourceIBMISPublicGatewayDelete,
		Exists:        resourceIBMISPublicGatewayExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isPublicGatewayName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_public_gateway", isPublicGatewayName),
				Description:  "Name of the Public gateway instance",
			},

			isPublicGatewayFloatingIP: {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: flex.ApplyOnce,
			},

			isPublicGatewayStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway instance status",
			},

			isPublicGatewayResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Public gateway resource group info",
			},

			isPublicGatewayVPC: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Public gateway VPC info",
			},

			isPublicGatewayZone: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Public gateway zone info",
			},

			isPublicGatewayTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_public_gateway", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Service tags for the public gateway instance",
			},

			isPublicGatewayAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_public_gateway", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
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
			isPublicGatewayCRN: {
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
		},
	}
}

func ResourceIBMISPublicGatewayValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isPublicGatewayName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISPublicGatewayResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_public_gateway", Schema: validateSchema}
	return &ibmISPublicGatewayResourceValidator
}

func resourceIBMISPublicGatewayCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	name := d.Get(isPublicGatewayName).(string)
	vpc := d.Get(isPublicGatewayVPC).(string)
	zone := d.Get(isPublicGatewayZone).(string)

	options := &vpcv1.CreatePublicGatewayOptions{
		Name: &name,
		VPC: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	}
	floatingipID := ""
	floatingipadd := ""
	if floatingipdataIntf, ok := d.GetOk(isPublicGatewayFloatingIP); ok && floatingipdataIntf != nil {
		fip := &vpcv1.PublicGatewayFloatingIPPrototype{}
		floatingipdata := floatingipdataIntf.(map[string]interface{})
		if floatingipidintf, ok := floatingipdata["id"]; ok && floatingipidintf != nil {
			floatingipID = floatingipidintf.(string)
			fip.ID = &floatingipID
		}
		if floatingipaddintf, ok := floatingipdata[isPublicGatewayFloatingIPAddress]; ok && floatingipaddintf != nil {
			floatingipadd = floatingipaddintf.(string)
			fip.Address = &floatingipadd
		}
		options.FloatingIP = fip
	}
	if grp, ok := d.GetOk(isPublicGatewayResourceGroup); ok {
		rg := grp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	publicgw, _, err := sess.CreatePublicGatewayWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*publicgw.ID)
	log.Printf("[INFO] PublicGateway : %s", *publicgw.ID)

	_, err = isWaitForPublicGatewayAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for PublicGateway available failed: %s", err.Error()), "ibm_is_public_gateway", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isPublicGatewayTags); ok || v != "" {
		oldList, newList := d.GetChange(isPublicGatewayTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicgw.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc public gateway (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isPublicGatewayAccessTags); ok {
		oldList, newList := d.GetChange(isPublicGatewayAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicgw.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc public gateway (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISPublicGatewayRead(context, d, meta)
}

func isWaitForPublicGatewayAvailable(publicgwC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for public gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isPublicGatewayProvisioning},
		Target:     []string{isPublicGatewayProvisioningDone, ""},
		Refresh:    isPublicGatewayRefreshFunc(publicgwC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPublicGatewayRefreshFunc(publicgwC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
			ID: &id,
		}
		publicgw, response, err := publicgwC.GetPublicGateway(getPublicGatewayOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting Public Gateway : %s\n%s", err, response)
		}

		if *publicgw.Status == isPublicGatewayProvisioningDone {
			return publicgw, isPublicGatewayProvisioningDone, nil
		}

		return publicgw, isPublicGatewayProvisioning, nil
	}
}

func resourceIBMISPublicGatewayRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()
	getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
		ID: &id,
	}
	publicGateway, response, err := sess.GetPublicGatewayWithContext(context, getPublicGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(publicGateway.Name) {
		if err = d.Set("name", publicGateway.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-name").GetDiag()
		}
	}
	if publicGateway.FloatingIP != nil {
		floatIP := map[string]interface{}{
			"id":                             *publicGateway.FloatingIP.ID,
			isPublicGatewayFloatingIPAddress: *publicGateway.FloatingIP.Address,
		}
		if err = d.Set("floating_ip", floatIP); err != nil {
			err = fmt.Errorf("Error setting floating_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-floating_ip").GetDiag()
		}
	}
	if err = d.Set("status", publicGateway.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-status").GetDiag()
	}
	if err = d.Set("zone", *publicGateway.Zone.Name); err != nil {
		err = fmt.Errorf("Error setting zone: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-zone").GetDiag()
	}
	if err = d.Set("vpc", *publicGateway.VPC.ID); err != nil {
		err = fmt.Errorf("Error setting vpc: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-vpc").GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *publicGateway.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of vpc public gateway (%s) tags: %s", id, err)
	}
	if err = d.Set("tags", tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-tags").GetDiag()
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *publicGateway.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of vpc public gateway (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set("access_tags", accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_public_gateway", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/publicGateways")
	if !core.IsNil(publicGateway.Name) {
		if err = d.Set("resource_name", publicGateway.Name); err != nil {
			err = fmt.Errorf("Error setting resource_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-resource_name").GetDiag()
		}
	}
	if err = d.Set("resource_crn", publicGateway.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("crn", publicGateway.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-crn").GetDiag()
	}
	if err = d.Set("resource_status", publicGateway.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-status").GetDiag()
	}
	if publicGateway.ResourceGroup != nil {
		if err = d.Set("resource_group", *publicGateway.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set("resource_group_name", *publicGateway.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "read", "set-resource_group_name").GetDiag()
		}
	}
	return nil
}

func resourceIBMISPublicGatewayUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	name := ""
	hasChanged := false

	if d.HasChange(isPublicGatewayName) {
		name = d.Get(isPublicGatewayName).(string)
		hasChanged = true
	}
	if d.HasChange(isPublicGatewayTags) {
		getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
			ID: &id,
		}
		publicgw, _, err := sess.GetPublicGatewayWithContext(context, getPublicGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isPublicGatewayTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicgw.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource Public Gateway (%s) tags: %s", id, err)
		}
	}

	if d.HasChange(isPublicGatewayAccessTags) {
		getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
			ID: &id,
		}
		publicgw, _, err := sess.GetPublicGatewayWithContext(context, getPublicGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isPublicGatewayTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicgw.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource Public Gateway (%s) access tags: %s", d.Id(), err)
		}
	}

	if hasChanged {
		updatePublicGatewayOptions := &vpcv1.UpdatePublicGatewayOptions{
			ID: &id,
		}
		PublicGatewayPatchModel := &vpcv1.PublicGatewayPatch{
			Name: &name,
		}
		PublicGatewayPatch, err := PublicGatewayPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PublicGatewayPatchModel.AsPatch failed: %s", err.Error()), "ibm_is_public_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updatePublicGatewayOptions.PublicGatewayPatch = PublicGatewayPatch
		_, _, err = sess.UpdatePublicGateway(updatePublicGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISPublicGatewayRead(context, d, meta)
}

func resourceIBMISPublicGatewayDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_public_gateway", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetPublicGatewayWithContext(context, getPublicGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePublicGatewayOptions := &vpcv1.DeletePublicGatewayOptions{
		ID: &id,
	}
	response, err = sess.DeletePublicGatewayWithContext(context, deletePublicGatewayOptions)
	if err != nil {
		if response.StatusCode == 409 && strings.Contains(strings.ToLower(err.Error()), strings.ToLower("The Public Gateway is in use by subnet")) {
			listSubnetsOptions := &vpcv1.ListSubnetsOptions{}
			subnets, _, _ := sess.ListSubnets(listSubnetsOptions)
			for _, s := range subnets.Subnets {
				if s.PublicGateway != nil && id == *s.PublicGateway.ID {
					unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
						ID: s.ID,
					}
					res, errSub := sess.UnsetSubnetPublicGatewayWithContext(context, unsetSubnetPublicGatewayOptions)
					if res.StatusCode == 204 {
						_, err = isWaitForSubnetPublicGatewayUnset(sess, *s.ID, d.Timeout(schema.TimeoutDelete))
						if err != nil {
							tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnsetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "delete")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
						response, err = sess.DeletePublicGatewayWithContext(context, deletePublicGatewayOptions)
						if err != nil {
							tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "delete")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
					} else {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error Unsetting Public Gateway : %s\n%s", errSub, res), "ibm_is_public_gateway", "delete")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}
		} else {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePublicGatewayWithContext failed: %s", err.Error()), "ibm_is_public_gateway", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	_, err = isWaitForPublicGatewayDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for PublicGateway deleted failed: %s", err.Error()), "ibm_is_public_gateway", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForPublicGatewayDeleted(pg *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for public gateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isPublicGatewayDeleting},
		Target:     []string{isPublicGatewayDeleted, ""},
		Refresh:    isPublicGatewayDeleteRefreshFunc(pg, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPublicGatewayDeleteRefreshFunc(pg *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] is pubic gateway delete function here")
		getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
			ID: &id,
		}
		pgw, response, err := pg.GetPublicGateway(getPublicGatewayOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return pgw, isPublicGatewayDeleted, nil
			}
			return nil, "", fmt.Errorf("[ERROR] The Public Gateway %s failed to delete: %s\n%s", id, err, response)
		}
		return pgw, isPublicGatewayDeleting, nil
	}
}

func resourceIBMISPublicGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	id := d.Id()
	getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetPublicGateway(getPublicGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Public Gateway: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForSubnetPublicGatewayUnset(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for public gateway (%s) to be unset.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", "wait"},
		Target:     []string{"done", ""},
		Refresh:    isSubnetPublicGatewayUnsetRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetPublicGatewayUnsetRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	log.Printf("Waiting for public gateway (%s) to be unset.", id)
	return func() (interface{}, string, error) {
		getSubnetPublicGatewayOptions := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &id,
		}
		subnetPublicGateway, response, err := subnetC.GetSubnetPublicGateway(getSubnetPublicGatewayOptions)
		if err != nil {
			if response.StatusCode == 404 {
				return subnetPublicGateway, "done", nil
			}
			return subnetPublicGateway, "", fmt.Errorf("[ERROR] Error getting Subnet PublicGateway : %s\n%s", err, response)
		}

		return subnetPublicGateway, "wait", nil
	}
}
