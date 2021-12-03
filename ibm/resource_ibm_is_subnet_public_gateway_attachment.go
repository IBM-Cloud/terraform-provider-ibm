// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

func resourceIBMISSubnetPublicGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSubnetPublicGatewayAttachmentCreate,
		Read:     resourceIBMISSubnetPublicGatewayAttachmentRead,
		Update:   resourceIBMISSubnetPublicGatewayAttachmentUpdate,
		Delete:   resourceIBMISSubnetPublicGatewayAttachmentDelete,
		Exists:   resourceIBMISSubnetPublicGatewayAttachmentExists,
		Importer: &schema.ResourceImporter{},
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
				Description: "The unique identifier of network ACL",
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

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISSubnetPublicGatewayAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
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

	pg, response, err := sess.SetSubnetPublicGateway(setSubnetPublicGatewayOptions)

	if err != nil {
		log.Printf("[DEBUG] Error while attaching public gateway(%s) to subnet(%s) %s\n%s", publicGateway, subnet, err, response)
		return fmt.Errorf("Error while attaching public gateway(%s) to subnet(%s) %s\n%s", publicGateway, subnet, err, response)
	}
	d.SetId(subnet)
	_, err = isWaitForSubnetPublicGatewayAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	log.Printf("[INFO] Public Gateway : %s", *pg.ID)
	log.Printf("[INFO] Subnet ID : %s", subnet)

	return resourceIBMISSubnetPublicGatewayAttachmentRead(d, meta)
}

func resourceIBMISSubnetPublicGatewayAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
		ID: &id,
	}
	pg, response, err := sess.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
	}
	d.Set(isPublicGatewayName, pg.Name)
	if pg.FloatingIP != nil {
		floatIP := map[string]interface{}{
			"id":                             *pg.FloatingIP.ID,
			isPublicGatewayFloatingIPAddress: *pg.FloatingIP.Address,
		}
		d.Set(isPublicGatewayFloatingIP, floatIP)
	}
	d.Set(isPublicGatewayStatus, pg.Status)
	if pg.ResourceGroup != nil {
		d.Set(isPublicGatewayResourceGroup, *pg.ResourceGroup.ID)
		d.Set(ResourceGroupName, *pg.ResourceGroup.Name)
	}
	d.Set(isPublicGatewayVPC, *pg.VPC.ID)
	d.Set(isPublicGatewayZone, *pg.Zone.Name)
	d.Set(IsPublicGatewayResourceType, pg.ResourceType)
	d.Set(isPublicGatewayCRN, pg.CRN)

	return nil
}

func resourceIBMISSubnetPublicGatewayAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
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

		pg, response, err := sess.SetSubnetPublicGateway(setSubnetPublicGatewayOptions)

		if err != nil || pg == nil {
			log.Printf("[DEBUG] Error while attaching public gateway(%s) to subnet(%s) %s\n%s", publicGateway, subnet, err, response)
			return fmt.Errorf("Error while attaching public gateway(%s) to subnet(%s) %s\n%s", publicGateway, subnet, err, response)
		}
		// log.Printf("[INFO] Updated subnet %s with public gateway(%s) : %s", subnet, publicGateway, *resultACL.ID)

		d.SetId(subnet)
		return resourceIBMISSubnetPublicGatewayAttachmentRead(d, meta)
	}

	return resourceIBMISSubnetPublicGatewayAttachmentRead(d, meta)
}

func resourceIBMISSubnetPublicGatewayAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	// Set the subnet with VPC default network ACL
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}

	// Construct an instance of the UnsetSubnetPublicGatewayOptions model
	unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
		ID: &id,
	}
	res, err := sess.UnsetSubnetPublicGateway(unsetSubnetPublicGatewayOptions)

	if err != nil {
		log.Printf("[DEBUG] Error while detaching public gateway to subnet %s\n%s", err, res)
		return fmt.Errorf("Error while detaching public gateway to subnet %s\n%s", err, res)
	}
	_, err = isWaitForSubnetPublicGatewayDelete(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISSubnetPublicGatewayAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
		ID: &id,
	}
	pg, response, err := sess.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)
	if err != nil || pg == nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
	}
	return true, nil
}

func isWaitForSubnetPublicGatewayAvailable(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) public gateway attachment to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{IsPublicGatewayAttachmentPending, IsPublicGatewayAttachmentDeleting},
		Target:     []string{IsPublicGatewayAttachmentAvailable, IsPublicGatewayAttachmentFailed, ""},
		Refresh:    isSubnetPublicGatewayRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetPublicGatewayRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &id,
		}
		pg, response, err := subnetC.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return pg, "", fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
			}
			return pg, "", fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
		}

		if *pg.Status == "failed" {
			return pg, IsPublicGatewayAttachmentFailed, fmt.Errorf("Error subnet (%s) public gateway attachment failed: %s\n%s", id, err, response)
		}

		return pg, *pg.Status, nil
	}
}

func isWaitForSubnetPublicGatewayDelete(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) public gateway attachment to be detached.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{IsPublicGatewayAttachmentPending, IsPublicGatewayAttachmentDeleting},
		Target:     []string{IsPublicGatewayAttachmentAvailable, IsPublicGatewayAttachmentFailed, ""},
		Refresh:    isSubnetPublicGatewayDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetPublicGatewayDeleteRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &id,
		}
		pg, response, err := subnetC.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return pg, "", fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
			}
			return pg, "", fmt.Errorf("Error getting subnet's (%s) attached public gateway: %s\n%s", id, err, response)
		}

		if *pg.Status == "failed" {
			return pg, IsPublicGatewayAttachmentFailed, fmt.Errorf("Error subnet (%s) public gateway attachment failed: %s\n%s", id, err, response)
		}

		return pg, *pg.Status, nil
	}
}
