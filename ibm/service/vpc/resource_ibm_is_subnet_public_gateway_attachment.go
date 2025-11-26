// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isPublicGatewayID                  = "public_gateway"
	IsPublicGatewayResourceType        = "resource_type"
	IsPublicGatewayAttachmentAvailable = "available"
	IsPublicGatewayAttachmentDeleting  = "deleting"
	IsPublicGatewayAttachmentFailed    = "failed"
	IsPublicGatewayAttachmentPending   = "pending"
)

func ResourceIBMISSubnetPublicGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSubnetPublicGatewayAttachmentCreate,
		ReadContext:   resourceIBMISSubnetPublicGatewayAttachmentRead,
		UpdateContext: resourceIBMISSubnetPublicGatewayAttachmentUpdate,
		DeleteContext: resourceIBMISSubnetPublicGatewayAttachmentDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isSubnetID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet identifier",
			},

			isPublicGatewayID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of public gateway",
			},

			isPublicGatewayName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Public gateway instance",
			},

			isPublicGatewayFloatingIP: {
				Type:     schema.TypeMap,
				Computed: true,
			},

			isPublicGatewayStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway instance status",
			},

			isPublicGatewayResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway resource group info",
			},

			isPublicGatewayVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway VPC info",
			},

			isPublicGatewayZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway zone info",
			},

			IsPublicGatewayResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			isPublicGatewayCRN: {
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

func resourceIBMISSubnetPublicGatewayAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	subnet := d.Get(isSubnetID).(string)
	publicGateway := d.Get(isPublicGatewayID).(string)

	publicGatewayIdentity := &vpcv1.PublicGatewayIdentity{
		ID: &publicGateway,
	}

	// Construct an instance of the SetSubnetPublicGatewayOptions
	setSubnetPublicGatewayOptions := &vpcv1.SetSubnetPublicGatewayOptions{
		ID:                    &subnet,
		PublicGatewayIdentity: publicGatewayIdentity,
	}

	pg, _, err := sess.SetSubnetPublicGatewayWithContext(context, setSubnetPublicGatewayOptions)

	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(subnet)
	_, err = isWaitForSubnetPublicGatewayAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetPublicGatewayAvailable failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	log.Printf("[INFO] Public Gateway : %s", *pg.ID)
	log.Printf("[INFO] Subnet ID : %s", subnet)

	return resourceIBMISSubnetPublicGatewayAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetPublicGatewayAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
		ID: &id,
	}
	pg, response, err := sess.GetSubnetPublicGatewayWithContext(context, getSubnetPublicGatewayOptionsModel)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isPublicGatewayName, pg.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-name").GetDiag()
	}
	if err = d.Set(isSubnetID, id); err != nil {
		err = fmt.Errorf("Error setting subnet: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-subnet").GetDiag()
	}

	if err = d.Set(isPublicGatewayID, pg.ID); err != nil {
		err = fmt.Errorf("Error setting public_gateway: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-public_gateway").GetDiag()
	}
	if pg.FloatingIP != nil {
		floatIP := map[string]interface{}{
			"id":                             *pg.FloatingIP.ID,
			isPublicGatewayFloatingIPAddress: *pg.FloatingIP.Address,
		}
		if err = d.Set(isPublicGatewayFloatingIP, floatIP); err != nil {
			err = fmt.Errorf("Error setting floating_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-floating_ip").GetDiag()
		}
	}

	if err = d.Set(isPublicGatewayStatus, pg.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-status").GetDiag()
	}
	if pg.ResourceGroup != nil {
		if err = d.Set(isPublicGatewayResourceGroup, *pg.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *pg.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-resource_group_name").GetDiag()
		}
	}
	if pg.VPC != nil {
		if err = d.Set(isPublicGatewayVPC, *pg.VPC.ID); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-vpc").GetDiag()
		}
	}
	if pg.Zone != nil {
		if err = d.Set(isPublicGatewayZone, *pg.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-zone").GetDiag()
		}
	}
	if err = d.Set(IsPublicGatewayResourceType, pg.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set(isPublicGatewayCRN, pg.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "read", "set-crn").GetDiag()
	}
	return nil
}

func resourceIBMISSubnetPublicGatewayAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isPublicGatewayID) {
		subnet := d.Get(isSubnetID).(string)
		publicGateway := d.Get(isPublicGatewayID).(string)

		publicGatewayIdentity := &vpcv1.PublicGatewayIdentity{
			ID: &publicGateway,
		}

		// Construct an instance of the SetSubnetPublicGatewayOptions
		setSubnetPublicGatewayOptions := &vpcv1.SetSubnetPublicGatewayOptions{
			ID:                    &subnet,
			PublicGatewayIdentity: publicGatewayIdentity,
		}

		pg, _, err := sess.SetSubnetPublicGatewayWithContext(context, setSubnetPublicGatewayOptions)

		if err != nil || pg == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with public gateway(%s)", subnet, publicGateway)

		d.SetId(subnet)
		return resourceIBMISSubnetPublicGatewayAttachmentRead(context, d, meta)
	}

	return resourceIBMISSubnetPublicGatewayAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetPublicGatewayAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_public_gateway_attachment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Get subnet details
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnetWithContext(context, getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Construct an instance of the UnsetSubnetPublicGatewayOptions model
	unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
		ID: &id,
	}
	_, err = sess.UnsetSubnetPublicGatewayWithContext(context, unsetSubnetPublicGatewayOptions)

	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnsetSubnetPublicGatewayWithContext failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForSubnetPublicGatewayDelete(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSubnetPublicGatewayDelete failed: %s", err.Error()), "ibm_is_subnet_public_gateway_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForSubnetPublicGatewayAvailable(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) public gateway attachment to be available.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{IsPublicGatewayAttachmentPending, IsPublicGatewayAttachmentDeleting},
		Target:     []string{IsPublicGatewayAttachmentAvailable, IsPublicGatewayAttachmentFailed, ""},
		Refresh:    isSubnetPublicGatewayRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetPublicGatewayRefreshFunc(subnetC *vpcv1.VpcV1, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &id,
		}
		pg, response, err := subnetC.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return pg, "", fmt.Errorf("[ERROR] Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
			}
			return pg, "", fmt.Errorf("[ERROR] Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
		}

		if *pg.Status == "failed" {
			return pg, IsPublicGatewayAttachmentFailed, fmt.Errorf("[ERROR] Error subnet (%s) public gateway attachment failed: %s\n%s", id, err, response)
		}

		return pg, *pg.Status, nil
	}
}

func isWaitForSubnetPublicGatewayDelete(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) public gateway attachment to be detached.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{IsPublicGatewayAttachmentPending, IsPublicGatewayAttachmentDeleting},
		Target:     []string{IsPublicGatewayAttachmentAvailable, IsPublicGatewayAttachmentFailed, ""},
		Refresh:    isSubnetPublicGatewayDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetPublicGatewayDeleteRefreshFunc(subnetC *vpcv1.VpcV1, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &id,
		}
		pg, response, err := subnetC.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return pg, "", nil
			}
			return pg, "", fmt.Errorf("[ERROR] Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
		}

		if *pg.Status == "failed" {
			return pg, IsPublicGatewayAttachmentFailed, fmt.Errorf("[ERROR] Error subnet (%s) public gateway attachment failed: %s\n%s", id, err, response)
		}

		return pg, *pg.Status, nil
	}
}
