// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIVolumesDetach() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumesDetachCreate,
		ReadContext:   resourceIBMPIVolumesDetachRead,
		DeleteContext: resourceIBMPIVolumesDetachDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
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
			Arg_DetachAllVolumes: {
				ConflictsWith: []string{Arg_VolumeIDs},
				Description:   "Indicates if all volumes, except primary boot volume, attached to the pi_instance should be detached (default=false); required if volumeIDs is not provided.",
				ForceNew:      true,
				Optional:      true,
				Type:          schema.TypeBool,
			},
			Arg_DetachPrimaryBootVolume: {
				Description: "Indicates if primary boot volume attached to the pi_instance should be detached (default=false).",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			Arg_VolumeIDs: {
				ConflictsWith: []string{Arg_DetachAllVolumes},
				Description:   "List of volumes to be detached from a  pi_instance; required if detachAllVolumes is not provided.",
				Elem:          &schema.Schema{Type: schema.TypeString},
				ForceNew:      true,
				Optional:      true,
				Type:          schema.TypeList,
			},
		},
	}
}

func resourceIBMPIVolumesDetachCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	pvmInstanceID := d.Get(Arg_InstanceID).(string)
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	body := &models.VolumesDetach{}
	var volumeIDs []string
	if _, ok := d.GetOk(Arg_VolumeIDs); ok {
		for _, v := range d.Get(Arg_VolumeIDs).([]interface{}) {
			volumeIDsItem := v.(string)
			volumeIDs = append(volumeIDs, volumeIDsItem)
		}
		body.VolumeIDs = volumeIDs
	}

	if _, ok := d.GetOk(Arg_DetachAllVolumes); ok {
		body.DetachAllVolumes = flex.PtrToBool(d.Get(Arg_DetachAllVolumes).(bool))
	}
	if _, ok := d.GetOk(Arg_DetachPrimaryBootVolume); ok {
		body.DetachPrimaryBootVolume = flex.PtrToBool(d.Get(Arg_DetachPrimaryBootVolume).(bool))
	}

	volClient := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volinfo, err := volClient.BulkVolumeDetach(pvmInstanceID, body)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Volumes detach accepted:  %s", *volinfo.Summary)
	for _, volumeID := range volumeIDs {
		_, err = isWaitForIBMPIVolumeDetach(ctx, volClient, volumeID, cloudInstanceID, pvmInstanceID, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, pvmInstanceID))

	return resourceIBMPIVolumesAttachRead(ctx, d, meta)
}

func resourceIBMPIVolumesDetachRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIBMPIVolumesDetachDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
