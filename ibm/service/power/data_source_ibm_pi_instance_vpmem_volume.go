// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMPIInstanceVpmemVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceVpmemVolumeRead,

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
			Arg_VPMEMVolumeID: {
				Description: "vPMEM volume ID.",
				Required:    true,
				Type:        schema.TypeString,
			},

			// Attributes
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
				Description: "Volume name.",
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
				Description: "Volume size (GB).",
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
	}
}

func dataSourceIBMPIInstanceVpmemVolumeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_vpmem_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	pvmInstanceID := d.Get(Arg_PVMInstanceID).(string)
	vpmemVolumeID := d.Get(Arg_VPMEMVolumeID).(string)

	client := instance.NewIBMPIVPMEMClient(ctx, sess, cloudInstanceID)
	vpmemVolume, err := client.GetPvmVpmemVolume(pvmInstanceID, vpmemVolumeID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPvmVpmemVolume failed: %s", err.Error()), "(Data) ibm_pi_instance_vpmem_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*vpmemVolume.UUID)
	d.Set(Attr_CreationDate, vpmemVolume.CreationDate.String())
	if vpmemVolume.Crn != "" {
		d.Set(Attr_CRN, vpmemVolume.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(vpmemVolume.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of vpmem(%s) user_tags: %s", *vpmemVolume.UUID, err)
		}
		d.Set(Attr_UserTags, tags)
	}
	d.Set(Attr_ErrorCode, vpmemVolume.ErrorCode)
	d.Set(Attr_Href, vpmemVolume.Href)
	d.Set(Attr_Name, vpmemVolume.Name)
	d.Set(Attr_PVMInstanceID, vpmemVolume.PvmInstanceID)
	d.Set(Attr_Reason, vpmemVolume.Reason)
	d.Set(Attr_Size, vpmemVolume.Size)
	d.Set(Attr_Status, vpmemVolume.Status)
	if vpmemVolume.UpdatedDate != nil {
		d.Set(Attr_UpdatedDate, vpmemVolume.UpdatedDate.String())
	}

	return nil
}
