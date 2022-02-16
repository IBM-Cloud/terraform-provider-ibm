// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

// Attributes and Arguments defined in data_source_ibm_pi_volume.go
func ResourceIBMPIVolume() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumeCreate,
		ReadContext:   resourceIBMPIVolumeRead,
		UpdateContext: resourceIBMPIVolumeUpdate,
		DeleteContext: resourceIBMPIVolumeDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Instance ID - This is the service_instance_id.",
			},
			PIVolumeName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Volume Name to create",
			},
			PIVolumeShareable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Flag to indicate if the volume can be shared across multiple instances?",
			},
			PIVolumeSize: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Size of the volume in GB",
			},
			PIVolumeType: {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validate.ValidateAllowedStringValues([]string{"ssd", "standard", "tier1", "tier3"}),
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Type of Disk, required if pi_affinity_policy and pi_volume_pool not provided, otherwise ignored",
			},
			PIVolumePool: {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Volume pool where the volume will be created; if provided then pi_volume_type and pi_affinity_policy values will be ignored",
			},
			PIVolumeAffinityPolicy: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Affinity policy for data volume being created; ignored if pi_volume_pool provided; for policy affinity requires one of pi_affinity_instance or pi_affinity_volume to be specified; for policy anti-affinity requires one of pi_anti_affinity_instances or pi_anti_affinity_volumes to be specified",
				ValidateFunc:     validate.InvokeValidator("ibm_pi_volume", "pi_affinity"),
			},
			PIVolumeAffinityVolume: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Volume (ID or Name) to base volume affinity policy against; required if requesting affinity and pi_affinity_instance is not provided",
				ConflictsWith:    []string{PIVolumeAffinityInstance},
			},
			PIVolumeAffinityInstance: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "PVM Instance (ID or Name) to base volume affinity policy against; required if requesting affinity and pi_affinity_volume is not provided",
				ConflictsWith:    []string{PIVolumeAffinityVolume},
			},
			PIVolumeAniAffinityVolumes: {
				Type:             schema.TypeList,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "List of volumes to base volume anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_instances is not provided",
				ConflictsWith:    []string{PIVolumeAntiAffinityInstances},
			},
			PIVolumeAntiAffinityInstances: {
				Type:             schema.TypeList,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "List of pvmInstances to base volume anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_volumes is not provided",
				ConflictsWith:    []string{PIVolumeAniAffinityVolumes},
			},

			// Computed Attributes
			VolumeID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume ID",
			},
			VolumeStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume status",
			},

			VolumeDeleteOnTermination: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Should the volume be deleted during termination",
			},
			VolumeWWN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "WWN Of the volume",
			},
		},
	}
}
func ResourceIBMPIVolumeValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "pi_affinity",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "affinity, anti-affinity"})
	ibmPIVolumeResourceValidator := validate.ResourceValidator{
		ResourceName: "ibm_pi_volume",
		Schema:       validateSchema}
	return &ibmPIVolumeResourceValidator
}

func resourceIBMPIVolumeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get(PIVolumeName).(string)
	size := float64(d.Get(PIVolumeSize).(float64))
	var shared bool
	if v, ok := d.GetOk(PIVolumeShareable); ok {
		shared = v.(bool)
	}
	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	body := &models.CreateDataVolume{
		Name:      &name,
		Shareable: &shared,
		Size:      &size,
	}
	if v, ok := d.GetOk(PIVolumeType); ok {
		volType := v.(string)
		body.DiskType = volType
	}
	if v, ok := d.GetOk(PIVolumePool); ok {
		volumePool := v.(string)
		body.VolumePool = volumePool
	}
	if ap, ok := d.GetOk(PIVolumeAffinityPolicy); ok {
		policy := ap.(string)
		body.AffinityPolicy = &policy

		if policy == "affinity" {
			if av, ok := d.GetOk(PIVolumeAffinityVolume); ok {
				afvol := av.(string)
				body.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(PIVolumeAffinityInstance); ok {
				afins := ai.(string)
				body.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(PIVolumeAniAffinityVolumes); ok {
				afvols := flex.ExpandStringList(avs.([]interface{}))
				body.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(PIVolumeAntiAffinityInstances); ok {
				afinss := flex.ExpandStringList(ais.([]interface{}))
				body.AntiAffinityPVMInstances = afinss
			}
		}

	}

	client := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	vol, err := client.CreateVolume(body)
	if err != nil {
		return diag.FromErr(err)
	}

	volumeid := *vol.VolumeID
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, volumeid))

	_, err = isWaitForIBMPIVolumeAvailable(ctx, client, volumeid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIVolumeRead(ctx, d, meta)
}

func resourceIBMPIVolumeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, volumeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)

	vol, err := client.Get(volumeID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set(PIVolumeName, vol.Name)
	d.Set(PIVolumeSize, vol.Size)
	if vol.Shareable != nil {
		d.Set(PIVolumeShareable, vol.Shareable)
	}
	d.Set(PIVolumeType, vol.DiskType)
	d.Set(PIVolumePool, vol.VolumePool)
	d.Set(VolumeStatus, vol.State)
	if vol.VolumeID != nil {
		d.Set(VolumeID, vol.VolumeID)
	}
	if vol.DeleteOnTermination != nil {
		d.Set(VolumeDeleteOnTermination, vol.DeleteOnTermination)
	}
	d.Set(VolumeWWN, vol.Wwn)
	d.Set(PICloudInstanceID, cloudInstanceID)

	return nil
}

func resourceIBMPIVolumeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, volumeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	name := d.Get(PIVolumeName).(string)
	size := float64(d.Get(PIVolumeSize).(float64))
	var shareable bool
	if v, ok := d.GetOk(PIVolumeShareable); ok {
		shareable = v.(bool)
	}

	body := &models.UpdateVolume{
		Name:      &name,
		Shareable: &shareable,
		Size:      size,
	}
	volrequest, err := client.UpdateVolume(volumeID, body)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = isWaitForIBMPIVolumeAvailable(ctx, client, *volrequest.VolumeID, d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIVolumeRead(ctx, d, meta)
}

func resourceIBMPIVolumeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, volumeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	err = client.DeleteVolume(volumeID)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = isWaitForIBMPIVolumeDeleted(ctx, client, volumeID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func isWaitForIBMPIVolumeAvailable(ctx context.Context, client *st.IBMPIVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", "creating"},
		Target:     []string{"available"},
		Refresh:    isIBMPIVolumeRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeRefreshFunc(client *st.IBMPIVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "available" || vol.State == "in-use" {
			return vol, "available", nil
		}

		return vol, "creating", nil
	}
}

func isWaitForIBMPIVolumeDeleted(ctx context.Context, client *st.IBMPIVolumeClient, id string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "creating"},
		Target:     []string{"deleted"},
		Refresh:    isIBMPIVolumeDeleteRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    timeout,
	}
	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeDeleteRefreshFunc(client *st.IBMPIVolumeClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
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
