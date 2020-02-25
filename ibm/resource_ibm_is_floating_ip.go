package ibm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isFloatingIPAddress       = "address"
	isFloatingIPName          = "name"
	isFloatingIPStatus        = "status"
	isFloatingIPZone          = "zone"
	isFloatingIPTarget        = "target"
	isFloatingIPResourceGroup = "resource_group"
	isFloatingIPTags          = "tags"

	isFloatingIPPending   = "pending"
	isFloatingIPAvailable = "available"
	isFloatingIPDeleting  = "deleting"
	isFloatingIPDeleted   = "done"
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
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isFloatingIPAddress: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isFloatingIPName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
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

			isFloatingIPResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isFloatingIPTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      resourceIBMVPCHash,
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
	var rg string
	rg = ""
	if grp, ok := d.GetOk(isFloatingIPResourceGroup); ok {
		rg = grp.(string)
	}
	floatingipC := network.NewFloatingIPClient(sess)
	floatingip, err := floatingipC.Create(name, zone, rg, target)
	if err != nil {
		log.Printf("[DEBUG] floating ip err %s", isErrorToString(err))
		return err
	}

	d.SetId(floatingip.ID.String())
	log.Printf("[INFO] Floating IP : %s[%s]", floatingip.ID.String(), floatingip.Address)
	_, err = isWaitForInstanceFloatingIP(floatingipC, d.Id(), d)
	if err != nil {
		return err
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isFloatingIPTags); ok || v != "" {
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, floatingip.Crn)
		if err != nil {
			log.Printf(
				"Error on create of vpc Floating IP (%s) tags: %s", d.Id(), err)
		}
	}

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
	tags, err := GetTagsUsingCRN(meta, floatingip.Crn)
	if err != nil {
		log.Printf(
			"Error on get of vpc Floating IP (%s) tags: %s", d.Id(), err)
	}
	d.Set(isFloatingIPTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/network/floatingIPs")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/network/floatingIPs")
	}
	d.Set(ResourceName, floatingip.Name)
	d.Set(ResourceCRN, floatingip.Crn)
	d.Set(ResourceStatus, floatingip.Status)
	if floatingip.ResourceGroup != nil {
		d.Set(ResourceGroupName, floatingip.ResourceGroup.Name)
		d.Set(isFloatingIPResourceGroup, floatingip.ResourceGroup.ID)

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

	if d.HasChange(isFloatingIPTags) {
		fip, err := floatingipC.Get(d.Id())
		if err != nil {
			return err
		}
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, fip.Crn)
		if err != nil {
			log.Printf(
				"Error on update of vpc Floating IP (%s) tags: %s", d.Id(), err)
		}
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
		Pending:    []string{isFloatingIPPending, isFloatingIPDeleting},
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
			return FloatingIP, FloatingIP.Status, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil, isFloatingIPDeleted, nil
			}
		}
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

func isWaitForInstanceFloatingIP(floatingipC *network.FloatingIPClient, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for floating IP (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isFloatingIPPending},
		Target:     []string{isFloatingIPAvailable},
		Refresh:    isInstanceFloatingIPRefreshFunc(floatingipC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceFloatingIPRefreshFunc(floatingipC *network.FloatingIPClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		instance, err := floatingipC.Get(id)
		if err != nil {
			return nil, "", err
		}

		if instance.Status == "available" {
			return instance, isFloatingIPAvailable, nil
		}

		return instance, isFloatingIPPending, nil
	}
}
