package ibm

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMPISnapshot() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPISnapshotCreate,
		Read:     resourceIBMPISnapshotRead,
		Update:   resourceIBMPISnapshotUpdate,
		Delete:   resourceIBMPISnapshotDelete,
		Exists:   resourceIBMPISnapshotExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			//Snapshots are created at the pvm instance level

			helpers.PISnapshotName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name of the snapshot",
			},

			helpers.PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance name / id of the pvm",
			},

			helpers.PIInstanceVolumeIds: {
				Type:             schema.TypeSet,
				Optional:         true,
				Computed:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "List of PI volumes",
			},

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Snapshot description",
			},
			// Computed Attributes

			helpers.PISnapshot: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the snapshot",
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_snapshots": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func resourceIBMPISnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	instanceid := d.Get(helpers.PIInstanceName).(string)
	volids := expandStringList((d.Get(helpers.PIInstanceVolumeIds).(*schema.Set)).List())
	name := d.Get(helpers.PISnapshotName).(string)
	description := d.Get("description").(string)
	if d.Get(description) == "" {
		description = "Testing from Terraform"
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	snapshotBody := &models.SnapshotCreate{Name: &name, Description: description}

	if len(volids) > 0 {
		snapshotBody.VolumeIds = volids
	}

	snapshotResponse, err := client.CreatePvmSnapShot(&p_cloud_p_vm_instances.PcloudPvminstancesSnapshotsPostParams{
		Body: snapshotBody,
	}, instanceid, powerinstanceid, createTimeOut)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, *snapshotResponse.SnapshotID))
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	//return nil
	return resourceIBMPISnapshotRead(d, meta)
}

func resourceIBMPISnapshotRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Calling the Snapshot Read function post create")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	powerinstanceid := parts[0]
	snapshot := st.NewIBMPISnapshotClient(sess, powerinstanceid)
	snapshotdata, err := snapshot.Get(parts[1], powerinstanceid, getTimeOut)

	if err != nil {
		return err
	}

	d.Set(helpers.PISnapshotName, snapshotdata.Name)
	d.Set(helpers.PISnapshot, *snapshotdata.SnapshotID)
	d.Set("status", snapshotdata.Status)
	d.Set("creation_date", snapshotdata.CreationDate.String())
	d.Set("volume_snapshots", snapshotdata.VolumeSnapshots)
	d.Set("last_update_date", snapshotdata.LastUpdateDate.String())

	return nil
}

func resourceIBMPISnapshotUpdate(d *schema.ResourceData, meta interface{}) error {

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

	volrequest, err := client.Update(parts[1], name, size, shareable, powerinstanceid, postTimeOut)
	if err != nil {
		return err
	}

	_, err = isWaitForIBMPIVolumeAvailable(client, *volrequest.VolumeID, powerinstanceid, d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPISnapshotDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]

	client := st.NewIBMPISnapshotClient(sess, powerinstanceid)

	snapshot, err := client.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		return err
	}

	log.Printf("The snapshot  to be deleted is in the following state .. %s", snapshot.Status)
	//_, err = isWaitForIBMPIVolumeAvailable(client, parts[1], powerinstanceid, d.Timeout(schema.TimeoutDelete))
	//if err != nil {
	//	return err
	//}
	snapshotdel_err := client.Delete(parts[1], powerinstanceid, deleteTimeOut)
	if snapshotdel_err != nil {
		return snapshotdel_err
	}

	d.SetId("")
	return nil
}
func resourceIBMPISnapshotExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	powerinstanceid := parts[0]
	client := st.NewIBMPISnapshotClient(sess, powerinstanceid)

	snapshotdelete, err := client.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	log.Printf("Calling the existing function.. %s", *(snapshotdelete.SnapshotID))

	volumeid := *snapshotdelete.SnapshotID
	return volumeid == parts[1], nil
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

*/
func isWaitForPIInstanceSnapshotAvailable(client *st.IBMPISnapshotClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance Snapshot (%s) to be available and active ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"creating_snapshot", "BUILD"},
		Target:     []string{"available", "ACTIVE"},
		Refresh:    isPIInstanceSnapshotRefreshFunc(client, id, powerinstanceid),
		Delay:      30 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    60 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceSnapshotRefreshFunc(client *st.IBMPISnapshotClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		snapshotInfo, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, "", err
		}

		//if pvm.Health.Status == helpers.PIInstanceHealthOk {
		if snapshotInfo.Status == "available" && snapshotInfo.PercentComplete == 100 {
			log.Printf("The snapshot is now available")
			return snapshotInfo, "available", nil

		}
		return snapshotInfo, "in_progress", nil
	}
}
