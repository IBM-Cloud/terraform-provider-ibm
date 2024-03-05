// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func ResourceIBMPIVolumesDelete() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumesDeleteCreate,
		ReadContext:   resourceIBMPIVolumesDeleteRead,
		DeleteContext: resourceIBMPIVolumesDeleteDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_VolumeIDs: {
				Description:  "List of volumes to be deleted.",
				Elem:         &schema.Schema{Type: schema.TypeString},
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeList,
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func resourceIBMPIVolumesDeleteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var volumeIDs []string
	for _, v := range d.Get(Arg_VolumeIDs).([]interface{}) {
		volumeIDsItem := v.(string)
		volumeIDs = append(volumeIDs, volumeIDsItem)
	}
	body := &models.VolumesDelete{
		VolumeIDs: volumeIDs,
	}
	volClient := st.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volinfo, err := volClient.BulkVolumeDelete(body)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Volumes delete accepted:  %s", volinfo.Summary)

	for _, v := range d.Get(Arg_VolumeIDs).([]interface{}) {
		_, err = isWaitForIBMPIVolumeDeleted(ctx, volClient, v.(string), d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, strings.Join(volumeIDs, "/")))

	return nil
}

func resourceIBMPIVolumesDeleteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIBMPIVolumesDeleteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
