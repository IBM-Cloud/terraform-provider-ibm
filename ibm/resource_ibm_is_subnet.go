package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isSubnetIpv4CidrBlock             = "ipv4_cidr_block"
	isSubnetIpv6CidrBlock             = "ipv6_cidr_block"
	isSubnetTotalIpv4AddressCount     = "total_ipv4_address_count"
	isSubnetIPVersion                 = "ip_version"
	isSubnetName                      = "name"
	isSubnetNetworkACL                = "network_acl"
	isSubnetPublicGateway             = "public_gateway"
	isSubnetStatus                    = "status"
	isSubnetVPC                       = "vpc"
	isSubnetZone                      = "zone"
	isSubnetAvailableIpv4AddressCount = "available_ipv4_address_count"

	isSubnetProvisioning     = "provisioning"
	isSubnetProvisioningDone = "done"
	isSubnetDeleting         = "deleting"
	isSubnetDeleted          = "done"
)

func resourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSubnetCreate,
		Read:     resourceIBMISSubnetRead,
		Update:   resourceIBMISSubnetUpdate,
		Delete:   resourceIBMISSubnetDelete,
		Exists:   resourceIBMISSubnetExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isSubnetIpv4CidrBlock: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetTotalIpv4AddressCount},
				ValidateFunc:  validateCIDR,
			},

			isSubnetIpv6CidrBlock: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetAvailableIpv4AddressCount: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetTotalIpv4AddressCount: {
				Type:          schema.TypeInt,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetIpv4CidrBlock},
			},

			isSubnetIPVersion: {
				Type:         schema.TypeString,
				ForceNew:     true,
				Default:      "ipv4",
				Optional:     true,
				ValidateFunc: validateIPVersion,
			},

			isSubnetName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isSubnetNetworkACL: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},

			isSubnetPublicGateway: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			isSubnetStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetVPC: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isSubnetZone: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceIBMISSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	name := d.Get(isSubnetName).(string)
	vpc := d.Get(isSubnetVPC).(string)
	zone := d.Get(isSubnetZone).(string)

	ipv4cidr := d.Get(isSubnetIpv4CidrBlock).(string)
	ipv4addrcount := d.Get(isSubnetTotalIpv4AddressCount).(int)
	if ipv4cidr == "" && ipv4addrcount == 0 {
		return fmt.Errorf("%s or %s need to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
	}

	if ipv4cidr != "" && ipv4addrcount != 0 {
		return fmt.Errorf("only one of %s or %s needs to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
	}

	acl := d.Get(isSubnetNetworkACL).(string)
	gw := d.Get(isSubnetPublicGateway).(string)

	subnetC := network.NewSubnetClient(sess)
	subnet, err := subnetC.Create(name, zone, vpc, acl, gw, "", "", ipv4cidr, ipv4addrcount)
	if err != nil {
		return err
	}

	d.SetId(subnet.ID.String())
	log.Printf("[INFO] Subnet : %s", subnet.ID.String())

	_, err = isWaitForSubnetAvailable(subnetC, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMISSubnetRead(d, meta)
}

func isWaitForSubnetAvailable(subnetC *network.SubnetClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetProvisioning},
		Target:     []string{isSubnetProvisioningDone},
		Refresh:    isSubnetRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetRefreshFunc(subnetC *network.SubnetClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		subnet, err := subnetC.Get(id)
		if err != nil {
			return nil, "", err
		}

		if subnet.Status == "available" || subnet.Status == "failed" {
			return subnet, isSubnetProvisioningDone, nil
		}

		return subnet, isSubnetProvisioning, nil
	}
}

func resourceIBMISSubnetRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	subnetC := network.NewSubnetClient(sess)

	subnet, err := subnetC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set("id", subnet.ID.String())
	d.Set(isSubnetName, subnet.Name)
	d.Set(isSubnetIPVersion, subnet.IPVersion)
	d.Set(isSubnetIpv4CidrBlock, subnet.IPV4CidrBlock)
	d.Set(isSubnetIpv6CidrBlock, subnet.IPV6CidrBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, subnet.AvailableIPV4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, subnet.TotalIPV4AddressCount)
	if subnet.NetworkACL != nil {
		d.Set(isSubnetNetworkACL, subnet.NetworkACL.ID.String())
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, subnet.PublicGateway.ID.String())
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, subnet.Status)
	d.Set(isSubnetZone, subnet.Zone.Name)
	d.Set(isSubnetVPC, subnet.Vpc.ID.String())

	return nil
}

func resourceIBMISSubnetUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	subnetC := network.NewSubnetClient(sess)

	name := ""
	acl := ""
	gw := ""
	if d.HasChange(isSubnetName) {
		name = d.Get(isSubnetName).(string)
	}
	if d.HasChange(isSubnetNetworkACL) {
		acl = d.Get(isSubnetNetworkACL).(string)
	}
	if d.HasChange(isSubnetPublicGateway) {
		gw = d.Get(isSubnetPublicGateway).(string)
		if gw == "" {
			err = subnetC.DetachPublicGateway(d.Id())
			if err != nil {
				return err
			}
			_, err = isWaitForSubnetAvailable(subnetC, d.Id(), d.Timeout(schema.TimeoutDelete))
			if err != nil {
				return err
			}

		}
	}

	_, err = subnetC.Update(d.Id(), name, acl, gw)
	if err != nil {
		return err
	}

	_, err = isWaitForSubnetAvailable(subnetC, d.Id(), d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	return resourceIBMISSubnetRead(d, meta)
}

func resourceIBMISSubnetDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	subnetC := network.NewSubnetClient(sess)

	subnet, err := subnetC.Get(d.Id())
	if err != nil {
		return err
	}

	if subnet.PublicGateway != nil {
		err = subnetC.DetachPublicGateway(d.Id())
		if err != nil {
			return err
		}
		_, err = isWaitForSubnetAvailable(subnetC, d.Id(), d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return err
		}

	}

	err = subnetC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForSubnetDeleted(subnetC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForSubnetDeleted(subnetC *network.SubnetClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetDeleting},
		Target:     []string{},
		Refresh:    isSubnetDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetDeleteRefreshFunc(subnetC *network.SubnetClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		subnet, err := subnetC.Get(id)
		if err == nil {
			return subnet, isSubnetDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("[DEBUG] returning deleted")
				return nil, isSubnetDeleted, nil
			}
		}
		log.Printf("[DEBUG] returning x")
		return nil, isSubnetDeleting, err
	}
}

func resourceIBMISSubnetExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	subnetC := network.NewSubnetClient(sess)

	_, err = subnetC.Get(d.Id())
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
