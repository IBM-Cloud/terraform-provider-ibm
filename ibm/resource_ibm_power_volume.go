package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	PowerVolumeName      = "name"
	PowerVolumeSize      = "size"
	PowerVolumeType      = "type"
	PowerVolumeShareable = "shareable"
	PowerVolumeId        = "volumeid"

	PowerVolumeStatus           = "status"
	PowerVolumeDeleting         = "deleting"
	PowerVolumeDeleted          = "done"
	PowerVolumeProvisioning     = "creating"
	PowerVolumeProvisioningDone = "available"
)

func resourceIBMPowerVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPowerVolumeCreate,
		Read:   resourceIBMPowerVolumeRead,
		Update: resourceIBMPowerVolumeUpdate,
		Delete: resourceIBMPowerVolumeDelete,
		//Exists:   resourceIBMPowerVolumeExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PowerVolumeId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			PowerVolumeName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PowerVolumeShareable: {
				Type:     schema.TypeBool,
				Required: true,
			},
			PowerVolumeSize: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			PowerVolumeType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"ssd", "shared"}),
			},

			PowerVolumeStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIBMPowerVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	name := d.Get(PowerVolumeName).(string)
	volType := d.Get(PowerVolumeType).(string)
	size := float64(d.Get(PowerVolumeSize).(float64))
	shared := d.Get(PowerVolumeShareable).(bool)

	client := st.NewPowerVolumeClient(sess)

	vol, err := client.Create(name, size, volType, shared)

	if err != nil {
		return err
	}

	volumeid := *vol.VolumeID
	d.SetId(volumeid)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForPowerVolumeAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil
	return resourceIBMPowerVolumeRead(d, meta)
}

func resourceIBMPowerVolumeRead(d *schema.ResourceData, meta interface{}) error {
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

func resourceIBMPowerVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the IBM Power Volume update call")
	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerVolumeClient(sess)

	//name := ""
	//if d.HasChange(PowerVolumeName) {
	name := d.Get(PowerVolumeName).(string)
	//}

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

func resourceIBMPowerVolumeDelete(d *schema.ResourceData, meta interface{}) error {

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
func isWaitForPowerVolumeAvailable(client *st.PowerVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", PowerVolumeProvisioning},
		Target:     []string{PowerVolumeProvisioningDone},
		Refresh:    isPowerVolumeRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPowerVolumeRefreshFunc(client *st.PowerVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "available" {
			return vol, PowerVolumeProvisioningDone, nil
		}

		return vol, PowerVolumeProvisioning, nil
	}
}
