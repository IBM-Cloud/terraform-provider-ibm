package ibm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isInstanceNICInstanceID         = "instance_id"
	isInstanceNICName               = "name"
	isInstanceNICStatus             = "status"
	isInstanceNICPortSpeed          = "port_speed"
	isInstanceNICPrimaryIPV4Address = "primary_ipv4_address"
	isInstanceNICPrimaryIPV6Address = "primary_ipv6_address"
	isInstanceNICSecondaryAddresses = "secondary_addresses"
	isInstanceNICSecurityGroups     = "security_groups"
	isInstanceNICSubnet             = "subnet"
)

func resourceIBMISInstanceNIC() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceNICCreate,
		Read:     resourceIBMISInstanceNICRead,
		Update:   resourceIBMISInstanceNICUpdate,
		Delete:   resourceIBMISInstanceNICDelete,
		Exists:   resourceIBMISInstanceNICExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isInstanceNICName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			isInstanceNICPortSpeed: {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			isInstanceNICSubnet: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			isInstanceNICInstanceID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isInstanceNICStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isInstanceNICPrimaryIPV4Address: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			isInstanceNICPrimaryIPV6Address: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			isInstanceNICSecondaryAddresses: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			isInstanceNICSecurityGroups: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMISInstanceNICRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instanceid, interfaceid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	nic, err := instanceC.GetInterface(instanceid, interfaceid)
	if err != nil {
		return err
	}

	d.Set(isInstanceNICInstanceID, instanceid)
	d.Set(isInstanceNICName, nic.Name)
	d.Set(isInstanceNICStatus, nic.Status)
	d.Set(isInstanceNICPortSpeed, nic.PortSpeed)
	d.Set(isInstanceNICSubnet, nic.Subnet.ID.String())
	d.Set(isInstanceNICPrimaryIPV4Address, nic.PrimaryIPV4Address)
	d.Set(isInstanceNICPrimaryIPV6Address, nic.PrimaryIPV6Address)
	d.Set(isInstanceNICSecondaryAddresses, nic.SecondaryAddresses)

	var secGroups []string
	secGroups = make([]string, len(nic.SecurityGroups), len(nic.SecurityGroups))
	if nic.SecurityGroups != nil {
		for i := 0; i < len(nic.SecurityGroups); i++ {
			if nic.SecurityGroups[i] != nil {
				secGroups[i] = nic.SecurityGroups[i].ID.String()
			}
		}
	}
	d.Set(isInstanceNICSecurityGroups, secGroups)

	return nil
}

func resourceIBMISInstanceNICDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instanceid, interfaceid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	err = instanceC.DeleteInterface(instanceid, interfaceid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISInstanceNICCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instanceid := d.Get(isInstanceNICInstanceID).(string)
	name := d.Get(isInstanceNICName).(string)
	portspeed := d.Get(isInstanceNICPortSpeed).(int)
	subnetid := d.Get(isInstanceNICSubnet).(string)
	v4address := d.Get(isInstanceNICPrimaryIPV4Address).(string)
	v6address := d.Get(isInstanceNICPrimaryIPV6Address).(string)
	secaddrs := expandStringList((d.Get(isInstanceNICSecondaryAddresses).(*schema.Set)).List())
	secgrps := expandStringList((d.Get(isInstanceNICSecurityGroups).(*schema.Set)).List())
	tags := []string{}

	nic, err := instanceC.AddInterface(instanceid, name, subnetid, portspeed,
		v4address, v6address, secaddrs, secgrps, tags)
	if err != nil {
		log.Printf("[DEBUG] instance NIC err %s", isErrorToString(err))
		return err
	}

	tid := makeTerraformRuleID(instanceid, nic.ID.String())

	d.SetId(tid)

	return resourceIBMISInstanceNICRead(d, meta)
}

func resourceIBMISInstanceNICUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)

	instanceid, interfaceid, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	name := ""
	portspeed := 0

	if d.HasChange(isInstanceNICName) {
		name = d.Get(isInstanceNICName).(string)
	}

	if d.HasChange(isInstanceNICPortSpeed) {
		portspeed = d.Get(isInstanceNICPortSpeed).(int)
	}

	_, err = instanceC.UpdateInterface(instanceid, interfaceid, name, portspeed)
	if err != nil {
		return err
	}

	return resourceIBMISInstanceNICRead(d, meta)

}

func resourceIBMISInstanceNICExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	instanceC := compute.NewInstanceClient(sess)

	instanceid, interfaceid, err := parseISTerraformID(d.Id())
	if err != nil {
		return false, err
	}

	_, err = instanceC.GetInterface(instanceid, interfaceid)
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
