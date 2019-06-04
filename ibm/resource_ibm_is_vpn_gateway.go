package ibm

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isVPNGatewayName             = "name"
	isVPNGatewayResourceGroup    = "resource_group"
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

		Schema: map[string]*schema.Schema{

			isVPNGatewayName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isVPNGatewaySubnet: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func resourceIBMISVPNGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] VPNGateway create")
	name := d.Get(isVPNGatewayName).(string)
	subnetID := d.Get(isVPNGatewaySubnet).(string)
	var rg string

	if grp, ok := d.GetOk(isVPNGatewayResourceGroup); ok {
		rg = grp.(string)
	}

	VPNGatewayC := vpn.NewVpnClient(sess)

	VPNGateway, err := VPNGatewayC.Create(name, subnetID, rg)
	if err != nil {
		log.Printf("[DEBUG] VPNGateway err %s", isErrorToString(err))
		return err
	}

	_, err = isWaitForVPNGatewayAvailable(VPNGatewayC, VPNGateway.ID.String(), d)
	if err != nil {
		return err
	}

	d.SetId(VPNGateway.ID.String())
	log.Printf("[INFO] VPNGateway : %s", VPNGateway.ID.String())
	return resourceIBMISVPNGatewayRead(d, meta)
}

func resourceIBMISVPNGatewayRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	VPNGatewayC := vpn.NewVpnClient(sess)

	VPNGateway, err := VPNGatewayC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isVPNGatewayName, VPNGateway.Name)
	d.Set(isVPNGatewaySubnet, VPNGateway.Subnet.ID)
	d.Set(isVPNGatewayResourceGroup, VPNGateway.ResourceGroup.ID)
	d.Set(isVPNGatewayStatus, VPNGateway.Status)
	d.Set(isVPNGatewayPublicIPAddress, VPNGateway.PublicIP.Address)
	return nil
}

func resourceIBMISVPNGatewayUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	VPNGatewayC := vpn.NewVpnClient(sess)

	if d.HasChange(isVPNGatewayName) {
		name := d.Get(isVPNGatewayName).(string)
		_, err := VPNGatewayC.Update(d.Id(), name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVPNGatewayRead(d, meta)
}

func resourceIBMISVPNGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	VPNGatewayC := vpn.NewVpnClient(sess)
	err = VPNGatewayC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForVPNGatewayDeleted(VPNGatewayC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForVPNGatewayAvailable(VPNGateway *vpn.VpnClient, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for vpn gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayProvisioning},
		Target:     []string{isVPNGatewayProvisioningDone},
		Refresh:    isVPNGatewayRefreshFunc(VPNGateway, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPNGatewayRefreshFunc(VPNGateway *vpn.VpnClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		instance, err := VPNGateway.Get(id)
		if err != nil {
			return nil, "", err
		}

		if instance.Status == "available" || instance.Status == "failed" || instance.Status == "running" {
			return instance, isInstanceProvisioningDone, nil
		}

		return instance, isInstanceProvisioning, nil
	}
}

func isWaitForVPNGatewayDeleted(VPNGateway *vpn.VpnClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayDeleting},
		Target:     []string{},
		Refresh:    isVPNGatewayDeleteRefreshFunc(VPNGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPNGatewayDeleteRefreshFunc(VPNGateway *vpn.VpnClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		VPNGateway, err := VPNGateway.Get(id)
		if err == nil {
			return VPNGateway, isVPNGatewayDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "vpn_gateway_not_found" {
				return nil, isVPNGatewayDeleted, nil
			}
		}
		return nil, isVPNGatewayDeleting, err
	}
}

func resourceIBMISVPNGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	VPNGatewayC := vpn.NewVpnClient(sess)

	_, err = VPNGatewayC.Get(d.Id())
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
