// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIVolumeBulk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumeBulkCreate,
		ReadContext:   resourceIBMPIVolumeBulkRead,
		DeleteContext: resourceIBMPIVolumeBulkDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_AffinityInstance: {
				ConflictsWith:    []string{Arg_AffinityVolume},
				Description:      "PVM Instance (ID or Name) to base volume affinity policy against; required if requesting 'affinity' and 'pi_affinity_volume' is not provided.",
				DiffSuppressFunc: flex.ApplyOnce,
				ForceNew:         true,
				Optional:         true,
				Type:             schema.TypeString,
			},
			Arg_AffinityPolicy: {
				Description:      "Affinity policy for data volume being created; ignored if 'pi_volume_pool' provided; for policy 'affinity' requires one of 'pi_affinity_instance' or 'pi_affinity_volume' to be specified; for policy 'anti-affinity' requires one of 'pi_anti_affinity_instances' or 'pi_anti_affinity_volumes' to be specified; Allowable values: 'affinity', 'anti-affinity'.",
				DiffSuppressFunc: flex.ApplyOnce,
				ForceNew:         true,
				Optional:         true,
				Type:             schema.TypeString,
				ValidateFunc:     validate.InvokeValidator("ibm_pi_volume", Arg_AffinityPolicy),
			},
			Arg_AffinityVolume: {
				ConflictsWith:    []string{Arg_AffinityInstance},
				Description:      "Volume (ID or Name) to base volume affinity policy against; required if requesting 'affinity' and 'pi_affinity_instance' is not provided.",
				DiffSuppressFunc: flex.ApplyOnce,
				ForceNew:         true,
				Optional:         true,
				Type:             schema.TypeString,
			},
			Arg_AntiAffinityInstances: {
				ConflictsWith:    []string{Arg_AntiAffinityVolumes},
				Description:      "List of pvmInstances to base volume anti-affinity policy against; required if requesting 'anti-affinity' and 'pi_anti_affinity_volumes' is not provided.",
				DiffSuppressFunc: flex.ApplyOnce,
				Elem:             &schema.Schema{Type: schema.TypeString},
				ForceNew:         true,
				Optional:         true,
				Type:             schema.TypeList,
			},
			Arg_AntiAffinityVolumes: {
				ConflictsWith:    []string{Arg_AntiAffinityInstances},
				Description:      "List of volumes to base volume anti-affinity policy against; required if requesting 'anti-affinity' and 'pi_anti_affinity_instances' is not provided.",
				DiffSuppressFunc: flex.ApplyOnce,
				Elem:             &schema.Schema{Type: schema.TypeString},
				ForceNew:         true,
				Optional:         true,
				Type:             schema.TypeList,
			},
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Count: {
				Default:      1,
				Description:  "Number of volumes to create. Default 1. Maximum is 500 for public workspaces, and 250 for private workspaces.",
				ForceNew:     true,
				Optional:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(1),
			},
			Arg_ReplicationEnabled: {
				Computed:    true,
				Description: "Indicates if the volume should be replication enabled or not.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			Arg_ReplicationSites: {
				Description: "List of replication sites for volume replication.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    true,
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			Arg_UserTags: {
				Description: "The user tags attached to this resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    true,
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			Arg_VolumeName: {
				Description:  "The shared prefix in the name of the volumes.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_VolumePool: {
				Computed:         true,
				Description:      "Volume pool where the volume will be created; if provided then 'pi_affinity_policy' values will be ignored.",
				DiffSuppressFunc: flex.ApplyOnce,
				Optional:         true,
				Type:             schema.TypeString,
			},
			Arg_VolumeShareable: {
				Description: "If set to true, the volume can be shared across Power Systems Virtual Server instances. If set to false, you can attach it only to one instance.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			Arg_VolumeSize: {
				Description:  "The size of the volume in GB.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeFloat,
				ValidateFunc: validation.FloatAtLeast(1),
			},
			Arg_VolumeType: {
				Computed:         true,
				Description:      "Type of disk, if diskType is not provided the disk type will default to 'tier3'",
				DiffSuppressFunc: flex.ApplyOnce,
				Optional:         true,
				Type:             schema.TypeString,
				ValidateFunc:     validate.ValidateAllowedStringValues([]string{"tier0", "tier1", "tier3", "tier5k"}),
			},

			// Attributes
			Attr_Volumes: {
				Computed:    true,
				Description: "List of volumes to create.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Auxiliary: {
							Computed:    true,
							Description: "Indicates if the volume is auxiliary or not.",
							Type:        schema.TypeBool,
						},
						Attr_AuxiliaryVolumeName: {
							Computed:    true,
							Description: "The auxiliary volume name.",
							Type:        schema.TypeString,
						},
						Attr_ConsistencyGroupName: {
							Computed:    true,
							Description: "The consistency group name if volume is a part of volume group.",
							Type:        schema.TypeString,
						},
						Attr_CRN: {
							Computed:    true,
							Description: "The CRN of this resource.",
							Type:        schema.TypeString,
						},
						Attr_DeleteOnTermination: {
							Computed:    true,
							Description: "Indicates if the volume should be deleted when the server terminates.",
							Type:        schema.TypeBool,
						},
						Attr_GroupID: {
							Computed:    true,
							Description: "The volume group id to which volume belongs.",
							Type:        schema.TypeString,
						},
						Attr_IOThrottleRate: {
							Computed:    true,
							Description: "Amount of iops assigned to the volume.",
							Type:        schema.TypeString,
						},
						Attr_MasterVolumeName: {
							Computed:    true,
							Description: "Indicates master volume name",
							Type:        schema.TypeString,
						},
						Attr_MirroringState: {
							Computed:    true,
							Description: "Mirroring state for replication enabled volume",
							Type:        schema.TypeString,
						},
						Attr_PrimaryRole: {
							Computed:    true,
							Description: "Indicates whether 'master'/'auxiliary' volume is playing the primary role.",
							Type:        schema.TypeString,
						},
						Attr_ReplicationStatus: {
							Computed:    true,
							Description: "The replication status of the volume.",
							Type:        schema.TypeString,
						},
						Attr_ReplicationSites: {
							Computed:    true,
							Description: "List of replication sites for volume replication.",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Type:        schema.TypeList,
						},
						Attr_ReplicationType: {
							Computed:    true,
							Description: "The replication type of the volume 'metro' or 'global'.",
							Type:        schema.TypeString,
						},
						Attr_VolumeID: {
							Computed:    true,
							Description: "The unique identifier of the volume.",
							Type:        schema.TypeString,
						},
						Attr_VolumeStatus: {
							Computed:    true,
							Description: "The status of the volume.",
							Type:        schema.TypeString,
						},
						Attr_WWN: {
							Computed:    true,
							Description: "The world wide name of the volume.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func resourceIBMPIVolumeBulkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get(Arg_VolumeName).(string)
	size := int64(d.Get(Arg_VolumeSize).(float64))
	var shared bool
	if v, ok := d.GetOk(Arg_VolumeShareable); ok {
		shared = v.(bool)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	body := &models.MultiVolumesCreate{
		Name:      &name,
		Shareable: &shared,
		Size:      &size,
	}
	body.Count = int64(d.Get(Arg_Count).(int))
	if v, ok := d.GetOk(Arg_VolumeType); ok {
		volType := v.(string)
		body.DiskType = volType
	}
	if v, ok := d.GetOk(Arg_VolumePool); ok {
		volumePool := v.(string)
		body.VolumePool = volumePool
	}
	if v, ok := d.GetOk(Arg_ReplicationEnabled); ok {
		replicationEnabled := v.(bool)
		body.ReplicationEnabled = &replicationEnabled
	}
	if v, ok := d.GetOk(Arg_ReplicationSites); ok {
		if d.Get(Arg_ReplicationEnabled).(bool) {
			body.ReplicationSites = flex.FlattenSet(v.(*schema.Set))
		} else {
			return diag.Errorf("Replication (%s) must be enabled if replication sites are specified.", Arg_ReplicationEnabled)
		}
	}
	if ap, ok := d.GetOk(Arg_AffinityPolicy); ok {
		policy := ap.(string)
		body.AffinityPolicy = &policy

		if policy == Affinity {
			if av, ok := d.GetOk(Arg_AffinityVolume); ok {
				afvol := av.(string)
				body.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(Arg_AffinityInstance); ok {
				afins := ai.(string)
				body.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(Arg_AntiAffinityVolumes); ok {
				afvols := flex.ExpandStringList(avs.([]interface{}))
				body.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(Arg_AntiAffinityInstances); ok {
				afinss := flex.ExpandStringList(ais.([]interface{}))
				body.AntiAffinityPVMInstances = afinss
			}
		}

	}
	if v, ok := d.GetOk(Arg_UserTags); ok {
		body.UserTags = flex.FlattenSet(v.(*schema.Set))
	}

	client := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	vols, err := client.CreateVolumeV2(body)
	if err != nil {
		return diag.FromErr(err)
	}

	// id is a combination of the cloud instance id and all of the volume ids
	id := cloudInstanceID
	for _, vol := range vols.Volumes {
		id += "/" + *vol.VolumeID
	}
	d.SetId(id)

	for _, vol := range vols.Volumes {
		volumeid := *vol.VolumeID
		_, err = isWaitForIBMPIVolumeBulkAvailable(ctx, client, volumeid, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
		if _, ok := d.GetOk(Arg_UserTags); ok {
			if vol.Crn != "" {
				oldList, newList := d.GetChange(Arg_UserTags)
				err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, string(vol.Crn), "", UserTagType)
				if err != nil {
					log.Printf("Error on update of volume (%s) pi_user_tags during creation: %s", volumeid, err)
				}
			}
		}
	}

	return resourceIBMPIVolumeBulkRead(ctx, d, meta)
}

func resourceIBMPIVolumeBulkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	idArr, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := idArr[0]
	client := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volumeIDs := idArr[1:]

	d.Set(Arg_CloudInstanceID, cloudInstanceID)

	// Set Arguments all volumes should have the same information
	// so just get it from one
	firstVolume, err := client.Get(volumeIDs[0])
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set(Arg_VolumePool, firstVolume.VolumePool)
	if firstVolume.Shareable != nil {
		d.Set(Arg_VolumeShareable, firstVolume.Shareable)
	}
	d.Set(Arg_VolumeSize, firstVolume.Size)
	d.Set(Arg_VolumeType, firstVolume.DiskType)
	d.Set(Arg_ReplicationEnabled, firstVolume.ReplicationEnabled)

	result := make([]map[string]interface{}, 0, len(volumeIDs))
	for _, volumeID := range volumeIDs {
		vol, err := client.Get(volumeID)
		if err != nil {
			return diag.FromErr(err)
		}
		l := map[string]interface{}{
			Attr_CRN:                  vol.Crn,
			Attr_Auxiliary:            vol.Auxiliary,
			Attr_AuxiliaryVolumeName:  vol.AuxVolumeName,
			Attr_ConsistencyGroupName: vol.ConsistencyGroupName,
			Attr_GroupID:              vol.GroupID,
			Attr_IOThrottleRate:       vol.IoThrottleRate,
			Attr_MasterVolumeName:     vol.MasterVolumeName,
			Attr_MirroringState:       vol.MirroringState,
			Attr_PrimaryRole:          vol.PrimaryRole,
			Attr_ReplicationSites:     vol.ReplicationSites,
			Attr_ReplicationStatus:    vol.ReplicationStatus,
			Attr_ReplicationType:      vol.ReplicationType,
			Attr_VolumeStatus:         vol.State,
			Attr_WWN:                  vol.Wwn,
		}
		if vol.DeleteOnTermination != nil {
			l[Attr_DeleteOnTermination] = vol.DeleteOnTermination
		}
		if vol.VolumeID != nil {
			l[Attr_VolumeID] = vol.VolumeID
		}
		result = append(result, l)
	}

	d.Set(Attr_Volumes, result)

	return nil
}

func resourceIBMPIVolumeBulkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	idArr, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := idArr[0]
	client := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)

	volumeIDs := idArr[1:]
	volumesDelete := models.VolumesDelete{
		VolumeIDs: volumeIDs,
	}

	volInfo, err := client.BulkVolumeDelete(&volumesDelete)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Volumes delete accepted: %s", volInfo.Summary)

	_, err = isWaitForIBMPIVolumeBulkDeleted(ctx, client, volumeIDs, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		d.SetId(cloudInstanceID + err.Error())
		err = errors.New("error deleting all volumes")
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForIBMPIVolumeBulkAvailable(ctx context.Context, client *instance.IBMPIVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Creating},
		Target:     []string{State_Available},
		Refresh:    isIBMPIVolumeBulkRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeBulkRefreshFunc(client *instance.IBMPIVolumeClient, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.State == State_Available || vol.State == State_InUse {
			return vol, State_Available, nil
		}

		return vol, State_Creating, nil
	}
}

func isWaitForIBMPIVolumeBulkDeleted(ctx context.Context, client *instance.IBMPIVolumeClient, volumeIDs []string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Deleting},
		Target:     []string{State_Deleted, State_Error},
		Refresh:    isIBMPIVolumeBulkDeleteRefreshFunc(client, volumeIDs),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    timeout,
	}
	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeBulkDeleteRefreshFunc(client *instance.IBMPIVolumeClient, volumeIDs []string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		// Every iteration find all volumes that were not deleted
		var leftoverVolumeIDs []string
		for _, volumeID := range volumeIDs {
			vol, err := client.Get(volumeID)
			if err == nil {
				leftoverVolumeIDs = append(leftoverVolumeIDs, *vol.VolumeID)
			}
		}

		// If all volumes were deleted then there are no leftovers
		if len(leftoverVolumeIDs) == 0 {
			return leftoverVolumeIDs, State_Deleted, nil
		} else {
			// Every volume that has not been deleted will be retried.
			volumesDelete := models.VolumesDelete{
				VolumeIDs: leftoverVolumeIDs,
			}
			_, err := client.BulkVolumeDelete(&volumesDelete)
			if err != nil {
				var leftoverIDs string
				for _, leftoverVolumeID := range leftoverVolumeIDs {
					leftoverIDs += "/" + leftoverVolumeID
				}
				leftoverError := errors.New(leftoverIDs)
				return leftoverVolumeIDs, State_Error, leftoverError
			}
			return leftoverVolumeIDs, State_Deleting, nil
		}
	}
}
