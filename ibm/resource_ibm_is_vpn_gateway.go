package ibm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isVPNGatewayName             = "name"
	isVPNGatewayResourceGroup    = "resource_group"
	isVPNGatewayTags             = "tags"
	isVPNGatewaySubnet           = "subnet"
	isVPNGatewayStatus           = "status"
	isVPNGatewayDeleting         = "deleting"
	isVPNGatewayDeleted          = "done"
	isVPNGatewayProvisioning     = "provisioning"
	isVPNGatewayProvisioningDone = "done"
	isVPNGatewayPublicIPAddress  = "public_ip_address"
)

func resourceIBMISVPNGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPNGatewayCreate,
		Read:     resourceIBMISVPNGatewayRead,
		Update:   resourceIBMISVPNGatewayUpdate,
		Delete:   resourceIBMISVPNGatewayDelete,
		Exists:   resourceIBMISVPNGatewayExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			isVPNGatewayName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "VPN Gateway instance name",
			},

			isVPNGatewaySubnet: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPNGateway subnet info",
			},

			isVPNGatewayResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isVPNGatewayStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPNGatewayPublicIPAddress: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPNGatewayTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "VPN Gateway tags list",
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISVPNGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] VPNGateway create")
	name := d.Get(isVPNGatewayName).(string)
	subnetID := d.Get(isVPNGatewaySubnet).(string)

	if userDetails.generation == 1 {
		err := classicVpngwCreate(d, meta, name, subnetID)
		if err != nil {
			return err
		}
	} else {
		err := vpngwCreate(d, meta, name, subnetID)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVPNGatewayRead(d, meta)
}

func classicVpngwCreate(d *schema.ResourceData, meta interface{}, name, subnetID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateVpnGatewayOptions{
		Subnet: &vpcclassicv1.SubnetIdentity{
			ID: &subnetID,
		},
		Name: &name,
	}

	if rgrp, ok := d.GetOk(isVPNGatewayResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcclassicv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	VPNGateway, response, err := sess.CreateVpnGateway(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create vpc VPN Gateway %s\n%s", err, response)
	}
	_, err = isWaitForClassicVpnGatewayAvailable(sess, *VPNGateway.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	d.SetId(*VPNGateway.ID)
	log.Printf("[INFO] VPNGateway : %s", *VPNGateway.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVPNGatewayTags); ok || v != "" {
		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *VPNGateway.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func vpngwCreate(d *schema.ResourceData, meta interface{}, name, subnetID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateVpnGatewayOptions{
		Subnet: &vpcv1.SubnetIdentity{
			ID: &subnetID,
		},
		Name: &name,
	}

	if rgrp, ok := d.GetOk(isVPNGatewayResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	VPNGateway, response, err := sess.CreateVpnGateway(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create vpc VPN Gateway %s\n%s", err, response)
	}
	_, err = isWaitForVpnGatewayAvailable(sess, *VPNGateway.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	d.SetId(*VPNGateway.ID)
	log.Printf("[INFO] VPNGateway : %s", *VPNGateway.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVPNGatewayTags); ok || v != "" {
		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *VPNGateway.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func isWaitForClassicVpnGatewayAvailable(vpnGateway *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for vpn gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayProvisioning},
		Target:     []string{isVPNGatewayProvisioningDone, ""},
		Refresh:    isClassicVpnGatewayRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicVpnGatewayRefreshFunc(vpnGateway *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
			ID: &id,
		}
		vpngw, response, err := vpnGateway.GetVpnGateway(getVpnGatewayOptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Vpn Gateway: %s\n%s", err, response)
		}

		if *vpngw.Status == "available" || *vpngw.Status == "failed" || *vpngw.Status == "running" {
			return vpngw, isVPNGatewayProvisioningDone, nil
		}

		return vpngw, isVPNGatewayProvisioning, nil
	}
}

func isWaitForVpnGatewayAvailable(vpnGateway *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for vpn gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayProvisioning},
		Target:     []string{isVPNGatewayProvisioningDone, ""},
		Refresh:    isVpnGatewayRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpnGatewayRefreshFunc(vpnGateway *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
			ID: &id,
		}
		vpngw, response, err := vpnGateway.GetVpnGateway(getVpnGatewayOptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Vpn Gateway: %s\n%s", err, response)
		}

		if *vpngw.Status == "available" || *vpngw.Status == "failed" || *vpngw.Status == "running" {
			return vpngw, isVPNGatewayProvisioningDone, nil
		}

		return vpngw, isVPNGatewayProvisioning, nil
	}
}

func resourceIBMISVPNGatewayRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	if userDetails.generation == 1 {
		err := classicVpngwGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := vpngwGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpngwGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
		ID: &id,
	}
	VPNGateway, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Vpn Gateway (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isVPNGatewayName, *VPNGateway.Name)
	d.Set(isVPNGatewaySubnet, *VPNGateway.Subnet.ID)
	d.Set(isVPNGatewayStatus, *VPNGateway.Status)
	d.Set(isVPNGatewayPublicIPAddress, *VPNGateway.PublicIp.Address)
	tags, err := GetTagsUsingCRN(meta, *VPNGateway.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVPNGatewayTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/vpngateways")
	d.Set(ResourceName, *VPNGateway.Name)
	d.Set(ResourceCRN, *VPNGateway.Crn)
	d.Set(ResourceStatus, *VPNGateway.Status)
	if VPNGateway.ResourceGroup != nil {
		d.Set(ResourceGroupName, *VPNGateway.ResourceGroup.ID)
		d.Set(isVPNGatewayResourceGroup, *VPNGateway.ResourceGroup.ID)
	}
	return nil
}

func vpngwGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
		ID: &id,
	}
	VPNGateway, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Vpn Gateway (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isVPNGatewayName, *VPNGateway.Name)
	d.Set(isVPNGatewaySubnet, *VPNGateway.Subnet.ID)
	d.Set(isVPNGatewayStatus, *VPNGateway.Status)
	d.Set(isVPNGatewayPublicIPAddress, *VPNGateway.PublicIp.Address)
	tags, err := GetTagsUsingCRN(meta, *VPNGateway.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVPNGatewayTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/vpngateways")
	d.Set(ResourceName, *VPNGateway.Name)
	d.Set(ResourceCRN, *VPNGateway.Crn)
	d.Set(ResourceStatus, *VPNGateway.Status)
	if VPNGateway.ResourceGroup != nil {
		d.Set(ResourceGroupName, *VPNGateway.ResourceGroup.Name)
		d.Set(isVPNGatewayResourceGroup, *VPNGateway.ResourceGroup.ID)
	}
	return nil
}

func resourceIBMISVPNGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	name := ""
	hasChanged := false

	if d.HasChange(isVPNGatewayName) {
		name = d.Get(isVPNGatewayName).(string)
		hasChanged = true
	}

	if userDetails.generation == 1 {
		err := classicVpngwUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := vpngwUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVPNGatewayRead(d, meta)
}

func classicVpngwUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isVPNGatewayTags) {
		getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
			ID: &id,
		}
		VPNGateway, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
		if err != nil {
			return fmt.Errorf("Error getting Volume : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *VPNGateway.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Vpn Gateway (%s) tags: %s", id, err)
		}
	}
	if hasChanged {
		options := &vpcclassicv1.UpdateVpnGatewayOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateVpnGateway(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc Vpn Gateway: %s\n%s", err, response)
		}
	}
	return nil
}

func vpngwUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isVPNGatewayTags) {
		getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
			ID: &id,
		}
		VPNGateway, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
		if err != nil {
			return fmt.Errorf("Error getting Volume : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *VPNGateway.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Vpn Gateway (%s) tags: %s", id, err)
		}
	}
	if hasChanged {
		options := &vpcv1.UpdateVpnGatewayOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateVpnGateway(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc Vpn Gateway: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISVPNGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicVpngwDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := vpngwDelete(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpngwDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Vpn Gateway (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	options := &vpcclassicv1.DeleteVpnGatewayOptions{
		ID: &id,
	}
	response, err = sess.DeleteVpnGateway(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Vpn Gateway : %s\n%s", err, response)
	}
	_, err = isWaitForClassicVpnGatewayDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func vpngwDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Vpn Gateway (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	options := &vpcv1.DeleteVpnGatewayOptions{
		ID: &id,
	}
	response, err = sess.DeleteVpnGateway(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Vpn Gateway : %s\n%s", err, response)
	}
	_, err = isWaitForVpnGatewayDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForClassicVpnGatewayDeleted(vpnGateway *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayDeleting},
		Target:     []string{isVPNGatewayDeleted, ""},
		Refresh:    isClassicVpnGatewayDeleteRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicVpnGatewayDeleteRefreshFunc(vpnGateway *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
			ID: &id,
		}
		vpngw, response, err := vpnGateway.GetVpnGateway(getVpnGatewayOptions)
		if err != nil && response.StatusCode != 404 {
			return vpngw, "", fmt.Errorf("Error Getting Vpn Gateway: %s\n%s", err, response)
		}
		if response.StatusCode == 404 {
			return vpngw, isVPNGatewayDeleted, nil
		}
		return nil, isVPNGatewayDeleting, err
	}
}

func isWaitForVpnGatewayDeleted(vpnGateway *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayDeleting},
		Target:     []string{isVPNGatewayDeleted, ""},
		Refresh:    isVpnGatewayDeleteRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpnGatewayDeleteRefreshFunc(vpnGateway *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
			ID: &id,
		}
		vpngw, response, err := vpnGateway.GetVpnGateway(getVpnGatewayOptions)
		if err != nil && response.StatusCode != 404 {
			return vpngw, "", fmt.Errorf("Error Getting Vpn Gateway: %s\n%s", err, response)
		}
		if response.StatusCode == 404 {
			return vpngw, isVPNGatewayDeleted, nil
		}
		return nil, isVPNGatewayDeleting, err
	}
}

func resourceIBMISVPNGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicVpngwExists(d, meta, id)
		if err != nil {
			return false, err
		}
	} else {
		err := vpngwExists(d, meta, id)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicVpngwExists(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getVpnGatewayOptions := &vpcclassicv1.GetVpnGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting Vpn Gatewa: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func vpngwExists(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getVpnGatewayOptions := &vpcv1.GetVpnGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVpnGateway(getVpnGatewayOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting Vpn Gatewa: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}
