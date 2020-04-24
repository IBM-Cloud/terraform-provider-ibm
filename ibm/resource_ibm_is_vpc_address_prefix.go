package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isVPCAddressPrefixPrefixName = "name"
	isVPCAddressPrefixZoneName   = "zone"
	isVPCAddressPrefixCIDR       = "cidr"
	isVPCAddressPrefixVPCID      = "vpc"
	isVPCAddressPrefixHasSubnets = "has_subnets"
)

func resourceIBMISVpcAddressPrefix() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVpcAddressPrefixCreate,
		Read:     resourceIBMISVpcAddressPrefixRead,
		Update:   resourceIBMISVpcAddressPrefixUpdate,
		Delete:   resourceIBMISVpcAddressPrefixDelete,
		Exists:   resourceIBMISVpcAddressPrefixExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isVPCAddressPrefixPrefixName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
			},
			isVPCAddressPrefixZoneName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPCAddressPrefixCIDR: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isVPCAddressPrefixVPCID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPCAddressPrefixHasSubnets: {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceIBMISVpcAddressPrefixCreate(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	prefixName := d.Get(isVPCAddressPrefixPrefixName).(string)
	zoneName := d.Get(isVPCAddressPrefixZoneName).(string)
	cidr := d.Get(isVPCAddressPrefixCIDR).(string)
	vpcID := d.Get(isVPCAddressPrefixVPCID).(string)

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	ibmMutexKV.Lock(isVPCAddressPrefixKey)
	defer ibmMutexKV.Unlock(isVPCAddressPrefixKey)

	if userDetails.generation == 1 {
		err := classicVpcAddressPrefixCreate(d, meta, prefixName, zoneName, cidr, vpcID)
		if err != nil {
			return err
		}
	} else {
		err := vpcAddressPrefixCreate(d, meta, prefixName, zoneName, cidr, vpcID)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVpcAddressPrefixRead(d, meta)
}

func classicVpcAddressPrefixCreate(d *schema.ResourceData, meta interface{}, name, zone, cidr, vpcID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateVpcAddressPrefixOptions{
		Name:  &name,
		VpcID: &vpcID,
		Cidr:  &cidr,
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: &zone,
		},
	}
	addrPrefix, response, err := sess.CreateVpcAddressPrefix(options)
	if err != nil {
		return fmt.Errorf("Error while creating VPC Address Prefix %s\n%s", err, response)
	}

	addrPrefixID := *addrPrefix.ID

	d.SetId(fmt.Sprintf("%s/%s", vpcID, addrPrefixID))
	return nil
}

func vpcAddressPrefixCreate(d *schema.ResourceData, meta interface{}, name, zone, cidr, vpcID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateVpcAddressPrefixOptions{
		Name:  &name,
		VpcID: &vpcID,
		Cidr:  &cidr,
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	}
	addrPrefix, response, err := sess.CreateVpcAddressPrefix(options)
	if err != nil {
		return fmt.Errorf("Error while creating VPC Address Prefix %s\n%s", err, response)
	}

	addrPrefixID := *addrPrefix.ID
	d.SetId(fmt.Sprintf("%s/%s", vpcID, addrPrefixID))
	return nil
}

func resourceIBMISVpcAddressPrefixRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	addrPrefixID := parts[1]
	if userDetails.generation == 1 {
		err := classicVpcAddressPrefixGet(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return err
		}
	} else {
		err := vpcAddressPrefixGet(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return err
		}
	}

	return nil
}

func classicVpcAddressPrefixGet(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getvpcAddressPrefixOptions := &vpcclassicv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	addrPrefix, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isVPCAddressPrefixPrefixName, *addrPrefix.Name)
	if addrPrefix.Zone != nil {
		d.Set(isVPCAddressPrefixZoneName, *addrPrefix.Zone.Name)
	}
	d.Set(isVPCAddressPrefixCIDR, *addrPrefix.Cidr)
	d.Set(isVPCAddressPrefixHasSubnets, *addrPrefix.HasSubnets)

	return nil
}

