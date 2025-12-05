// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMPIInstanceVpmemVolumes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceVpmemVolumesRead,

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_PVMInstanceID: {
				Description: "PCloud PVM instance ID.",
				Required:    true,
				Type:        schema.TypeString,
			},

			// Attributes
			Attr_Volumes: vpmemVolumeSchema(),
		},
	}
}

func dataSourceIBMPIInstanceVpmemVolumesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_vpmem_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	pvmInstanceID := d.Get(Arg_PVMInstanceID).(string)
	client := instance.NewIBMPIVPMEMClient(ctx, sess, cloudInstanceID)
	vpmemVolumes, err := client.GetAllPvmVpmemVolumes(pvmInstanceID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAllPvmVpmemVolumes failed: %s", err.Error()), "(Data) ibm_pi_instance_vpmem_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	volumes := []map[string]any{}
	if vpmemVolumes.Volumes != nil {
		for _, volume := range vpmemVolumes.Volumes {
			vpmemVol := dataSourceIBMPIVPMEMVolumeToMap(volume, meta)
			volumes = append(volumes, vpmemVol)
		}
	}
	d.Set(Attr_Volumes, volumes)

	return nil
}

func dataSourceIBMPIVPMEMVolumeToMap(volume *models.VPMemVolumeReference, meta any) map[string]any {
	vpmemVol := make(map[string]any)
	vpmemVol[Attr_CreationDate] = volume.CreationDate.String()
	if volume.Crn != "" {
		vpmemVol[Attr_CRN] = volume.Crn
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(volume.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of vpmem (%s) user_tags: %s", *volume.UUID, err)
		}
		vpmemVol[Attr_UserTags] = tags
	}
	vpmemVol[Attr_ErrorCode] = volume.ErrorCode
	vpmemVol[Attr_Href] = volume.Href
	vpmemVol[Attr_Name] = volume.Name
	vpmemVol[Attr_PVMInstanceID] = volume.PvmInstanceID
	vpmemVol[Attr_Reason] = volume.Reason
	vpmemVol[Attr_Size] = volume.Size
	vpmemVol[Attr_Status] = volume.Status
	if volume.UpdatedDate != nil {
		vpmemVol[Attr_UpdatedDate] = volume.UpdatedDate.String()
	}
	vpmemVol[Attr_VolumeID] = volume.UUID
	return vpmemVol
}

func vpmemVolumeSchema() *schema.Schema {
	return &schema.Schema{
		Computed:    true,
		Description: "List of vPMEM volumes.",
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				Attr_CreationDate: {
					Computed:    true,
					Description: "The date and time when the volume was created.",
					Type:        schema.TypeString,
				},
				Attr_CRN: {
					Computed:    true,
					Description: "The CRN for this resource.",
					Type:        schema.TypeString,
				},
				Attr_ErrorCode: {
					Computed:    true,
					Description: "Error code for the vPMEM volume.",
					Type:        schema.TypeString,
				},
				Attr_Href: {
					Computed:    true,
					Description: "Link to vPMEM volume resource.",
					Type:        schema.TypeString,
				},
				Attr_Name: {
					Computed:    true,
					Description: "Volume Name.",
					Type:        schema.TypeString,
				},
				Attr_PVMInstanceID: {
					Computed:    true,
					Description: "PVM Instance ID which the volume is attached to.",
					Type:        schema.TypeString,
				},
				Attr_Reason: {
					Computed:    true,
					Description: "Reason for error.",
					Type:        schema.TypeString,
				},
				Attr_Size: {
					Computed:    true,
					Description: "Volume Size (GB).",
					Type:        schema.TypeFloat,
				},
				Attr_Status: {
					Computed:    true,
					Description: "Status of the volume.",
					Type:        schema.TypeString,
				},
				Attr_UpdatedDate: {
					Computed:    true,
					Description: "The date and time when the volume was updated.",
					Type:        schema.TypeString,
				},
				Attr_UserTags: {
					Computed:    true,
					Description: "List of user tags.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Set:  schema.HashString,
					Type: schema.TypeSet,
				},
				Attr_VolumeID: {
					Computed:    true,
					Description: "Volume ID.",
					Type:        schema.TypeString,
				},
			},
		},
	}
}
