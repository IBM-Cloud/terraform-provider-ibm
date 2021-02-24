package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	/* Power Volume creation depends on response from PowerVC */
	volPostTimeOut   = 180 * time.Second
	volGetTimeOut    = 180 * time.Second
	volDeleteTimeOut = 180 * time.Second
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
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"volume_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume ID",
			},

			helpers.PIVolumeName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Volume Name to create",
			},

			helpers.PIVolumeShareable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Flag to indicate if the volume can be shared across multiple instances?",
			},
			helpers.PIVolumeSize: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Size of the volume in GB",
			},
			helpers.PIVolumeType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"ssd", "standard", "tier1", "tier3"}),
				Description:  "Volume type",
			},

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			// Computed Attributes

			"volume_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume status",
			},

			"delete_on_termination": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Should the volume be deleted during termination",
			},
			"wwn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "WWN Of the volume",
			},
		},
	}
}

func resourceIBMPIVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	name := d.Get(helpers.PIVolumeName).(string)
	volType := d.Get(helpers.PIVolumeType).(string)
	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	shared := d.Get(helpers.PIVolumeShareable).(bool)
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Create(name, size, volType, shared, powerinstanceid, volPostTimeOut)

	if err != nil {
		return fmt.Errorf("Failed to Create the volume %v", err)
	}

	volumeid := *vol.VolumeID
	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, volumeid))
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return err
	}
	_, err = isWaitForIBMPIVolumeAvailable(client, volumeid, powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil
	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPIVolumeRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Get(parts[1], powerinstanceid, volGetTimeOut)
	if err != nil {
		return fmt.Errorf("Failed to get the volume %v", err)

	}

	d.Set(helpers.PIVolumeName, vol.Name)
	d.Set(helpers.PIVolumeSize, vol.Size)
	d.Set(helpers.PIVolumeShareable, vol.Shareable)
	d.Set(helpers.PIVolumeType, vol.DiskType)
	d.Set("volume_status", vol.State)
	d.Set("volume_id", vol.VolumeID)
	d.Set(helpers.PICloudInstanceId, powerinstanceid)
	d.Set("delete_on_termination", vol.DeleteOnTermination)
	d.Set("wwn", vol.Wwn)

	return nil
}

func resourceIBMPIVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the IBM Power Volume update call")
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	//name := ""
	//if d.HasChange(helpers.PIVolumeName) {
	name := d.Get(helpers.PIVolumeName).(string)
	//}

	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	shareable := bool(d.Get(helpers.PIVolumeShareable).(bool))

	volrequest, err := client.Update(parts[1], name, size, shareable, powerinstanceid, volPostTimeOut)
	if err != nil {
		return err
	}

	_, err = isWaitForIBMPIVolumeAvailable(client, *volrequest.VolumeID, powerinstanceid, d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPIVolumeDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Get(parts[1], powerinstanceid, volDeleteTimeOut)
	if err != nil {
		return err
	}

	log.Printf("The volume to be deleted is in the following state .. %s", vol.State)
	_, err = isWaitForIBMPIVolumeAvailable(client, parts[1], powerinstanceid, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	voldelete_err := client.Delete(parts[1], powerinstanceid, deleteTimeOut)
	if voldelete_err != nil {
		return voldelete_err
	}

	d.SetId("")
	return nil
}
func resourceIBMPIVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	powerinstanceid := parts[0]
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	vol, err := client.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	log.Printf("Calling the existing function.. %s", *(vol.VolumeID))

	volumeid := *vol.VolumeID
	return volumeid == parts[1], nil
}

func isWaitForIBMPIVolumeAvailable(client *st.IBMPIVolumeClient, id, powerinstanceid string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIVolumeProvisioning},
		Target:     []string{helpers.PIVolumeProvisioningDone},
		Refresh:    isIBMPIVolumeRefreshFunc(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isIBMPIVolumeRefreshFunc(client *st.IBMPIVolumeClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id, powerinstanceid, volGetTimeOut)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "available" {
			return vol, helpers.PIVolumeProvisioningDone, nil
		}

		return vol, helpers.PIVolumeProvisioning, nil
	}
}