func vpcAddressPrefixGet(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getvpcAddressPrefixOptions := &vpcv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	addrPrefix, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isVPCAddressPrefixPrefixName, *addrPrefix.Name)
	if addrPrefix.Zone != nil {
		d.Set(isVPCAddressPrefixZoneName, *addrPrefix.Zone.Name)
	}
	d.Set(isVPCAddressPrefixCIDR, *addrPrefix.Cidr)
	d.Set(isVPCAddressPrefixHasSubnets, *addrPrefix.HasSubnets)

	return nil
}

func resourceIBMISVpcAddressPrefixUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := ""
	hasChanged := false

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	ibmMutexKV.Lock(isVPCAddressPrefixKey)
	defer ibmMutexKV.Unlock(isVPCAddressPrefixKey)

	if d.HasChange(isVPCAddressPrefixPrefixName) {
		name = d.Get(isVPCAddressPrefixPrefixName).(string)
		hasChanged = true
	}

	if userDetails.generation == 1 {
		err := classicVpcAddressPrefixUpdate(d, meta, vpcID, addrPrefixID, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := vpcAddressPrefixUpdate(d, meta, vpcID, addrPrefixID, name, hasChanged)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVpcAddressPrefixRead(d, meta)
}

func classicVpcAddressPrefixUpdate(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if hasChanged {
		updatevpcAddressPrefixoptions := &vpcclassicv1.UpdateVpcAddressPrefixOptions{
			VpcID: &vpcID,
			ID:    &addrPrefixID,
			Name:  &name,
		}
		_, response, err := sess.UpdateVpcAddressPrefix(updatevpcAddressPrefixoptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC Address Prefix: %s\n%s", err, response)
		}
	}
	return nil
}

func vpcAddressPrefixUpdate(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if hasChanged {
		updatevpcAddressPrefixoptions := &vpcv1.UpdateVpcAddressPrefixOptions{
			VpcID: &vpcID,
			ID:    &addrPrefixID,
			Name:  &name,
		}
		_, response, err := sess.UpdateVpcAddressPrefix(updatevpcAddressPrefixoptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC Address Prefix: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISVpcAddressPrefixDelete(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	isVPCAddressPrefixKey := "vpc_address_prefix_key_" + vpcID
	ibmMutexKV.Lock(isVPCAddressPrefixKey)
	defer ibmMutexKV.Unlock(isVPCAddressPrefixKey)

	if userDetails.generation == 1 {
		err := classicVpcAddressPrefixDelete(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return err
		}
	} else {
		err := vpcAddressPrefixDelete(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil
}

func classicVpcAddressPrefixDelete(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getvpcAddressPrefixOptions := &vpcclassicv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	deletevpcAddressPrefixOptions := &vpcclassicv1.DeleteVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	response, err = sess.DeleteVpcAddressPrefix(deletevpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Deleting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	d.SetId("")
	return nil
}

func vpcAddressPrefixDelete(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getvpcAddressPrefixOptions := &vpcv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	deletevpcAddressPrefixOptions := &vpcv1.DeleteVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	response, err = sess.DeleteVpcAddressPrefix(deletevpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Deleting VPC Address Prefix (%s): %s\n%s", addrPrefixID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	d.SetId("")
	return nil
}

func resourceIBMISVpcAddressPrefixExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	vpcID := parts[0]
	addrPrefixID := parts[1]

	if userDetails.generation == 1 {
		err := classicVpcAddressPrefixExists(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return false, err
		}
	} else {
		err := vpcAddressPrefixExists(d, meta, vpcID, addrPrefixID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicVpcAddressPrefixExists(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getvpcAddressPrefixOptions := &vpcclassicv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting VPC Address Prefix: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func vpcAddressPrefixExists(d *schema.ResourceData, meta interface{}, vpcID, addrPrefixID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getvpcAddressPrefixOptions := &vpcv1.GetVpcAddressPrefixOptions{
		VpcID: &vpcID,
		ID:    &addrPrefixID,
	}
	_, response, err := sess.GetVpcAddressPrefix(getvpcAddressPrefixOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting VPC Address Prefix: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}
