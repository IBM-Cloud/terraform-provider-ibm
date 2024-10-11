// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIVolumesAttach() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumesAttachCreate,
		ReadContext:   resourceIBMPIVolumesAttachRead,
		DeleteContext: resourceIBMPIVolumesAttachDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_InstanceID: {
				Description:  "The unique identifier of the instance.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_VolumeIDs: {
				Description: "List of volumes to be detached from a pi_instance; required if detachAllVolumes is not provided.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeList,
			},
			Arg_BootVolumeID: {
				Description: "Primary Boot Volume Id.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPIVolumesAttachCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	body := &models.VolumesAttach{}
	pvmInstanceID := d.Get(Arg_InstanceID).(string)
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)

	var volumeIDs []string
	if _, ok := d.GetOk(Arg_VolumeIDs); ok {
		for _, v := range d.Get(Arg_VolumeIDs).([]interface{}) {
			volumeIDsItem := v.(string)
			volumeIDs = append(volumeIDs, volumeIDsItem)
		}
		body.VolumeIDs = volumeIDs
	}

	if _, ok := d.GetOk(Arg_BootVolumeID); ok {
		body.BootVolumeID = d.Get(Arg_BootVolumeID).(string)
	}
	volClient := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volinfo, err := volClient.BulkVolumeAttach(pvmInstanceID, body)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Volumes attach accepted:  %s", *volinfo.Summary)

	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, pvmInstanceID, strings.Join(volumeIDs, "/")))
	for _, volumeID := range volumeIDs {
		_, err = isWaitForIBMPIVolumeAttachAvailable(ctx, volClient, volumeID, cloudInstanceID, pvmInstanceID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceIBMPIVolumesAttachRead(ctx, d, meta)
}

func resourceIBMPIVolumesAttachRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIBMPIVolumesAttachDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	ids, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID, pvmInstanceID := ids[0], ids[1]
	volClient := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	body := &models.VolumesDetach{
		VolumeIDs: ids[2:],
	}
	volinfo, err := volClient.BulkVolumeDetach(pvmInstanceID, body)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[DEBUG] Volumes delete accepted:  %s", *volinfo.Summary)
	for _, volumeID := range ids[2:] {
		_, err = isWaitForIBMPIVolumeDetach(ctx, volClient, volumeID, cloudInstanceID, pvmInstanceID, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
