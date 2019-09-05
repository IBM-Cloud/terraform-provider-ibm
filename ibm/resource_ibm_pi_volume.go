package ibm

import (
	_ "github.com/IBM-Cloud/bluemix-go/bmxerror"
	_ "github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	IBMPIVolumeName      = "name"
	IBMPIVolumeSize      = "size"
	IBMPIVolumeType      = "type"
	IBMPIVolumeShareable = "shareable"
	IBMPIVolumeId        = "volume_id"
	IBMPIInstanceId      = "power_instance_id"

	IBMPIVolumeStatus           = "status"
	IBMPIVolumeDeleting         = "deleting"
	IBMPIVolumeDeleted          = "done"
	IBMPIVolumeProvisioning     = "creating"
	IBMPIVolumeProvisioningDone = "available"
)

func resourceIBMPIVolume() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIVolumeCreate,
		Read:     resourceIBMPIVolumeRead,
		Update:   resourceIBMPIVolumeUpdate,
		Delete:   resourceIBMPIVolumeDelete,
		Exists:   resourceIBMPIVolumeExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			IBMPIVolumeId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			IBMPIVolumeName: {
				Type:     schema.TypeString,
				Required: true,
			},

			IBMPIVolumeShareable: {
				Type:     schema.TypeBool,
				Required: true,
			},
			IBMPIVolumeSize: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			IBMPIVolumeType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"ssd", "shared"}),
			},

			IBMPIInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed Attributes

			IBMPIVolumeStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIBMPIVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	name := d.Get(IBMPIVolumeName).(string)
	volType := d.Get(IBMPIVolumeType).(string)
	size := float64(d.Get(IBMPIVolumeSize).(float64))
	shared := d.Get(IBMPIVolumeShareable).(bool)
	powerinstanceid := d.Get(IBMPIInstanceId).(string)

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Create(name, size, volType, shared, powerinstanceid)

	if err != nil {
		return err
	}

	volumeid := *vol.VolumeID
	d.SetId(volumeid)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForIBMPIVolumeAvailable(client, d.Id(), powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil
	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPIVolumeRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	//d.SetId(vol.ID.String())
	d.Set(IBMPIVolumeName, vol.Name)
	d.Set(IBMPIVolumeSize, vol.Size)
	d.Set(IBMPIVolumeShareable, vol.Shareable)
	return nil
}

func resourceIBMPIVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the IBM Power Volume update call")
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	//name := ""
	//if d.HasChange(IBMPIVolumeName) {
	name := d.Get(IBMPIVolumeName).(string)
	//}

	size := float64(d.Get(IBMPIVolumeSize).(float64))
	shareable := bool(d.Get(IBMPIVolumeShareable).(bool))

	volrequest, err := client.Update(d.Id(), name, size, shareable, powerinstanceid)
	if err != nil {
		return err
	}

	_, err = isWaitForIBMPIVolumeAvailable(client, *volrequest.VolumeID, powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPIVolumeDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)
	err := client.Delete(d.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
func resourceIBMPIVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return vol.VolumeID == &id, nil
}

/*
func isWaitForIBMPIVolumeDeleted(vol *st.IBMPIVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeDeleting},
		Target:     []string{},
		Refresh:    isIBMPIVolumeDeleteRefreshFunc(vol, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}*/

/*func isIBMPIVolumeDeleteRefreshFunc(vol *st.IBMPIVolumeClient, id string) resource.StateRefreshFunc {
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

func resourceIBMIBMPIVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewIBMPIVolumeClient(sess)

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
func isWaitForIBMPIVolumeAvailable(client *st.IBMPIVolumeClient, id, powerinstanceid string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", IBMPIVolumeProvisioning},
		Target:     []string{IBMPIVolumeProvisioningDone},
		Refresh:    isIBMPIVolumeRefreshFunc(client, id, powerinstanceid),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isIBMPIVolumeRefreshFunc(client *st.IBMPIVolumeClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "available" {
			return vol, IBMPIVolumeProvisioningDone, nil
		}

		return vol, IBMPIVolumeProvisioning, nil
	}
}

/*

func resourceIBMOrgExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return false, err
	}
	id := d.Id()
	org, err := cfClient.Organizations().Get(id)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return org.Metadata.GUID == id, nil
}
*/
