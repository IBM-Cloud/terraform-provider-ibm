package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
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
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)

	prefixName := d.Get(isVPCAddressPrefixPrefixName).(string)
	zoneName := d.Get(isVPCAddressPrefixZoneName).(string)
	cidr := d.Get(isVPCAddressPrefixCIDR).(string)
	vpcID := d.Get(isVPCAddressPrefixVPCID).(string)

	params := &models.PostVpcsVpcIDAddressPrefixesParamsBody{
		Cidr: cidr,
		Name: prefixName,
		Zone: &models.PostVpcsVpcIDAddressPrefixesParamsBodyZone{
			Name: zoneName,
		},
	}

	addrPrefix, err := vpcClient.CreateAddressPrefix(params, vpcID)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", vpcID, addrPrefix.ID))
	return resourceIBMISVpcAddressPrefixRead(d, meta)

}

func resourceIBMISVpcAddressPrefixRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	addrPrefixID := parts[1]
	addrPrefix, err := vpcClient.GetAddressPrefix(vpcID, addrPrefixID)
	if err != nil {
		return err
	}

	d.Set(isVPCAddressPrefixPrefixName, addrPrefix.Name)
	if addrPrefix.Zone != nil {
		d.Set(isVPCAddressPrefixZoneName, addrPrefix.Zone.Name)
	}

	d.Set(isVPCAddressPrefixCIDR, addrPrefix.Cidr)
	d.Set(isVPCAddressPrefixVPCID, vpcID)
	d.Set(isVPCAddressPrefixHasSubnets, addrPrefix.HasSubnets)

	return nil
}

func resourceIBMISVpcAddressPrefixUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)
	hasChanged := false
	params := &models.PatchVpcsVpcIDAddressPrefixesIDParamsBody{}
	if d.HasChange(isVPCAddressPrefixCIDR) {
		params.Cidr = d.Get(isVPCAddressPrefixCIDR).(string)
		hasChanged = true
	}

	if d.HasChange(isVPCAddressPrefixPrefixName) {
		params.Name = d.Get(isVPCAddressPrefixPrefixName).(string)
		hasChanged = true
	}

	if hasChanged {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}

		vpcID := parts[0]
		addrPrefixID := parts[1]
		_, err = vpcC.UpdateAddressPrefix(params, vpcID, addrPrefixID)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVpcAddressPrefixRead(d, meta)
}

func resourceIBMISVpcAddressPrefixDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	addrPrefixID := parts[1]
	err = vpcClient.DeleteAddressPrefix(vpcID, addrPrefixID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISVpcAddressPrefixExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	vpcID := parts[0]
	addrPrefixID := parts[1]
	_, err = vpcClient.GetAddressPrefix(vpcID, addrPrefixID)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
