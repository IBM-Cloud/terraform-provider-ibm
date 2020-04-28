package ibm

import (
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPISnapshot() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPISnapshotsRead,
		Schema: map[string]*schema.Schema{

			helpers.PIInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "PVM Instance Name",
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"snapshot_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_createdate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_lastupdate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_percent_complete": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"snapshot_volumes": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPISnapshotsRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	powerC := instance.NewIBMPISnapshotClient(sess, powerinstanceid)
	snapshotdata, err := powerC.Get(d.Get(helpers.PIInstanceName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	d.SetId(*snapshotdata.SnapshotID)

	d.Set("snapshot_status", snapshotdata.Status)
	d.Set("snapshot_description", snapshotdata.Description)
	d.Set("snapshot_action", snapshotdata.Action)
	d.Set("snapshot_createDate", snapshotdata.CreationDate)
	d.Set("snapshot_lastupdate", snapshotdata.LastUpdateDate)
	d.Set("snapshot_percent_complete", snapshotdata.PercentComplete)
	d.Set("snapshot_volumes", snapshotdata.VolumeSnapshots)

	return nil

}
