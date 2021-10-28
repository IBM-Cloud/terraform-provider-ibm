// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

const (
	/* Power Volume creation depends on response from PowerVC */
	volPostTimeOut          = 180 * time.Second
	volGetTimeOut           = 180 * time.Second
	volDeleteTimeOut        = 180 * time.Second
	PIAffinityPolicy        = "pi_affinity_policy"
	PIAffinityVolume        = "pi_affinity_volume"
	PIAffinityInstance      = "pi_affinity_instance"
	PIAntiAffinityInstances = "pi_anti_affinity_instances"
	PIAntiAffinityVolumes   = "pi_anti_affinity_volumes"
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
			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Instance ID - This is the service_instance_id.",
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
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"ssd", "standard", "tier1", "tier3"}),
				Description:  "Type of Disk, required if pi_affinity_policy and pi_volume_pool not provided, otherwise ignored",
			},
			helpers.PIVolumePool: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Volume pool where the volume will be created; if provided then pi_volume_type and pi_affinity_policy values will be ignored",
			},
			PIAffinityPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Affinity policy for data volume being created; ignored if pi_volume_pool provided; for policy affinity requires one of pi_affinity_instance or pi_affinity_volume to be specified; for policy anti-affinity requires one of pi_anti_affinity_instances or pi_anti_affinity_volumes to be specified",
				ValidateFunc: InvokeValidator("ibm_pi_volume", "pi_affinity"),
			},
			PIAffinityVolume: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Volume (ID or Name) to base volume affinity policy against; required if requesting affinity and pi_affinity_instance is not provided",
				ConflictsWith: []string{PIAffinityInstance},
			},
			PIAffinityInstance: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "PVM Instance (ID or Name) to base volume affinity policy against; required if requesting affinity and pi_affinity_volume is not provided",
				ConflictsWith: []string{PIAffinityVolume},
			},
			PIAntiAffinityVolumes: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of volumes to base volume anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_instances is not provided",
				ConflictsWith: []string{PIAntiAffinityInstances},
			},
			PIAntiAffinityInstances: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of pvmInstances to base volume anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_volumes is not provided",
				ConflictsWith: []string{PIAntiAffinityVolumes},
			},

			// Computed Attributes
			"volume_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume ID",
			},
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
func resourceIBMPIVolumeValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 0)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "pi_affinity",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "affinity, anti-affinity"})
	ibmPIVolumeResourceValidator := ResourceValidator{
		ResourceName: "ibm_pi_volume",
		Schema:       validateSchema}
	return &ibmPIVolumeResourceValidator
}
func resourceIBMPIVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	name := d.Get(helpers.PIVolumeName).(string)
	volType := d.Get(helpers.PIVolumeType).(string)
	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	var shared bool
	if v, ok := d.GetOk(helpers.PIVolumeShareable); ok {
		shared = v.(bool)
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	body := models.CreateDataVolume{
		Name:      &name,
		DiskType:  volType,
		Shareable: &shared,
		Size:      &size,
	}
	if v, ok := d.GetOk(helpers.PIVolumePool); ok {
		volumePool := v.(string)
		body.VolumePool = volumePool
	}
	if ap, ok := d.GetOk(PIAffinityPolicy); ok {
		policy := ap.(string)
		body.AffinityPolicy = &policy

		if policy == "affinity" {
			if av, ok := d.GetOk(PIAffinityVolume); ok {
				afvol := av.(string)
				body.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(PIAffinityInstance); ok {
				afins := ai.(string)
				body.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(PIAntiAffinityVolumes); ok {
				afvols := expandStringList(avs.([]interface{}))
				body.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(PIAntiAffinityInstances); ok {
				afinss := expandStringList(ais.([]interface{}))
				body.AntiAffinityPVMInstances = afinss
			}
		}

	}

	resquestParams := p_cloud_volumes.PcloudCloudinstancesVolumesPostParams{
		Body:            &body,
		CloudInstanceID: powerinstanceid,
	}

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)
	vol, err := client.CreateVolume(&resquestParams, powerinstanceid, volPostTimeOut)
	if err != nil {
		return fmt.Errorf("Failed to Create the volume %v", err)
	}

	volumeid := *vol.VolumeID
	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, volumeid))

	_, err = isWaitForIBMPIVolumeAvailable(client, volumeid, powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

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
	if vol.Shareable != nil {
		d.Set(helpers.PIVolumeShareable, vol.Shareable)
	}
	d.Set(helpers.PIVolumeType, vol.DiskType)
	d.Set(helpers.PIVolumePool, vol.VolumePool)
	d.Set("volume_status", vol.State)
	if vol.VolumeID != nil {
		d.Set("volume_id", vol.VolumeID)
	}
	if vol.DeleteOnTermination != nil {
		d.Set("delete_on_termination", vol.DeleteOnTermination)
	}
	d.Set("wwn", vol.Wwn)
	d.Set(helpers.PICloudInstanceId, powerinstanceid)

	return nil
}

func resourceIBMPIVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)
	name := d.Get(helpers.PIVolumeName).(string)
	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	var shareable bool
	if v, ok := d.GetOk(helpers.PIVolumeShareable); ok {
		shareable = v.(bool)
	}

	body := models.UpdateVolume{
		Name:      &name,
		Shareable: &shareable,
		Size:      size,
	}
	updateParams := p_cloud_volumes.PcloudCloudinstancesVolumesPutParams{
		Body:            &body,
		CloudInstanceID: powerinstanceid,
	}
	volrequest, err := client.UpdateVolume(&updateParams, parts[1], powerinstanceid, volPostTimeOut)
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
	voldeleteErr := client.DeleteVolume(parts[1], powerinstanceid, deleteTimeOut)
	if voldeleteErr != nil {
		return voldeleteErr
	}
	_, err = isWaitForIBMPIVolumeDeleted(client, parts[1], powerinstanceid, d.Timeout(schema.TimeoutDelete))
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
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) < 2 {
		return false, fmt.Errorf("Incorrect ID %s: Id should be a combination of powerInstanceID/VolumeID", d.Id())
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

func isWaitForIBMPIVolumeDeleted(client *st.IBMPIVolumeClient, id, powerinstanceid string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", helpers.PIVolumeProvisioning},
		Target:     []string{"deleted"},
		Refresh:    isIBMPIVolumeDeleteRefreshFunc(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    30 * time.Minute,
	}
	return stateConf.WaitForState()
}

func isIBMPIVolumeDeleteRefreshFunc(client *st.IBMPIVolumeClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id, powerinstanceid, volGetTimeOut)
		if err != nil {
			if strings.Contains(err.Error(), "Resource not found") {
				return vol, "deleted", nil
			}
			return nil, "", err
		}
		if vol == nil {
			return vol, "deleted", nil
		}
		return vol, "deleting", nil
	}
}
