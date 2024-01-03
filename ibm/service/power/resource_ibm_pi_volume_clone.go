// Copyright IBM Corp. 2023, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMPIVolumeClone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumeCloneCreate,
		ReadContext:   resourceIBMPIVolumeCloneRead,
		DeleteContext: resourceIBMPIVolumeCloneDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloud Instance ID - This is the service_instance_id.",
			},
			PIVolumeCloneName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Base name of the new cloned volume(s)",
			},
			PIVolumeIds: {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "List of volumes to be cloned",
			},
			PITargetStorageTier: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Target storage tier for the cloned volumes.",
			},
			helpers.PIReplicationEnabled: {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "By default, cloned volume replication is determined by the source volume unless explicitly specified",
			},

			// Computed attributes
			"volume_clone_task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Clone task ID",
			},
			"cloned_volumes": clonedVolumesSchema(),
			"volume_clone_failure_reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reason the clone volumes task has failed",
			},
			"volume_clone_percent_complete": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Clone task completion percentage",
			},
			"volume_clone_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the clone volumes task",
			},
		},
	}
}

func resourceIBMPIVolumeCloneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	vcName := d.Get(PIVolumeCloneName).(string)
	volids := flex.ExpandStringList((d.Get(PIVolumeIds).(*schema.Set)).List())

	body := &models.VolumesCloneAsyncRequest{
		Name:      &vcName,
		VolumeIDs: volids,
	}

	if v, ok := d.GetOk(PITargetStorageTier); ok {
		body.TargetStorageTier = v.(string)
	}

	if !d.GetRawConfig().GetAttr(helpers.PIReplicationEnabled).IsNull() {
		body.TargetReplicationEnabled = flex.PtrToBool(d.Get(helpers.PIReplicationEnabled).(bool))
	}

	client := st.NewIBMPICloneVolumeClient(ctx, sess, cloudInstanceID)
	volClone, err := client.Create(body)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *volClone.CloneTaskID))

	_, err = isWaitForIBMPIVolumeCloneCompletion(ctx, client, *volClone.CloneTaskID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIVolumeCloneRead(ctx, d, meta)
}

func resourceIBMPIVolumeCloneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, vcTaskID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPICloneVolumeClient(ctx, sess, cloudInstanceID)
	volCloneTask, err := client.Get(vcTaskID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("volume_clone_task_id", vcTaskID)
	if volCloneTask.Status != nil {
		d.Set("volume_clone_status", *volCloneTask.Status)
	}
	d.Set("volume_clone_failure_reason", volCloneTask.FailedReason)
	if volCloneTask.PercentComplete != nil {
		d.Set("volume_clone_percent_complete", *volCloneTask.PercentComplete)
	}
	d.Set("cloned_volumes", flattenClonedVolumes(volCloneTask.ClonedVolumes))

	return nil
}

func resourceIBMPIVolumeCloneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// There is no delete or unset concept for volume clone
	d.SetId("")
	return nil
}

func flattenClonedVolumes(list []*models.ClonedVolume) (cloneVolumes []map[string]interface{}) {
	if list != nil {
		cloneVolumes := make([]map[string]interface{}, len(list))
		for i, data := range list {
			l := map[string]interface{}{
				"clone_volume_id":  data.ClonedVolumeID,
				"source_volume_id": data.SourceVolumeID,
			}
			cloneVolumes[i] = l
		}
		return cloneVolumes
	}
	return
}

func isWaitForIBMPIVolumeCloneCompletion(ctx context.Context, client *st.IBMPICloneVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume clone (%s) to be completed.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{PIVolumeCloneRunning},
		Target:     []string{PIVolumeCloneCompleted},
		Refresh:    isIBMPIVolumeCloneRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeCloneRefreshFunc(client *st.IBMPICloneVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		volClone, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if *volClone.Status == PIVolumeCloneCompleted {
			return volClone, PIVolumeCloneCompleted, nil
		}

		return volClone, PIVolumeCloneRunning, nil
	}
}

func clonedVolumesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"clone_volume_id": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "ID of the new cloned volume",
				},
				"source_volume_id": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "ID of the source volume to be cloned",
				},
			},
		},
	}
}
