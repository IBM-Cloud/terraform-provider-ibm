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
	isFloatingIPAddress = "address"
	isFloatingIPName    = "name"
	isFloatingIPStatus  = "status"
	isFloatingIPZone    = "zone"
	isFloatingIPTarget  = "target"

	isFloatingIPProvisioning     = "provisioning"
	isFloatingIPProvisioningDone = "done"
	isFloatingIPDeleting         = "deleting"
	isFloatingIPDeleted          = "done"
)

func resourceIBMISFloatingIP() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISFloatingIPCreate,
		Read:     resourceIBMISFloatingIPRead,
		Update:   resourceIBMISFloatingIPUpdate,
		Delete:   resourceIBMISFloatingIPDelete,
		Exists:   resourceIBMISFloatingIPExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isFloatingIPAddress: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isFloatingIPName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isFloatingIPStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isFloatingIPZone: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPTarget},
			},

			isFloatingIPTarget: {
				Type:          schema.TypeString,
				ForceNew:      false,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPZone},
			},
		},
	}
}

func resourceIBMISFloatingIPCreate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	name := d.Get(isFloatingIPName).(string)
	var zone, target string

	if zn, ok := d.GetOk(isFloatingIPZone); ok {
		zone = zn.(string)
	}

	if tgt, ok := d.GetOk(isFloatingIPTarget); ok {
		target = tgt.(string)
	}

	if zone == "" && target == "" {
		return fmt.Errorf("%s or %s need to be provided", isFloatingIPZone, isFloatingIPTarget)
	}

	floatingipC := network.NewFloatingIPClient(sess)
	floatingip, err := floatingipC.Create(name, zone, "", target)
	if err != nil {
		log.Printf("[DEBUG] floating ip err %s", isErrorToString(err))
		return err
	}

	d.SetId(floatingip.ID.String())
	log.Printf("[INFO] Floating IP : %s[%s]", floatingip.ID.String(), floatingip.Address)

	return resourceIBMISFloatingIPRead(d, meta)
}

func resourceIBMISFloatingIPRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	floatingipC := network.NewFloatingIPClient(sess)

	floatingip, err := floatingipC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isFloatingIPName, floatingip.Name)
	d.Set(isFloatingIPAddress, floatingip.Address)
	d.Set(isFloatingIPStatus, floatingip.Status)
	d.Set(isFloatingIPZone, floatingip.Zone.Name)
	if floatingip.Target != nil && &floatingip.Target.ID != nil {
		d.Set(isFloatingIPTarget, floatingip.Target.ID.String())
	}

	return nil
}

func resourceIBMISFloatingIPUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	floatingipC := network.NewFloatingIPClient(sess)

	name := ""
	target := ""
	hasChange := false

	if d.HasChange(isFloatingIPName) {
		name = d.Get(isFloatingIPName).(string)
		hasChange = true
	}

	if d.HasChange(isFloatingIPTarget) {
		target = d.Get(isFloatingIPTarget).(string)
		hasChange = true
	}

	if hasChange {

		_, err := floatingipC.Update(d.Id(), name, target)
		if err != nil {
			return err
		}

	}

	return resourceIBMISFloatingIPRead(d, meta)
}

func resourceIBMISFloatingIPDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	floatingipC := network.NewFloatingIPClient(sess)

	err = floatingipC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForFloatingIPDeleted(floatingipC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForFloatingIPDeleted(fip *network.FloatingIPClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for FloatingIP (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isFloatingIPDeleting},
		Target:     []string{},
		Refresh:    isFloatingIPDeleteRefreshFunc(fip, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isFloatingIPDeleteRefreshFunc(fip *network.FloatingIPClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		FloatingIP, err := fip.Get(id)
		if err == nil {
			return FloatingIP, isFloatingIPDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("[DEBUG] returning deleted")
				return nil, isFloatingIPDeleted, nil
			}
		}
		log.Printf("[DEBUG] returning x")
		return nil, isFloatingIPDeleting, err
	}
}

func resourceIBMISFloatingIPExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	floatingipC := network.NewFloatingIPClient(sess)

	_, err = floatingipC.Get(d.Id())
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
