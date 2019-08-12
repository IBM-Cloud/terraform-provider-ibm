package ibm

import (
	"errors"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	PowerVolumeAttachName             = "volumename"
	PowerInstanceName                 = "servername"
	PowerVolumeAllowableAttachStatus  = "in-use"
	PowerVolumeAttachStatus           = "status"
	PowerVolumeAttachDeleting         = "deleting"
	PowerVolumeAttachProvisioning     = "creating"
	PowerVolumeAttachProvisioningDone = "available"
)

func resourceIBMPowerVolumeAttach() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPowerVolumeAttachCreate,
		Read:   resourceIBMPowerVolumeAttachRead,
		Update: resourceIBMPowerVolumeAttachUpdate,
		Delete: resourceIBMPowerVolumeAttachDelete,
		//Exists:   resourceIBMPowerVolumeExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"volumeattachid": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			PowerVolumeAttachName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PowerInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PowerVolumeStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			PowerVolumeShareable: {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIBMPowerVolumeAttachCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	name := d.Get(PowerVolumeAttachName).(string)

	servername := d.Get(PowerInstanceName).(string)

	client := st.NewPowerVolumeClient(sess)
	log.Print("Now doing a get with the volumename %s  ", name)
	volinfo, err := client.Get(name)

	if err != nil {
		return errors.New("The volume cannot be attached since it's not available")
		log.Printf(" The volume that is being attached is not available ")
	}
	log.Print("The volume info is %s", volinfo)

	if volinfo.State == PowerVolumeAllowableAttachStatus {

		return errors.New("The volume cannot be attached in the current state. The volume must be in the *available* state. No other states are permissible")
	}

	resp, err := client.Attach(servername, name)

	if err != nil {
		return err
	}
	log.Printf("Printing the resp %+v", resp)

	d.SetId(*volinfo.VolumeID)
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForPowerVolumeAttachAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil
	return resourceIBMPowerVolumeAttachRead(d, meta)
}

func resourceIBMPowerVolumeAttachRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerVolumeClient(sess)

	vol, err := client.Get(d.Id())
	if err != nil {
		return err
	}

	//d.SetId(vol.ID.String())
	d.Set(PowerVolumeName, vol.Name)
	d.Set(PowerVolumeSize, vol.Size)
	d.Set(PowerVolumeShareable, vol.Shareable)
	return nil
}

func resourceIBMPowerVolumeAttachUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerVolumeClient(sess)

	name := ""
	if d.HasChange(PowerVolumeName) {
		name = d.Get(PowerVolumeName).(string)
	}

	size := float64(d.Get(PowerVolumeSize).(float64))
	shareable := bool(d.Get(PowerVolumeShareable).(bool))

	volrequest, err := client.Update(d.Id(), name, size, shareable)
	if err != nil {
		return err
	}

	_, err = isWaitForPowerVolumeAvailable(client, *volrequest.VolumeID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMPowerVolumeRead(d, meta)
}

func resourceIBMPowerVolumeAttachDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerVolumeClient(sess)
	err := client.Delete(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

/*
func isWaitForPowerVolumeDeleted(vol *st.PowerVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeDeleting},
		Target:     []string{},
		Refresh:    isPowerVolumeDeleteRefreshFunc(vol, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}*/

/*func isPowerVolumeDeleteRefreshFunc(vol *st.PowerVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := vol.Get(id)
		if err == nil {
			return vol, isVolumeDeleting, nil
		}

		iserror, ok := err.(iserrors.Power)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "volume_not_found" {
				return nil, isVolumeDeleted, nil
			}
		}
		return nil, isVolumeDeleting, err
	}
}

func resourceIBMPowerVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerVolumeClient(sess)

	_, err := client.Get(d.Id())
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
*/
func isWaitForPowerVolumeAttachAvailable(client *st.PowerVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", PowerVolumeAttachProvisioningDone},
		Target:     []string{PowerVolumeAllowableAttachStatus},
		Refresh:    isPowerVolumeAttachRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPowerVolumeAttachRefreshFunc(client *st.PowerVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "in-use" {
			return vol, PowerVolumeAllowableAttachStatus, nil
		}

		return vol, PowerVolumeAttachProvisioningDone, nil
	}
}
